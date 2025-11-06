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

func UpdateLikesFromPostToProto(userId, postId string) (*pbPost.UpdateLikesFromPostRequest, error) {
	return &pbPost.UpdateLikesFromPostRequest{
		UserId: userId,
		PostId: postId,
	}, nil
}

func AddCommentToPostToProto(userId, postId, content string) (*pbPost.AddCommentToPostRequest, error) {
	return &pbPost.AddCommentToPostRequest{
		UserId:  userId,
		PostId:  postId,
		Content: content,
	}, nil
}

func RemoveCommentFromPostToProto(userId, postId, commentId string) (*pbPost.RemoveCommentFromPostRequest, error) {
	return &pbPost.RemoveCommentFromPostRequest{
		UserId:    userId,
		PostId:    postId,
		CommentId: commentId,
	}, nil
}
