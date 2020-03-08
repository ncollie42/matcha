package userService

import (
	"context"
	"errors"
	"fmt"
	pg "github.com/go-pg/pg"
	generated "hazuki/generated"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"time"
	"unsafe"
)

type UserService struct{
	DB *pg.DB
}
//Altered version of the genereated struct from protofile
type PendingUser struct {	//change to pendingUser, add a toString?
	Id                  int `sql:",:gen_random_uuid()"`
	UserName            string `sql:",unique"`
	FirstName           string
	LastName            string
	Password            string
	Email               string `sql:",unique"`
	Hash				string
}

type User struct {
	Id                  int `sql:",:gen_random_uuid()"`
	UserName            string `sql:",unique"`
	FirstName           string
	LastName            string
	Password            string
	Email               string `sql:",unique"`
	Gender              int32
	Preference          int32
	Bio                 string
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
//TODO: make use of error return
func replyError(location, message string) (*generated.Reply, error) {
	log.Println(location,": ", message)
	return &generated.Reply{Message: message, Status: false}, nil
}
func (db *UserService) VerifyUser(_ context.Context, req *generated.VerifyRequest) (*generated.Reply, error) {
	user := new(PendingUser)
	err := db.DB.Model(user).Where("email = ?", req.GetEmail()).Select()
	if err != nil {
		return replyError("On Verification", "Can't find email")
	}
	if user.Hash == req.Hash {
		//remove from tmp, add to real user
		err = db.DB.Insert(&User{
			UserName:			  user.UserName,
			FirstName:            user.FirstName,
			LastName:             user.LastName,
			Password:             user.Password,
			Email:                user.Email,
		})
		if err != nil {
			return replyError("On Verification", "Can't insert user")
		}
		_, err = db.DB.Model(user).Where("email = ?", req.GetEmail()).Delete()
		if err != nil {
			return replyError("On Verification", "Account has already been verified")
		}
	} else {
		return replyError("On Verification", "Hash values aren't the same")
	}
	log.Println("Email ", req.GetEmail()," has been verified.")
	return &generated.Reply{Message: "You're good", Status: true}, nil
}

//TODO: change toEmail to actual email
func sendMail(email, hash string) {
	gmailSMTP := "smtp.gmail.com:587"
	fromEmail := "42.matcha.project@gmail.com"
	pass := os.Getenv("PASS")
	toEmail := "42.matcha.project@gmail.com" //Update this later to email from arg
	//toEmail := newUser.GetEmail()
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: Matcha: email confirmation!\n\n" +
		"\r\n" +
		"Please click this link to activate your account:\n" +
		"http://localhost:3000/verify?" + "email=" + email + "&hash=" + hash +
		".\r\n")
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


func (db *UserService) CreateUser(ctx context.Context, req *generated.CreateRequest) (*generated.Reply, error) {
	if !userDataisValid(req) {
		return &generated.Reply{Message: "invalid data", Status: false}, nil
	}
	userName := req.GetUserName()
	email := req.GetEmail()
	firstName := req.GetFirstName()
	lastName := req.GetLastName()
	pass := req.GetPassword()
	tmp := fmt.Sprintf("Username: %s\nCreated user: %s %s\nEmail: %s\nPassword: %s", userName, firstName, lastName, email, pass)
	fmt.Println(tmp)
	hash :=randomString(30)
	//Add user
	err := db.DB.Insert(&PendingUser{
		UserName: userName ,
		FirstName: firstName,
		LastName: lastName,
		Password:  pass,
		Email:     email,
		Hash: hash,
	})
	if err != nil {
		log.Println("Error while adding to creatUser Table", err)
	}
	//Send email to user for verification
	sendMail(email, hash)
	return &generated.Reply{Message: tmp, Status: true}, nil
}
func (db *UserService) LoginUser(_ context.Context, login *generated.LoginRequest) (*generated.Reply, error) {
	if login.GetUserName() == "1" {
		err := errors.New("THis is my special error")
		return &generated.Reply{Message: "FAILURE test", Status: false}, err
	}
	return &generated.Reply{Message: "Okay test", Status: true}, nil
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



const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var src = rand.NewSource(time.Now().UnixNano())
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)
/*
	generates a random string for pending users
*/
func randomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}