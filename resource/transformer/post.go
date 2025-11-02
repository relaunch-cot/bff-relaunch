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
