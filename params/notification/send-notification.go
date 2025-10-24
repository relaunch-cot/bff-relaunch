package notification

type SendNotificationPOST struct {
	ReceiverId string `json:"receiverId"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Type       string `json:"type"`
}
