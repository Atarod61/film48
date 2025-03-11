package user

type User2 struct {
	id       int
	name     string
	lastName string
	email    string
}

func SerUser2(id int, name string, email string) User2 {
	return User2{id: id, name: name, email: email}
}
