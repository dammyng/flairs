package rest

import (
	"auth/grpc/authclient"
	"auth/libs/persistence"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"shared/events"
	amqp "shared/events/amqp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

	"golang.org/x/crypto/bcrypt"

	"shared/helper"
	"shared/models/appuser"
	"strings"

	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)



// ServiceHandler represent routes dependencies
type ServiceHandler struct {
	EventEmitter amqp.EventEmitter
	RedisConn    redis.Conn
}



type Claims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

// NewServiceHandler : Service handler constructor
func NewServiceHandler(eventEmitter amqp.EventEmitter, redisConn redis.Conn) ServiceHandler {
	return ServiceHandler{
		EventEmitter: eventEmitter,
		RedisConn:    redisConn,
	}
}

//Register a new user
func (serviceHandler ServiceHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload EmailPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		msg := InvalidRequest
		fmt.Println(msg)
		helper.DisplayAppError(w, err, msg, http.StatusUnprocessableEntity)
		return
	}
	payload.Email = strings.ToLower(payload.Email)

	// Find user
	u, err := authclient.GetAUser(&appuser.UserArg{UserPayload: &appuser.User{Email: payload.Email}})
	if u != nil {
		helper.DisplayAppError(
			w,
			fmt.Errorf("Duplicate entry Error"),
			"A Flairs account with this email already exists. Please try logging in.",
			http.StatusUnauthorized,
		)
		return
	}
	if u == nil && err.Error() != "rpc error: code = Unknown desc = record not found" {
		helper.DisplayAppError(
			w, err,
			ProcessingRequestError,
			http.StatusInternalServerError,
		)
		return
	}

	// User was not found so we create a new user instance
	ID := uuid.NewV4()
	tempPass := helper.RandString(20)

	hashedPass, err := getPasswordHash(tempPass)
	if err != nil {
		msg := ProcessingRequestError
		helper.DisplayAppError(w, err, msg, http.StatusInternalServerError)
	}
	user := appuser.User{
		ID:       ID.String(),
		Email:    payload.Email,
		Password: hashedPass,
	}

	// Check for referral in link
	ref := r.URL.Query().Get("ref")

	refcode := helper.RandString(15)

	user.RefCode = refcode
	if len(ref) > 0 {

		user.Referrer = ref
	}
	if _, err := authclient.AddNewUser(&appuser.UserArg{UserPayload: &user}); err != nil {
		helper.DisplayAppError(w, err, UserCreationError, http.StatusBadRequest)
		return
	}
	// Send user the message token to verify email
	token := helper.RandInt(6)
	log.Println(token)
	_, err = serviceHandler.RedisConn.Do("HMSET", "email:verification", user.Email, token)

	// TODO how to handle this error
	if err != nil {
		log.Fatal(err)
	}
	_, err = serviceHandler.RedisConn.Do("HMSET", "password:reset", user.Email, token)
	if err != nil {
		log.Fatal(err)
	}
	// Do it with Rabbit

	data := helper.HttpResponse{
		Message: UserCreationSuccessful,
		Code:    http.StatusCreated,
	}
	helper.WriteJsonResponse(w, data, http.StatusOK)
}

func getPasswordHash(password string) ([]byte, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return b, err
}

//Register a new user
func (serviceHandler ServiceHandler) Login(w http.ResponseWriter, r *http.Request) {

	var payload AuthPayload
	var err error
	if err = json.NewDecoder(r.Body).Decode(&payload); err != nil {
		helper.DisplayAppError(
			w,
			err,
			InvalidRequest,
			http.StatusUnprocessableEntity,
		)
		return
	}

	payload.Email = strings.ToLower(payload.Email)

	user, err := authclient.GetAUser(&appuser.UserArg{UserPayload: &appuser.User{Email: payload.Email}})
	if user == nil {
		helper.DisplayAppError(
			w,
			err,
			"An account with this email does exist on Flairs. Please try logging in.",
			http.StatusUnauthorized,
		)
		return
	}
	if user == nil && err != nil {
		helper.DisplayAppError(w, err, InvalidCredentialError, http.StatusUnauthorized)
		return
	}

	verifiedAt := user.EmailVerifiedAt
	emailVerfiedAt, _ := time.Parse(time.RFC3339, verifiedAt)

	if emailVerfiedAt.IsZero() {
		helper.DisplayAppError(w, fmt.Errorf("User unverified"), UnverifiedEmailError, http.StatusFailedDependency)
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(payload.Password))
	if err != nil {
		helper.DisplayAppError(
			w,
			err,
			InvalidPassword,
			http.StatusUnauthorized,
		)
		return
	}

	expirationTime := time.Now().Add(24 * 60 * time.Minute)

	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWTKey")))
	//fmt.Println("token :", tokenString)
	if err != nil {
		helper.DisplayAppError(w, err, InternalServerError, http.StatusUnauthorized)
		return
	}
	resp := struct {
		Token string       `json:"token"`
		User  appuser.User `json:"user"`
	}{
		tokenString,
		persistence.CleanJson(*user),
	}
	helper.WriteJsonResponse(w, resp, http.StatusOK)
}

func (serviceHandler ServiceHandler) Find(w http.ResponseWriter, r *http.Request) {

	responce, err := authclient.Connect()

	if err != nil {
		msg := InvalidRequest
		helper.DisplayAppError(w, err, msg, http.StatusUnprocessableEntity)
	}
	helper.WriteJsonResponse(w, responce, http.StatusOK)
	ID := uuid.NewV4()

	msg := events.UserCreatedEvent{
		ID:    hex.EncodeToString(ID.Bytes()),
		Host:  "http://localhost:15672/",
		Email: "user.Email",
		Token: "Token",
	}

	serviceHandler.EventEmitter.Emit(&msg, "auth")
}

func (serviceHandler ServiceHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	email, ok := vars["email"]
	email = strings.ToLower(email)
	if !ok {
		helper.DisplayAppError(w, fmt.Errorf("invalid url email argument not supplied"), " Invalid request",
			http.StatusBadRequest)
		return
	}
	requestToken, ok := vars["token"]
	if !ok {
		helper.DisplayAppError(w, fmt.Errorf("Invalid url token argument not supplied"), "error processing request",
			http.StatusBadRequest)
		return
	}
	token, err := redis.String(serviceHandler.RedisConn.Do("HGET", "email:verification", email))
	if err != nil {
		if err == redis.ErrNil {
			helper.DisplayAppError(w, fmt.Errorf("Invalid or incorrect token"),
				"error processing request", http.StatusBadRequest)
			return

		}

		log.Fatal(err)
	}
	if token == "" {
		helper.DisplayAppError(w, fmt.Errorf("Invalid or incorrect token"),
			"error processing request", http.StatusBadRequest)
		return
	}

	//log.Printf("found token is %v \n", token)
	//fmt.Printf("requestToken token is %v \n", requestToken)

	if match := token == requestToken; match {
		user, err := authclient.GetAUser(&appuser.UserArg{UserPayload: &appuser.User{Email: email}})

		if err != nil {
			helper.DisplayAppError(w, err, "error fetching user record ", http.StatusExpectationFailed)

			return
		}

		currentTime := time.Now()
		//fmt.Println("current time is %v", currentTime)
		_, err = authclient.UpdateUser(&appuser.UpdateArg{NewObj: &appuser.User{EmailVerifiedAt: currentTime.Format(time.RFC3339)}, OldUser: user})
		// TODO handle error
		if err != nil {
			log.Fatal(err)
		}
		redis.Int(serviceHandler.RedisConn.Do("HDEL", "email:verification", email))
		helper.WriteJsonResponse(w, map[string]interface{}{"message": "successfully verified email address"}, http.StatusOK)

	} else {
		helper.DisplayAppError(w, fmt.Errorf("Not Found"), UserNotFound, http.StatusNotFound)
	}
}

func (serviceHandler ServiceHandler) SetPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token, ok := vars["token"]
	if !ok {
		helper.DisplayAppError(w, fmt.Errorf("Invalid Request"), InvalidRequest, http.StatusBadRequest)
		return
	}
	email, ok := vars["email"]
	if !ok {
		helper.DisplayAppError(w, fmt.Errorf("Invalid Request"), InvalidRequest, http.StatusBadRequest)
		return
	}

	email = strings.ToLower(email)

	var payload AuthPayload
	var err error
	if err = json.NewDecoder(r.Body).Decode(&payload); err != nil {
		helper.DisplayAppError(
			w,
			err,
			InvalidRequest,
			http.StatusBadRequest,
		)
		return
	}
	if len(payload.Password) < 8 {
		helper.DisplayAppError(w, fmt.Errorf("Invalid request data"), InvalidPassword, http.StatusBadRequest)
		return
	}
	payload.Email = strings.ToLower(payload.Email)
	user, err := authclient.GetAUser(&appuser.UserArg{UserPayload: &appuser.User{Email: payload.Email}})

	if err != nil {
		helper.DisplayAppError(
			w,
			err,
			InvalidCredentialError,
			http.StatusUnauthorized,
		)
		return
	}

	verifiedAt := user.EmailVerifiedAt
	emailVerfiedAt, _ := time.Parse(time.RFC3339, verifiedAt)

	if emailVerfiedAt.IsZero() {
		helper.DisplayAppError(w, fmt.Errorf("User unverified"), UnverifiedEmailError, http.StatusFailedDependency)
		return
	}

	ptoken, err := redis.String(serviceHandler.RedisConn.Do("HGET", "password:reset", email))
	if err != nil {
		if err == redis.ErrNil {
			helper.DisplayAppError(w, fmt.Errorf("Invalid or incorrect token"),
				"error processing request", http.StatusBadRequest)
			return

		}
		log.Fatal(err)
	}
	if ptoken == "" {
		helper.DisplayAppError(w, fmt.Errorf("Invalid or incorrect token"),
			"error processing request", http.StatusBadRequest)
		return
	}
	if ptoken != token {
		//TODO add error message
		helper.DisplayAppError(w, fmt.Errorf("Invalid or incorrect token"),
			"You supplied an incorret token", http.StatusBadRequest)

		return
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.DisplayAppError(w, err, InternalServerError, http.StatusBadRequest)
		return
	}
	_, err = authclient.UpdateUser(&appuser.UpdateArg{OldUser: user, NewObj: &appuser.User{Password: hashedPass}})
	if err != nil {
		helper.DisplayAppError(w, err, ProcessingRequestError, http.StatusBadRequest)
		return
	}
	redis.Int(serviceHandler.RedisConn.Do("HDEL", "password:reset", email))
	helper.WriteJsonResponse(w, map[string]interface{}{"message": RecordUpdateSuccessful}, http.StatusOK)
}


func (serviceHandler ServiceHandler) UserData(w http.ResponseWriter, r *http.Request) {

	userId, ok := r.Context().Value("user").(string)

	if !ok {
		helper.DisplayAppError(w,
			fmt.Errorf("Invalid request"),
			InvalidRequest,
			http.StatusForbidden,
		)
		return
	}
	user, err := authclient.GetAUser(&appuser.UserArg{UserPayload: &appuser.User{ID: userId}})
	if err != nil {
		helper.DisplayAppError(w,
			err,
			InternalServerError,
			http.StatusBadRequest,
		)
		return
	}

	if r.Method == "GET" {
		helper.WriteJsonResponse(w, persistence.CleanJson(*user), http.StatusOK)
		return
	}

	var payload UpdateUserDataPayload
	err = helper.DecodeRequestData(w, r, &payload)
	if err != nil {
		helper.DisplayAppError(w,
			err,
			InvalidRequest,
			http.StatusBadRequest,
		)

		return
	}
	userData := appuser.User{}
	b, err := json.Marshal(payload)
	if err != nil {
		helper.DisplayAppError(w, err, InvalidRequest, http.StatusBadRequest)
		return

		//log.Fatalf(err.Error())
	}
	err = json.Unmarshal(b, &userData)
	if err != nil {
		helper.DisplayAppError(w, err, InvalidRequest, http.StatusBadRequest)
		return

		//log.Fatalf(err.Error())
	}

	dob := payload.DOB
	if dob != "" {
		//layout := "2006-01-02"
		t, err := time.Parse(time.RFC3339, dob)

		if err != nil {
			helper.DisplayAppError(w, err, "Invalid date", http.StatusUnprocessableEntity)
			return

		}
		userData.DOB =  t.Format(time.RFC3339)
	}

	phoneVerifiedAt := payload.PhoneVerifiedAt

	if phoneVerifiedAt != "" {
		t, err := time.Parse(time.RFC3339, phoneVerifiedAt)
		if err != nil {
			helper.DisplayAppError(w, err, "Invalid timestamp", http.StatusUnprocessableEntity)
			return
		}
		userData.PhoneVerifiedAt = t.Format(time.RFC3339)
	}

	//userData.LastCardRequest = cardReqId
	if len(payload.Pin) > 0 {
		if len(payload.Pin) < 4 {
			helper.DisplayAppError(w,
				errors.New("Invalid input length"),
				"Pin must be four or more digits",
				http.StatusBadRequest,
			)
			return

		}
		hashedPin, err := getPasswordHash(payload.Pin)
		if err != nil {
			log.Fatalf(err.Error())
		}
		userData.Pin = hashedPin

	}
	_, err = authclient.UpdateUser(&appuser.UpdateArg{NewObj: &userData, OldUser: user})
	if err != nil {
		helper.DisplayAppError(w,
			err,
			UpdateRecordError,
			http.StatusBadRequest,
		)
		return
	}
	isProfileComplete := persistence.IsProfileComplete(*user)

	_, err = authclient.UpdateUser(&appuser.UpdateArg{NewObj: &appuser.User{IsProfileCompleted: isProfileComplete} , OldUser: user }) 
	
	if err != nil {
		helper.DisplayAppError(w,
			err,
			UpdateRecordError,
			http.StatusBadRequest,
		)
		return
	}

	helper.WriteJsonResponse(w, map[string]interface{}{"message": RecordUpdateSuccessful, "data": persistence.CleanJson(*user)}, http.StatusOK)

}


func (serviceHandler ServiceHandler) CardRequest(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("user").(string)

	if !ok {
		helper.DisplayAppError(w, fmt.Errorf("Invalid request"), ProcessingRequestError, http.StatusForbidden)
		return
	}
	user, err := authclient.GetAUser(&appuser.UserArg{UserPayload: &appuser.User{ID: userId}})
	if err != nil {
		helper.DisplayAppError(w,
			err,
			InternalServerError,
			http.StatusBadRequest,
		)
		return
	}
	if r.Method == "GET" {
		cr, err := authclient.GetUserCardRequest(&appuser.CardRequest{UserId: userId})
		if err != nil {
			helper.DisplayAppError(
				w, err, InternalServerError, http.StatusBadRequest)
		}
		helper.WriteJsonResponse(w, map[string]interface{}{"card_requests": cr}, http.StatusOK)
		return
	}
	var payload CardRequestPayload

	err = helper.DecodeRequestData(w, r, &payload)
	if err != nil {
		helper.DisplayAppError(w,
			err,
			InvalidRequest,
			http.StatusBadRequest,
		)

		return
	}

	cardReqId := uuid.NewV4().String()

	cardReq := appuser.CardRequest{
		ID:       cardReqId,
		Color:    payload.Color,
		Currency: payload.Currency,
		UserId:   userId,
	}
	_, err = authclient.CreateCardRequest(&cardReq)
	if err != nil {
		helper.DisplayAppError(w, err, ProcessingRequestError, http.StatusBadRequest)
		return

	}
	userData := appuser.User{
		LastCardRequested: cardReqId,
	}

	_, err = authclient.UpdateUser(&appuser.UpdateArg{NewObj: &userData, OldUser: user})

	helper.WriteJsonResponse(w, map[string]interface{}{"message": CardRequestSuccessful, "card_request": cardReq}, http.StatusOK)

}


func (serviceHandler ServiceHandler) Wallet(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("user").(string)

	if !ok {
		helper.DisplayAppError(w, fmt.Errorf("Invalid request"), ProcessingRequestError, http.StatusForbidden)
		return
	}
	_, err := authclient.GetAUser(&appuser.UserArg{UserPayload: &appuser.User{ID: userId}})
	if err != nil {
		helper.DisplayAppError(w,
			err,
			InternalServerError,
			http.StatusBadRequest,
		)
		return
	}
	if r.Method == "GET" {
		cr, err := authclient.GetUserWallets(&appuser.WalletArg{UserId: userId})
		if err != nil {
			helper.DisplayAppError(
				w, err, InternalServerError, http.StatusBadRequest)
		}
		helper.WriteJsonResponse(w, map[string]interface{}{"card_requests": cr}, http.StatusOK)
		return
	}
	var payload WalletPayload

	err = helper.DecodeRequestData(w, r, &payload)
	if err != nil {
		helper.DisplayAppError(w,
			err,
			InvalidRequest,
			http.StatusBadRequest,
		)

		return
	}

	walletId := uuid.NewV4().String()
	
	newWallet := appuser.Wallet{
		ID: walletId,
		UserId:   userId,
	}
	_, err = authclient.CreateNewWallet(&newWallet)
	if err != nil {
		helper.DisplayAppError(w, err, ProcessingRequestError, http.StatusBadRequest)
		return

	}
	
	helper.WriteJsonResponse(w, map[string]interface{}{"message": "Done"}, http.StatusOK)

}