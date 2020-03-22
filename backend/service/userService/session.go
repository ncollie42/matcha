package userService

import (
	"fmt"
	helper "hazuki/service/Helpers"
	generated "hazuki/generated"
	"strconv"
	"time"
)

type SessionInfo struct {
	sessionID string
	userID int32
	loginTime time.Time
}

func (s SessionInfo) String() string {
	return fmt.Sprintln("session: ",s.sessionID,
	"\nID:", s.userID,
	"\nLoginTIme", s.loginTime)
}
//type session map[string]*sessionInfo


//Start session if it doesn't already exsist, else update current session
func (U *UserService) startSession(user *User) *generated.Session {
	ok := false
	sesh := SessionInfo{}
	if sesh, ok = U.Session[strconv.Itoa(int(user.Id))]; ok {
		sesh.loginTime = time.Now()
		//update time? location?
		//session is already here, update it and return sessionID
	} else {
		sessionID := helper.RandomString(40)
		sesh = SessionInfo{
			sessionID: sessionID,
			userID: user.Id,
			loginTime: time.Now(),
		}
		U.Session[strconv.Itoa(int(user.Id))] = sesh
	}
	tmp := &generated.Session{
		UserID:               int32(sesh.userID),
		SessionID:            sesh.sessionID,
	}
	return tmp
}

