package main

import "fmt"

type OrderProcessor struct{}

func (op OrderProcessor) CalculateTax(price float64, state string) float64 {
	taxRate := 0.0

	if state == "CA" {
		taxRate = 0.0825
	} else if state == "NY" {
		taxRate = 0.04
	} else if state == "TX" {
		taxRate = 0.0625
	}

	return price * taxRate
}

func (op OrderProcessor) CalculateShipping(weight float64, distance float64) float64 {
	baseRate := 5.0
	weightRate := weight * 0.5
	distanceRate := distance * 0.1

	return baseRate + weightRate + distanceRate
}

type InvoiceGenerator struct{}

func (ig InvoiceGenerator) CalculateTax(price float64, state string) float64 {
	taxRate := 0.0

	if state == "CA" {
		taxRate = 0.0825
	} else if state == "NY" {
		taxRate = 0.04
	} else if state == "TX" {
		taxRate = 0.0625
	}

	return price * taxRate
}

func (ig InvoiceGenerator) CalculateShipping(weight float64, distance float64) float64 {
	baseRate := 5.0
	weightRate := weight * 0.5
	distanceRate := distance * 0.1

	return baseRate + weightRate + distanceRate
}

type QuoteGenerator struct{}

func (qg QuoteGenerator) CalculateTax(price float64, state string) float64 {
	taxRate := 0.0

	if state == "CA" {
		taxRate = 0.0825
	} else if state == "NY" {
		taxRate = 0.04
	} else if state == "TX" {
		taxRate = 0.0625
	}

	return price * taxRate
}

func (qg QuoteGenerator) CalculateShipping(weight float64, distance float64) float64 {
	baseRate := 5.0
	weightRate := weight * 0.5
	distanceRate := distance * 0.1

	return baseRate + weightRate + distanceRate
}

func main() {
	// Example usage
	op := OrderProcessor{}
	fmt.Printf("Tax for $100 in CA: $%.2f\n", op.CalculateTax(100, "CA"))
	fmt.Printf("Shipping for 10lbs over 50 miles: $%.2f\n", op.CalculateShipping(10, 50))
}
