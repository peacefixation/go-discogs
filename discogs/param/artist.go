package param

// ArtistSort the sort parameter for the GetArtistReleases request
type ArtistSort int

const (
	ArtistSortYear ArtistSort = iota
	ArtistSortTitle
	ArtistSortFormat
)

func (as ArtistSort) Key() string {
	return "sort"
}

func (as ArtistSort) Value() string {
	return []string{"year", "title", "format"}[as]
}

// ArtistSortOrder the sort_order parameter request
type ArtistSortOrder int

const (
	ArtistSortOrderAsc ArtistSortOrder = iota
	ArtistSortOrderDesc
)

func (as ArtistSortOrder) Key() string {
	return "sort_order"
}

func (as ArtistSortOrder) Value() string {
	return []string{"asc", "desc"}[as]
}
