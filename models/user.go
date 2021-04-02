package models

type User struct {
	ID int
	Age int
	FirstName string
	LastName string
	Email string
}
type UserList struct {
	Users []User
}
