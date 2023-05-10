package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"xmlconverter/mappers"
	"xmlconverter/models/warmix"
)

func main() {

	input := flag.String("input", "", "ścieżka do pliku wejściowego")
	mode := flag.String("mode", "warmix", "typ formatera")
	output := flag.String("output", "fv.xml", "ścieżka do pliku wyjściowego")
	flag.Parse()

	xmlFile, err := os.Open(*input)
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var fv warmix.DocumentInvoice
	xml.Unmarshal(byteValue, &fv)

	if *mode == "warmix" {
		response := mappers.GenerateFromWarmix(fv)
		binary, _ := xml.MarshalIndent(response, "", "   ")

		f, _ := os.Create(*output)
		f.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
		f.Write(binary)
		log.Println("OK WARMIX")
	} else {
		log.Println("nieobsługiwany format")
	}

}
