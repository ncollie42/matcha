package userService

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"os"

	pg "github.com/go-pg/pg"
	generated "hazuki/generated"
)

type UserService struct{
	DB *pg.DB
}
//Altered version of the genereated struct from protofile
type CreateRequest struct {
	Id                   int64    `sql:"id,pk"`
	UserName             string   `sqk:"name,unique"`
	FirsName             string   ``
	LastName             string   ``
	Password             string   ``
	Email                string   ``
}

//Return error?
func userDataisValid(req *generated.CreateRequest) bool {
	//makesure it's a valid email
	if req.GetEmail() == "" {
		return false;
	}
	if req.GetFirstName() == "" {
		return false;
	}
	if req.GetLastName() == "" {
		return false;
	}
	if req.GetPassword() == "" {
		return false;
	}
	return true;
}

func (db *UserService) VerifyUser(_ context.Context, req *generated.VerifyRequest) (*generated.Reply, error) {
	log.Println("Verifying:")
	log.Println(req.Email)
	log.Println(req.Hash)
	return &generated.Reply{Message: "You're good"}, nil
}

func (db *UserService) DBTest() {
	//db.DB
}
//Add Has field
//add the ID field after account creation for postgress? remove from proto
//make email / password env var
func sendMail(newUser *generated.CreateRequest) {
	gmailSMTP := "smtp.gmail.com:587"
	fromEmail := "nicocollie@gmail.com"
	pass := os.Getenv("PASS")
	toEmail := "hazuki.miyake@gmail.com"
	//toEmail := newUser.GetEmail()
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: Matcha: email confirmation!\n\n" +
		"\r\n" +
		"Please click this link to activate your account:" +
		"http://localhost:3000/verify?" + "email=UserEmail" + "&hash=TheHash" +
		".\r\n")
	auth := smtp.PlainAuth(
		"",
		fromEmail,
		pass,
		"smtp.gmail.com",
	)
	err := smtp.SendMail(gmailSMTP, auth, "ncollie42@gmail.com", []string{toEmail}, msg)
	if err != nil {
		log.Println("Error sending email", err)
	}
}


func (db *UserService) CreateUser(ctx context.Context, req *generated.CreateRequest) (*generated.Reply, error) {
	if !userDataisValid(req) {
		return &generated.Reply{Message: "invalid data"}, nil
	}
	userName := req.GetUserName()
	email := req.GetEmail()
	firstName := req.GetFirstName()
	lastName := req.GetLastName()
	pass := req.GetPassword()
	tmp := fmt.Sprintf("Username: %s\nCreated user: %s %s\nEmail: %s\nPassword: %s", userName, firstName, lastName, email, pass)
	fmt.Println(tmp)
	//Add user
	err := db.DB.Insert(req)
	if err != nil {
		log.Println("Error while adding to creatUser Table", err)
	}
	//Send email to user for verification
	sendMail(req)
	return &generated.Reply{Message: tmp}, nil
}

func (db *UserService) GetUsers(_ *generated.CreateRequest,stream generated.Account_GetUsersServer) error {
	var tmp []generated.CreateRequest //For now it's createRequest
	db.DB.Model(&tmp).Select()
	log.Println("Getting", len(tmp), "users:", tmp)
	for _, person := range tmp {
		ret := &generated.User{
			FirstName:            person.FirstName,
			LastName:             person.LastName,
			Password:             person.Password,
			Email:                person.Email,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
		stream.Send(ret)
	}
	return nil
}
////TODO: validate input (no empty field)
//TODO: Hash password
//TODO:
//