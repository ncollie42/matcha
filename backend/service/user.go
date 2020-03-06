package userService

import (
	"context"
	"fmt"
	"log"

	pg "github.com/go-pg/pg"
	generated "hazuki/generated"
)

type UserService struct{
	DB *pg.DB
}
//Return error?
func userDataisValid(req *generated.CreateRequest) bool {
	if req.GetEmail() == "" {
		return false;
	}
	if req.GetFirsName() == "" {
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

func (db *UserService) CreateUser(ctx context.Context, req *generated.CreateRequest) (*generated.Reply, error) {
	if !userDataisValid(req) {
		return &generated.Reply{Message: "invalid data"}, nil
	}
	userName := req.GetUserName()
	email := req.GetEmail()
	firstName := req.GetFirsName()
	lastName := req.GetLastName()
	pass := req.GetPassword()
	tmp := fmt.Sprintf("Username: %s\nCreated user: %s %s\nEmail: %s\nPassword: %s", userName, firstName, lastName, email, pass)
	fmt.Println(tmp)
	//Add user
	err := db.DB.Insert(req)
	if err != nil {
		log.Println("Error while adding to creatUser Table", err)
	}
	return &generated.Reply{Message: tmp}, nil
}

func (db *UserService) GetUsers(_ *generated.CreateRequest,stream generated.Account_GetUsersServer) error {
	var tmp []generated.User
	db.DB.Model(&tmp).Select()
	for _, person := range tmp {
		stream.Send(&person)
	}
	return nil
}
////TODO: validate input (no empty field)
//TODO: Hash password
//TODO:
//