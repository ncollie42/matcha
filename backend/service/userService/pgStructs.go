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
	Id                  int `sql:",:gen_random_uuid()"`
	UserName            string `sql:",unique"`
	FirstName           string
	LastName            string
	Password            string
	Email               string `sql:",unique"`
	Gender             	string
	Preference          string
	Bio                 string
	Hash                string
	Tags                []string
	Tags2               []string `sql:",array"`
}
