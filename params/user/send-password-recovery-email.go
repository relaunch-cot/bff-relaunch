package user

type SendPasswordRecoveryEmailPOST struct {
	Email        string `json:"email" form:"email"`
	RecoveryLink string `json:"recovery-link" form:"recovery-link"`
}
