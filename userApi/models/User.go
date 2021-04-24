package models

type User struct {
	FirstName string
	ID        int
	Interest  []Interest
	LastName  string
	Profile   string
	Username  string
}
