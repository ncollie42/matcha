package userService

import (
	pg "github.com/go-pg/pg"
	"log"
)

type UserService struct {
	DB      *pg.DB
	Session map[string]SessionInfo
}

func FindUserBy(DB *pg.DB, by, login string) (*User, error) {
	user := new(User)
	err := DB.Model(user).Where(by+"= ?", login).Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func isError (err error) bool {
	if err != nil {
		log.Println(err.Error())
	}
	return (err != nil)
}