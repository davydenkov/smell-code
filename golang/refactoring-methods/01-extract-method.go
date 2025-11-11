package main

import (
	"errors"
	"fmt"
)

// 1. Method Extraction (Extract Method)
//
// BEFORE: A method contains too much logic, making it hard to understand
type OrderProcessorBefore struct{}

func (op OrderProcessorBefore) ProcessOrder(order map[string]interface{}) (float64, error) {
	// Validate order
	if total, ok := order["total"].(float64); !ok || total <= 0 {
		return 0, errors.New("invalid order total")
	}

	// Calculate tax
	subtotal := order["subtotal"].(float64)
	tax := subtotal * 0.08

	// Calculate shipping
	weight := order["weight"].(float64)
	var shipping float64
	if weight > 10 {
		shipping = 15.00
	} else {
		shipping = 5.00
	}

	// Calculate total
	total := subtotal + tax + shipping

	// Save to database
	op.saveOrder(order, total)

	return total, nil
}

func (op OrderProcessorBefore) saveOrder(order map[string]interface{}, total float64) {
	// Database save logic would go here
	fmt.Printf("Saving order with total: $%.2f\n", total)
}

// AFTER: Extract methods to separate concerns
type OrderProcessorAfter struct{}

func (op OrderProcessorAfter) ProcessOrder(order map[string]interface{}) (float64, error) {
	if err := op.validateOrder(order); err != nil {
		return 0, err
	}

	tax := op.calculateTax(order)
	shipping := op.calculateShipping(order)
	total := op.calculateTotal(order, tax, shipping)

	op.saveOrder(order, total)

	return total, nil
}

func (op OrderProcessorAfter) validateOrder(order map[string]interface{}) error {
	if total, ok := order["total"].(float64); !ok || total <= 0 {
		return errors.New("invalid order total")
	}
	return nil
}

func (op OrderProcessorAfter) calculateTax(order map[string]interface{}) float64 {
	subtotal := order["subtotal"].(float64)
	return subtotal * 0.08
}

func (op OrderProcessorAfter) calculateShipping(order map[string]interface{}) float64 {
	weight := order["weight"].(float64)
	if weight > 10 {
		return 15.00
	}
	return 5.00
}

func (op OrderProcessorAfter) calculateTotal(order map[string]interface{}, tax, shipping float64) float64 {
	subtotal := order["subtotal"].(float64)
	return subtotal + tax + shipping
}

func (op OrderProcessorAfter) saveOrder(order map[string]interface{}, total float64) {
	// Database save logic would go here
	fmt.Printf("Saving order with total: $%.2f\n", total)
}

// 2. Embedding a method (Inline Method)
//
// BEFORE: A method is too simple and adds no value
type UserBefore struct {
	firstName string
	lastName  string
}

func (u UserBefore) GetFullName() string {
	return u.getFirstName() + " " + u.getLastName()
}

func (u UserBefore) getFirstName() string {
	return u.firstName
}

func (u UserBefore) getLastName() string {
	return u.lastName
}

// AFTER: Inline the simple method
type UserAfter struct {
	firstName string
	lastName  string
}

func (u UserAfter) GetFullName() string {
	return u.firstName + " " + u.lastName
}

// Example usage for extract method refactoring
func demonstrateExtractMethod() {
	fmt.Println("=== 1. Extract Method Refactoring ===")

	// Before refactoring - all logic in one method
	processorBefore := OrderProcessorBefore{}
	order := map[string]interface{}{
		"total":    150.0,
		"subtotal": 100.0,
		"weight":   5.0,
	}

	if total, err := processorBefore.ProcessOrder(order); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Order processed with total: $%.2f\n", total)
	}

	// After refactoring - logic separated into focused methods
	processorAfter := OrderProcessorAfter{}
	if total, err := processorAfter.ProcessOrder(order); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Order processed with total: $%.2f\n", total)
	}

	fmt.Println("\n=== 2. Inline Method Refactoring ===")

	// Before - using separate methods
	userBefore := UserBefore{firstName: "John", lastName: "Doe"}
	fmt.Printf("Full name (before): %s\n", userBefore.GetFullName())

	// After - inlined methods
	userAfter := UserAfter{firstName: "John", lastName: "Doe"}
	fmt.Printf("Full name (after): %s\n", userAfter.GetFullName())
}
