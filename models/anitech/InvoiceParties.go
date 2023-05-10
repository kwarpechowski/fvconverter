package anitech

type InvoiceParties struct {
	Seller Seller `xml:"Seller"`
	Buyer  Buyer  `xml:"Buyer"`
}
