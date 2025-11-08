package transformer

import pbPost "github.com/relaunch-cot/lib-relaunch-cot/proto/post"

func CreatePostToProto(userId, title, content, postType, urlImagePost string) (*pbPost.CreatePostRequest, error) {
	return &pbPost.CreatePostRequest{
		UserId:       userId,
		Title:        title,
		Content:      content,
		Type:         postType,
		UrlImagePost: urlImagePost,
	}, nil
}

func GetPostToProto(postId string) (*pbPost.GetPostRequest, error) {
	return &pbPost.GetPostRequest{
		PostId: postId,
	}, nil
}

func UpdatePostToProto(userId, postId, title, content, urlImagePost string) (*pbPost.UpdatePostRequest, error) {
	return &pbPost.UpdatePostRequest{
		UserId:       userId,
		PostId:       postId,
		Title:        title,
		Content:      content,
		UrlImagePost: urlImagePost,
	}, nil
}

func DeletePostToProto(userId, postId string) (*pbPost.DeletePostRequest, error) {
	return &pbPost.DeletePostRequest{
		UserId: userId,
		PostId: postId,
	}, nil
}

func GetAllPostsFromUserToProto(userId string) (*pbPost.GetAllPostsFromUserRequest, error) {
	return &pbPost.GetAllPostsFromUserRequest{
		UserId: userId,
	}, nil
}

func GetAllLikesFromPostToProto(userId, postId string) (*pbPost.GetAllLikesFromPostRequest, error) {
	return &pbPost.GetAllLikesFromPostRequest{
		PostId: postId,
		UserId: userId,
	}, nil
}

func UpdateLikesFromPostToProto(userId, postId, likeType, commentId string) (*pbPost.UpdateLikesFromPostOrCommentRequest, error) {
	return &pbPost.UpdateLikesFromPostOrCommentRequest{
		UserId:    userId,
		PostId:    postId,
		Type:      likeType,
		CommentId: commentId,
	}, nil
}

func CreateCommentOrReplyToProto(userId, postId, content, commentType, parentCommentId string) (*pbPost.CreateCommentOrReplyRequest, error) {
	return &pbPost.CreateCommentOrReplyRequest{
		UserId:          userId,
		PostId:          postId,
		Content:         content,
		Type:            commentType,
		ParentCommentId: parentCommentId,
	}, nil
}

func DeleteCommentOrReplyToProto(userId, replyId, commentId, commentType string) (*pbPost.DeleteCommentOrReplyRequest, error) {
	return &pbPost.DeleteCommentOrReplyRequest{
		UserId:    userId,
		ReplyId:   replyId,
		CommentId: commentId,
		Type:      commentType,
	}, nil
}

func GetAllCommentsFromPostToProto(userId, postId string) (*pbPost.GetAllCommentsFromPostRequest, error) {
	return &pbPost.GetAllCommentsFromPostRequest{
		PostId: postId,
		UserId: userId,
	}, nil
}
