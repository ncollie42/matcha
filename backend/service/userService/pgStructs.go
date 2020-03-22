package userService

//Altered version of the genereated struct from protofile, specially for postgress
type PendingUser struct {	//change to pendingUser, add a toString?
	Id                  int `sql:",:gen_random_uuid()"`
	UserName            string `sql:",unique"`
	FirstName           string
	LastName            string
	Password            string
	Email               string `sql:",unique"`
	Hash				string
}

type User struct {
	Id                  int32 `sql:",:gen_random_uuid()"`
	UserName            string `sql:",unique"`
	FirstName           string
	LastName            string
	Password            string
	Email               string `sql:",unique"`
	NewEmail            string `sql:",unique"`
	Gender             	string
	Preference          string
	Bio                 string
	Hash                string
	Tags                []string `sql:",array"`
	SeenHistory         []int32 `sql:",array"`
	PeopleLiked         []int32 `sql:",array"`
	PeopleBlocked       []int32 `sql:",array"`
	FameRating			int32 `sql:"default:0"`
}
