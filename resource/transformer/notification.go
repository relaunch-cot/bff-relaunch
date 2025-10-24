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
