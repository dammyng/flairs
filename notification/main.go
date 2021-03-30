package main

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	_ "github.com/GeertJohan/go.rice"
	"github.com/joho/godotenv"

	"notification/helper"
	"shared/events"
	event_amqp "shared/events/amqp"

	"github.com/streadway/amqp"
)

func loadEnv() {
	log.Println("env loading...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func main() {
	if os.Getenv("APP_ENV") != "production" {
		loadEnv()
	}

	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		log.Fatal("could not establish amqp: ", err.Error())
	}
	defer conn.Close()
	// amqp component

	_, err = event_amqp.NewAMQPEventEmitter(conn, "auth")
	if err != nil {
		log.Fatal("could not establish amqp connection :" + err.Error())
	}

	eventListener, err := event_amqp.NewAMQPEventListener(conn, "auth")
	if err != nil {
		log.Fatalf("setup ampq", err.Error())
	}
	go ProcessEvents(eventListener)
	c := make(chan int)
	<-c
}

func ProcessEvents(eventListener events.EventListener) error {
	received, errors, err := eventListener.Listen("auth", "user.created", "user.defaultwallet", "user.transact", "user.reset_password","user.welcome")
	if err != nil {
		log.Fatalf("event listenner error %v", err.Error())
	}

	//templateBox, err := rice.FindBox("html")
	//if err != nil {
	//	log.Fatal(err)
	//}

	for {
		select {
		case evt := <-received:
			fmt.Printf("got event %s ", evt.EventName())

			switch e := evt.(type) {
			case *events.PasswordReset:
				subject := "Flairs Password Reset"

				textContent := fmt.Sprintf("Your password reset token is %s. Kindly use to reset your account password", e.Token)

				t := template.Must(template.New("password_reset").Parse(`Your password reset token is {{.Token}}. Kindly use to reset your account password`))
				out := new(bytes.Buffer)
				data := struct {
					Token string
				}{
					e.Token,
				}
				err = t.Execute(out, data)
				if err != nil {
					log.Fatal(err)
				}

				htmlBytes := out.Bytes()
				htmlContent := string(htmlBytes)

				msg := helper.EmailMessage{
					subject,
					e.Email,
					textContent,
					htmlContent,
				}
				helper.SendMail(msg, os.Getenv("AlphaAdmin"), os.Getenv("SendGridKey"))
			
			case *events.WelcomeUserEvent:
				subject := "Welcome to the world of Flairs!"
				//tString, err := templateBox.String("welcome.html")
				//t := template.Must(template.New("welcome_user").ParseFiles("html/welcome.html"))
				//t := template.Must(template.New("welcome_user").Parse(tString))
				textContent := fmt.Sprint(`%s, Now that you have successfully created your Flairs account you are welcome to a world of limitless opportunities:
					- Do all your banking operations all in one app
					- Use your Flairs VISA card on all channels across the globe!
					- Same money on  whooping Flairs deals,
					- Pay for bills and other utilities,
					- Get expect finance advice from Lola your Personal Finance Advisor
					- Multiply your funds with Flairs Wealth Manager & lots more!

				Welcome aboard!
				With,
				Lola
				Your PFA`, e.Username)
				out := new(bytes.Buffer)
				data := struct {
					Username string
				}{
					e.Username,
				}
				t := template.Must(template.New("welcome_user").Parse(textContent))

				err = t.Execute(out, data)
				if err != nil {
					log.Fatal(err)
				}

				htmlBytes := out.Bytes()
				htmlContent := string(htmlBytes)
				msg := helper.EmailMessage{
					subject,
					e.Email,
					html.EscapeString(htmlContent),
					htmlContent,
				}
				helper.SendMail(msg, os.Getenv("AlphaAdmin"), os.Getenv("SendGridKey"))
			
			case *events.UserCreatedEvent:
				subject := "Your Flairs Email Verification Code"
				textContent := fmt.Sprintf("You're on your way! Your email verification code is %s", e.Token)
				t := template.Must(template.New("email_confirm").Parse(`
					You're on your way! Your email verification code is {{.Token}}`))
				out := new(bytes.Buffer)
				data := struct {
					Token string
				}{
					e.Token,
				}
				err = t.Execute(out, data)
				if err != nil {
					log.Fatal(err)
				}

				htmlBytes := out.Bytes()
				htmlContent := string(htmlBytes)
				msg := helper.EmailMessage{
					subject,
					e.Email,
					textContent,
					htmlContent,
				}
				helper.SendMail(msg, os.Getenv("AlphaAdmin"), os.Getenv("SendGridKey"))
				fmt.Println("got here")
			case *events.CreateDefWallet:
				reqURL, _ := url.Parse(e.URL)

				// create request body
				bodyContent := fmt.Sprintf(
					`{
					"name":"default",
					"memo":"default flairs wallet",
					"userId":"`+e.UserID+`",
					"walletType": 211,
					"currency": "NG",
					"accountBal":0.00,
					"ledgerBal":0.00,
					"status":"active"
					}`, "userid")

				reqBody := ioutil.NopCloser(strings.NewReader(bodyContent))

				req := &http.Request{
					Method: "POST",
					URL:    reqURL,
					Header: map[string][]string{
						"Content-Type":  {"application/json; charset=UTF-8"},
						"Authorization": {e.Token},
					},
					Body: reqBody,
				}
				helper.HttpReq(req)
			case *events.PerformTransaction:
				reqURL, _ := url.Parse(os.Getenv("TransactionURLv1") + e.WalletID)

				// create request body
				bodyContent := fmt.Sprintf("{\"amount\":%v}", e.Amount)

				reqBody := ioutil.NopCloser(strings.NewReader(bodyContent))

				req := &http.Request{
					Method: "POST",
					URL:    reqURL,
					Header: map[string][]string{
						"Content-Type":  {"application/json; charset=UTF-8"},
						"Authorization": {e.Token},
					},
					Body: reqBody,
				}
				helper.HttpReq(req)
			default:
				log.Printf("unknown event: %t", e)
			}
		case err = <-errors:
			log.Printf(" recieved error while processing msg: %s", err)
		}
	}
}
