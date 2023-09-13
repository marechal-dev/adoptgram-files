package requests

type UploadFileRequest struct {
	PostID string `form:"postId" binding:"required"`
	files  string `form:"postId" binding:"required"`
}
