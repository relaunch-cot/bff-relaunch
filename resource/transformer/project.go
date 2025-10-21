package transformer

import (
	"github.com/goccy/go-json"
	params "github.com/relaunch-cot/bff-relaunch/params/project"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/project"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateProjectToProto(userId, freelancerId, name, description, category, projectDeliveryDeadline string, amount float32) (*pb.CreateProjectRequest, error) {
	return &pb.CreateProjectRequest{
		UserId:                  userId,
		FreelancerId:            freelancerId,
		Name:                    name,
		Description:             description,
		Category:                category,
		ProjectDeliveryDeadline: projectDeliveryDeadline,
		Amount:                  amount,
	}, nil
}

func GetProjectToProto(projectId string) (*pb.GetProjectRequest, error) {
	return &pb.GetProjectRequest{
		ProjectId: projectId,
	}, nil
}
func GetAllProjectsFromUserToProto(userId, userType string) (*pb.GetAllProjectsFromUserRequest, error) {
	return &pb.GetAllProjectsFromUserRequest{
		UserId:   userId,
		UserType: userType,
	}, nil
}

func UpdateProjectToProto(projectId string, in *params.UpdateProjectPUT) (*pb.UpdateProjectRequest, error) {
	updateProjectRequest := &pb.UpdateProjectRequest{}

	b, err := json.Marshal(in)
	if err != nil {
		return nil, status.Error(codes.Internal, "error marshalling params. Details: "+err.Error())
	}

	err = json.Unmarshal(b, &updateProjectRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, "error unmarshalling params. Details: "+err.Error())
	}

	updateProjectRequest.ProjectId = projectId

	return updateProjectRequest, nil
}
