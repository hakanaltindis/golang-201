package models

type User struct {
	Id        int
	Username  string
	FirstName string
	LastName  string
	Profile   string
	Interests []Interest
}
