package param

// Param a URL parameter for Discogs API requests that accept parameters
type Param interface {
	Key() string
	Value() string
}
