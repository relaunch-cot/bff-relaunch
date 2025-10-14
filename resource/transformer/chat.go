package transformer

import pb "github.com/relaunch-cot/lib-relaunch-cot/proto/chat"

func CreateNewChatToProto(userIds []int64, createdBy int64) (*pb.CreateNewChatRequest, error) {
	return &pb.CreateNewChatRequest{
		UserIds:   userIds,
		CreatedBy: createdBy,
	}, nil
}

func SendMessageToProto(chatId, senderId int64, messageContent string) (*pb.SendMessageRequest, error) {
	return &pb.SendMessageRequest{
		ChatId:         chatId,
		SenderId:       senderId,
		MessageContent: messageContent,
	}, nil
}

func GetAllMessagesFromChatToProto(chatId int64) (*pb.GetAllMessagesFromChatRequest, error) {
	return &pb.GetAllMessagesFromChatRequest{ChatId: chatId}, nil
}
