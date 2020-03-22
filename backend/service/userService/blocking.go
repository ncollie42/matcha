package userService

import (
	"context"
	"errors"
	generated "hazuki/generated"
)

func (db *UserService) BlockUser(ctx context.Context, request *generated.UserID) (*generated.Reply, error) {
	//session, err := checkSession(ctx, db)
	//if isError(err) {
	//	return nil, err
	//}
	session := SessionInfo{userID: 1} //Tmp for testing
	if request.UserID == session.userID {
		return nil, errors.New("Can't block self")
	}
	self := &User{Id: session.userID}
	err := db.DB.Select(self)
	if isError(err) {
		return nil, errors.New("Could not get user (self)")
	}
	requestedUser := &User{Id: request.UserID}
	err = db.DB.Select(requestedUser)
	if isError(err) {
		return nil, errors.New("Could not get user (requested)")
	}
	if ok := find(self.PeopleBlocked, request.UserID); !ok {
		self.PeopleBlocked = append(self.PeopleBlocked, request.UserID)
		err = db.DB.Update(self)
		if isError(err) {
			return nil, errors.New("Could not update blocked list for user")
		}
	}
	return &generated.Reply{Message: "Haven't done anything here yet"}, nil
}

