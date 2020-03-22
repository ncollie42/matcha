package helper

import (
	"log"
	"net/smtp"
	"os"
)

const WebsiteIP = "http://localhost:3000/"

type Email struct {
	ToEmail string
	Subject string
	Name string
	Message string
	Link string
	Footer string
}

//TODO: change toEmail to actual email
// (confirmation email: new account
func SendMail(emailInfo Email) {
	gmailSMTP := "smtp.gmail.com:587"
	fromEmail := "42.matcha.project@gmail.com"
	pass := os.Getenv("PASS")
	toEmail := "42.matcha.project@gmail.com" //Update this later to email from arg

	msg := []byte("To: "+ emailInfo.ToEmail + "\r\n" +
				"Subject: Matcha: " + emailInfo.Subject + "\n\n" +
				"\r" +
				emailInfo.Message +
				"\r\n\n" +
				emailInfo.Link +
				"\r\n\n" +
				emailInfo.Footer +
				"Best wishes," +
				"Matcha team.")

	auth := smtp.PlainAuth(
		"",
		fromEmail,
		pass,
		"smtp.gmail.com",
	)
	err := smtp.SendMail(gmailSMTP, auth, fromEmail, []string{toEmail}, msg)
	if isError(err) {
		return
	}
}

func isError (err error) bool {
	if err != nil {
		log.Println(err.Error())
	}
	return (err != nil)
}
