package models

type User struct {
	Name     string `json:"username"`
	Email    string
	Password string `json:"-"`
}
