package userService

import (
	"context"
	"log"
	generated "hazuki/generated"
)


func (db *UserService) Test (_ context.Context, login *generated.Reply) (*generated.Reply, error) {
	return &generated.Reply{Message: "I've done nothing here yet"}, nil
}


func arrayTestSearch(db *UserService) {
	//user3 := User{
	//	Id:         0,
	//	UserName:   "",
	//	FirstName:  "",
	//	LastName:   "",
	//	Password:   "",
	//	Email:      "",
	//	Gender:     0,
	//	Preference: 0,
	//	Bio:        "",
	//	Hash:       "",
	//	Tags:       []string{"vegan", "other", "ok"},
	//	Tags2:      []string{"vegan", "this", "that"},
	//}
	//err := db.DB.Insert(&user3)
	//if err != nil {
	//	log.Println("0",err)
	//}
	user := &[]User{}
	err := db.DB.Model(user).Where("tags = ?", []string{"vegan","this","that"}).Select()
	if err != nil {
		log.Println("1",err)
	}
	log.Println("the size of result:", len(*user))
	log.Println(user)

	user2 := &[]User{}
	err = db.DB.Model(user2).Where("tags2 = ?", []string{"vegan","this","that"}).Select()
	if err != nil {
		log.Println("2",err)
	}


	log.Println("This is what we got:", len(*user2))
}

