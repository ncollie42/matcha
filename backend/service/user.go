package userService

import (
	"context"
	pg "github.com/go-pg/pg"
	generated "hazuki/generated"
	"io/ioutil"
	"log"
)

type UserService struct{
	DB *pg.DB
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

func (db *UserService) ForgotPassword(_ context.Context, req *generated.ResetPassRequest) (*generated.Reply, error) {
	ErrorLocation := "On sending reset pass email"
	email := req.GetEmail()
	user, err := findUserBy(db,"email", email)
	if err != nil {
		return replyError(ErrorLocation, "Can't find email", err)
	}
	if req.GetHash() == user.Hash {
		res, err := db.DB.Model(user).Set("password = ?", req.GetNewPass()).Where("id = ?id").Update()
		if err != nil {
			return replyError(ErrorLocation, "Can't update hash", err)
		}
		log.Println(res)
	} else {
		return replyError(ErrorLocation, "Hash does not match", err)
	}
	return &generated.Reply{Message: "Password was changed"}, nil
}

func (db *UserService) UpdateUser(_ context.Context, user *generated.User) (*generated.Reply, error) {
	return &generated.Reply{Message: "I've done nothing here yet"}, nil
}

func (db *UserService) SendPassResetEmail(_ context.Context,req *generated.SendEmailRequest) (*generated.Reply, error) {
	email := req.GetEmail()
	user, err := findUserBy(db,"email", email)
	if err != nil {
		return replyError("On sending reset pass email", "Can't find email", err)
	}
	hash := randomString(40)
	res, err := db.DB.Model(user).Set("hash = ?", hash).Where("id = ?id").Update()
	if err != nil {
		return replyError("On sending reset pass email", "Can't update hash", err)
	}
	log.Println(res)
	sendMail(email, hash, "reset")
	return &generated.Reply{Message: "Reset email sent"}, nil
}

func (db *UserService) VerifyUser(_ context.Context, req *generated.VerifyRequest) (*generated.Reply, error) {
	pendingUser := new(PendingUser)
	err := db.DB.Model(pendingUser).Where("email = ?", req.GetEmail()).Select()
	if err != nil {
		return replyError("On Verification", "Can't find email", err)
	}
	//moves user from pending table to user table
	if pendingUser.Hash == req.Hash {
		err = db.DB.Insert(&User{
			UserName:			  pendingUser.UserName,
			FirstName:            pendingUser.FirstName,
			LastName:             pendingUser.LastName,
			Password:             pendingUser.Password,
			Email:                pendingUser.Email,
		})
		if err != nil {
			return replyError("On Verification", "Can't insert user", err)
		}
		_, err = db.DB.Model(pendingUser).Where("email = ?", req.GetEmail()).Delete()
		if err != nil {
			return replyError("On Verification", "Account has already been verified", err)
		}
	} else {
		return replyError("On Verification", "Hash values aren't the same", err)
	}
	log.Println("Email ", req.GetEmail()," has been verified.")
	return &generated.Reply{Message: "You're verified"}, nil
}

func (db *UserService) CreateUser(ctx context.Context, req *generated.CreateRequest) (*generated.Reply, error) {
	ErrorLocation := "On user creation"
	if !userDataisValid(req) {
		return replyError(ErrorLocation, "A field was left empty", nil)
	}
	_, err := findUserBy(db,"user_name", req.GetUserName()) //Only checks actual users, not pending
	if err == nil {
		return replyError(ErrorLocation, "Username already taken", err)
	}
	password, err := HashPassword(req.GetPassword())
	if err != nil {
		return replyError(ErrorLocation, "Had problem hasing password", err)
	}

	hash :=randomString(40)
	//Add user
	err = db.DB.Insert(&PendingUser{
		UserName: req.GetUserName() ,
		FirstName: req.GetFirstName(),
		LastName: req.GetLastName(),
		Password:  password,
		Email:     req.GetEmail(),
		Hash: hash,
	})
	if err != nil {
		return replyError(ErrorLocation,"Error adding user to table", err)
	}
	//Send email to user for verification
	sendMail(req.GetEmail(), hash, "verify")
	return &generated.Reply{Message: "user was created"}, nil
}

func (db *UserService) LoginUser(_ context.Context, login *generated.LoginRequest) (*generated.Reply, error) {
	ErrorLocation := "On Login"
	user, err := findUserBy(db,"user_name", login.GetUserName())
	if err != nil {
		return replyError(ErrorLocation, "invalid username or password", err)
	}
	if !CheckPasswordHash(login.GetPassword(), user.Password)  {
		return replyError(ErrorLocation, "invalid password", err)
	}
	return &generated.Reply{Message: "Username and password match"}, nil
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
	return nil
}
