package warmix

type InvoiceParties struct {
	Buyer  Buyer  `xml:"Buyer"`
	Seller Seller `xml:"Seller"`
}
