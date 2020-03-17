package userService

import (
	helper "hazuki/service/Helpers"
	"log"
	generated "hazuki/generated"
	"strconv"
	"time"
)

type SessionInfo struct {
	sessionID string
	userID int
	loginTime time.Time
}

//type session map[string]*sessionInfo


//Start session if it doesn't already exsist, else update current session
func (U *UserService) startSession(user *User) *generated.Session {
	sesh := new(SessionInfo)
	if sesh, ok := U.Session[strconv.Itoa(user.Id)]; ok {
		sesh.loginTime = time.Now()
		log.Println("updated time")
		//update time? location?
		//session is already here, update it and return sessionID
	} else {
		sessionID := helper.RandomString(40)
		sesh = &SessionInfo{
			sessionID: sessionID,
			loginTime: time.Now(),
		}
		U.Session[strconv.Itoa(user.Id)] = sesh
	}
	log.Println("Made a session and put it ")
	log.Println("All sessions:\n", U.Session)
	return &generated.Session{
		UserID:               int32(sesh.userID),
		SessionID:            sesh.sessionID,
	}
}

