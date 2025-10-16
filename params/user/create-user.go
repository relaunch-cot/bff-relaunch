package user

import (
	"github.com/relaunch-cot/bff-relaunch/models"
	libModels "github.com/relaunch-cot/lib-relaunch-cot/models"
)

type CreateUserPOST struct {
	Name     string              `json:"name"`
	Email    string              `json:"email"`
	Password string              `json:"password"`
	Settings models.UserSettings `json:"settings"`
	Type     string              `json:"type"`
}

func GetUserModelFromCreate(in *CreateUserPOST) *libModels.User {
	settings := libModels.UserSettings{
		Phone:       in.Settings.Phone,
		Cpf:         in.Settings.Cpf,
		DateOfBirth: in.Settings.DateOfBirth,
	}

	return &libModels.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
		Settings: settings,
		Type:     in.Type,
	}
}
