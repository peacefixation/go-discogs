package discogs

import (
	"context"
	"net/http"

	"github.com/peacefixation/go-discogs/discogs/param"
)

// SearchResponse the response to a discogs database search
type SearchResponse struct {
	Pagination Pagination     `json:"pagination"`
	Results    []SearchResult `json:"results"`
}

// SearchResult details of a release returned for a search
type SearchResult struct {
	Style       []string              `json:"style"`
	Thumb       string                `json:"thumb"`
	Title       string                `json:"title"`
	Country     string                `json:"country"`
	Format      []string              `json:"format"`
	URI         string                `json:"uri"`
	Community   SearchResultCommunity `json:"community"`
	Label       []string              `json:"label"`
	CatNo       string                `json:"catno"`
	Year        string                `json:"year"`
	Genre       []string              `json:"genre"`
	ResourceURL string                `json:"resource_url"`
	Type        string                `json:"type"`
	ID          int                   `json:"id"`
}

// SearchResultCommunity the count of users that have a release in their collection or wantlist
type SearchResultCommunity struct {
	Want int `json:"want"`
	Have int `json:"have"`
}

// Search search the discogs database for a string
// Optionally provide additional parameters (see param/search.go)
// This endpoint accepts pagination parameters (see param/pagination.go)
func (client *Client) Search(ctx context.Context, searchTerm string, params []param.Param) (*SearchResponse, error) {
	path := "/database/search"

	if params == nil {
		params = make([]param.Param, 0)
	}

	params = append(params, param.NewSearchTerm(searchTerm))

	url, err := encodeURL(client.BaseURL, path, params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp SearchResponse

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
