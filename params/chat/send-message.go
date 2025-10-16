package chat

type SendMessagePOST struct {
	ChatId         string `json:"chatId"`
	MessageContent string `json:"messageContent"`
}
