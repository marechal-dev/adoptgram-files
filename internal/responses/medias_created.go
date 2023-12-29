package responses

type MediasCreatedResponse struct {
	Urls []string `json:"urls"`
}

type MediaCreatedResponse struct {
	Url string `json:"url"`
}

func NewMediasCreatedResponse(urls []string) *MediasCreatedResponse {
	return &MediasCreatedResponse{
		Urls: urls,
	}
}

func NewMediaCreatedResponse(url string) *MediaCreatedResponse {
	return &MediaCreatedResponse{
		Url: url,
	}
}

func (mcr *MediasCreatedResponse) GetUrls() []string {
	return mcr.Urls
}

func (mcr *MediaCreatedResponse) GetUrl() string {
	return mcr.Url
}
