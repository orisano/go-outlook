package outlook

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Client struct description
type Client struct {
	URL        *url.URL
	httpClient *http.Client

	token string

	logger *log.Logger
}

func NewClient(urlStr string, logger *log.Logger) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, err
	}
	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	return &Client{
		URL:        parsedURL,
		httpClient: http.DefaultClient,

		logger: logger,
	}
}
