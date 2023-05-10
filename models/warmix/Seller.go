package warmix

type Seller struct {
	ILN             string `xml:"ILN"`
	TaxID           string `xml:"TaxID"`
	AccountNumber   string `xml:"AccountNumber"`
	Name            string `xml:"Name"`
	StreetAndNumber string `xml:"StreetAndNumber"`
	CityName        string `xml:"CityName"`
	PostalCode      string `xml:"PostalCode"`
	Country         string `xml:"Country"`
}
