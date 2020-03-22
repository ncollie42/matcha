package userService

import (
	"context"
	helper "hazuki/service/Helpers"

	//"github.com/go-pg/pg"
	generated "hazuki/generated"
	"log"
)


func (db *UserService) Test (_ context.Context, login *generated.Testing) (*generated.Reply, error) {

	switch login.Message {
	case	"hey":
		log.Println("Yo what's up")
	case	"ok":
		log.Println("diff test")
	case 	"add":
		addUser(db)
	case "check":
		arrayTestSearch(db, login.Arg1, login.Arg2)
	}
	return &generated.Reply{Message: "I've done nothing here yet"}, nil
}


func addUser(db *UserService) {
	user3 := User{
		Id:         0,
		UserName:   helper.RandomString(5),
		FirstName:  "",
		LastName:   "",
		Password:   helper.RandomString(10),
		Email:      "",
		Gender:     "dude",
		Preference: "chick",
		Bio:        "",
		Hash:       "",
		Tags:      []string{"vegan", "this", "that"},
	}
	err := db.DB.Insert(&user3)
	if err != nil {
		log.Println("0",err)
	}
}

func arrayTestSearch(db *UserService, tmp1, tmp2 string) {
	//filer the args for sql injections
	var users []User
	_, err := db.DB.Query(&users, "Select distinct id FROM (SELECT id, unnest(tags) AS tag from users) as t WHERE tag like '"+tmp1+"' OR tag like '"+tmp2+"'")
	if isError(err) {
		return
	}

	log.Println(users)
}

