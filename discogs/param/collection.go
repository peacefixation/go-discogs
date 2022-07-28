package param

// CollectionSort the sort parameter for the GetCollectionItemsByFolder request
type CollectionSort int

const (
	CollectionSortLabel CollectionSort = iota
	CollectionSortArtist
	CollectionSortTitle
	CollectionSortCatNo
	CollectionSortFormat
	CollectionSortRating
	CollectionSortAdded
	CollectionSortYear
)

func (as CollectionSort) Key() string {
	return "sort"
}

func (as CollectionSort) Value() string {
	return []string{"label", "artist", "title", "catno", "format", "rating", "added", "year"}[as]
}

// CollectionSortOrder the sort_order parameter request
type CollectionSortOrder int

const (
	CollectionSortOrderAsc CollectionSortOrder = iota
	CollectionSortOrderDesc
)

func (as CollectionSortOrder) Key() string {
	return "sort_order"
}

func (as CollectionSortOrder) Value() string {
	return []string{"asc", "desc"}[as]
}
