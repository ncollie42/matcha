package userService

import (
	helper "hazuki/service/Helpers"
	"log"
)

type Sessions struct {
	Current map[string]int //For now it's just int but make it into a other struct for more info
}

func (s *Sessions) startSession(user *User) string {
	sessionID := helper.RandomString(15)
	s.Current[sessionID] = user.Id
	log.Println("Made a session and put it ")
	return sessionID
}

