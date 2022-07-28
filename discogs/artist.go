package discogs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/peacefixation/go-discogs/discogs/param"
)

// Artist an artist (or group)
type Artist struct {
	Aliases        []Alias  `json:"aliases"`
	DataQuality    string   `json:"data_quality"`
	ID             int      `json:"id"`
	Images         []Image  `json:"images"`
	Members        []Member `json:"members"`
	Name           string   `json:"name"`
	NameVariations []string `json:"namevariations"` // TODO: Is this deprecated? Doesn't appear in response and seems to be replaced by Aliases
	Profile        string   `json:"profile"`
	ReleasesURL    string   `json:"releases_url"`
	ResourceURL    string   `json:"resource_url"`
	URI            string   `json:"uri"`
	URLs           []string `json:"urls"`
}

type Alias struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ResourceURL  string `json:"resource_url"`
	ThumbnailURL string `json:"thumbnail_url"`
}

// Member member of an (artist) group
type Member struct {
	Active      bool   `json:"active"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}

// ArtistReleases releases for an artist
type ArtistReleases struct {
	Pagination Pagination
	Releases   []ArtistRelease
}

// ArtistRelease release for an artist
type ArtistRelease struct {
	Artist      string `json:"artist"`
	Format      string `json:"format"`
	ID          int    `json:"id"`
	Label       string `json:"label"`
	MainRelease int    `json:"main_release"`
	ResourceURL string `json:"resource_url"`
	Role        string `json:"role"`
	Status      string `json:"status"`
	Stats       Stats  `json:"stats"`
	Thumb       string `json:"thumb"`
	Title       string `json:"title"`
	TrackInfo   string `json:"trackinfo"`
	Type        string `json:"type"`
	Year        int    `json:"year"`
}

// GetArtist get an artist by ID
func (client *Client) GetArtist(ctx context.Context, artistID int) (*Artist, error) {
	path := fmt.Sprintf("/artists/%d", artistID)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp Artist

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetArtistReleases get the releases for an artist by ID
// Optionally provide additional parameters (see param/artist.go)
// This endpoint accepts pagination parameters (see param/pagination.go)
func (client *Client) GetArtistReleases(ctx context.Context, artistID int, params []param.Param) (*ArtistReleases, error) {
	path := fmt.Sprintf("/artists/%d/releases", artistID)

	url, err := encodeURL(client.BaseURL, path, params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp ArtistReleases

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
