package responses

import "net/http"

type UploadFailedResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewUploadFailedResponse(reason string) *UploadFailedResponse {
	return &UploadFailedResponse{
		Status:  http.StatusBadRequest,
		Message: reason,
	}
}
