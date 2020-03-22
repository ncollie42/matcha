package createAccountService

import (
	"context"
	"errors"
	pg "github.com/go-pg/pg"
	generated "hazuki/generated"
	helper "hazuki/service/Helpers"
	"hazuki/service/userService"
	"log"
)
type CreateAccountService struct {
	DB *pg.DB
}

func (db *CreateAccountService) Create(_ context.Context, req *generated.CreateRequest) (*generated.Reply, error) {
	ErrorLocation := "On user creation"
	if !helper.UserDataisValid(req) {
		return helper.ReplyError(ErrorLocation, "A field was left empty", nil)
	}
	user := &userService.User{}
	err := db.DB.Model(user).Where("user_name = ?", req.GetUserName()).Select()
	if err == nil {
		return helper.ReplyError(ErrorLocation, "Username already taken", nil)
	}
	password, err := helper.HashPassword(req.GetPassword())
	if err != nil {
		return helper.ReplyError(ErrorLocation, "Had problem hasing password", err)
	}
	hash := helper.RandomString(40)
	//Add user
	err = db.DB.Insert(&userService.PendingUser{
		UserName: req.GetUserName() ,
		FirstName: req.GetFirstName(),
		LastName: req.GetLastName(),
		Password:  password,
		Email:     req.GetEmail(),
		Hash: hash,
	})
	if err != nil {
		return helper.ReplyError(ErrorLocation,"Error adding user to table", err)
	}
	//Send email to user for verification
	email := req.GetEmail()
	emailBody := helper.Email{
		ToEmail: email,
		Subject: "Email confirmation",
		Name:    user.FirstName,
		Message: "Please click this link to activate your account:",
		Link:    helper.WebsiteIP + "verify" + "?email="+email+"&hash="+hash,
		Footer:  "",
	}
	helper.SendMail(emailBody)
	return &generated.Reply{Message: "user was created"}, nil
}

func (db *CreateAccountService) Verify(_ context.Context, req *generated.VerifyRequest) (*generated.Reply, error) {
	pendingUser := new(userService.PendingUser)
	err := db.DB.Model(pendingUser).Where("email = ?", req.GetEmail()).Select()
	if isError(err) {
		return nil, errors.New("Can't find email")
	}
	//moves user from pending table to user table
	if pendingUser.Hash == req.Hash {
		err = db.DB.Insert(&userService.User{
			UserName:			  pendingUser.UserName,
			FirstName:            pendingUser.FirstName,
			LastName:             pendingUser.LastName,
			Password:             pendingUser.Password,
			Email:                pendingUser.Email,
		})
		if isError(err) {
			return nil, errors.New("Can't insert user")
		}
		_, err = db.DB.Model(pendingUser).Where("email = ?", req.GetEmail()).Delete()
		if isError(err) {
			return nil, errors.New("Accout has already been verified")
		}
	} else {
		return helper.ReplyError("On Verification", "Hash values aren't the same", err)
	}
	log.Println("Email ", req.GetEmail()," has been verified.")
	return &generated.Reply{Message: "You're verified"}, nil
}

func isError (err error) bool {
	if err != nil {
		log.Println(err.Error())
	}
	return (err != nil)
}