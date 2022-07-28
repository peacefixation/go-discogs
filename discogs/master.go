package discogs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/peacefixation/go-discogs/discogs/param"
)

// MasterRelease The Master resource represents a set of similar Releases.
// Masters (also known as "master releases") have a "main release" which is
// often the chronologically earliest.
type MasterRelease struct {
	Artists              []ReleaseArtist `json:"artists"`
	DataQuality          string          `json:"data_quality"`
	Genres               []string        `json:"genres"`
	ID                   int             `json:"id"`
	Images               []Image         `json:"images"`
	LowestPrice          float32         `json:"lowest_price"`
	MainRelease          int             `json:"main_release"`
	MainReleaseURL       string          `json:"main_release_url"`
	MostRecentRelease    int             `json:"most_recent_release"`
	MostRecentReleaseURL string          `json:"most_recent_release_url"`
	NumForSale           int             `json:"num_for_sale"`
	ResourceURL          string          `json:"resource_url"`
	Styles               []string        `json:"styles"`
	Title                string          `json:"title"`
	Tracklist            []Track         `json:"tracklist"`
	URI                  string          `json:"uri"`
	VersionsURL          string          `json:"versions_url"`
	Videos               []Video         `json:"videos"`
	Year                 int             `json:"year"`
}

// Video details of a video for a master release
type Video struct {
	Duration    int    `json:"duration"`
	Description string `json:"description"`
	Embed       bool   `json:"embed"`
	URI         string `json:"uri"`
	Title       string `json:"title"`
}

// MasterReleaseVersions releases that are versions of the master release
type MasterReleaseVersions struct {
	Pagination Pagination
	Versions   []MasterReleaseVersion
}

// MasterReleaseVersion a release that is a version of the master release
type MasterReleaseVersion struct {
	Status       string   `json:"status"`
	Stats        Stats    `json:"stats"`
	Thumb        string   `json:"thumb"`
	Format       string   `json:"format"`
	Country      string   `json:"country"`
	Title        string   `json:"title"`
	Label        string   `json:"label"`
	Released     string   `json:"released"`
	MajorFormats []string `json:"major_formats"`
	CatNo        string   `json:"catno"`
	ResourceURL  string   `json:"resource_url"`
	ID           int      `json:"id"`
}

// Stats count of users that have a version of the master release in their collection and wantlist
type Stats struct {
	User      UserStats      `json:"user"`
	Community CommunityStats `json:"community"`
}

// UserStats TODO: is this the stats for the authenticated user? API docs don't specify
type UserStats struct {
	InCollection int `json:"in_collection"`
	InWantlist   int `json:"in_wantlist"`
}

// CommunityStats count of users that have a version of the master release in their collection and wantlist
type CommunityStats struct {
	InCollection int `json:"in_collection"`
	InWantlist   int `json:"in_wantlist"`
}

// GetMasterRelease get the master release by ID
func (client *Client) GetMasterRelease(ctx context.Context, masterID int) (*MasterRelease, error) {
	path := fmt.Sprintf("/masters/%d", masterID)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp MasterRelease

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetMasterReleaseVersions get the master release versions for a master by ID
// This endpoint accepts pagination parameters (see param/pagination.go)
func (client *Client) GetMasterReleaseVersions(ctx context.Context, masterID int, params []param.Param) (*MasterReleaseVersions, error) {
	path := fmt.Sprintf("/masters/%d/versions", masterID)

	url, err := encodeURL(client.BaseURL, path, params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp MasterReleaseVersions

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
