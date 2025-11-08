package post

type DeleteCommentOrReplyDELETE struct {
	CommentId string `json:"commentId"`
	ReplyId   string `json:"replyId"`
}
