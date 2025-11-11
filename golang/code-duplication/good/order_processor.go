package main

import "fmt"

type TaxCalculator struct {
	taxRates map[string]float64
}

func NewTaxCalculator() *TaxCalculator {
	return &TaxCalculator{
		taxRates: map[string]float64{
			"CA": 0.0825,
			"NY": 0.04,
			"TX": 0.0625,
		},
	}
}

func (tc TaxCalculator) CalculateTax(price float64, state string) float64 {
	taxRate := tc.taxRates[state]
	return price * taxRate
}

type ShippingCalculator struct {
	baseRate               float64
	weightRatePerUnit      float64
	distanceRatePerUnit    float64
}

func NewShippingCalculator() *ShippingCalculator {
	return &ShippingCalculator{
		baseRate:            5.0,
		weightRatePerUnit:   0.5,
		distanceRatePerUnit: 0.1,
	}
}

func (sc ShippingCalculator) CalculateShipping(weight float64, distance float64) float64 {
	weightRate := weight * sc.weightRatePerUnit
	distanceRate := distance * sc.distanceRatePerUnit

	return sc.baseRate + weightRate + distanceRate
}

type OrderProcessor struct {
	taxCalculator      *TaxCalculator
	shippingCalculator *ShippingCalculator
}

func NewOrderProcessor() *OrderProcessor {
	return &OrderProcessor{
		taxCalculator:      NewTaxCalculator(),
		shippingCalculator: NewShippingCalculator(),
	}
}

func (op OrderProcessor) CalculateTax(price float64, state string) float64 {
	return op.taxCalculator.CalculateTax(price, state)
}

func (op OrderProcessor) CalculateShipping(weight float64, distance float64) float64 {
	return op.shippingCalculator.CalculateShipping(weight, distance)
}

type InvoiceGenerator struct {
	taxCalculator      *TaxCalculator
	shippingCalculator *ShippingCalculator
}

func NewInvoiceGenerator() *InvoiceGenerator {
	return &InvoiceGenerator{
		taxCalculator:      NewTaxCalculator(),
		shippingCalculator: NewShippingCalculator(),
	}
}

func (ig InvoiceGenerator) CalculateTax(price float64, state string) float64 {
	return ig.taxCalculator.CalculateTax(price, state)
}

func (ig InvoiceGenerator) CalculateShipping(weight float64, distance float64) float64 {
	return ig.shippingCalculator.CalculateShipping(weight, distance)
}

type QuoteGenerator struct {
	taxCalculator      *TaxCalculator
	shippingCalculator *ShippingCalculator
}

func NewQuoteGenerator() *QuoteGenerator {
	return &QuoteGenerator{
		taxCalculator:      NewTaxCalculator(),
		shippingCalculator: NewShippingCalculator(),
	}
}

func (qg QuoteGenerator) CalculateTax(price float64, state string) float64 {
	return qg.taxCalculator.CalculateTax(price, state)
}

func (qg QuoteGenerator) CalculateShipping(weight float64, distance float64) float64 {
	return qg.shippingCalculator.CalculateShipping(weight, distance)
}

func main() {
	// Example usage
	op := NewOrderProcessor()
	fmt.Printf("Tax for $100 in CA: $%.2f\n", op.CalculateTax(100, "CA"))
	fmt.Printf("Shipping for 10lbs over 50 miles: $%.2f\n", op.CalculateShipping(10, 50))
}
