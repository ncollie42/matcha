package userService

import (
	"errors"
	"context"
	"google.golang.org/grpc/metadata"
	helper "hazuki/service/Helpers"
	generated "hazuki/generated"
)

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

// Move to helper?
func checkSession(ctx context.Context, U *UserService) (SessionInfo, error) {
	headers, _ := metadata.FromIncomingContext(ctx)
	var ok bool
	var sessionID []string
	var userID []string
	sesh := SessionInfo{}
	//check sessionID
	if sessionID, ok = headers["session_id"]; !ok {
		return sesh, errors.New("no 'session_id' given")
	}
	//check userID
	if userID, ok = headers["user_id"]; !ok {
		return sesh, errors.New("no 'user_id' given")
	}
	//set sesh from saved session
	if sesh, ok = U.Session[userID[0]]; !ok {
		return sesh, errors.New("No session found for user")
	}
	//check if given sessionID and stored ID are equal
	if sesh.sessionID != sessionID[0] {
		return sesh, errors.New("session id is not valid")
	}
	return sesh, nil
}