package chat

type CreateNewChatPOST struct {
	UserIds   []string `json:"userIds" form:"userIds,omitempty"`
	CreatedBy string   `json:"createdBy" form:"createdBy,omitempty"`
}
