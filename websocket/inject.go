package websocket

var WSManager *Manager

func InitializeWebSocket() {
	WSManager = NewManager()
	go WSManager.Run()
}
