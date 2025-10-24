package handler

import (
	"context"

	"github.com/relaunch-cot/bff-relaunch/grpc"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"
)

type INotification interface {
	SendNotification(ctx *context.Context, in *pb.SendNotificationRequest) error
}

type notificationResource struct {
	grpcClient *grpc.Grpc
}

func (r *notificationResource) SendNotification(ctx *context.Context, in *pb.SendNotificationRequest) error {
	err := r.grpcClient.NotificationGRPC.SendNotification(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func NewNotificationHandler(grpcClient *grpc.Grpc) INotification {
	return &notificationResource{
		grpcClient: grpcClient,
	}
}
