package models

type User struct {
	UserId   string `json:"userId" form:"userId"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserSettings struct {
	Phone       string   `json:"phone" form:"phone"`
	Cpf         string   `json:"cpf" form:"cpf"`
	DateOfBirth string   `json:"dateOfBirth" form:"dateOfBirth"`
	Biography   string   `json:"biography" form:"biography"`
	Skills      []string `json:"skills" form:"skills"`
}
