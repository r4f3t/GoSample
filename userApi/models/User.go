package models

type User struct {
	FirstName string
	ID        int64
	Interest  []Interest
	LastName  string
	Profile   string
	Username  string
}
