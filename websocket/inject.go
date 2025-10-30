package websocket

var NotificationManager *Manager

var ChatManager *Manager

func InitializeWebSocket() {
	NotificationManager = NewManager()
	ChatManager = NewManager()

	go NotificationManager.Run()
	go ChatManager.Run()
}
