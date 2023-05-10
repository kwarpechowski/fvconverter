package anitech

type InvoiceSummary struct {
	TotalNetAmount   string          `xml:"TotalNetAmount"`
	TotalTaxAmount   string          `xml:"TotalTaxAmount"`
	TotalGrossAmount string          `xml:"TotalGrossAmount"`
	TaxSummaryLines  TaxSummaryLines `xml:"TaxSummaryLines"`
}
