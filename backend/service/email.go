package userService

import (
	"log"
	"net/smtp"
	"os"
)

//TODO: change toEmail to actual email
// Endpoint -> verify -> Password
func sendMail(email, hash, endpoint string) {
	gmailSMTP := "smtp.gmail.com:587"
	fromEmail := "42.matcha.project@gmail.com"
	pass := os.Getenv("PASS")
	toEmail := "42.matcha.project@gmail.com" //Update this later to email from arg
	//toEmail := newUser.GetEmail()
	msg := []byte("What endpoint did you end up on??")
	if endpoint == "reset" {
		msg = []byte("To: recipient@example.net\r\n" +
			"Subject: Matcha: email confirmation!\n\n" +
			"\r\n" +
			"Hi,\n\nWe got a request to reset your Matcha password.\n" +
			"http://localhost:3000/"+ endpoint +"?" + "email=" + email + "&hash=" + hash +
			"If you ignore this message, your password will not be changed. If you didn't request a password reset, let us know" +
			".\r\n")
	} else if endpoint == "verify" {
		msg = []byte("To: recipient@example.net\r\n" +
			"Subject: Matcha: email confirmation!\n\n" +
			"\r\n" +
			"Please click this link to activate your account:\n" +
			"http://localhost:3000/"+ endpoint +"?" + "email=" + email + "&hash=" + hash +
			".\r\n")
	}

	auth := smtp.PlainAuth(
		"",
		fromEmail,
		pass,
		"smtp.gmail.com",
	)
	err := smtp.SendMail(gmailSMTP, auth, fromEmail, []string{toEmail}, msg)
	if err != nil {
		log.Println("Error sending email", err)
	}
}
/*
*
Hi nc.nico.25,

We got a request to reset your Instagram password.


Reset Password



.
 */

//Once at this page front end will ask for new Pass twice,
//Send me, Email, hash, new Pass