package anitech

type DocumentInvoice struct {
	InvoiceHeader  InvoiceHeader  `xml:"InvoiceHeader"`
	InvoiceParties InvoiceParties `xml:"InvoiceParties"`
	InvoiceLines   InvoiceLines   `xml:"InvoiceLines"`
	InvoiceSummary InvoiceSummary `xml:"InvoiceSummary"`
}
