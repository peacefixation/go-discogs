package discogs

import (
	"context"
	"net/http"
)

// Identity basic information of the authenticated user
type Identity struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	ResourceURL  string `json:"resource_url"`
	ConsumerName string `json:"consumer_name"`
}

// GetIdentity get basic information of authenticated user
//
// You can use this resource to find out who you’re authenticated as, and it
// also doubles as a good sanity check to ensure that you’re using OAuth correctly.
//
// For more detailed information, make another request for the user’s Profile.
func (client *Client) GetIdentity(ctx context.Context) (*Identity, error) {
	path := "/oauth/identity"

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp Identity

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
