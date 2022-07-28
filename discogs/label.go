package discogs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/peacefixation/go-discogs/discogs/param"
)

//  Label the label for a release
type Label struct {
	ContactInfo string      `json:"contact_info"`
	DataQuality string      `json:"data_quality"`
	ID          int         `json:"id"`
	Images      []Image     `json:"images"`
	Name        string      `json:"name"`
	ParentLabel ParentLabel `json:"parent_label"`
	Profile     string      `json:"profile"`
	ReleasesURL string      `json:"releases_url"`
	ResourceURL string      `json:"resource_url"`
	SubLabels   []SubLabel  `json:"sublabels"`
	URI         string      `json:"uri"`
	URLs        []string    `json:"urls"`
}

type ParentLabel struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}

// SubLabel a sub-label of a label
type SubLabel struct {
	ResourceURL string `json:"resource_url"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
}

// LabelReleases the releases for a label
type LabelReleases struct {
	Pagination Pagination     `json:"pagination"`
	Releases   []LabelRelease `json:"releases"`
}

// LabelRelease a release for a label
type LabelRelease struct {
	Artist      string `json:"artist"`
	CatNo       string `json:"catno"`
	Format      string `json:"format"`
	ID          int    `json:"id"`
	ResourceURL string `json:"resource_url"`
	Status      string `json:"status"`
	Stats       Stats  `json:"stats"`
	Thumb       string `json:"thumb"`
	Title       string `json:"title"`
	Year        int    `json:"year"`
}

// GetLabel get a label by ID
func (client *Client) GetLabel(ctx context.Context, labelID int) (*Label, error) {
	path := fmt.Sprintf("/labels/%d", labelID)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp Label

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetLabelReleases get the releases for an label by ID
// This endpoint accepts pagination parameters (see param/pagination.go)
func (client *Client) GetLabelReleases(ctx context.Context, labelID int, params []param.Param) (*LabelReleases, error) {
	path := fmt.Sprintf("/labels/%d/releases", labelID)

	url, err := encodeURL(client.BaseURL, path, params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp LabelReleases

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
