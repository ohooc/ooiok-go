package simple_factory

type User struct {
	Id   int
	Name string
}

func NewUser() *User {
	return &User{}
}

type Admin struct {
	Id   int
	Name string
	Role string
}

func NewAdmin() *Admin {
	return &Admin{}
}

type UserGenFunc func(id int, name string) interface{}

func GenUser() UserGenFunc {
	return func(id int, name string) interface{} {
		return &User{Id: id, Name: name}
	}
}

func GenAdmin() UserGenFunc {
	return func(id int, name string) interface{} {
		return &Admin{Id: id, Name: name, Role: "admin"}
	}
}
