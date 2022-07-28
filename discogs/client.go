package discogs

import (
	"net/http"
	"time"

	"github.com/peacefixation/go-discogs/discogs/auth"
)

// BaseURL base URL for Discogs API
const BaseURL = "https://api.discogs.com"

// DefaultHTTPClientTimeOut the timeout for the default HTTP client
const DefaultHTTPClientTimeOut = 1 * time.Minute

// Client a Discogs API client
type Client struct {
	UserAgent  string
	Auth       auth.Auth
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient create a new Discogs API client
// A default HTTP client is provided
func NewClient(baseURL, userAgent string, auth auth.Auth) *Client {
	return &Client{
		UserAgent: userAgent,
		Auth:      auth,
		BaseURL:   baseURL,
		HTTPClient: &http.Client{
			Timeout: DefaultHTTPClientTimeOut,
		},
	}
}
