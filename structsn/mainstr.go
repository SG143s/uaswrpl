package structsn

type Items struct {
	Indexs    int    `json:indexs`
	Urls      string `json:url`
	Name      string `json:name`
	Sku       string `json:sku`
	SellP     int    `json:sellp`
	OrP       string `json:orp`
	Cur       string `json:cur`
	Avail     string `json:ava`
	Col       string `json:col`
	Cat       string `json:cat`
	Sour      string `json:sour`
	WebSou    string `json:wsour`
	Breadc    string `json:breadc`
	Desc      string `json:desc`
	Brand     string `json:br`
	Images    string `json:img`
	Country   string `json:country`
	Lang      string `json:lang`
	AvgRate   int    `json:avgrate`
	RevC      int    `json:revc`
	Crawledat string `json:crat`
}
