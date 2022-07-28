package param

// SearchTerm the search term for the Search request
type SearchTerm string

func NewSearchTerm(value string) SearchTerm {
	return SearchTerm(value)
}

func (st SearchTerm) Key() string {
	return "q"
}

func (st SearchTerm) Value() string {
	return string(st)
}

// SearchType the sort parameter for the Search request
type SearchType int

const (
	Release SearchType = iota
	Master
	Artist
	Label
)

func (st SearchType) Key() string {
	return "type"
}

func (st SearchType) Value() string {
	return []string{"release", "master", "artist", "label"}[st]
}

// SearchTitle the title parameter for the Search request
// search by combined “Artist Name - Release Title” title field
type SearchTitle struct {
	value string
}

func NewSearchTitle(value string) SearchTitle {
	return SearchTitle{
		value: value,
	}
}

func (st SearchTitle) Key() string {
	return "title"
}

func (st SearchTitle) Value() string {
	return st.value
}

// SearchReleaseTitle the release_title parameter for the Search request
// search by combined "Artist Name - Release Title" title field
type SearchReleaseTitle struct {
	value string
}

func NewSearchReleaseTitle(value string) SearchReleaseTitle {
	return SearchReleaseTitle{
		value: value,
	}
}

func (srt SearchReleaseTitle) Key() string {
	return "release_title"
}

func (srt SearchReleaseTitle) Value() string {
	return srt.value
}

// SearchCredit the credit parameter for the Search request
type SearchCredit struct {
	value string
}

func NewSearchCredit(value string) SearchCredit {
	return SearchCredit{
		value: value,
	}
}

func (sc SearchCredit) Key() string {
	return "credit"
}

func (sc SearchCredit) Value() string {
	return sc.value
}

// SearchArtist the artist parameter for the Search request
type SearchArtist struct {
	value string
}

func NewSearchArtist(value string) SearchArtist {
	return SearchArtist{
		value: value,
	}
}

func (sa SearchArtist) Key() string {
	return "artist"
}

func (sa SearchArtist) Value() string {
	return sa.value
}

// SearchANV the anv parameter for the Search request
type SearchANV struct {
	value string
}

func NewSearchANV(value string) SearchANV {
	return SearchANV{
		value: value,
	}
}

func (sa SearchANV) Key() string {
	return "anv"
}

func (sa SearchANV) Value() string {
	return sa.value
}

// SearchLabel the label parameter for the Search request
type SearchLabel struct {
	value string
}

func NewSearchLabel(value string) SearchLabel {
	return SearchLabel{
		value: value,
	}
}

func (sl SearchLabel) Key() string {
	return "label"
}

func (sl SearchLabel) Value() string {
	return sl.value
}

// SearchGenre the genre parameter for the Search request
type SearchGenre struct {
	value string
}

func NewSearchGenre(value string) SearchGenre {
	return SearchGenre{
		value: value,
	}
}

func (sg SearchGenre) Key() string {
	return "genre"
}

func (sg SearchGenre) Value() string {
	return sg.value
}

// SearchStyle the style parameter for the Search request
type SearchStyle struct {
	value string
}

func NewSearchStyle(value string) SearchStyle {
	return SearchStyle{
		value: value,
	}
}

func (ss SearchStyle) Key() string {
	return "style"
}

func (ss SearchStyle) Value() string {
	return ss.value
}

// SearchCountry the country parameter for the Search request
type SearchCountry struct {
	value string
}

func NewSearchCountry(value string) SearchCountry {
	return SearchCountry{
		value: value,
	}
}

func (sc SearchCountry) Key() string {
	return "country"
}

func (sc SearchCountry) Value() string {
	return sc.value
}

// SearchYear the year parameter for the Search request
type SearchYear struct {
	value string
}

func NewSearchYear(value string) SearchYear {
	return SearchYear{
		value: value,
	}
}

func (sy SearchYear) Key() string {
	return "year"
}

func (sy SearchYear) Value() string {
	return sy.value
}

// SearchFormat the format parameter for the Search request
type SearchFormat struct {
	value string
}

func NewSearchFormat(value string) SearchFormat {
	return SearchFormat{
		value: value,
	}
}

func (sf SearchFormat) Key() string {
	return "format"
}

func (sf SearchFormat) Value() string {
	return sf.value
}

// SearchCatNo the catno parameter for the Search request
type SearchCatNo struct {
	value string
}

func NewSearchCatNo(value string) SearchCatNo {
	return SearchCatNo{
		value: value,
	}
}

func (scn SearchCatNo) Key() string {
	return "catno"
}

func (scn SearchCatNo) Value() string {
	return scn.value
}

// SearchBarcode the barcode parameter for the Search request
type SearchBarcode struct {
	value string
}

func NewSearchBarcode(value string) SearchBarcode {
	return SearchBarcode{
		value: value,
	}
}

func (sb SearchBarcode) Key() string {
	return "barcode"
}

func (sb SearchBarcode) Value() string {
	return sb.value
}

// SearchTrack the track parameter for the Search request
type SearchTrack struct {
	value string
}

func NewSearchTrack(value string) SearchTrack {
	return SearchTrack{
		value: value,
	}
}

func (st SearchTrack) Key() string {
	return "track"
}

func (st SearchTrack) Value() string {
	return st.value
}

// SearchSubmitter the submitter parameter for the Search request
type SearchSubmitter struct {
	value string
}

func NewSearchSubmitter(value string) SearchSubmitter {
	return SearchSubmitter{
		value: value,
	}
}

func (ss SearchSubmitter) Key() string {
	return "submitter"
}

func (ss SearchSubmitter) Value() string {
	return ss.value
}

// SearchContributor the contributor parameter for the Search request
type SearchContributor struct {
	value string
}

func NewSearchContributor(value string) SearchContributor {
	return SearchContributor{
		value: value,
	}
}

func (sc SearchContributor) Key() string {
	return "contributor"
}

func (sc SearchContributor) Value() string {
	return sc.value
}
