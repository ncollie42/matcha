package resetPasswordService

import (
	"context"
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
	ErrorLocation := "On sending reset pass email"
	email := req.GetEmail()
	user, err := userService.FindUserBy(db.DB,"email", email)
	if err != nil {
		return helper.ReplyError(ErrorLocation, "Can't find email", err)
	}
	hash := helper.RandomString(40)
	res, err := db.DB.Model(user).Set("hash = ?", hash).Where("id = ?id").Update()
	if err != nil {
		return helper.ReplyError(ErrorLocation, "Can't update hash", err)
	}
	log.Println(res)
	helper.SendMail(email, hash, "reset")
	return &generated.Reply{Message: "Reset email sent"}, nil
}

func (db *ResetPasswordService) ResetPassword(_ context.Context, req *generated.ResetPassRequest) (*generated.Reply, error) {
	ErrorLocation := "On sending reset pass email"
	email := req.GetEmail()
	user, err := userService.FindUserBy(db.DB,"email", email)
	if err != nil {
		return helper.ReplyError(ErrorLocation, "Can't find email", err)
	}
	if req.GetHash() == user.Hash {
		password, err := helper.HashPassword(req.GetNewPass())
		if err != nil {
			return helper.ReplyError(ErrorLocation, "Had problem hasing password", err)
		}
		hash := helper.RandomString(40)
		res, err := db.DB.Model(user).Set("hash = ?", hash).Where("id = ?id").Update()
		if err != nil {
			return helper.ReplyError(ErrorLocation, "Can't update hash", err)
		}
		res, err = db.DB.Model(user).Set("password = ?", password).Where("id = ?id").Update()
		if err != nil {
			return helper.ReplyError(ErrorLocation, "Can't update hash", err)
		}
		log.Println(res)
	} else {
		return helper.ReplyError(ErrorLocation, "Hash does not match", err)
	}
	return &generated.Reply{Message: "Password was changed"}, nil
}