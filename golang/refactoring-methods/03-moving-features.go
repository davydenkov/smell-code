package main

import (
	"fmt"
	"time"
)

// 9. Substitute Algorithm
//
// BEFORE: Complex algorithm that can be simplified
type PricingServiceBefore struct{}

func (ps PricingServiceBefore) CalculatePrice(items []map[string]interface{}) float64 {
	total := 0.0
	for _, item := range items {
		itemType := item["type"].(string)
		price := item["price"].(float64)

		if itemType == "book" {
			total += price * 0.9 // 10% discount for books
		} else if itemType == "electronics" {
			total += price * 1.1 // 10% markup for electronics
		} else {
			total += price
		}
	}
	return total
}

// AFTER: Substitute with a simpler algorithm
type PricingServiceAfter struct {
	discounts map[string]float64
}

func NewPricingServiceAfter() *PricingServiceAfter {
	return &PricingServiceAfter{
		discounts: map[string]float64{
			"book":        0.9,
			"electronics": 1.1,
			"default":     1.0,
		},
	}
}

func (ps PricingServiceAfter) CalculatePrice(items []map[string]interface{}) float64 {
	total := 0.0
	for _, item := range items {
		itemType := item["type"].(string)
		price := item["price"].(float64)

		multiplier, exists := ps.discounts[itemType]
		if !exists {
			multiplier = ps.discounts["default"]
		}

		total += price * multiplier
	}
	return total
}

// 10. Moving functions between objects (Move Method)
//
// BEFORE: Method in wrong class
type AccountBefore struct {
	balance float64
}

func NewAccountBefore(balance float64) *AccountBefore {
	return &AccountBefore{balance: balance}
}

func (a AccountBefore) GetBalance() float64 {
	return a.balance
}

// This method belongs in Bank, not Account
func (a *AccountBefore) TransferTo(target *AccountBefore, amount float64) bool {
	if a.balance >= amount {
		a.balance -= amount
		target.balance += amount
		return true
	}
	return false
}

// AFTER: Move method to appropriate class
type AccountAfter struct {
	balance float64
}

func NewAccountAfter(balance float64) *AccountAfter {
	return &AccountAfter{balance: balance}
}

func (a AccountAfter) GetBalance() float64 {
	return a.balance
}

func (a *AccountAfter) DecreaseBalance(amount float64) {
	a.balance -= amount
}

func (a *AccountAfter) IncreaseBalance(amount float64) {
	a.balance += amount
}

type Bank struct{}

func (b Bank) Transfer(from, to *AccountAfter, amount float64) bool {
	if from.GetBalance() >= amount {
		from.DecreaseBalance(amount)
		to.IncreaseBalance(amount)
		return true
	}
	return false
}

// 11. Moving the field (Move Field)
//
// BEFORE: Field in wrong class
type CustomerBefore struct {
	name    string
	address map[string]string // This should be in Address struct
}

func NewCustomerBefore(name, street, city, zipCode string) *CustomerBefore {
	return &CustomerBefore{
		name: name,
		address: map[string]string{
			"street":  street,
			"city":    city,
			"zipCode": zipCode,
		},
	}
}

func (c CustomerBefore) GetAddress() string {
	return fmt.Sprintf("%s, %s %s",
		c.address["street"], c.address["city"], c.address["zipCode"])
}

// AFTER: Move field to dedicated class
type Address struct {
	street  string
	city    string
	zipCode string
}

func NewAddress(street, city, zipCode string) *Address {
	return &Address{
		street:  street,
		city:    city,
		zipCode: zipCode,
	}
}

func (a Address) GetFullAddress() string {
	return fmt.Sprintf("%s, %s %s", a.street, a.city, a.zipCode)
}

type CustomerAfter struct {
	name    string
	address *Address
}

func NewCustomerAfter(name string, address *Address) *CustomerAfter {
	return &CustomerAfter{
		name:    name,
		address: address,
	}
}

func (c CustomerAfter) GetAddress() string {
	return c.address.GetFullAddress()
}

// 12. Class Allocation (Extract Class)
//
// BEFORE: Class has too many responsibilities
type PersonBefore struct {
	name             string
	phoneNumber      string
	officeAreaCode   string
	officeNumber     string
}

func (p PersonBefore) GetTelephoneNumber() string {
	return fmt.Sprintf("(%s) %s", p.officeAreaCode, p.officeNumber)
}

// AFTER: Extract telephone number to separate class
type TelephoneNumber struct {
	areaCode string
	number   string
}

func NewTelephoneNumber(areaCode, number string) *TelephoneNumber {
	return &TelephoneNumber{
		areaCode: areaCode,
		number:   number,
	}
}

func (tn TelephoneNumber) GetTelephoneNumber() string {
	return fmt.Sprintf("(%s) %s", tn.areaCode, tn.number)
}

type PersonAfter struct {
	name             string
	phoneNumber      string
	officeTelephone  *TelephoneNumber
}

func NewPersonAfter(name string) *PersonAfter {
	return &PersonAfter{name: name}
}

func (p PersonAfter) GetOfficeTelephone() string {
	if p.officeTelephone != nil {
		return p.officeTelephone.GetTelephoneNumber()
	}
	return ""
}

func (p *PersonAfter) SetOfficeTelephone(telephone *TelephoneNumber) {
	p.officeTelephone = telephone
}

// 13. Embedding a class (Inline Class)
//
// BEFORE: Unnecessary class with single responsibility
type OrderValidator struct{}

func (ov OrderValidator) IsValid(order map[string]interface{}) bool {
	if total, ok := order["total"].(float64); ok {
		return total > 0
	}
	return false
}

type OrderProcessorBefore struct {
	validator *OrderValidator
}

func NewOrderProcessorBefore() *OrderProcessorBefore {
	return &OrderProcessorBefore{
		validator: &OrderValidator{},
	}
}

func (op OrderProcessorBefore) Process(order map[string]interface{}) {
	if op.validator.IsValid(order) {
		fmt.Println("Processing order...")
	}
}

// AFTER: Inline the class
type OrderProcessorAfter struct{}

func (op OrderProcessorAfter) Process(order map[string]interface{}) {
	if op.isValidOrder(order) {
		fmt.Println("Processing order...")
	}
}

func (op OrderProcessorAfter) isValidOrder(order map[string]interface{}) bool {
	if total, ok := order["total"].(float64); ok {
		return total > 0
	}
	return false
}

// 14. Hiding delegation (Hide Delegate)
//
// BEFORE: Client has to know about delegation
type DepartmentBefore struct {
	manager *Person
}

func NewDepartmentBefore(manager *Person) *DepartmentBefore {
	return &DepartmentBefore{manager: manager}
}

func (d DepartmentBefore) GetManager() *Person {
	return d.manager
}

type Person struct {
	department *DepartmentBefore
}

func (p Person) GetDepartment() *DepartmentBefore {
	return p.department
}

// AFTER: Hide the delegation
type DepartmentAfter struct {
	manager *PersonAfter
}

func NewDepartmentAfter(manager *PersonAfter) *DepartmentAfter {
	return &DepartmentAfter{manager: manager}
}

func (d DepartmentAfter) GetManager() *PersonAfter {
	return d.manager
}

type PersonAfter struct {
	department *DepartmentAfter
}

func (p PersonAfter) GetDepartment() *DepartmentAfter {
	return p.department
}

func (p PersonAfter) GetManager() *PersonAfter {
	return p.department.GetManager()
}

// 16. Introduction of an external method (Introduce Foreign Method)
//
// BEFORE: Using external time method in wrong place
type ReportGeneratorBefore struct{}

func (rg ReportGeneratorBefore) GenerateReport() {
	date := time.Now()
	nextMonth := date.AddDate(0, 1, 0) // Foreign method usage

	fmt.Printf("Generating report for: %s\n", nextMonth.Format("2006-01"))
}

// AFTER: Introduce foreign method
type ReportGeneratorAfter struct{}

func (rg ReportGeneratorAfter) GenerateReport() {
	date := time.Now()
	nextMonth := rg.nextMonth(date)

	fmt.Printf("Generating report for: %s\n", nextMonth.Format("2006-01"))
}

func (rg ReportGeneratorAfter) nextMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, 0)
}

// 17. The introduction of local extension (Introduce Local Extension)
//
// AFTER: Create local extension (Go equivalent using composition)
type DateTimeExtension struct {
	time.Time
}

func NewDateTimeExtension(t time.Time) *DateTimeExtension {
	return &DateTimeExtension{Time: t}
}

func (dte DateTimeExtension) NextMonth() time.Time {
	return dte.AddDate(0, 1, 0)
}

func (dte DateTimeExtension) PreviousMonth() time.Time {
	return dte.AddDate(0, -1, 0)
}

// Example usage for moving features
func demonstrateMovingFeatures() {
	fmt.Println("=== 9. Substitute Algorithm ===")

	items := []map[string]interface{}{
		{"type": "book", "price": 20.0},
		{"type": "electronics", "price": 100.0},
		{"type": "clothing", "price": 50.0},
	}

	serviceBefore := PricingServiceBefore{}
	serviceAfter := NewPricingServiceAfter()

	fmt.Printf("Price before: $%.2f\n", serviceBefore.CalculatePrice(items))
	fmt.Printf("Price after: $%.2f\n", serviceAfter.CalculatePrice(items))

	fmt.Println("\n=== 10. Move Method ===")

	account1Before := NewAccountBefore(1000.0)
	account2Before := NewAccountBefore(500.0)

	account1After := NewAccountAfter(1000.0)
	account2After := NewAccountAfter(500.0)
	bank := Bank{}

	fmt.Printf("Transfer before: Account1=$%.0f, Account2=$%.0f\n",
		account1Before.GetBalance(), account2Before.GetBalance())
	account1Before.TransferTo(account2Before, 200)
	fmt.Printf("After transfer: Account1=$%.0f, Account2=$%.0f\n",
		account1Before.GetBalance(), account2Before.GetBalance())

	fmt.Printf("Transfer after: Account1=$%.0f, Account2=$%.0f\n",
		account1After.GetBalance(), account2After.GetBalance())
	bank.Transfer(account1After, account2After, 200)
	fmt.Printf("After transfer: Account1=$%.0f, Account2=$%.0f\n",
		account1After.GetBalance(), account2After.GetBalance())

	fmt.Println("\n=== 11. Move Field ===")

	customerBefore := NewCustomerBefore("John Doe", "123 Main St", "Anytown", "12345")
	address := NewAddress("123 Main St", "Anytown", "12345")
	customerAfter := NewCustomerAfter("John Doe", address)

	fmt.Printf("Address before: %s\n", customerBefore.GetAddress())
	fmt.Printf("Address after: %s\n", customerAfter.GetAddress())

	fmt.Println("\n=== 12. Extract Class ===")

	personBefore := PersonBefore{
		name:           "Jane Smith",
		officeAreaCode: "555",
		officeNumber:   "123-4567",
	}

	personAfter := NewPersonAfter("Jane Smith")
	phone := NewTelephoneNumber("555", "123-4567")
	personAfter.SetOfficeTelephone(phone)

	fmt.Printf("Phone before: %s\n", personBefore.GetTelephoneNumber())
	fmt.Printf("Phone after: %s\n", personAfter.GetOfficeTelephone())

	fmt.Println("\n=== 13. Inline Class ===")

	order := map[string]interface{}{"total": 150.0}

	processorBefore := NewOrderProcessorBefore()
	processorAfter := OrderProcessorAfter{}

	fmt.Print("Processing with validator class: ")
	processorBefore.Process(order)

	fmt.Print("Processing with inlined validation: ")
	processorAfter.Process(order)

	fmt.Println("\n=== 16. Introduce Foreign Method ===")

	reportBefore := ReportGeneratorBefore{}
	reportAfter := ReportGeneratorAfter{}

	fmt.Print("Report generation before: ")
	reportBefore.GenerateReport()

	fmt.Print("Report generation after: ")
	reportAfter.GenerateReport()

	fmt.Println("\n=== 17. Introduce Local Extension ===")

	now := time.Now()
	extendedTime := NewDateTimeExtension(now)

	fmt.Printf("Current time: %s\n", now.Format("2006-01-02"))
	fmt.Printf("Next month: %s\n", extendedTime.NextMonth().Format("2006-01-02"))
	fmt.Printf("Previous month: %s\n", extendedTime.PreviousMonth().Format("2006-01-02"))
}
