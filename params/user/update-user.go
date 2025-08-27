package user

type UpdateUserPasswordPATCH struct {
	Email       string `json:"email" form:"email"`
	CurrentUser string `json:"currentUser" form:"currentUser"`
	NewUser     string `json:"newUser" form:"newUser"`
}
