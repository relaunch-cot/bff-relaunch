package transformer

import pbNotification "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"

func SendNotificationToProto(senderId, receiverId, title, content, notificationType string) (*pbNotification.SendNotificationRequest, error) {
	return &pbNotification.SendNotificationRequest{
		SenderId:   senderId,
		ReceiverId: receiverId,
		Title:      title,
		Content:    content,
		Type:       notificationType,
	}, nil
}

func GetNotificationToProto(notificationId string) (*pbNotification.GetNotificationRequest, error) {
	return &pbNotification.GetNotificationRequest{
		NotificationId: notificationId,
	}, nil
}

func GetAllNotificationsFromUserToProto(userId string) (*pbNotification.GetAllNotificationsFromUserRequest, error) {
	return &pbNotification.GetAllNotificationsFromUserRequest{
		UserId: userId,
	}, nil
}
