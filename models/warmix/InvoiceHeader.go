package warmix

type InvoiceHeader struct {
	InvoiceNumber         string   `xml:"InvoiceNumber"`
	InvoiceDate           string   `xml:"InvoiceDate"`
	SalesDate             string   `xml:"SalesDate"`
	InvoiceCurrency       string   `xml:"InvoiceCurrency"`
	InvoicePaymentDueDate string   `xml:"InvoicePaymentDueDate"`
	InvoicePaymentTerms   string   `xml:"InvoicePaymentTerms"`
	InvoicePaymentMeans   string   `xml:"InvoicePaymentMeans"`
	InvoicePostDate       string   `xml:"InvoicePostDate"`
	DocumentFunctionCode  string   `xml:"DocumentFunctionCode"`
	Order                 Order    `xml:"Order"`
	Delivery              Delivery `xml:"Delivery"`
}
