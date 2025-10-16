package user

import (
	"github.com/relaunch-cot/bff-relaunch/models"
	libModels "github.com/relaunch-cot/lib-relaunch-cot/models"
)

type UpdateUserPUT struct {
	Name     string              `json:"name" form:"name"`
	Email    string              `json:"email" form:"email"`
	Password string              `json:"password" form:"password"`
	Settings models.UserSettings `json:"settings" form:"settings"`
	Type     string              `json:"type" form:"type"`
}

func GetUserModelFromUpdate(in *UpdateUserPUT) *libModels.User {
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
