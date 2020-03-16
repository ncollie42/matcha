package userService

import (
	"context"
	pg "github.com/go-pg/pg"
	generated "hazuki/generated"
	"hazuki/service/Helpers"
	"io/ioutil"
	"log"
)

type UserService struct{
	DB *pg.DB
	*Sessions
}

func (db *UserService)ImageTest(_ context.Context, tmp *generated.ImageData) (*generated.ImageData, error) {
	log.Println("Testing image")
	bb, err := ioutil.ReadFile("./service/pic.png")
	if err != nil {
		log.Println("Error with image,", err)
	}
	return &generated.ImageData{
		Image:                bb,
	}, nil
}
//
//func (db *UserService) ForgotPassword(_ context.Context, req *generated.ResetPassRequest) (*generated.Reply, error) {
//	ErrorLocation := "On sending reset pass email"
//	email := req.GetEmail()
//	user, err := FindUserBy(db.DB,"email", email)
//	if err != nil {
//		return helper.ReplyError(ErrorLocation, "Can't find email", err)
//	}
//	if req.GetHash() == user.Hash {
//		res, err := db.DB.Model(user).Set("password = ?", req.GetNewPass()).Where("id = ?id").Update()
//		if err != nil {
//			return helper.ReplyError(ErrorLocation, "Can't update hash", err)
//		}
//		log.Println(res)
//	} else {
//		return helper.ReplyError(ErrorLocation, "Hash does not match", err)
//	}
//	return &generated.Reply{Message: "Password was changed"}, nil
//}

func (db *UserService) UpdateUser(_ context.Context, user *generated.User) (*generated.Reply, error) {
	return &generated.Reply{Message: "I've done nothing here yet"}, nil
}
//
//func (db *UserService) SendPassResetEmail(_ context.Context,req *generated.SendEmailRequest) (*generated.Reply, error) {
//	email := req.GetEmail()
//	user, err := FindUserBy(db.DB,"email", email)
//	if err != nil {
//		return helper.ReplyError("On sending reset pass email", "Can't find email", err)
//	}
//	hash := helper.RandomString(40)
//	res, err := db.DB.Model(user).Set("hash = ?", hash).Where("id = ?id").Update()
//	if err != nil {
//		return helper.ReplyError("On sending reset pass email", "Can't update hash", err)
//	}
//	log.Println(res)
//	helper.SendMail(email, hash, "reset")
//	return &generated.Reply{Message: "Reset email sent"}, nil
//}
//
//func (db *UserService) VerifyUser(_ context.Context, req *generated.VerifyRequest) (*generated.Reply, error) {
//	ErrorLocation := "On Verification"
//	pendingUser := new(PendingUser)
//	err := db.DB.Model(pendingUser).Where("email = ?", req.GetEmail()).Select()
//	if err != nil {
//		return helper.ReplyError(ErrorLocation, "Can't find email", err)
//	}
//	//moves user from pending table to user table
//	if pendingUser.Hash == req.Hash {
//		err = db.DB.Insert(&User{
//			UserName:			  pendingUser.UserName,
//			FirstName:            pendingUser.FirstName,
//			LastName:             pendingUser.LastName,
//			Password:             pendingUser.Password,
//			Email:                pendingUser.Email,
//		})
//		if err != nil {
//			return helper.ReplyError("On Verification", "Can't insert user", err)
//		}
//		_, err = db.DB.Model(pendingUser).Where("email = ?", req.GetEmail()).Delete()
//		if err != nil {
//			return helper.ReplyError("On Verification", "Account has already been verified", err)
//		}
//	} else {
//		return helper.ReplyError("On Verification", "Hash values aren't the same", err)
//	}
//	log.Println("Email ", req.GetEmail()," has been verified.")
//	return &generated.Reply{Message: "You're verified"}, nil
//}
//
//func (db *UserService) CreateUser(ctx context.Context, req *generated.CreateRequest) (*generated.Reply, error) {
//	ErrorLocation := "On user creation"
//	if !helper.UserDataisValid(req) {
//		return helper.ReplyError(ErrorLocation, "A field was left empty", nil)
//	}
//	_, err := FindUserBy(db.DB,"user_name", req.GetUserName()) //Only checks actual users, not pending
//	if err == nil {
//		return helper.ReplyError(ErrorLocation, "Username already taken", err)
//	}
//	password, err := helper.HashPassword(req.GetPassword())
//	if err != nil {
//		return helper.ReplyError(ErrorLocation, "Had problem hasing password", err)
//	}
//
//	hash := helper.RandomString(40)
//	//Add user
//	err = db.DB.Insert(&PendingUser{
//		UserName: req.GetUserName() ,
//		FirstName: req.GetFirstName(),
//		LastName: req.GetLastName(),
//		Password:  password,
//		Email:     req.GetEmail(),
//		Hash: hash,
//	})
//	if err != nil {
//		return helper.ReplyError(ErrorLocation,"Error adding user to table", err)
//	}
//	//Send email to user for verification
//	helper.SendMail(req.GetEmail(), hash, "verify")
//	return &generated.Reply{Message: "user was created"}, nil
//}



func (db *UserService) LoginUser(_ context.Context, login *generated.LoginRequest) (*generated.Reply, error) {
	ErrorLocation := "On Login"
	user, err := FindUserBy(db.DB,"user_name", login.GetUserName())
	if err != nil {
		return helper.ReplyError(ErrorLocation, "invalid username or password", err)
	}
	if !helper.CheckPasswordHash(login.GetPassword(), user.Password)  {
		return helper.ReplyError(ErrorLocation, "invalid password", err)
	}
	//Start session by giving a string / hash
	sessionID := db.startSession(user)

	return &generated.Reply{Message: sessionID}, nil
}

func (db *UserService) GetUsers(_ *generated.CreateRequest,stream generated.Account_GetUsersServer) error {
	var tmp []User //For now it's createRequest
	db.DB.Model(&tmp).Select()
	log.Println("Getting", len(tmp), "users:", tmp)
	for _, person := range tmp {
		ret := &generated.User{
			FirstName:            person.FirstName,
			LastName:             person.LastName,
			Password:             person.Password,
			Email:                person.Email,
		}
		stream.Send(ret)
	}
	arrayTestSearch(db) // remove from here and insert into testing rpc call
	return nil
}


func FindUserBy(DB *pg.DB,by, login string) (*User, error) {
	user := new(User)
	err := DB.Model(user).Where(by + "= ?",login).Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}
