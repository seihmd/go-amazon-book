package amazonbook

type Response struct {
	IsValid string `xml:"Items>Request>IsValid"`
	Items   []Item `xml:"Items>Item"`
}

type Item struct {
	ASIN               string        `xml:"ASIN"`
	DetailPageURL      string        `xml:"DetailPageURL"`
	SalesRank          int           `xml:"SalesRank"`
	SmallImageURL      string        `xml:"SmallImage>URL"`
	MediumImageURL     string        `xml:"MediumImage>URL"`
	LargeImageURL      string        `xml:"LargeImage>URL"`
	ItemAttributes     ItemAttribute `xml:"ItemAttributes"`
	CustomerReviewsURL string        `xml:"CustomerReviews>IFrameURL"`
	HasReviews         bool          `xml:"CustomerReviews>HasReviews"`
}

type ItemAttribute struct {
	Author                  []string `xml:"Author"`
	Binding                 string   `xml:"Binding"`
	Edition                 string   `xml:"Edition"`
	Format                  string   `xml:"Format"`
	EISBN                   string   `xml:"EISBN"`
	IsAdultProduct          string   `xml:"IsAdultProduct"`
	ISBN                    string   `xml:"ISBN"`
	Languages               string   `xml:"Languages>Language>Name"`
	ListPriceAmout          int      `xml:"ListPrice>Amount"`
	ListPriceCurrencyCode   string   `xml:"ListPrice>CurrencyCode"`
	ListPriceFormattedPrice string   `xml:"ListPrice>FormattedPrice"`
	NumberOfPages           int      `xml:"NumberOfPages"`
	ProductGroup            string   `xml:"ProductGroup"`
	PublicationDate         string   `xml:"PublicationDate"`
	Publisher               string   `xml:"Publisher"`
	Title                   string   `xml:"Title"`
}
