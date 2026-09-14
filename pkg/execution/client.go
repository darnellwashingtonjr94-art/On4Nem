package execution

import (
	"net/http"
	"time"
)

type BookClient struct {
	HTTPClient *http.Client
	BaseURL    string
}

func NewBookClient(url string) *BookClient {
	return &BookClient{
		HTTPClient: &http.Client{Timeout: 3 * time.Second},
		BaseURL:    url,
	}
}
