package user

type CreateNewChatPOST struct {
	UserIds   []int64 `json:"userIds" form:"userIds,omitempty"`
	CreatedBy int64   `json:"createdBy" form:"createdBy,omitempty"`
}
