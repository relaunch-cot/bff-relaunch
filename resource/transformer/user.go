package transformer

import (
	"encoding/json"

	libModels "github.com/relaunch-cot/lib-relaunch-cot/models"
	pbBaseModels "github.com/relaunch-cot/lib-relaunch-cot/proto/base_models"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

func CreateUserToProto(in *libModels.User) (*pb.CreateUserRequest, error) {
	settings := &pbBaseModels.UserSettings{
		Phone:       in.Settings.Phone,
		Cpf:         in.Settings.Cpf,
		DateOfBirth: in.Settings.DateOfBirth,
	}

	return &pb.CreateUserRequest{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
		Settings: settings,
		Type:     in.Type,
	}, nil
}

func LoginUserToProto(in *libModels.User) (*pb.LoginUserRequest, error) {
	return &pb.LoginUserRequest{
		Email:    in.Email,
		Password: in.Password,
	}, nil
}

func UpdateUserToProto(in *libModels.User) (*pb.UpdateUserRequest, error) {
	baseModelsSettings := pbBaseModels.UserSettings{
		Phone:       in.Settings.Phone,
		Cpf:         in.Settings.Cpf,
		DateOfBirth: in.Settings.DateOfBirth,
		Biography:   in.Settings.Biography,
		Skills:      in.Settings.Skills,
	}

	newUser := &pbBaseModels.User{
		UserId:   in.UserId,
		Name:     in.Name,
		Email:    in.Email,
		Settings: &baseModelsSettings,
	}

	return &pb.UpdateUserRequest{
		UserId:  in.UserId,
		NewUser: newUser,
	}, nil
}

func UpdateUserPasswordToProto(userId, newPassword string) (*pb.UpdateUserPasswordRequest, error) {
	return &pb.UpdateUserPasswordRequest{
		UserId:      userId,
		NewPassword: newPassword,
	}, nil
}

func DeleteUserToProto(email, password string) (*pb.DeleteUserRequest, error) {
	return &pb.DeleteUserRequest{
		Email:    email,
		Password: password,
	}, nil
}

func ReportDataToProto(reportData *libModels.ReportData) (*pb.GenerateReportRequest, error) {
	jsonBytes, err := json.Marshal(reportData)
	if err != nil {
		return nil, err
	}

	return &pb.GenerateReportRequest{
		JsonData: string(jsonBytes),
	}, nil
}

func SendPasswordRecoveryEmailToProto(email, recoveryLink string) (*pb.SendPasswordRecoveryEmailRequest, error) {
	return &pb.SendPasswordRecoveryEmailRequest{
		Email:        email,
		RecoveryLink: recoveryLink,
	}, nil
}

func GetUserProfileToProto(userId string) (*pb.GetUserProfileRequest, error) {
	return &pb.GetUserProfileRequest{
		UserId: userId,
	}, nil
}

func GetUserTypeToProto(userId string) (*pb.GetUserTypeRequest, error) {
	return &pb.GetUserTypeRequest{
		UserId: userId,
	}, nil
}
