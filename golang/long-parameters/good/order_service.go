package main

import (
	"fmt"
	"time"
)

type Address struct {
	street  string
	city    string
	state   string
	zipCode string
}

func NewAddress(street, city, state, zipCode string) *Address {
	return &Address{
		street:  street,
		city:    city,
		state:   state,
		zipCode: zipCode,
	}
}

type Customer struct {
	id             int
	name           string
	email          string
	phone          string
	shippingAddress *Address
	billingAddress  *Address
}

func NewCustomer(id int, name, email, phone string, shippingAddress, billingAddress *Address) *Customer {
	return &Customer{
		id:              id,
		name:            name,
		email:           email,
		phone:           phone,
		shippingAddress: shippingAddress,
		billingAddress:  billingAddress,
	}
}

type Product struct {
	id    int
	name  string
	price float64
}

func NewProduct(id int, name string, price float64) *Product {
	return &Product{
		id:    id,
		name:  name,
		price: price,
	}
}

type OrderDetails struct {
	product         *Product
	quantity        float64
	taxRate         float64
	discountPercent float64
	shippingMethod  string
	paymentMethod   string
	notes           string
}

func NewOrderDetails(product *Product, quantity, taxRate, discountPercent float64, shippingMethod, paymentMethod, notes string) *OrderDetails {
	return &OrderDetails{
		product:         product,
		quantity:        quantity,
		taxRate:         taxRate,
		discountPercent: discountPercent,
		shippingMethod:  shippingMethod,
		paymentMethod:   paymentMethod,
		notes:           notes,
	}
}

type OrderCalculator struct{}

func (oc OrderCalculator) CalculateTotals(orderDetails *OrderDetails) map[string]float64 {
	subtotal := orderDetails.product.price * orderDetails.quantity
	discountAmount := subtotal * (orderDetails.discountPercent / 100)
	taxableAmount := subtotal - discountAmount
	taxAmount := taxableAmount * (orderDetails.taxRate / 100)

	// Simplified shipping cost calculation
	shippingCost := 9.99
	if orderDetails.shippingMethod == "express" {
		shippingCost = 19.99
	}

	total := taxableAmount + taxAmount + shippingCost

	return map[string]float64{
		"subtotal":       subtotal,
		"discount_amount": discountAmount,
		"tax_amount":     taxAmount,
		"shipping_cost":  shippingCost,
		"total":          total,
	}
}

type OrderService struct {
	calculator OrderCalculator
}

func NewOrderService() *OrderService {
	return &OrderService{
		calculator: OrderCalculator{},
	}
}

func (os OrderService) CreateOrder(customer *Customer, orderDetails *OrderDetails) map[string]interface{} {
	totals := os.calculator.CalculateTotals(orderDetails)

	orderData := map[string]interface{}{
		"customer_id":       customer.id,
		"customer_name":     customer.name,
		"customer_email":    customer.email,
		"customer_phone":    customer.phone,
		"customer_address":  customer.shippingAddress.street,
		"customer_city":     customer.shippingAddress.city,
		"customer_state":    customer.shippingAddress.state,
		"customer_zip":      customer.shippingAddress.zipCode,
		"product_id":        orderDetails.product.id,
		"product_name":      orderDetails.product.name,
		"product_price":     orderDetails.product.price,
		"quantity":          orderDetails.quantity,
		"subtotal":          totals["subtotal"],
		"discount_percent":  orderDetails.discountPercent,
		"discount_amount":   totals["discount_amount"],
		"tax_rate":          orderDetails.taxRate,
		"tax_amount":        totals["tax_amount"],
		"shipping_method":   orderDetails.shippingMethod,
		"shipping_cost":     totals["shipping_cost"],
		"payment_method":    orderDetails.paymentMethod,
		"billing_address":   customer.billingAddress.street,
		"billing_city":      customer.billingAddress.city,
		"billing_state":     customer.billingAddress.state,
		"billing_zip":       customer.billingAddress.zipCode,
		"total":             totals["total"],
		"notes":             orderDetails.notes,
		"created_at":        time.Now().Format("2006-01-02 15:04:05"),
	}

	// In a real application, this would save to database
	return orderData
}

func (os OrderService) UpdateOrder(orderId int, customer *Customer, orderDetails *OrderDetails) map[string]interface{} {
	totals := os.calculator.CalculateTotals(orderDetails)

	orderData := map[string]interface{}{
		"id":                 orderId,
		"customer_id":        customer.id,
		"customer_name":      customer.name,
		"customer_email":     customer.email,
		"customer_phone":     customer.phone,
		"customer_address":   customer.shippingAddress.street,
		"customer_city":      customer.shippingAddress.city,
		"customer_state":     customer.shippingAddress.state,
		"customer_zip":       customer.shippingAddress.zipCode,
		"product_id":         orderDetails.product.id,
		"product_name":       orderDetails.product.name,
		"product_price":     orderDetails.product.price,
		"quantity":          orderDetails.quantity,
		"subtotal":          totals["subtotal"],
		"discount_percent":  orderDetails.discountPercent,
		"discount_amount":   totals["discount_amount"],
		"tax_rate":          orderDetails.taxRate,
		"tax_amount":        totals["tax_amount"],
		"shipping_method":   orderDetails.shippingMethod,
		"shipping_cost":     totals["shipping_cost"],
		"payment_method":    orderDetails.paymentMethod,
		"billing_address":   customer.billingAddress.street,
		"billing_city":      customer.billingAddress.city,
		"billing_state":     customer.billingAddress.state,
		"billing_zip":       customer.billingAddress.zipCode,
		"total":             totals["total"],
		"notes":             orderDetails.notes,
		"updated_at":        time.Now().Format("2006-01-02 15:04:05"),
	}

	return orderData
}

func main() {
	os := NewOrderService()

	// Create objects instead of long parameter lists
	shippingAddr := NewAddress("123 Main St", "Anytown", "CA", "12345")
	billingAddr := NewAddress("123 Main St", "Anytown", "CA", "12345")
	customer := NewCustomer(1, "John Doe", "john@example.com", "555-1234", shippingAddr, billingAddr)

	product := NewProduct(101, "Widget", 29.99)
	orderDetails := NewOrderDetails(product, 2, 8.25, 10, "standard", "credit_card", "Handle with care")

	// Much cleaner method calls!
	order := os.CreateOrder(customer, orderDetails)

	fmt.Printf("Order total: $%.2f\n", order["total"])
	fmt.Println("Long parameter lists have been replaced with objects:")
	fmt.Println("- Customer object contains customer data and addresses")
	fmt.Println("- OrderDetails object contains order-specific data")
	fmt.Println("- Product object contains product information")
}
