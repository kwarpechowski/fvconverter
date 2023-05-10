package warmix

type Order struct {
	BuyerOrderNumber    string `xml:"BuyerOrderNumber"`
	SupplierOrderNumber string `xml:"SupplierOrderNumber"`
	BuyerOrderDate      string `xml:"BuyerOrderDate"`
}
