package handler

import (
	"context"

	"github.com/relaunch-cot/bff-relaunch/grpc"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"
)

type INotification interface {
	SendNotification(ctx *context.Context, in *pb.SendNotificationRequest) error
	GetNotification(ctx *context.Context, in *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error)
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

func (r *notificationResource) GetNotification(ctx *context.Context, in *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {
	getNotificationResponse, err := r.grpcClient.NotificationGRPC.GetNotification(ctx, in)
	if err != nil {
		return nil, err
	}

	return getNotificationResponse, nil
}

func NewNotificationHandler(grpcClient *grpc.Grpc) INotification {
	return &notificationResource{
		grpcClient: grpcClient,
	}
}
