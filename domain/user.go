package domain

type User struct {
	Base
	Name     string
	Email    string
	Password string
	Token    string
}