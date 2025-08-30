package user

type DeleteUserDELETE struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
