package mappers

import (
	"github.com/google/uuid"
	"regexp"
	"strings"
	"xmlconverter/models/anitech"
	"xmlconverter/models/warmix"
)

func GenerateFromWarmix(fv warmix.DocumentInvoice) anitech.DocumentInvoice {

	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	var taxSummaryLines []anitech.TaxSummaryLine

	for _, item := range fv.InvoiceSummary.TaxSummary.TaxSummaryLine {
		taxSummaryLines = append(
			taxSummaryLines,
			anitech.TaxSummaryLine{
				TaxRate:         strings.Replace(item.TaxRate, ".00", "", -1),
				TaxSymbol:       strings.Replace(item.TaxRate, ".00", "", -1) + "%",
				TaxCategoryCode: item.TaxCategoryCode,
				TaxAmount:       item.TaxAmount,
				TaxableAmount:   item.TaxableAmount,
				GrossAmount:     item.GrossAmount,
			})
	}

	var items []anitech.Line

	for _, item := range fv.InvoiceLines.Line {
		items = append(
			items,
			anitech.Line{
				Detail: anitech.Detail{
					LineNumber:             item.Detail.LineNumber,
					EAN:                    item.Detail.EAN,
					SupplierItemCode:       item.Detail.SupplierItemCode,
					Description:            item.Detail.ItemDescription,
					Type:                   item.Detail.ItemType,
					Quantity:               item.Detail.InvoiceQuantity,
					UnitOfMeasure:          item.Detail.UnitOfMeasure,
					UnitPrice:              item.Detail.InvoiceUnitNetPrice,
					PercentDiscount:        "0,00",
					UnitPriceAfterDiscount: item.Detail.InvoiceUnitNetPrice,
					TaxRate:                item.Detail.TaxRate,
					TaxSymbol:              strings.Replace(item.Detail.TaxRate, ".00", "", -1) + "%",
					TaxCategoryCode:        item.Detail.TaxCategoryCode,
					TaxReference: anitech.TaxReference{
						ReferenceType:   item.Detail.TaxReference.ReferenceType,
						ReferenceNumber: item.Detail.TaxReference.ReferenceNumber,
					},
					TaxAmount: item.Detail.TaxAmount,
					NetAmount: item.Detail.NetAmount,
				},
			},
		)
	}

	return anitech.DocumentInvoice{
		InvoiceHeader: anitech.InvoiceHeader{
			Application: anitech.Application{
				Name:          "Generator",
				Version:       "1.0.0",
				FormatVersion: "1",
				OperatorName:  "Przykładowy operator",
			},
			DocumentID:           uuid.NewString(),
			PriceType:            "NETTO",
			Number:               fv.InvoiceHeader.InvoiceNumber,
			InvoiceDate:          fv.InvoiceHeader.InvoiceDate,
			SalesDate:            fv.InvoiceHeader.SalesDate,
			Currency:             fv.InvoiceHeader.InvoiceCurrency,
			PaymentDueDate:       fv.InvoiceHeader.InvoicePaymentDueDate,
			PaymentTerms:         fv.InvoiceHeader.InvoicePaymentTerms,
			DocumentType:         "FV",
			DocumentFunctionCode: "0",
			Delivery: anitech.Delivery{
				DeliveryDate: fv.InvoiceHeader.Delivery.DeliveryDate,
			},
		},
		InvoiceParties: anitech.InvoiceParties{
			Seller: anitech.Seller{
				Company: anitech.Company{
					ILN:           fv.InvoiceParties.Seller.ILN,
					TaxID:         strings.Replace(fv.InvoiceParties.Seller.TaxID, "-", "", -1),
					AccountNumber: reg.ReplaceAllString(fv.InvoiceParties.Seller.AccountNumber, ""),
					Name:          fv.InvoiceParties.Seller.Name,
					Address1:      fv.InvoiceParties.Seller.StreetAndNumber,
					CityName:      fv.InvoiceParties.Seller.CityName,
					PostalCode:    fv.InvoiceParties.Seller.PostalCode,
					Country:       fv.InvoiceParties.Seller.Country,
				},
			},
			Buyer: anitech.Buyer{
				TaxID:      fv.InvoiceParties.Buyer.TaxID,
				Name:       fv.InvoiceParties.Buyer.Name,
				Address1:   fv.InvoiceParties.Buyer.StreetAndNumber,
				CityName:   fv.InvoiceParties.Buyer.CityName,
				PostalCode: fv.InvoiceParties.Buyer.PostalCode,
			},
		},
		InvoiceLines: anitech.InvoiceLines{
			Line: items,
		},
		InvoiceSummary: anitech.InvoiceSummary{
			TotalNetAmount:   fv.InvoiceSummary.TotalNetAmount,
			TotalTaxAmount:   fv.InvoiceSummary.TotalTaxAmount,
			TotalGrossAmount: fv.InvoiceSummary.TotalGrossAmount,
			TaxSummaryLines: anitech.TaxSummaryLines{
				TaxSummaryLines: taxSummaryLines,
			},
		},
	}
}
