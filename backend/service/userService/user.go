package userService

import (
	"context"
	pg "github.com/go-pg/pg"
	"google.golang.org/grpc/metadata"
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

func (db *UserService) UpdateUser(_ context.Context, user *generated.User) (*generated.Reply, error) {
	if user.Password != "" {
		//update
	}
	return &generated.Reply{Message: "I've done nothing here yet"}, nil
}


func (db *UserService) LoginUser(ctx context.Context, login *generated.LoginRequest) (*generated.Reply, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	log.Println(ok, headers)
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

func (db *UserService) GetUser(ctx context.Context, user *generated.User) (*generated.User, error) {
	headers, _ := metadata.FromIncomingContext(ctx)
	if session_id, ok := headers["session_ID"]; ok {
		if user_id, ok := db.Current[session_id[0]]; ok {
			log.Println("Got user ID: ", user_id)
		}
	}
	return &generated.User{}, nil
}

func (db *UserService) GetSelf(ctx context.Context, _ *generated.Empty) (*generated.User, error) {
	headers, _ := metadata.FromIncomingContext(ctx)
	if session_id, ok := headers["session_ID"]; ok {
		if user_id, ok := db.Current[session_id[0]]; ok {
			log.Println("Got user ID: ", user_id)
		}
	}
	return &generated.User{}, nil
}


func (db *UserService) GetUsers(_ *generated.CreateRequest,stream generated.Account_GetUsersServer) error {
	headers, ok := metadata.FromIncomingContext(stream.Context())
	log.Println(ok, headers)
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


func FindUserBy(DB *pg.DB,by, login string) (*User, error) {
	user := new(User)
	err := DB.Model(user).Where(by + "= ?",login).Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}
