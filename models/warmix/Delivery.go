package warmix

type Delivery struct {
	DeliveryLocationNumber string `xml:"DeliveryLocationNumber"`
	DeliveryDate           string `xml:"DeliveryDate"`
	DespatchNumber         string `xml:"DespatchNumber"`
}
