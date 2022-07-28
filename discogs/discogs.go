package discogs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/peacefixation/go-discogs/discogs/param"
)

// sendRequest send a request to the Discogs API
func (client *Client) sendRequest(req *http.Request, respObj any, successCode int) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", client.UserAgent)

	if client.Auth.String() != "" {
		req.Header.Set("Authorization", client.Auth.String())
	}

	var err error

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == successCode {
		if respObj != nil {
			err = json.NewDecoder(resp.Body).Decode(&respObj)
		}
	} else {
		err = decodeError(resp.Body, resp.StatusCode)
	}

	return err
}

// encodeURL encode the URL path including any parameters
func encodeURL(baseURL string, path string, params []param.Param) (*url.URL, error) {
	url, err := url.Parse(formatPath(baseURL, path))
	if err != nil {
		return nil, fmt.Errorf("URL parse error: %v", err)
	}

	query := url.Query()

	for _, param := range params {
		query.Add(param.Key(), param.Value())
	}

	url.RawQuery = query.Encode()

	return url, nil
}

// formatPath format a Discogs API path by preprending the baseURL to the path
// the path should begin with a forward slash '/'
func formatPath(baseURL, path string) string {
	return fmt.Sprintf("%s%s", baseURL, path)
}
