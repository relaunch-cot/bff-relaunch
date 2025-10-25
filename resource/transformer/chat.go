package transformer

import pb "github.com/relaunch-cot/lib-relaunch-cot/proto/chat"

func CreateNewChatToProto(userIds []string, createdBy string) (*pb.CreateNewChatRequest, error) {
	return &pb.CreateNewChatRequest{
		UserIds:   userIds,
		CreatedBy: createdBy,
	}, nil
}

func SendMessageToProto(chatId, senderId, messageContent string) (*pb.SendMessageRequest, error) {
	return &pb.SendMessageRequest{
		ChatId:         chatId,
		SenderId:       senderId,
		MessageContent: messageContent,
	}, nil
}

func GetAllMessagesFromChatToProto(chatId string) (*pb.GetAllMessagesFromChatRequest, error) {
	return &pb.GetAllMessagesFromChatRequest{ChatId: chatId}, nil
}

func GetAllChatsFromUserToProto(userId string) (*pb.GetAllChatsFromUserRequest, error) {
	return &pb.GetAllChatsFromUserRequest{UserId: userId}, nil
}

func GetChatFromUsersToProto(userIds []string) (*pb.GetChatFromUsersRequest, error) {
	return &pb.GetChatFromUsersRequest{
		UserIds: userIds,
	}, nil
}
