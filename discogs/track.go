package discogs

// Track a track is part of a tracklist for a release
type Track struct {
	Duration     string          `json:"duration"`
	Position     string          `json:"position"`
	Type         string          `json:"type_"` // TODO: why the trailing underscore?
	Title        string          `json:"title"`
	Artists      []ReleaseArtist `json:"artists"`
	ExtraArtists []ReleaseArtist `json:"extraartists"`
}
