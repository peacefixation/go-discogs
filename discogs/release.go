package discogs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/peacefixation/go-discogs/discogs/param"
)

// Release a release
type Release struct {
	ArtistsSort       string          `json:"artists_sort"`
	Title             string          `json:"title"`
	ID                int             `json:"id"`
	Artists           []ReleaseArtist `json:"artists"`
	DataQuality       string          `json:"data_quality"`
	Thumb             string          `json:"thumb"`
	Community         Community       `json:"community"`
	Companies         []Company       `json:"companies"`
	Country           string          `json:"country"`
	DateAdded         string          `json:"date_added"`
	DateChanged       string          `json:"date_changed"`
	EstimatedWeight   int             `json:"estimated_weight"`
	ExtraArtists      []ReleaseArtist `json:"extraartists"`
	FormatQuantity    int             `json:"format_quantity"`
	Formats           []Format        `json:"formats"`
	Genres            []string        `json:"genres"`
	Identifiers       []Identifier    `json:"identifiers"`
	Images            []Image         `json:"images"`
	Labels            []ReleaseLabel  `json:"labels"`
	LowestPrice       float32         `json:"lowest_price"`
	MasterID          int             `json:"master_id"`
	MasterURL         string          `json:"master_url"`
	Notes             string          `json:"notes"`
	NumForSale        int             `json:"num_for_sale"`
	Released          string          `json:"released"`
	ReleasedFormatted string          `json:"released_formatted"`
	ResourceURL       string          `json:"resource_url"`
	Series            []string        `json:"series"` // TODO: array of what?
	Status            string          `json:"status"`
	Styles            []string        `json:"styles"`
	Tracklist         []Track         `json:"tracklist"`
	URI               string          `json:"uri"`
	Videos            []Video         `json:"videos"`
	Year              int             `json:"year"`
	BlockedFromSale   bool            `json:"blocked_from_sale"`
}

// ReleaseArtist an artist associated with a release
type ReleaseArtist struct {
	ANV          string `json:"anv"`
	ID           int    `json:"id"`
	Join         string `json:"join"`
	Name         string `json:"name"`
	ResourceURL  string `json:"resource_url"`
	Role         string `json:"role"`
	ThumbnailURL string `json:"thumbnail_url"`
	Tracks       string `json:"tracks"`
}

// Community community details associated with a release
type Community struct {
	Contributors []Contributor `json:"contributors"`
	DataQuality  string        `json:"data_quality"`
	Have         int           `json:"have"`
	Rating       Rating        `json:"rating"`
	Status       string        `json:"status"`
	Submitter    Submitter     `json:"submitter"`
	Want         int           `json:"want"`
}

// Contributor a user that contributed to a release listing
type Contributor struct {
	ResourceURL string `json:"resource_url"`
	Username    string `json:"username"`
}

// Rating release rating details including average rating and total count of user ratings
type Rating struct {
	Average float32 `json:"average"`
	Count   int     `json:"count"`
}

// Submitter the submitter of the release details
type Submitter struct {
	ResourceURL string `json:"resource_url"`
	Username    string `json:"username"`
}

// Company a company associated with a release
type Company struct {
	CatNo          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ResourceURL    string `json:"resource_url"`
}

// Format a release format
type Format struct {
	Descriptions []string `json:"descriptions"`
	Name         string   `json:"name"`
	Quantity     string   `json:"qty"`
	Text         string   `json:"text"`
}

// Identifier an identifier of a release (i.e. barcode)
type Identifier struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// ReleaseLabel label details for a release
type ReleaseLabel struct {
	CatNo          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ResourceURL    string `json:"resource_url"`
	ThumbnailURL   string `json:"thumbnail_url"`
}

// UserReleaseRating user rating for a release
type UserReleaseRating struct {
	ReleaseID int    `json:"release_id"`
	Rating    int    `json:"rating"`
	Username  string `json:"username"`
}

// CommunityReleaseRating community release ratings including average rating and total count of user ratings
type CommunityReleaseRating struct {
	Rating    Rating `json:"rating"`
	ReleaseID int    `json:"release_id"`
}

// ReleaseStats count of users that have a release in their collection and count of users that have a release in their wantlist
type ReleaseStats struct {
	NumHave int `json:"num_have"`
	NumWant int `json:"num_want"`
}

// GetRelease get the release by ID
// optionally provide additional parameters (see param/release.go)
func (client *Client) GetRelease(ctx context.Context, releaseID int, params []param.Param) (*Release, error) {
	path := fmt.Sprintf("/releases/%d", releaseID)

	url, err := encodeURL(client.BaseURL, path, params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp Release

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetReleaseRatingByUser get the release rating for a user
func (client *Client) GetReleaseRatingByUser(ctx context.Context, releaseID int, username string) (*UserReleaseRating, error) {
	path := fmt.Sprintf("/releases/%d/rating/%s", releaseID, username)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp UserReleaseRating

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateReleaseRating update the release rating for a user
// Authentication as the user is required
func (client *Client) UpdateReleaseRating(ctx context.Context, releaseID int, username string, rating int) (*UserReleaseRating, error) {
	path := fmt.Sprintf("/releases/%d/rating/%s", releaseID, username)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	body := struct {
		Rating int `json:"rating"`
	}{
		Rating: rating,
	}

	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var resp UserReleaseRating

	err = client.sendRequest(req, &resp, http.StatusCreated)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// DeleteReleaseRating delete the release rating for a user
// Authentication as the user is required
func (client *Client) DeleteReleaseRating(ctx context.Context, releaseID int, username string) error {
	path := fmt.Sprintf("/releases/%d/rating/%s", releaseID, username)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url.String(), nil)
	if err != nil {
		return err
	}

	err = client.sendRequest(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}

	return nil
}

// GetCommunityReleaseRating get average rating and total number of user ratings for a release
func (client *Client) GetCommunityReleaseRating(ctx context.Context, releaseID int) (*CommunityReleaseRating, error) {
	path := fmt.Sprintf("/releases/%d/rating", releaseID)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp CommunityReleaseRating

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetReleaseStats get the total number of haves (in community collections)
// and wants (in community wantlists) for a release
func (client *Client) GetReleaseStats(ctx context.Context, releaseID int) (*ReleaseStats, error) {
	path := fmt.Sprintf("/releases/%d/stats", releaseID)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp ReleaseStats

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
