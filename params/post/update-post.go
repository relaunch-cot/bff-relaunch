package post

type UpdatePostPUT struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	UrlImagePost string `json:"urlImagePost" `
}
