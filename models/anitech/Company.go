package anitech

type Company struct {
	ILN                 string `xml:"ILN,omitempty"`
	TaxID               string `xml:"TaxID"`
	IDByBuyer           string `xml:"IDByBuyer,omitempty"`
	BankName            string `xml:"BankName,omitempty"`
	AccountNumber       string `xml:"AccountNumber"`
	BankNameVATPLN      string `xml:"BankNameVATPLN,omitempty"`
	AccountNumberVATPLN string `xml:"AccountNumberVATPLN,omitempty"`
	Name                string `xml:"Name"`
	Address1            string `xml:"Address1"`
	Address2            string `xml:"Address2,omitempty"`
	CityName            string `xml:"CityName"`
	PostalCode          string `xml:"PostalCode"`
	Country             string `xml:"Country"`
	CSK                 string `xml:"CSK,omitempty"`
	WEEERegisterNumber  string `xml:"WEEERegisterNumber,omitempty"`
	RegisterInformation string `xml:"RegisterInformation,omitempty"`
}
