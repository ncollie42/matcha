package userService

import (
	"context"
	"errors"
	generated "hazuki/generated"
	helper "hazuki/service/Helpers"
	"log"
)
const seenHistoryLimit = 10
//get user by id, id 0 will return calling user
//getProfile
func (db *UserService) GetUser(ctx context.Context, request *generated.UserID) (*generated.User, error) {
	//session, err := checkSession(ctx, db)
	//if isError(err) {
	//	return nil, err
	//}
	session := SessionInfo{userID: 1} //Tmp for testing
	user := &User{Id: request.UserID}
	err := db.DB.Select(user)
	if isError(err) {
		return nil, errors.New("Could not get user")
	}
	if session.userID != request.UserID  {
		//If contains Don't add?
		user.SeenHistory = append(user.SeenHistory, session.userID)
		if len(user.SeenHistory) > seenHistoryLimit {
			user.SeenHistory = user.SeenHistory[1:]
		}
		user.FameRating += 1
		db.DB.Update(user)
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
		Tags:		user.Tags,
		SeenHistory:user.SeenHistory,
		Age:        user.Age,
		Location: 	user.Location,
	}, nil
}

//Non empty values will be updated on the given userID
func (db *UserService) UpdateUser(ctx context.Context, request *generated.User) (*generated.Reply, error) {
	session, err := checkSession(ctx, db)
	if isError(err) {
		return nil, err
	}
	//make sure only session and request match inorder to update request
	newData := &User{Id: session.userID}

	err = db.DB.Select(newData)
	if isError(err) {
		return nil , errors.New("Can't update request: error selecting request from DB")
	}
	if request.Gender != "" {
		newData.Gender = request.Gender
	}
	if request.Preference != "" {
		newData.Preference = request.Preference
	}
	if request.Bio != "" {
		newData.Bio = request.Bio
	}
	if len(request.Tags) != 0 {
		newData.Tags = request.Tags
	}
	if request.FirstName != "" {
		newData.FirstName = request.FirstName
	}
	if request.LastName != "" {
		newData.LastName = request.LastName
	}
	if request.Location != "" {
		newData.Location = request.Location
	}
	if request.Email != "" {
		newData.NewEmail = request.Email
		hash := helper.RandomString(40)
		newData.Hash = hash

		emailBody := helper.Email{
			ToEmail: request.Email,
			Subject: "Email Reset confirmation",
			Name:    newData.FirstName,
			Message: "We got a request to reset your Matcha password, please click this link:",
			Link:    helper.WebsiteIP + "updateEmail" + "?email="+ newData.Email+"&hash="+hash,
			Footer:  "If you ignore this message, your password will not be changed. If you didn't request a password reset, let us know",
		}
		helper.SendMail(emailBody)
	}

	db.DB.Update(newData)
	return &generated.Reply{Message: "I've done nothing here yet"}, nil
}

func (db *UserService) VerifyNewEmail(ctx context.Context, request *generated.VerifyRequest) (*generated.Reply, error) {
	user := &User{Email: request.Email}
	log.Println(request.Email)
	err := db.DB.Model(user).Where("email =?email").Select()
	if isError(err) {
		//log.Println("bad 1")
		return nil, errors.New("Can't find user by this email")
	}
	res, err := db.DB.Model(user).Set("email = ?", user.NewEmail).Where("id = ?id").Update()
	log.Println(res)
	return &generated.Reply{Message: "New email has been changed"}, nil
}