package discogs

// Image an image for a release
type Image struct {
	Height      int    `json:"height"`
	ResourceURL string `json:"resource_url"`
	Type        string `json:"type"`
	URI         string `json:"uri"`
	URI150      string `json:"uri150"`
	Width       int    `json:"width"`
}
