package userService

//Altered version of the genereated struct from protofile
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
	Id                  int `sql:",:gen_random_uuid()"`
	UserName            string `sql:",unique"`
	FirstName           string
	LastName            string
	Password            string
	Email               string `sql:",unique"`
	Gender              int32
	Preference          int32
	Bio                 string
	Hash                string
}
