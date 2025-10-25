package handler

import (
	"context"

	"github.com/relaunch-cot/bff-relaunch/grpc"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"
)

type INotification interface {
	SendNotification(ctx *context.Context, in *pb.SendNotificationRequest) error
	GetNotification(ctx *context.Context, in *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error)
	GetAllNotificationsFromUser(ctx *context.Context, in *pb.GetAllNotificationsFromUserRequest) (*pb.GetAllNotificationsFromUserResponse, error)
	DeleteNotification(ctx *context.Context, in *pb.DeleteNotificationRequest) error
	DeleteAllNotificationsFromUser(ctx *context.Context, in *pb.DeleteAllNotificationsFromUserRequest) error
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

func (r *notificationResource) GetAllNotificationsFromUser(ctx *context.Context, in *pb.GetAllNotificationsFromUserRequest) (*pb.GetAllNotificationsFromUserResponse, error) {
	getAllNotificationsFromUserResponse, err := r.grpcClient.NotificationGRPC.GetAllNotificationsFromUser(ctx, in)
	if err != nil {
		return nil, err
	}

	return getAllNotificationsFromUserResponse, nil
}

func (r *notificationResource) DeleteNotification(ctx *context.Context, in *pb.DeleteNotificationRequest) error {
	err := r.grpcClient.NotificationGRPC.DeleteNotification(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *notificationResource) DeleteAllNotificationsFromUser(ctx *context.Context, in *pb.DeleteAllNotificationsFromUserRequest) error {
	err := r.grpcClient.NotificationGRPC.DeleteAllNotificationsFromUser(ctx, in)
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
