package user

import (
	"github.com/relaunch-cot/bff-relaunch/models"
	libModels "github.com/relaunch-cot/lib-relaunch-cot/models"
)

type UpdateUserPUT struct {
	Name         string              `json:"name" form:"name"`
	Email        string              `json:"email" form:"email"`
	Settings     models.UserSettings `json:"settings" form:"settings"`
	Type         string              `json:"type" form:"type"`
	UrlImageUser string              `json:"urlImageUser" form:"urlImageUser"`
}

func GetUserModelFromUpdate(in *UpdateUserPUT, userId string) *libModels.User {
	settings := libModels.UserSettings{
		Phone:       in.Settings.Phone,
		Cpf:         in.Settings.Cpf,
		DateOfBirth: in.Settings.DateOfBirth,
		Biography:   in.Settings.Biography,
		Skills:      in.Settings.Skills,
	}

	return &libModels.User{
		UserId:       userId,
		Name:         in.Name,
		Email:        in.Email,
		Settings:     settings,
		Type:         in.Type,
		UrlImageUser: in.UrlImageUser,
	}
}
