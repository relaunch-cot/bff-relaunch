package transformer

import (
	"encoding/json"

	models "github.com/relaunch-cot/bff-relaunch/models/user"
	pbBaseModels "github.com/relaunch-cot/lib-relaunch-cot/proto/base_models"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

func CreateUserToProto(in *models.User) (*pb.CreateUserRequest, error) {
	return &pb.CreateUserRequest{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}, nil
}

func LoginUserToProto(in *models.User) (*pb.LoginUserRequest, error) {
	return &pb.LoginUserRequest{
		Email:    in.Email,
		Password: in.Password,
	}, nil
}

func UpdateUserToProto(in *models.User) (*pb.UpdateUserRequest, error) {
	newUser := &pbBaseModels.User{
		Name:  in.Name,
		Email: in.Email,
	}

	return &pb.UpdateUserRequest{
		UserId:   in.UserId,
		Password: in.Password,
		NewUser:  newUser,
	}, nil
}

func UpdateUserPasswordToProto(userId int64, newPassword string) (*pb.UpdateUserPasswordRequest, error) {
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

func ReportDataToProto(reportData *models.ReportData) (*pb.GenerateReportRequest, error) {
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

func GetUserProfileToProto(userId int64) (*pb.GetUserProfileRequest, error) {
	return &pb.GetUserProfileRequest{
		UserId: userId,
	}, nil
}
