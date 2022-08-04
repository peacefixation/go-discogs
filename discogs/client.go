package discogs

import (
	"net/http"
	"time"

	"github.com/peacefixation/go-discogs/discogs/auth"
)

// DefaultHTTPClientTimeOut the timeout for the default HTTP client
const defaultHTTPClientTimeOut = 1 * time.Minute

// Client a Discogs API client
type Client struct {
	UserAgent  string
	Auth       auth.Auth
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient create a new Discogs API client
// A default HTTP client is provided
func NewClient(userAgent string, auth auth.Auth) *Client {
	return &Client{
		UserAgent: userAgent,
		Auth:      auth,
		BaseURL:   "https://api.discogs.com",
		HTTPClient: &http.Client{
			Timeout: defaultHTTPClientTimeOut,
		},
	}
}

// newTestClient create a new Discogs API client for use in tests
// A default HTTP client is provided
func newTestClient(baseURL, userAgent string, auth auth.Auth) *Client {
	return &Client{
		UserAgent: userAgent,
		Auth:      auth,
		BaseURL:   baseURL,
		HTTPClient: &http.Client{
			Timeout: defaultHTTPClientTimeOut,
		},
	}
}
