package responses

type MediasCreatedResponse struct {
	Urls []string `json:"urls"`
}

func NewMediasCreatedResponse(urls []string) *MediasCreatedResponse {
	return &MediasCreatedResponse{
		Urls: urls,
	}
}

func (mcr *MediasCreatedResponse) GetUrls() []string {
	return mcr.Urls
}
