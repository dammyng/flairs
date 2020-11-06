package helper

import (
	"io/ioutil"
	"fmt"
	"log"
	"net/http"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailMessage struct {
	Subject string
	To      string
	Text    string
	Html    string
}

func SendMail(msg EmailMessage, sender, key string) {
	log.Println("send mail called")
	from := mail.NewEmail("Flairs Support", sender)
	subject := msg.Subject
	to := mail.NewEmail("Recipient", msg.To)
	plainTextContent := msg.Text
	htmlContent := msg.Html
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(key)
	response, err := client.Send(message)
	if err != nil {
		log.Print(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)

	}
}

func HttpReq(req *http.Request) {


	// send an HTTP request using `req` object
	res, err := http.DefaultClient.Do(req)

	// check for response error
	if err != nil {
		log.Fatal("Error:", err)
	}
	// read response body
	data, _ := ioutil.ReadAll(res.Body)

	// close response body
res.Body.Close()

	// print response status and body
	log.Println("status: %d\n", res.StatusCode)
	log.Println("body: %s\n", string(data))
}
