package notification

import (
	"context"

	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"
)

type INotificationGRPC interface {
	SendNotification(ctx *context.Context, in *pb.SendNotificationRequest) error
}

type resource struct {
	grpcClient pb.NotificationServiceClient
}

func (r *resource) SendNotification(ctx *context.Context, in *pb.SendNotificationRequest) error {
	_, err := r.grpcClient.SendNotification(*ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func NewNotificationGrpcClient(grpcClient pb.NotificationServiceClient) INotificationGRPC {
	return &resource{
		grpcClient: grpcClient,
	}
}
