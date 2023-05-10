package warmix

type TaxSummaryLine struct {
	TaxRate         string `xml:"TaxRate"`
	TaxCategoryCode string `xml:"TaxCategoryCode"`
	TaxAmount       string `xml:"TaxAmount"`
	TaxableBasis    string `xml:"TaxableBasis"`
	TaxableAmount   string `xml:"TaxableAmount"`
	GrossAmount     string `xml:"GrossAmount"`
}

type TaxSummary struct {
	TaxSummaryLine []TaxSummaryLine `xml:"Tax-Summary-Line"`
}
