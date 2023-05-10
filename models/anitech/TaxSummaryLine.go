package anitech

type TaxSummaryLine struct {
	TaxRate         string `xml:"TaxRate"`
	TaxSymbol       string `xml:"TaxSymbol"`
	TaxCategoryCode string `xml:"TaxCategoryCode"`
	TaxAmount       string `xml:"TaxAmount"`
	TaxableAmount   string `xml:"TaxableAmount"`
	GrossAmount     string `xml:"GrossAmount"`
}
