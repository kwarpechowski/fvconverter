package anitech

type Application struct {
	Name          string `xml:"Name"`
	Version       string `xml:"Version"`
	FormatVersion string `xml:"FormatVersion"`
	OperatorName  string `xml:"OperatorName"`
}
