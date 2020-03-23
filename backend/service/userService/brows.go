package userService

import (
	"errors"
	generated "hazuki/generated"
	pg "github.com/go-pg/pg"
	"log"
	"github.com/go-pg/pg/orm"
)

//sort result before sending??
//Just need to add location to search, add .Order("id ASC") ext AND remove people from blocked list
//TODO: remove blocked users
// filter
func (db *UserService) Feed(request *generated.BrowseRequest, stream generated.Account_FeedServer) error {
	//session, err := checkSession(ctx, db)
	//if isError(err) {
	//	return nil, err
	//}
	session := SessionInfo{userID: 1} //Tmp for testing
	user := &User{Id: session.userID}
	err := db.DB.Select(user)
	if isError(err) {
		return errors.New("Could not get user from session")
	}
	//Set search values based on user
	minAge := user.Age - 15
	maxAge := user.Age + 15
	minFame := user.FameRating - 20
	maxFame := user.FameRating + 20
	tags := make([]string, 5)
	//location := user.Location

	//override values based on request filter
	if request.MinAge != 0 {
		minAge = request.MinAge
	}
	if request.MaxAge != 0 {
		maxAge = request.MaxAge
	}
	if request.MinFameRating != 0 {
		minFame = request.MinFameRating
	}
	if request.MaxFameRating != 0{
		maxFame = request.MaxFameRating
	}

	size := len(request.Tags)
	log.Println("size: ", len(request.Tags),"stuff: ", request.Tags)
	for size > 0 {
		tags[size -1] = request.Tags[size -1]
		size--
	}

	//base query, age fameRaiting, sex preference
	baseQuery := func ( query *orm.Query) (*orm.Query, error) {
		query.Where("age > ?", minAge).
			Where("age < ?", maxAge).
			Where("fame_rating > ?", minFame).
			Where("fame_rating < ?", maxFame).
			Where("id != ?", session.userID).
			WhereGroup(func ( query *orm.Query) (*orm.Query, error) {
				if user.Preference == "male" {
					query = query.Where("gender = ?", "male")
				} else if user.Preference == "female" {
					query = query.Where("gender = ?", "female")
				} else {
					query = query.WhereOr("gender = ?", "male").
						WhereOr("gender = ?", "female")
				}
				return query, nil
			})
		return query, nil
	}
	var users []User
	if len(request.Tags) >= 1{
		ids := idFromTags(db, session.userID, tags[0], tags[1], tags[2], tags[3], tags[4])
		db.DB.Model(&users).
			WhereGroup(baseQuery).
			Where("id in (?)", pg.In(ids)).
			Select()
	} else {
		db.DB.Model(&users).
			WhereGroup(baseQuery).
			Select()
	}

	users = removeBlockedUsers(users, user.PeopleBlocked)
	for _, person := range users {
		ret := &generated.User{
			Id:                   int32(person.Id),
			UserName:             person.UserName,
			FirstName:            person.FirstName,
			LastName:             person.LastName,
			Gender:               person.Gender,
			Bio:                  person.Bio,
			Age:                  person.Age,
			Location: 			  person.Location,
			Tags:                 nil,
		}
		stream.Send(ret)
	}
	return nil
}

func removeBlockedUsers(users []User, blocked []int32) []User{
	newUserList := []User{}
	check := map[int32]bool{}
 	for _, val := range blocked {
 		check[val] = true
	}
	for _, user := range users {
		if _, ok := check[user.Id]; !ok {
			newUserList = append(newUserList, user)
		}
	}
	return newUserList
}


func idFromTags(db *UserService, userID int32 , tag1, tag2, tag3, tag4, tag5 string) (list []int32){
	//filer the args for sql injections

	var users []User
	_, err := db.DB.Query(&users, "Select distinct id FROM (SELECT id, unnest(tags) AS tag from users) as t " +
		"WHERE tag like '"+tag1+
		"' OR tag like '"+tag2+
		"' OR tag like '"+tag3+
		"' OR tag like '"+tag4+
		"' OR tag like '"+tag5+"'")
	if isError(err) {
		return list
	}
	for _, u := range users {
		list = append(list, u.Id)
	}
	log.Println("1:", tag1, "2:", tag2, "3:", tag3, "4:", tag4, "5:", tag5)
	log.Println("From tags:",list)
	return list
}