package warmix

type DocumentInvoice struct {
	InvoiceHeader  InvoiceHeader  `xml:"Invoice-Header"`
	InvoiceParties InvoiceParties `xml:"Invoice-Parties"`
	InvoiceLines   InvoiceLines   `xml:"Invoice-Lines"`
	InvoiceSummary InvoiceSummary `xml:"Invoice-Summary"`
}
