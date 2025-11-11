package main

import (
	"fmt"
	"time"
)

type OrderService struct{}

func (os OrderService) CreateOrder(
	customerId int,
	customerName, customerEmail, customerPhone string,
	customerAddress, customerCity, customerState, customerZipCode string,
	productId int,
	productName string,
	productPrice, quantity float64,
	taxRate, discountPercent float64,
	shippingMethod string,
	shippingCost float64,
	paymentMethod string,
	billingAddress, billingCity, billingState, billingZipCode string,
	notes string,
) map[string]interface{} {
	// Calculate totals
	subtotal := productPrice * quantity
	discountAmount := subtotal * (discountPercent / 100)
	taxableAmount := subtotal - discountAmount
	taxAmount := taxableAmount * (taxRate / 100)
	total := taxableAmount + taxAmount + shippingCost

	// Create order record
	orderData := map[string]interface{}{
		"customer_id":       customerId,
		"customer_name":     customerName,
		"customer_email":    customerEmail,
		"customer_phone":    customerPhone,
		"customer_address":  customerAddress,
		"customer_city":     customerCity,
		"customer_state":    customerState,
		"customer_zip":      customerZipCode,
		"product_id":        productId,
		"product_name":      productName,
		"product_price":     productPrice,
		"quantity":          quantity,
		"subtotal":          subtotal,
		"discount_percent":  discountPercent,
		"discount_amount":   discountAmount,
		"tax_rate":          taxRate,
		"tax_amount":        taxAmount,
		"shipping_method":   shippingMethod,
		"shipping_cost":     shippingCost,
		"payment_method":    paymentMethod,
		"billing_address":   billingAddress,
		"billing_city":      billingCity,
		"billing_state":     billingState,
		"billing_zip":       billingZipCode,
		"total":             total,
		"notes":             notes,
		"created_at":        time.Now().Format("2006-01-02 15:04:05"),
	}

	// In a real application, this would save to database
	return orderData
}

func (os OrderService) UpdateOrder(
	orderId int,
	customerId int,
	customerName, customerEmail, customerPhone string,
	customerAddress, customerCity, customerState, customerZipCode string,
	productId int,
	productName string,
	productPrice, quantity float64,
	taxRate, discountPercent float64,
	shippingMethod string,
	shippingCost float64,
	paymentMethod string,
	billingAddress, billingCity, billingState, billingZipCode string,
	notes string,
) map[string]interface{} {
	// Calculate totals
	subtotal := productPrice * quantity
	discountAmount := subtotal * (discountPercent / 100)
	taxableAmount := subtotal - discountAmount
	taxAmount := taxableAmount * (taxRate / 100)
	total := taxableAmount + taxAmount + shippingCost

	orderData := map[string]interface{}{
		"id":                 orderId,
		"customer_id":        customerId,
		"customer_name":      customerName,
		"customer_email":     customerEmail,
		"customer_phone":     customerPhone,
		"customer_address":   customerAddress,
		"customer_city":      customerCity,
		"customer_state":     customerState,
		"customer_zip":       customerZipCode,
		"product_id":         productId,
		"product_name":       productName,
		"product_price":      productPrice,
		"quantity":           quantity,
		"subtotal":           subtotal,
		"discount_percent":   discountPercent,
		"discount_amount":    discountAmount,
		"tax_rate":           taxRate,
		"tax_amount":         taxAmount,
		"shipping_method":    shippingMethod,
		"shipping_cost":      shippingCost,
		"payment_method":     paymentMethod,
		"billing_address":    billingAddress,
		"billing_city":       billingCity,
		"billing_state":      billingState,
		"billing_zip":        billingZipCode,
		"total":              total,
		"notes":              notes,
		"updated_at":         time.Now().Format("2006-01-02 15:04:05"),
	}

	return orderData
}

func main() {
	os := OrderService{}

	// Too many parameters - long parameter list smell!
	order := os.CreateOrder(
		1, "John Doe", "john@example.com", "555-1234",
		"123 Main St", "Anytown", "CA", "12345",
		101, "Widget", 29.99, 2,
		8.25, 10,
		"standard", 9.99,
		"credit_card",
		"123 Main St", "Anytown", "CA", "12345",
		"Handle with care",
	)

	fmt.Printf("Order total: $%.2f\n", order["total"])
}
