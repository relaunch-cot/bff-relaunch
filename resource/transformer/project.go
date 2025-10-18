package transformer

import pb "github.com/relaunch-cot/lib-relaunch-cot/proto/project"

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
