package param

import "fmt"

// Page the page parameter for requests that accept pagination parameters
type Page struct {
	value int
}

func NewPaginationPage(page int) Page {
	return Page{
		value: page,
	}
}

func (p Page) Key() string {
	return "page"
}

func (p Page) Value() string {
	return fmt.Sprintf("%d", p.value)
}

// PerPage the per_page parameter for requests that accept pagination parameters
type PerPage struct {
	value int
}

func NewPerPage(perPage int) PerPage {
	return PerPage{
		value: perPage,
	}
}

func (pp PerPage) Key() string {
	return "per_page"
}

func (pp PerPage) Value() string {
	return fmt.Sprintf("%d", pp.value)
}
