package v1

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/metadata"

	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1internals "auth/internals/v1"
	v1 "auth/pkg/api/v1"
	redisconn "auth/redis"

	"auth/libs/setup"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var testDb *gorm.DB
var testRedis redis.Conn

func TestMain(m *testing.M) {

	initDB()
	initRedis()
	code := m.Run()
	//clearDB()
	os.Exit(code)
}
func initDB() {
	var err error
	dsn := fmt.Sprintf("root:password@tcp(127.0.0.1)/alpha_plus?charset=utf8&parseTime=True&loc=Local")
	testDb, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Panicf("flairServiceServer.AddNewUser() error = %v,", err)
	}
	testDb.Exec(setup.SQLMode)
}

func initRedis() {
	redisPool := redisconn.NewPool("localhost:6379")
	testRedis = redisPool.Get()
}

func TestAddUser_ok(t *testing.T) {
	clearUsersTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewFlairsServiceServer(sqlLayer, testRedis)

	req := &v1.AddNewUserRequest{
		Email: "someone@flairs.com",
		Ref:   "dddddd",
	}

	got, err := s.AddNewUser(ctx, req)

	if err != nil {
		t.Errorf("flairServiceServer.AddNewUser() error = %v, wantErr %v", err, "f")
		return
	}

	var u v1internals.User
	testDb.Where("email = ?", req.Email).Last(&u)

	if err == nil && !reflect.DeepEqual(got.ID, u.ID) {
		t.Errorf("flairServiceServer.AddNewUser() = %v, want %v", got.ID, u.ID)
	}
}

func TestAddUser_duplicate_email(t *testing.T) {
	clearUsersTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewFlairsServiceServer(sqlLayer, testRedis)

	req := &v1.AddNewUserRequest{
		Email: "someone@flairs.com",
		Ref:   "dddddd",
	}

	_, err := s.AddNewUser(ctx, req)
	_, duplicateErr := s.AddNewUser(ctx, req)

	if err != nil {
		t.Errorf("flairServiceServer.AddNewUser() error = %v, wantErr %v", err, "f")
		return
	}

	if duplicateErr == nil {
		t.Errorf("flairServiceServer.AddNewUser() should have returned a duplicate error  = %v, wantErr %v", err, "f")
		return
	}

	if (duplicateErr != nil) && duplicateErr.Error() != status.Error(codes.AlreadyExists, "A Flairs account with this email already exists. Please try logging in.").Error() {
		t.Errorf("flairServiceServer.AddNewUser() should return a duplicate error but got = %v,", duplicateErr.Error())
	}
}

func TestAddUser_invalid_entry(t *testing.T) {
	clearUsersTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewFlairsServiceServer(sqlLayer, testRedis)

	req := &v1.AddNewUserRequest{
		Ref: "dddddd",
	}

	_, err := s.AddNewUser(ctx, req)

	if err == nil {
		t.Errorf("flairServiceServer.AddNewUser() should have returned an internal server error  = %v", err)
		return
	}

	if (err != nil) && err.Error() != status.Error(codes.InvalidArgument, "Invalid entry").Error() {
		t.Errorf("flairServiceServer.AddNewUser() Wrong error! should return a internal error but got = %v,", err.Error())
	}
}

func TestAddUser_ok_no_ref(t *testing.T) {
	clearUsersTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewFlairsServiceServer(sqlLayer, testRedis)

	req := &v1.AddNewUserRequest{
		Email: "someone@flairs.com",
	}

	got, err := s.AddNewUser(ctx, req)

	if err != nil {
		t.Errorf("flairServiceServer.AddNewUser() error = %v, wantErr %v", err, "f")
		return
	}

	var u v1internals.User
	testDb.Where("email = ?", req.Email).Last(&u)

	if err == nil && !reflect.DeepEqual(got.ID, u.ID) {
		t.Errorf("flairServiceServer.AddNewUser() = %v, want %v", got.ID, u.ID)
	}
}

func TestVerifyEmail_ok(t *testing.T) {

	clearUsersTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewFlairsServiceServer(sqlLayer, testRedis)

	uReq := &v1.AddNewUserRequest{
		Email: "someone@flairs.com",
	}

	got, err := s.AddNewUser(ctx, uReq)

	if err != nil {
		t.Errorf("flairServiceServer.ValidateUserEmail() failed User account could not be created failed with = %v", err.Error())
	}

	token, err := redis.String(testRedis.Do("HGET", "email:verification", uReq.Email))
	if err != nil {
		t.Errorf("flairServiceServer.ValidateUserEmail() Redis token was not saved when user got created with = %v", err.Error())
	}
	req := &v1.ValidateEmailRequest{
		Token: token,
		Email: uReq.Email,
	}

	vGot, err := s.ValidateUserEmail(ctx, req)

	if err != nil {
		t.Errorf("flairServiceServer.ValidateUserEmail() failed with = %v", err.Error())
	}

	want := &v1.CustomResponse{
		Message: "Successfully verified email",
		Request: "verify_email",
	}

	if err == nil && !reflect.DeepEqual(vGot, want) {
		t.Errorf("flairServiceServer.ValidateUserEmail() returned = %v, want %v", got, want)
	}

}

func TestVerifyEmail_wrongtoken(t *testing.T) {

	clearUsersTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewFlairsServiceServer(sqlLayer, testRedis)

	uReq := &v1.AddNewUserRequest{
		Email: "someone@flairs.com",
	}

	_, err := s.AddNewUser(ctx, uReq)

	if err != nil {
		t.Errorf("flairServiceServer.ValidateUserEmail() failed User account could not be created failed with = %v", err.Error())
	}

	token, err := redis.String(testRedis.Do("HGET", "email:verification", uReq.Email))
	if err != nil {
		t.Errorf("flairServiceServer.ValidateUserEmail() Redis token was not saved when user got created with = %v", err.Error())
	}
	req := &v1.ValidateEmailRequest{
		Token: token + "xxx",
		Email: uReq.Email,
	}

	vGot, err := s.ValidateUserEmail(ctx, req)

	if err == nil {
		t.Errorf("flairServiceServer.ValidateUserEmail( wrong token) is expected to return an error but got = %v", vGot.Message)
	}

	if err != nil && !reflect.DeepEqual(err.Error(), status.Error(codes.InvalidArgument, "Wrong token string").Error()) {
		t.Errorf("flairServiceServer.ValidateUserEmail(wrong token) expected a wrong token error but returned = %v, want %v", err.Error(), status.Error(codes.InvalidArgument, "Wrong token string").Error())
	}

}

func TestAddPassword_ok(t *testing.T) {
	clearUsersTable()
	ctx := context.Background()

	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewFlairsServiceServer(sqlLayer, testRedis)

	uReq := &v1.AddNewUserRequest{
		Email: "someone@flairs.com",
	}

	_, err := s.AddNewUser(ctx, uReq)
	if err != nil {
		t.Errorf("flairServiceServer.AddPassword() failed User account could not be created failed with = %v", err.Error())
	}

	var u v1internals.User
	testDb.Where("email = ?", uReq.Email).Last(&u)

	err = sqlLayer.UpdateUser(&v1.User{ID: u.ID, Email: u.Email}, &v1.User{EmailVerifiedAt: time.Now().Format(time.RFC3339)})

	if err != nil {
		t.Errorf("flairServiceServer.Addpassword() failed - Could not verify email failed with = %v", err.Error())
	}

	token, err := redis.String(testRedis.Do("HGET", "password:reset", uReq.Email))
	if err != nil {
		t.Errorf("flairServiceServer.AddPassword() Redis token was not saved when user got created with = %v", err.Error())
	}
	req := &v1.SetPasswordRequest{
		Email:    uReq.Email,
		Password: "Password",
	}

	md := metadata.Pairs("authorization", token)
	ctx = metadata.NewIncomingContext(ctx, md)

	vGot, err := s.SetUserPassword(ctx, req)

	if err != nil {
		t.Errorf("flairServiceServer.AddPassword(ok) failed with error = %v", err)
	}

	want := &v1.CustomResponse{
		Message: "Successfully Add Password",
		Request: "add_password",
	}

	if err == nil && !reflect.DeepEqual(vGot, want) {
		t.Errorf("flairServiceServer.AddPassword() returned = %v, want %v", vGot, want)
	}

}

func TestLogin_ok(t *testing.T) {
	clearUsersTable()
	ctx := context.Background()

	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewFlairsServiceServer(sqlLayer, testRedis)

	uReq := &v1.AddNewUserRequest{
		Email: "someone@flairs.com",
	}

	_, err := s.AddNewUser(ctx, uReq)
	if err != nil {
		t.Errorf("flairServiceServer.Login() failed User account could not be created failed with = %v", err.Error())
	}

	var u v1internals.User
	testDb.Where("email = ?", uReq.Email).Last(&u)
	hashedPass, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	err = sqlLayer.UpdateUser(&v1.User{ID: u.ID, Email: u.Email}, &v1.User{EmailVerifiedAt: time.Now().Format(time.RFC3339), Password: hashedPass})

	if err != nil {
		t.Errorf("flairServiceServer.Login() failed - Could not verify email && password failed with = %v", err.Error())
	}

	req := &v1.LoginRequest{
		Email:    uReq.Email,
		Password: "password",
	}

	vGot, err := s.LoginUser(ctx, req)

	if err != nil {
		t.Errorf("flairServiceServer.Login(ok) failed with error = %v", err)
		return
	}

	claims := &Claims{}
	err = DecodeJwt(vGot.Token, claims)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			t.Errorf("flairServiceServer.Login() returned = an invalid token - could not verufy tokenv %v", vGot.Token)
			return
		}
		t.Errorf("flairServiceServer.Login() returned = an could not decode token %v", err)
	}
	if claims.Valid() != nil {
		t.Errorf("flairServiceServer.Login() returned a wrong token string id= %v", vGot.Token)
	}
	if claims.UserID != vGot.User.ID {
		t.Errorf("flairServiceServer.Login() returned an invalid user with id= %v, want %v", claims.UserID, vGot.User.ID)
	}
}

func TestUpdateUser_ok(t *testing.T) {
	clearUsersTable()
	ctx := context.Background()

	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewFlairsServiceServer(sqlLayer, testRedis)

	uReq := &v1.AddNewUserRequest{
		Email: "someone@flairs.com",
	}

	_, err := s.AddNewUser(ctx, uReq)
	if err != nil {
		t.Errorf("flairServiceServer.UpdateUser() failed User account could not be created failed with = %v", err.Error())
	}

	var u v1internals.User
	testDb.Where("email = ?", uReq.Email).Last(&u)
	hashedPass, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	err = sqlLayer.UpdateUser(&v1.User{ID: u.ID, Email: u.Email}, &v1.User{EmailVerifiedAt: time.Now().Format(time.RFC3339), Password: hashedPass})

	if err != nil {
		t.Errorf("flairServiceServer.UpdateUser() failed - Could not verify email && password failed with = %v", err.Error())
	}
	req := &v1.LoginRequest{
		Email:    uReq.Email,
		Password: "password",
	}

	vGot, err := s.LoginUser(ctx, req)

	if err != nil {
		t.Errorf("flairServiceServer.UpdateUser(ok) Users could not login = %v", err)
		return
	}
	
	uReqt := &v1.UpdateUserRequest{
		Id: u.ID,
		Profile: &v1.UpdateBody{
			City: "Lagos",
			FirstName: "Dami",
			DOB: "2020-10-19T02:11:07+00:00",
		},
	}

	md := metadata.Pairs("authorization", vGot.Token)
	ctx = metadata.NewIncomingContext(ctx, md)

	_, err = s.UpdateUserProfile(ctx, uReqt)

	if err != nil {
		t.Errorf("flairServiceServer.UpdateUser(ok) failed with = %v", err)
		return
	}

	var uu v1internals.User
	testDb.Where("email = ?", uReq.Email).Last(&uu)

	if uu.City == u.City{
		t.Errorf("flairServiceServer.UpdateUser(ok) failed City is supposed to be %v, got %v", uReqt.Profile.City, uu.City)
	}
	if reflect.DeepEqual(uu, u){
		t.Errorf("flairServiceServer.UpdateUser(ok) failed with")
	}
	
}

func clearUsersTable() {
	testDb.Exec(setup.ClearUserTable)
}
