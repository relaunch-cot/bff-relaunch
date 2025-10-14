package chat

type SendMessagePOST struct {
	ChatId         int64  `json:"chatId"`
	MessageContent string `json:"messageContent"`
}
