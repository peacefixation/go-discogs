package param

// CurrencyAbbreviation the curr_abbr parameter for the Release request
type CurrencyAbbreviation int

const (
	USD CurrencyAbbreviation = iota
	GBP
	EUR
	CAD
	AUD
	JPY
	CHF
	MXN
	BRL
	NZD
	SEK
	ZAR
)

func (st CurrencyAbbreviation) Key() string {
	return "curr_abbr"
}

func (st CurrencyAbbreviation) Value() string {
	return []string{"USD", "GBP", "EUR", "CAD", "AUD", "JPY", "CHF", "MXN", "BRL", "NZD", "SEK", "ZAR"}[st]
}
