package resetPasswordService

import (
	"context"
	"errors"
	pg "github.com/go-pg/pg"
	generated "hazuki/generated"
	helper "hazuki/service/Helpers"
	"hazuki/service/userService"
	"log"
)

type ResetPasswordService struct{
	DB *pg.DB
}

func (db *ResetPasswordService) SendEmail(_ context.Context,req *generated.SendEmailRequest) (*generated.Reply, error) {
	email := req.GetEmail()
	user, err := userService.FindUserBy(db.DB,"email", email)
	if isError(err) {
		return nil, errors.New("Can't find email")
	}
	hash := helper.RandomString(40)
	_, err = db.DB.Model(user).Set("hash = ?", hash).Where("id = ?id").Update()
	if isError(err) {
		return nil, errors.New("Can't update hash")
	}
	emailBody := helper.Email{
		ToEmail: email,
		Subject: "Email confirmation",
		Name:    user.FirstName,
		Message: "We got a request to reset your Matcha password, please click this link:",
		Link:    helper.WebsiteIP + "reset" + "?email="+email+"&hash="+hash,
		Footer:  "If you ignore this message, your password will not be changed. If you didn't request a password reset, let us know",
	}
	helper.SendMail(emailBody)
	return &generated.Reply{Message: "Reset email sent"}, nil
}

func (db *ResetPasswordService) ResetPassword(_ context.Context, req *generated.ResetPassRequest) (*generated.Reply, error) {
	email := req.GetEmail()
	user, err := userService.FindUserBy(db.DB,"email", email)
	if isError(err) {
		return nil, errors.New("Can't find email")
	}
	if req.GetHash() == user.Hash {
		password, err := helper.HashPassword(req.GetNewPass())
		if isError(err) {
			return nil, errors.New("Had problem hasing password")
		}
		hash := helper.RandomString(40)
		res, err := db.DB.Model(user).Set("hash = ?", hash).Where("id = ?id").Update()
		if isError(err) {
			return nil, errors.New("Can't update hash")
		}
		res, err = db.DB.Model(user).Set("password = ?", password).Where("id = ?id").Update()
		if isError(err) {
			return nil, errors.New("Can't update password")
		}
		log.Println(res)
	} else {
		return nil, errors.New("Hash does not match")
	}
	return &generated.Reply{Message: "Password was changed"}, nil
}

func isError (err error) bool {
	if err != nil {
		log.Println(err.Error())
	}
	return (err != nil)
}
