package warmix

type LineItem struct {
	LineNumber            string       `xml:"LineNumber"`
	OrderLineNumber       string       `xml:"OrderLineNumber"`
	EAN                   string       `xml:"EAN"`
	BuyerItemCode         string       `xml:"BuyerItemCode"`
	SupplierItemCode      string       `xml:"SupplierItemCode"`
	ItemDescription       string       `xml:"ItemDescription"`
	ItemType              string       `xml:"ItemType"`
	InvoiceQuantity       string       `xml:"InvoiceQuantity"`
	UnitOfMeasure         string       `xml:"UnitOfMeasure"`
	InvoiceUnitNetPrice   string       `xml:"InvoiceUnitNetPrice"`
	InvoiceUnitGrossPrice string       `xml:"InvoiceUnitGrossPrice"`
	TaxRate               string       `xml:"TaxRate"`
	TaxCategoryCode       string       `xml:"TaxCategoryCode"`
	TaxReference          TaxReference `xml:"TaxReference"`
	TaxAmount             string       `xml:"TaxAmount"`
	NetAmount             string       `xml:"NetAmount"`
}
