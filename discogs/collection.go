package discogs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/peacefixation/go-discogs/discogs/param"
)

type CollectionFolder struct {
	Count       int    `json:"count"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}

type CollectionFolders struct {
	Folders []CollectionFolder `json:"folders"`
}

type CollectionItems struct {
	Pagination Pagination
	Releases   []CollectionRelease
}

type CollectionRelease struct {
	BasicInformation BasicInformation `json:"basic_information"`
	ID               int              `json:"id"`
	InstanceID       int              `json:"instance_id"`
	DateAdded        string           `json:"date_added"`
	FolderID         int              `json:"folder_id"`
	Notes            []Note           `json:"notes"`
	Rating           int              `json:"rating"`
	ResourceURL      string           `json:"resource_url"`
}

type BasicInformation struct {
	Artists     []ReleaseArtist `json:"artists"`
	CoverImage  string          `json:"cover_image"`
	Formats     []Format        `json:"formats"`
	Genres      []string        `json:"genres"`
	Labels      []ReleaseLabel  `json:"labels"`
	ID          int             `json:"id"`
	MasterID    int             `json:"master_id"`
	MasterURL   string          `json:"master_url"`
	ResourceURL string          `json:"resource_url"`
	Styles      []string        `json:"styles"`
	Thumb       string          `json:"thumb"`
	Title       string          `json:"title"`
	Year        int             `json:"year"`
}

type Note struct {
	FieldID int
	Value   string
}

type ListFields struct {
	Fields []ListField
}

type ListField struct {
	ID       int      `json:"id"`
	Lines    int      `json:"lines"`
	Name     string   `json:"name"`
	Options  []string `json:"options"`
	Position int      `json:"position"`
	Public   bool     `json:"public"`
	Type     string   `json:"type"`
}

type CollectionValue struct {
	Maximum string
	Median  string
	Minimum string
}

// GetCollectionFolders retrieves a list of folders in a user's collection.
//
// If the collection has been made private by its owner, authentication as the
// collection owner is required. If you are not authenticated as the collection
// owner, only the folder with ID 0 (the "All" folder) will be visible (if the
// requested user's collection is public).
func (client *Client) GetCollectionFolders(ctx context.Context, username string) (*CollectionFolders, error) {
	path := fmt.Sprintf("/users/%s/collection/folders", username)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp CollectionFolders

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateCollectionFolder creates a new folder in a user’s collection.
// Authentication as the collection owner is required.
func (client *Client) CreateCollectionFolder(ctx context.Context, username, folderName string) (*CollectionFolder, error) {
	path := fmt.Sprintf("/users/%s/collection/folders", username)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	body := struct {
		Name string `json:"name"`
	}{
		Name: folderName,
	}

	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var resp CollectionFolder

	err = client.sendRequest(req, &resp, http.StatusCreated)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetCollectionFolder retrieves metadata about a folder in a user's collection.
// If folderID is not 0 (the "All" folder), authentication as the collection owner
// is required.
func (client *Client) GetCollectionFolder(ctx context.Context, username string, folderID int) (*CollectionFolder, error) {
	path := fmt.Sprintf("/users/%s/collection/folders/%d", username, folderID)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp CollectionFolder

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateCollectionFolder updates the name of a folder in the user's collection.
// Folders 0 (the "All" folder) and 1 (the "Uncategorized" folder) cannot be renamed.
// Authentication as the collection owner is required.
func (client *Client) UpdateCollectionFolder(ctx context.Context, username string, folderID int, newFolderName string) (*CollectionFolder, error) {
	path := fmt.Sprintf("/users/%s/collection/folders/%d", username, folderID)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	body := struct {
		Name string `json:"name"`
	}{
		Name: newFolderName,
	}

	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var resp CollectionFolder

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// DeleteCollectionFolder deletes a folder from a user's collection.
// A folder must be empty before it can be deleted.
// Authentication as the collection owner is required.
func (client *Client) DeleteCollectionFolder(ctx context.Context, username string, folderID int) error {
	path := fmt.Sprintf("/users/%s/collection/folders/%d", username, folderID)

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

// GetCollectionItemsByRelease retrieve the user's collection folders which contain
// a specified release.
// This will also show information about each release instance.
// The releaseID must be non-zero.
// Authentication as the collection owner is required if the owner's collection is private.
func (client *Client) GetCollectionItemsByRelease(ctx context.Context, username string, releaseID int) (*CollectionItems, error) {
	path := fmt.Sprintf("/users/%s/collection/releases/%d", username, releaseID)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp CollectionItems

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Returns the list of item in a folder in a user's collection.
// Optionally provide additional parameters (see param/collection.go)
// This endpoint accepts Pagination parameters (see param/pagination.go)
// Basic information about each release is provided, suitable for display in a list.
// For detailed information, make another API call to fetch the corresponding release.
// If folderID is not 0 (the "All" folder), or the collection has been made private by
// its owner, authentication as the collection owner is required.
// If you are not authenticated as the collection owner, only public notes fields will
// be visible.
func (client *Client) GetCollectionItemsByFolder(ctx context.Context, username string, folderID int, params []param.Param) (*CollectionItems, error) {
	path := fmt.Sprintf("/users/%s/collection/folders/%d/releases", username, folderID)

	url, err := encodeURL(client.BaseURL, path, params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp CollectionItems

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddReleaseToCollectionFolder adds a release to a folder in a user's collection.
// The folderID must be non-zero – you can use 1 (the "Uncategorized" folder).
// Authentication as the collection owner is required.
func (client *Client) AddReleaseToCollectionFolder(ctx context.Context, username string, folderID, releaseID int) (*CollectionRelease, error) {
	path := fmt.Sprintf("/users/%s/collection/folders/%d/releases/%d", username, folderID, releaseID)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp CollectionRelease

	err = client.sendRequest(req, &resp, http.StatusCreated)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeRatingOfReleaseAndOrMoveReleaseToFolder
//
// TODO: This method is not implemented yet and it looks gross!
// 		 We already expose UpdateReleaseRating. Is rating an instance
//		 of a release different to rating the release?
//       Perhaps we can just expose the MoveReleaseToFolder functionality.
//
// Change the rating on a release and/or move the instance to another folder.
// This endpoint potentially takes 2 folder_id parameters: one in the URL (which
// is the folder you are requesting, and is required), and one in the request
// body (representing the folder you want to move the instance to, which is optional)
// Authentication as the collection owner is required.
func (client *Client) ChangeRatingOfReleaseAndOrMoveReleaseToFolder(
	ctx context.Context,
	username string,
	folderID int,
	releaseID int,
	instanceID int,
	toFolderID *int, // nil -> don't move the release to a new folder
	rating *int, // nil -> don't change the release rating
) error {
	// /users/{username}/collection/folders/{folder_id}/releases/{release_id}/instances/{instance_id}
	return errors.New("not implemented")
}

// DeleteReleaseFromCollectionFolder
// Remove an instance of a release from a user's collection folder.
// Authentication as the collection owner is required.
func (client *Client) DeleteReleaseFromCollectionFolder(ctx context.Context, username string, folderID int, releaseID int, instanceID int) error {
	path := fmt.Sprintf("/users/%s/collection/folders/%d/releases/%d/instances/%d", username, folderID, releaseID, instanceID)

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

// ListCustomFields retrieves a list of user-defined collection notes fields.
// These fields are available on every release in the collection.
// If the collection has been made private by its owner, authentication as the
// collection owner is required.
// If you are not authenticated as the collection owner, only fields with public
// set to true will be visible.
func (client *Client) ListCustomFields(ctx context.Context, username string) (*ListFields, error) {
	path := fmt.Sprintf("/users/%s/collection/fields", username)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp ListFields

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// EditFieldsInstance edit a field on an instance of a release in a user's
// collection folder
func (client *Client) EditFieldsInstance(ctx context.Context, username string, folderID int, releaseID int, instanceID int, fieldID int, value string) error {
	path := fmt.Sprintf("/users/%s/collection/folders/%d/releases/%d/instances/%d/fields/%d", username, folderID, releaseID, instanceID, fieldID)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return err
	}

	body := struct {
		Value string `json:"value"`
	}{
		Value: value,
	}

	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	err = client.sendRequest(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) GetCollectionValue(ctx context.Context, username string) (*CollectionValue, error) {
	path := fmt.Sprintf("/users/%s/collection/value", username)

	url, err := encodeURL(client.BaseURL, path, nil)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp CollectionValue

	err = client.sendRequest(req, &resp, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return &resp, err
}
