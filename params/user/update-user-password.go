package user

type UpdateUserPasswordPATCH struct {
	Email           string `json:"email" form:"email"`
	CurrentPassword string `json:"currentPassword" form:"currentPassword"`
	NewPassword     string `json:"newPassword" form:"newPassword"`
}
