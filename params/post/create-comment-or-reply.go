package post

type CreateCommentOrReplyPOST struct {
	Content         string `json:"content"`
	ParentCommentId string `json:"parentCommentId"`
}
