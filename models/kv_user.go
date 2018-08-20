package models

type User struct {
	ID       int    `json:"id" from:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email form:"email"`
	Password string `json:"password" form:"password"`
	Tel      string `jsom:"tel" form:"tel"`
}
