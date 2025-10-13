package user

type UpdateUserPasswordPATCH struct {
	UserId      int64  `json:"userId"`
	NewPassword string `json:"newPassword"`
}
