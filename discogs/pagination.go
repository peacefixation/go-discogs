package discogs

// Pagination details returned by requests that accept pagination parameters
type Pagination struct {
	Page    int            `json:"page"`
	Pages   int            `json:"pages"`
	Items   int            `json:"items"`
	PerPage int            `json:"per_page"`
	URLs    PaginationURLs `json:"urls"`
}

// PaginationURLs pagination URLs
type PaginationURLs struct {
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}
