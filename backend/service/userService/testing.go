package userService

import (
	"context"
	helper "hazuki/service/Helpers"

	//"github.com/go-pg/pg"
	generated "hazuki/generated"
	"github.com/go-pg/pg/orm"
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
		arrayTestSearch(db, login.Arg1, login.Arg2, login.Arg3, login.Arg4, login.Arg5)
	case "sex":
		test2(db, login.Arg1, login.Arg2, login.Arg3)
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
		Gender:     "female",
		Preference: "both",
		Bio:        "",
		Hash:       "",
		Tags:      []string{"vegan", "this", "that"},
	}
	err := db.DB.Insert(&user3)
	if err != nil {
		log.Println("0",err)
	}
}

func q( query *orm.Query) (*orm.Query, error) {
	query = query.WhereOr("gender = ?", "male").
		WhereOr("gender = ?", "female")
	return query, nil
}

func test2(db *UserService, pref, arg1, arg2 string) {
	list := []int32{}
	users := &[]User{}
	//str := "male"
	err := db.DB.Model(users).
		WhereGroup(func ( query *orm.Query) (*orm.Query, error) {
		if pref == "male" {
			query = query.Where("gender = ?", "male")
		} else if pref == "female" {
			query = query.Where("gender = ?", "female")
		} else {
			query = query.WhereOr("gender = ?", "male").
			WhereOr("gender = ?", "female")
		}
		return query, nil
		}).WhereGroup(func ( query *orm.Query) (*orm.Query, error) {
			query.Where("age > ?", 20).Where("age < ?", 50)
		return query, nil
	}).WhereGroup(func ( query *orm.Query) (*orm.Query, error) {
		var users []User
		query.Query(&users, "Select distinct id FROM (SELECT id, unnest(tags) AS tag from users) as t " +
			"WHERE tag like '"+arg1+
			"' OR tag like '"+arg2+"'")
		return query, nil
	}).Select()
	//err := db.DB.Model(legal).Where("gender = ?", str).WhereOr("gender = ?", "female").Select()
	if isError(err){

	}
	for _, u := range *users {
		list = append(list, u.Id)
	}
	log.Println(list)
}

func arrayTestSearch(db *UserService, tmp1, tmp2, tmp3, tmp4, tmp5 string) {
	//filer the args for sql injections
	var users []User
	_, err := db.DB.Query(&users, "Select distinct id FROM (SELECT id, unnest(tags) AS tag from users) as t " +
		"WHERE tag like '"+tmp1+
		"' OR tag like '"+tmp2+
		"' OR tag like '"+tmp3+
		"' OR tag like '"+tmp4+
		"' OR tag like '"+tmp5+"'")
	if isError(err) {
		return
	}

	log.Println(users)
}

