package anitech

type InvoiceHeader struct {
	Application          Application `xml:"Application"`
	DocumentID           string      `xml:"DocumentID"`
	PriceType            string      `xml:"PriceType"`
	Number               string      `xml:"Number"`
	InvoiceDate          string      `xml:"InvoiceDate"`
	SalesDate            string      `xml:"SalesDate"`
	Currency             string      `xml:"Currency"`
	PaymentDueDate       string      `xml:"PaymentDueDate"`
	PaymentTerms         string      `xml:"PaymentTerms"`
	DocumentType         string      `xml:"DocumentType"`
	DocumentFunctionCode string      `xml:"DocumentFunctionCode"`
	PaymentName          string      `xml:"PaymentName,omitempty"`
	Delivery             Delivery    `xml:"Delivery"`
}
