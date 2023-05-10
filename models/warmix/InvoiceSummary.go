package warmix

type InvoiceSummary struct {
	TotalLines         string     `xml:"TotalLines"`
	TotalNetAmount     string     `xml:"TotalNetAmount"`
	TotalTaxableBasis  string     `xml:"TotalTaxableBasis"`
	TotalTaxAmount     string     `xml:"TotalTaxAmount"`
	TotalGrossAmount   string     `xml:"TotalGrossAmount"`
	GrossAmountInWords string     `xml:"GrossAmountInWords"`
	TaxSummary         TaxSummary `xml:"Tax-Summary"`
}
