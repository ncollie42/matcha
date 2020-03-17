package userService

import (
	"context"
	"errors"
	generated "hazuki/generated"
	helper "hazuki/service/Helpers"
	"io/ioutil"
	"log"

	pg "github.com/go-pg/pg"
	"google.golang.org/grpc/metadata"
)

type UserService struct {
	DB      *pg.DB
	Session map[string]*SessionInfo
}

func (db *UserService) ImageTest(_ context.Context, tmp *generated.ImageData) (*generated.ImageData, error) {
	log.Println("Testing image")
	bb, err := ioutil.ReadFile("./service/pic.png")
	if err != nil {
		log.Println("Error with image,", err)
	}
	return &generated.ImageData{
		Image: bb,
	}, nil
}

func (db *UserService) UpdateUser(_ context.Context, user *generated.User) (*generated.Reply, error) {
	if user.Password != "" {
		//update
	}
	return &generated.Reply{Message: "I've done nothing here yet"}, nil
}

//Logs user in and returns a session ID
func (db *UserService) LoginUser(ctx context.Context, login *generated.LoginRequest) (*generated.Session, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	log.Println(ok, headers)
	user, err := FindUserBy(db.DB, "user_name", login.GetUserName())
	if err != nil {
		return nil, errors.New("invalid username or password")
	}
	if !helper.CheckPasswordHash(login.GetPassword(), user.Password) {
		return nil, errors.New("invalid username or password")
	}

	session := db.startSession(user)
	return session, nil
}

func (db *UserService) GetUser(ctx context.Context, user *generated.User) (*generated.User, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	log.Println(ok, headers)
	if session_id, ok := headers["session_ID"]; ok {
		if user_id, ok := db.Session["0"]; ok {
			log.Println("Got user ID: ", user_id, session_id)
		}
	}
	return &generated.User{}, nil
}

func (db *UserService) GetSelf(ctx context.Context, _ *generated.Empty) (*generated.User, error) {
	session, err := checkSession(ctx, db)
	if err != nil {
		return nil, err
	}
	user := &User{Id: session.userID}
	err = db.DB.Select(user)
	if err != nil {
		return nil, errors.New("Could not get user")
	}
	return &generated.User{
		UserName:   user.UserName,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Gender:     user.Gender,
		Preference: user.Preference,
		Bio:        user.Bio,
	}, nil
}

//make cleaner, return early - not nested
func checkSession(ctx context.Context, U *UserService) (*SessionInfo, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	log.Println(ok, headers)
	sesh := new(SessionInfo)
	if sessionID, ok := headers["session_id"]; ok {
		if userID, ok := headers["user_id"]; ok {
			if sesh, ok = U.Session[userID[0]]; ok {
				if sesh.sessionID == sessionID[0] {
					log.Println("Got user ID: ", sessionID)
				} else {
					return nil, errors.New("session id is not valid")
				}
			} else {
				return nil, errors.New("No session found for user")
			}
		} else {
			return nil, errors.New("no 'user_id' given")
		}
	} else {
		return nil, errors.New("no 'session_id' given")
	}
	return sesh, nil
}

func (db *UserService) GetUsers(_ *generated.CreateRequest, stream generated.Account_GetUsersServer) error {
	headers, ok := metadata.FromIncomingContext(stream.Context())
	log.Println(ok, headers)
	var tmp []User //For now it's createRequest
	db.DB.Model(&tmp).Select()
	log.Println("Getting", len(tmp), "users:", tmp)
	for _, person := range tmp {
		ret := &generated.User{
			FirstName: person.FirstName,
			LastName:  person.LastName,
			Password:  person.Password,
			Email:     person.Email,
		}
		stream.Send(ret)
	}
	return nil
}

func FindUserBy(DB *pg.DB, by, login string) (*User, error) {
	user := new(User)
	err := DB.Model(user).Where(by+"= ?", login).Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}
