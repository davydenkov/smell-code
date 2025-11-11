package main

import (
	"fmt"
	"strings"
)

// 42. Renaming a method (Rename Method)
//
// BEFORE: Poorly named method
type CalculatorBefore struct{}

func (c CalculatorBefore) Calc(a, b int) int { // Unclear name
	return a + b
}

// AFTER: Rename method to be more descriptive
type CalculatorAfter struct{}

func (c CalculatorAfter) Add(a, b int) int {
	return a + b
}

// 43. Adding a parameter (Add Parameter)
//
// BEFORE: Method missing required parameter
type EmailSenderBefore struct{}

func (es EmailSenderBefore) SendEmail(to, subject, body string) {
	// Send email with default priority
	priority := "normal"
	fmt.Printf("Sending email to %s with priority %s\n", to, priority)
}

// AFTER: Add parameter
type EmailSenderAfter struct{}

func (es EmailSenderAfter) SendEmail(to, subject, body, priority string) {
	if priority == "" {
		priority = "normal"
	}
	fmt.Printf("Sending email to %s with priority %s\n", to, priority)
}

// 44. Deleting a parameter (Remove Parameter)
//
// BEFORE: Unnecessary parameter
type ReportGeneratorBefore struct{}

func (rg ReportGeneratorBefore) GenerateReport(data string, format string, includeHeader bool) {
	if format == "html" {
		// Always include header for HTML
		includeHeader = true
	}
	fmt.Printf("Generating %s report with header: %v\n", format, includeHeader)
}

// AFTER: Remove unnecessary parameter
type ReportGeneratorAfter struct{}

func (rg ReportGeneratorAfter) GenerateReport(data, format string) {
	includeHeader := (format == "html")
	fmt.Printf("Generating %s report with header: %v\n", format, includeHeader)
}

// 45. Separation of Query and Modifier (Separate Query from Modifier)
//
// BEFORE: Method that both queries and modifies
type BankAccountBefore struct {
	balance float64
}

func (ba *BankAccountBefore) Withdraw(amount float64) bool {
	if ba.balance >= amount {
		ba.balance -= amount
		return true
	}
	return false
}

// AFTER: Separate query from modifier
type BankAccountAfter struct {
	balance float64
}

func (ba BankAccountAfter) CanWithdraw(amount float64) bool {
	return ba.balance >= amount
}

func (ba *BankAccountAfter) Withdraw(amount float64) bool {
	if ba.CanWithdraw(amount) {
		ba.balance -= amount
		return true
	}
	return false
}

// 46. Parameterization of the method (Parameterize Method)
//
// BEFORE: Similar methods with different values
type ReportGeneratorParamBefore struct{}

func (rg ReportGeneratorParamBefore) GenerateWeeklyReport() {
	rg.generateReport(7)
}

func (rg ReportGeneratorParamBefore) GenerateMonthlyReport() {
	rg.generateReport(30)
}

func (rg ReportGeneratorParamBefore) GenerateQuarterlyReport() {
	rg.generateReport(90)
}

func (rg ReportGeneratorParamBefore) generateReport(days int) {
	fmt.Printf("Generating report for %d days\n", days)
}

// AFTER: Parameterize method
type ReportGeneratorParamAfter struct{}

func (rg ReportGeneratorParamAfter) GenerateReport(days int) {
	fmt.Printf("Generating report for %d days\n", days)
}

func (rg ReportGeneratorParamAfter) GenerateWeeklyReport() {
	rg.GenerateReport(7)
}

func (rg ReportGeneratorParamAfter) GenerateMonthlyReport() {
	rg.GenerateReport(30)
}

func (rg ReportGeneratorParamAfter) GenerateQuarterlyReport() {
	rg.GenerateReport(90)
}

// 47. Replacing a parameter with explicit methods (Replace Parameter with Explicit Methods)
//
// BEFORE: Parameter determines behavior
type EmployeeBefore struct {
	name   string
	salary float64
}

func (e EmployeeBefore) CalculateBonus(bonusType string) float64 {
	switch bonusType {
	case "performance":
		return e.salary * 0.1
	case "yearly":
		return e.salary * 0.05
	case "special":
		return e.salary * 0.15
	default:
		return 0
	}
}

// AFTER: Replace parameter with explicit methods
type EmployeeAfter struct {
	name   string
	salary float64
}

func (e EmployeeAfter) CalculatePerformanceBonus() float64 {
	return e.salary * 0.1
}

func (e EmployeeAfter) CalculateYearlyBonus() float64 {
	return e.salary * 0.05
}

func (e EmployeeAfter) CalculateSpecialBonus() float64 {
	return e.salary * 0.15
}

// 48. Save the Whole Object
//
// BEFORE: Passing individual fields
type OrderItemBefore struct {
	name     string
	price    float64
	quantity int
}

type OrderProcessorBefore struct{}

func (op OrderProcessorBefore) CalculateItemTotal(name string, price float64, quantity int) float64 {
	return price * float64(quantity)
}

// AFTER: Save the whole object
type OrderItemAfter struct {
	name     string
	price    float64
	quantity int
}

type OrderProcessorAfter struct{}

func (op OrderProcessorAfter) CalculateItemTotal(item OrderItemAfter) float64 {
	return item.price * float64(item.quantity)
}

// 49. Replacing a parameter with a method call (Replace Parameter with Method)
//
// BEFORE: Parameter passed that could be calculated
type CustomerBefore struct {
	name  string
	level string
}

type DiscountCalculatorBefore struct{}

func (dc DiscountCalculatorBefore) CalculateDiscount(customer CustomerBefore, purchaseAmount float64) float64 {
	discountRate := 0.0
	switch customer.level {
	case "gold":
		discountRate = 0.1
	case "silver":
		discountRate = 0.05
	case "bronze":
		discountRate = 0.02
	}
	return purchaseAmount * discountRate
}

// AFTER: Replace parameter with method call
type CustomerAfter struct {
	name  string
	level string
}

func (c CustomerAfter) GetDiscountRate() float64 {
	switch c.level {
	case "gold":
		return 0.1
	case "silver":
		return 0.05
	case "bronze":
		return 0.02
	default:
		return 0.0
	}
}

type DiscountCalculatorAfter struct{}

func (dc DiscountCalculatorAfter) CalculateDiscount(customer CustomerAfter, purchaseAmount float64) float64 {
	return purchaseAmount * customer.GetDiscountRate()
}

// 50. Introduction of the boundary object (Introduce Parameter Object)
//
// BEFORE: Multiple parameters
type ReservationServiceBefore struct{}

func (rs ReservationServiceBefore) MakeReservation(date, startTime, endTime, customerName, customerEmail string, partySize int) {
	fmt.Printf("Reservation for %s (%s) on %s from %s to %s for %d people\n",
		customerName, customerEmail, date, startTime, endTime, partySize)
}

// AFTER: Introduce parameter object
type ReservationDetails struct {
	Date          string
	StartTime     string
	EndTime       string
	CustomerName  string
	CustomerEmail string
	PartySize     int
}

type ReservationServiceAfter struct{}

func (rs ReservationServiceAfter) MakeReservation(details ReservationDetails) {
	fmt.Printf("Reservation for %s (%s) on %s from %s to %s for %d people\n",
		details.CustomerName, details.CustomerEmail, details.Date,
		details.StartTime, details.EndTime, details.PartySize)
}

// 51. Removing the Value Setting Method (Remove Setting Method)
//
// BEFORE: Setter that shouldn't exist
type ImmutableObjectBefore struct {
	value string
}

func (io *ImmutableObjectBefore) SetValue(value string) {
	io.value = value // This breaks immutability
}

func (io ImmutableObjectBefore) GetValue() string {
	return io.value
}

// AFTER: Remove setting method
type ImmutableObjectAfter struct {
	value string
}

func NewImmutableObjectAfter(value string) *ImmutableObjectAfter {
	return &ImmutableObjectAfter{value: value}
}

func (io ImmutableObjectAfter) GetValue() string {
	return io.value
}

// 52. Hiding a method (Hide Method)
//
// BEFORE: Public method that should be private
type DataProcessorBefore struct{}

func (dp DataProcessorBefore) ValidateInput(input string) bool {
	return len(strings.TrimSpace(input)) > 0
}

func (dp DataProcessorBefore) ProcessInput(input string) string {
	if dp.ValidateInput(input) {
		return strings.ToUpper(input)
	}
	return ""
}

// AFTER: Hide method
type DataProcessorAfter struct{}

func (dp DataProcessorAfter) ProcessInput(input string) string {
	if dp.validateInput(input) {
		return strings.ToUpper(input)
	}
	return ""
}

func (dp DataProcessorAfter) validateInput(input string) bool {
	return len(strings.TrimSpace(input)) > 0
}

// 53. Replacing the constructor with the factory method (Replace Constructor with Factory Method)
//
// BEFORE: Constructor with complex logic
type DatabaseConnectionBefore struct {
	host     string
	port     int
	database string
}

func NewDatabaseConnectionBefore(connectionString string) *DatabaseConnectionBefore {
	// Complex parsing logic in constructor
	parts := strings.Split(connectionString, ":")
	host := parts[0]
	port := 5432 // default
	database := "default"

	if len(parts) > 1 {
		// Parse port
		fmt.Sscanf(parts[1], "%d", &port)
	}
	if len(parts) > 2 {
		database = parts[2]
	}

	return &DatabaseConnectionBefore{
		host:     host,
		port:     port,
		database: database,
	}
}

// AFTER: Replace constructor with factory method
type DatabaseConnectionAfter struct {
	host     string
	port     int
	database string
}

func NewDatabaseConnectionAfter(connectionString string) *DatabaseConnectionAfter {
	return CreateDatabaseConnection(connectionString)
}

func CreateDatabaseConnection(connectionString string) *DatabaseConnectionAfter {
	// Complex parsing logic in factory method
	parts := strings.Split(connectionString, ":")
	host := parts[0]
	port := 5432 // default
	database := "default"

	if len(parts) > 1 {
		// Parse port
		fmt.Sscanf(parts[1], "%d", &port)
	}
	if len(parts) > 2 {
		database = parts[2]
	}

	return &DatabaseConnectionAfter{
		host:     host,
		port:     port,
		database: database,
	}
}

// Example usage for method calls
func demonstrateMethodCalls() {
	fmt.Println("=== 42. Rename Method ===")

	calcBefore := CalculatorBefore{}
	calcAfter := CalculatorAfter{}

	fmt.Printf("Result before: %d\n", calcBefore.Calc(5, 3))
	fmt.Printf("Result after: %d\n", calcAfter.Add(5, 3))

	fmt.Println("\n=== 43. Add Parameter ===")

	senderBefore := EmailSenderBefore{}
	senderAfter := EmailSenderAfter{}

	senderBefore.SendEmail("user@example.com", "Test", "Body")
	senderAfter.SendEmail("user@example.com", "Test", "Body", "high")

	fmt.Println("\n=== 44. Remove Parameter ===")

	reportBefore := ReportGeneratorBefore{}
	reportAfter := ReportGeneratorAfter{}

	reportBefore.GenerateReport("data", "html", false) // includeHeader ignored for HTML
	reportAfter.GenerateReport("data", "html")         // No unnecessary parameter

	fmt.Println("\n=== 45. Separate Query from Modifier ===")

	accountBefore := &BankAccountBefore{balance: 1000.0}
	accountAfter := &BankAccountAfter{balance: 1000.0}

	fmt.Printf("Can withdraw $500 before: %v\n", accountBefore.balance >= 500)
	successBefore := accountBefore.Withdraw(500)
	fmt.Printf("Withdrawal success before: %v, Balance: $%.2f\n", successBefore, accountBefore.balance)

	fmt.Printf("Can withdraw $500 after: %v\n", accountAfter.CanWithdraw(500))
	successAfter := accountAfter.Withdraw(500)
	fmt.Printf("Withdrawal success after: %v, Balance: $%.2f\n", successAfter, accountAfter.balance)

	fmt.Println("\n=== 46. Parameterize Method ===")

	genBefore := ReportGeneratorParamBefore{}
	genAfter := ReportGeneratorParamAfter{}

	fmt.Print("Weekly report before: ")
	genBefore.GenerateWeeklyReport()

	fmt.Print("Weekly report after: ")
	genAfter.GenerateWeeklyReport()

	fmt.Println("\n=== 47. Replace Parameter with Explicit Methods ===")

	employeeBefore := EmployeeBefore{name: "John", salary: 50000}
	employeeAfter := EmployeeAfter{name: "John", salary: 50000}

	fmt.Printf("Performance bonus before: $%.2f\n", employeeBefore.CalculateBonus("performance"))
	fmt.Printf("Performance bonus after: $%.2f\n", employeeAfter.CalculatePerformanceBonus())

	fmt.Println("\n=== 48. Preserve Whole Object ===")

	processorBefore := OrderProcessorBefore{}
	processorAfter := OrderProcessorAfter{}

	itemBefore := OrderItemBefore{name: "Widget", price: 10.0, quantity: 5}
	itemAfter := OrderItemAfter{name: "Widget", price: 10.0, quantity: 5}

	fmt.Printf("Total before: $%.2f\n", processorBefore.CalculateItemTotal(itemBefore.name, itemBefore.price, itemBefore.quantity))
	fmt.Printf("Total after: $%.2f\n", processorAfter.CalculateItemTotal(itemAfter))

	fmt.Println("\n=== 50. Introduce Parameter Object ===")

	serviceBefore := ReservationServiceBefore{}
	serviceAfter := ReservationServiceAfter{}

	// Before - many parameters
	serviceBefore.MakeReservation("2024-01-15", "19:00", "21:00", "John Doe", "john@example.com", 4)

	// After - single parameter object
	details := ReservationDetails{
		Date:          "2024-01-15",
		StartTime:     "19:00",
		EndTime:       "21:00",
		CustomerName:  "John Doe",
		CustomerEmail: "john@example.com",
		PartySize:     4,
	}
	serviceAfter.MakeReservation(details)
}
