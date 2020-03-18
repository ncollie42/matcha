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
	Session map[string]SessionInfo
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
	//make sure only session and user match inorder to update user
	if user.Password != "" {
		//update
	}
	return &generated.Reply{Message: "I've done nothing here yet"}, nil
}

//Logs user in and returns a session ID
func (db *UserService) LoginUser(ctx context.Context, login *generated.LoginRequest) (*generated.Session, error) {
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

func (db *UserService) GetImages(request *generated.ImageRequest, stream generated.Account_GetImagesServer) error {
	_, err := checkSession(stream.Context(), db)
	if err != nil {
		return err
	}
	count := request.GetCount()
	bb, err := ioutil.ReadFile("./service/pic.png")
	if err != nil {
		log.Println("Error with image,", err)
	}
	for count > 0 {
		//getImages(stream, userID, count) -> goes to user's file, opens file -> send -> count-- -> loop
		stream.Send(&generated.ImageData{Image: bb})
		count--
	}
	return nil
}

//get user by id, id 0 will return calling user
func (db *UserService) GetUser(ctx context.Context, request *generated.User) (*generated.User, error) {
	session, err := checkSession(ctx, db)
	if err != nil {
		return nil, err
	}
	user := &User{Id: int(request.Id)}
	err = db.DB.Select(user)
	if err != nil {
		return nil, errors.New("Could not get user")
	}
	if session.userID == int(request.Id) {
		log.Println("Getting self")
		//Don't add seen stuff and what not
	}
	return &generated.User{
		Id:			int32(user.Id),
		UserName:   user.UserName,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Gender:     user.Gender,
		Preference: user.Preference,
		Bio:        user.Bio,
	}, nil
}


//get user by id, id 0 will return calling user
func (db *UserService) GetUser2(ctx context.Context, request *generated.User2) (*generated.User2, error) {
	session, err := checkSession(ctx, db)
	if err != nil {
		return nil, err
	}
	images := new([][]byte)
	user := &User{Id: int(request.Id)}
	err = db.DB.Select(user)
	if err != nil {
		return nil, errors.New("Could not get user")
	}
	if session.userID == int(request.Id) {
		log.Println("Getting self")
		//Don't add seen stuff and what not
	}
	//bb, err := ioutil.ReadFile("./service/pic.png")
	//if err != nil {
	//	log.Println("Error with image,", err)
	//}
	//images[] = append(images[], bb)//
	return &generated.User2{
		Id:			int32(user.Id),
		UserName:   user.UserName,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Gender:     user.Gender,
		Preference: user.Preference,
		Bio:        user.Bio,
		Image:		*images,
	}, nil
}


//Removing this function for getuser later
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
func checkSession(ctx context.Context, U *UserService) (SessionInfo, error) {
	headers, _ := metadata.FromIncomingContext(ctx)
	var ok bool
	var sessionID []string
	var userID []string
	sesh := SessionInfo{}
	if sessionID, ok = headers["session_id"]; !ok {
		return sesh, errors.New("no 'session_id' given")
	}
	if userID, ok = headers["user_id"]; !ok {
		return sesh, errors.New("no 'user_id' given")
	}
	if sesh, ok = U.Session[userID[0]]; !ok {
		return sesh, errors.New("No session found for user")
	}
	if sesh.sessionID != sessionID[0] {
		return sesh, errors.New("session id is not valid")
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
