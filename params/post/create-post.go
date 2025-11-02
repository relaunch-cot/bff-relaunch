package post

type CreatePostPOST struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	Type         string `json:"type"`
	UrlImagePost string `json:"urlImagePost"`
}
