package anitech

type Detail struct {
	LineNumber             string       `xml:"LineNumber"`
	EAN                    string       `xml:"EAN"`
	SupplierItemCode       string       `xml:"SupplierItemCode"`
	Description            string       `xml:"Description"`
	Type                   string       `xml:"Type"`
	Quantity               string       `xml:"Quantity"`
	UnitOfMeasure          string       `xml:"UnitOfMeasure"`
	UnitPrice              string       `xml:"UnitPrice"`
	PercentDiscount        string       `xml:"PercentDiscount"`
	UnitPriceAfterDiscount string       `xml:"UnitPriceAfterDiscount"`
	TaxRate                string       `xml:"TaxRate"`
	TaxSymbol              string       `xml:"TaxSymbol"`
	TaxCategoryCode        string       `xml:"TaxCategoryCode"`
	TaxReference           TaxReference `xml:"TaxReference"`
	TaxAmount              string       `xml:"TaxAmount"`
	NetAmount              string       `xml:"NetAmount"`
}
