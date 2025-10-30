package websocket

var NotificationManager *Manager

var ChatManager *Manager

var PresenceManager *Manager

func InitializeWebSocket() {
	NotificationManager = NewManager()
	ChatManager = NewManager()
	PresenceManager = NewManager()

	go NotificationManager.Run()
	go ChatManager.Run()
	go PresenceManager.Run()
}
