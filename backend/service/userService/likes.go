package userService

import (
	"context"
	"errors"
	generated "hazuki/generated"
)

func (db *UserService) GetLikedStatus(ctx context.Context, request *generated.UserID) (*generated.LikedStatus, error) {

	//session, err := checkSession(ctx, db)
	//if isError(err) {
	//	return nil, err
	//}
	session := SessionInfo{userID: 1} //Tmp for testing
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
	status := likeStatus(self, requestedUser)
	return status, nil
}

func (db *UserService) LikeUser(ctx context.Context, request *generated.UserID) (*generated.LikedStatus, error) {
	//session, err := checkSession(ctx, db)
	//if isError(err) {
	//	return nil, err
	//}
	session := SessionInfo{userID: 1} //Tmp for testing
	if request.UserID == session.userID {
		return nil, errors.New("Can't like self")
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
	if ok := find(self.PeopleLiked, request.UserID); !ok {
		self.PeopleLiked = append(self.PeopleLiked, request.UserID)
		err = db.DB.Update(self)
		if isError(err) {
			return nil, errors.New("Could not update likes for user")
		}
	}
	status := likeStatus(self, requestedUser)
	return status, nil
}


func likeStatus(self, requestedUser *User) *generated.LikedStatus {
	Ilike := false
	theyLike := false
	bothLike := false
	if ok := find(self.PeopleLiked, requestedUser.Id); ok {
		Ilike = true
	}
	if ok := find(requestedUser.PeopleLiked, self.Id); ok {
		theyLike = true
	}
	if theyLike && Ilike {
		bothLike = true
	}
	return &generated.LikedStatus{
		BothLiked:            bothLike,
		ILiked:               Ilike,
	}
}

func find(list []int32, target int32) bool {
	for _, val := range list {
		if val == target {
			return true
		}
	}
	return false
}