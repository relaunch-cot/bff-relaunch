package notification

import (
	"context"

	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"
)

type INotificationGRPC interface {
	SendNotification(ctx *context.Context, in *pb.SendNotificationRequest) error
	GetNotification(ctx *context.Context, in *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error)
	GetAllNotificationsFromUser(ctx *context.Context, in *pb.GetAllNotificationsFromUserRequest) (*pb.GetAllNotificationsFromUserResponse, error)
	DeleteNotification(ctx *context.Context, in *pb.DeleteNotificationRequest) error
	DeleteAllNotificationsFromUser(ctx *context.Context, in *pb.DeleteAllNotificationsFromUserRequest) error
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

func (r *resource) GetNotification(ctx *context.Context, in *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {
	getNotificationResponse, err := r.grpcClient.GetNotification(*ctx, in)
	if err != nil {
		return nil, err
	}

	return getNotificationResponse, nil
}

func (r *resource) GetAllNotificationsFromUser(ctx *context.Context, in *pb.GetAllNotificationsFromUserRequest) (*pb.GetAllNotificationsFromUserResponse, error) {
	getAllNotificationsFromUserResponse, err := r.grpcClient.GetAllNotificationsFromUser(*ctx, in)
	if err != nil {
		return nil, err
	}

	return getAllNotificationsFromUserResponse, nil
}

func (r *resource) DeleteNotification(ctx *context.Context, in *pb.DeleteNotificationRequest) error {
	_, err := r.grpcClient.DeleteNotification(*ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *resource) DeleteAllNotificationsFromUser(ctx *context.Context, in *pb.DeleteAllNotificationsFromUserRequest) error {
	_, err := r.grpcClient.DeleteAllNotificationsFromUser(*ctx, in)
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
