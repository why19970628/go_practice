package main

type User struct {
	Id   int
	Sex  int
	Name string
}

func SerId(u *User) {
	u.Id = 123
}

func NewUser(f func(u *User)) *User {
	u := new(User)
	f(u)
	return u
}

func WithUserId(id int) func(u *User) {
	return func(u *User) {
		u.Id = id
		u.Name = "wang"
		u.Sex = 1
	}
}
