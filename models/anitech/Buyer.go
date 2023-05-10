package anitech

type Buyer struct {
	ILN        string `xml:"ILN,omitempty"`
	TaxID      string `xml:"TaxID"`
	Name       string `xml:"Name"`
	Address1   string `xml:"Address1"`
	Address2   string `xml:"Address2,omitempty"`
	CityName   string `xml:"CityName,omitempty"`
	PostalCode string `xml:"PostalCode,omitempty"`
	CSK        string `xml:"CSK,omitempty"`
}
