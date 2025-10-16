package user

type UpdateUserPasswordPATCH struct {
	UserId      string `json:"userId"`
	NewPassword string `json:"newPassword"`
}
