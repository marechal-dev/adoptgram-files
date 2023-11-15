package responses

import "net/http"

type CouldNotParseFileResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewCouldNotParseFileResponse() *CouldNotParseFileResponse {
	return &CouldNotParseFileResponse{
		Status:  http.StatusInternalServerError,
		Message: "Could not upload a file",
	}
}
