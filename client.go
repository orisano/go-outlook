package outlook

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"runtime"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

const (
	version = "0.1.0"
)

var (
	userAgent = fmt.Sprintf("Go/%s (%s-%s) go-outlook/%s", runtime.Version(), runtime.GOARCH, runtime.GOOS, version)
)

// Client struct description
type Client struct {
	baseURL     *url.URL
	httpClient  *http.Client
	tokenSource oauth2.TokenSource

	UserAgent string

	logger *log.Logger
}

func NewClient(urlStr string, tokenSource oauth2.TokenSource) (*Client, error) {
	if tokenSource == nil {
		return nil, errors.New("missing tokenSource")
	}

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.Wrapf(err, "url parse failed: %s", urlStr)
	}

	defaultLogger := log.New(ioutil.Discard, "", log.LstdFlags)
	return &Client{
		baseURL:     parsedURL,
		httpClient:  http.DefaultClient,
		tokenSource: tokenSource,

		UserAgent: userAgent,

		logger: defaultLogger,
	}, nil
}

func (c *Client) SetLogger(logger *log.Logger) {
	c.logger = logger
}
