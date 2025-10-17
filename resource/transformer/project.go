package transformer

import pb "github.com/relaunch-cot/lib-relaunch-cot/proto/project"

func CreateProjectToProto(userId, developerId, category, projectDeliveryDeadline string, amount float32) (*pb.CreateProjectRequest, error) {
	return &pb.CreateProjectRequest{
		UserId:                  userId,
		DeveloperId:             developerId,
		Category:                category,
		ProjectDeliveryDeadline: projectDeliveryDeadline,
		Amount:                  amount,
	}, nil
}
