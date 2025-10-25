package chat

type GetChatFromUsersGET struct {
	User1Id string `json:"user1Id" form:"user1Id"`
	User2Id string `json:"user2Id" form:"user2Id"`
}
