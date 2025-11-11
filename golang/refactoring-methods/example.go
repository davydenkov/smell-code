package main

import "fmt"

// Example demonstrating various refactoring techniques translated from PHP to Go

// 01 - Extract Method
type OrderProcessor struct{}

func (op OrderProcessor) ProcessOrder(orderId int, items []map[string]interface{}) {
	fmt.Printf("Processing order %d\n", orderId)

	// Extracted method for validation
	if !op.validateOrder(items) {
		fmt.Println("Order validation failed")
		return
	}

	// Extracted method for calculation
	total := op.calculateTotal(items)
	fmt.Printf("Order total: $%.2f\n", total)

	// Extracted method for processing
	op.processItems(items)
}

func (op OrderProcessor) validateOrder(items []map[string]interface{}) bool {
	for _, item := range items {
		if quantity, ok := item["quantity"].(int); !ok || quantity <= 0 {
			return false
		}
		if price, ok := item["price"].(float64); !ok || price <= 0 {
			return false
		}
	}
	return true
}

func (op OrderProcessor) calculateTotal(items []map[string]interface{}) float64 {
	total := 0.0
	for _, item := range items {
		quantity := item["quantity"].(int)
		price := item["price"].(float64)
		total += float64(quantity) * price
	}
	return total
}

func (op OrderProcessor) processItems(items []map[string]interface{}) {
	for _, item := range items {
		fmt.Printf("Processing item: %s\n", item["name"])
	}
}

// 02 - Variable Refactoring
type DataProcessor struct{}

func (dp DataProcessor) ProcessRawData(rawData string) map[string]interface{} {
	// Extract variables for clarity
	lines := dp.splitData(rawData)
	header := dp.parseHeader(lines[0])
	records := dp.parseRecords(lines[1:])

	return map[string]interface{}{
		"header":  header,
		"records": records,
	}
}

func (dp DataProcessor) splitData(data string) []string {
	// Implementation would split data into lines
	return []string{"header,line1", "data1,value1", "data2,value2"}
}

func (dp DataProcessor) parseHeader(headerLine string) map[string]string {
	// Parse header logic
	return map[string]string{"columns": "id,name,value"}
}

func (dp DataProcessor) parseRecords(recordLines []string) []map[string]string {
	// Parse records logic
	return []map[string]string{
		{"id": "1", "name": "item1", "value": "100"},
	}
}

// 03 - Moving Features between Objects
type User struct {
	id    int
	name  string
	email string
}

func NewUser(id int, name, email string) *User {
	return &User{id: id, name: name, email: email}
}

func (u User) GetId() int {
	return u.id
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetEmail() string {
	return u.email
}

type UserService struct{}

func (us UserService) GetUserDisplayInfo(user *User) string {
	// Moved formatting logic from User to UserService
	return fmt.Sprintf("User: %s (%s)", user.GetName(), user.GetEmail())
}

// 04 - Data Organization
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

func (a Address) ToString() string {
	return fmt.Sprintf("%s, %s, %s %s", a.street, a.city, a.state, a.zipCode)
}

// 05 - Conditional Expressions
type Account struct {
	balance float64
}

func (a Account) CanWithdraw(amount float64) bool {
	// Simplified conditional expression
	return a.balance >= amount && amount > 0
}

func (a Account) GetAccountType() string {
	// Replaced complex conditional with method
	if a.balance < 0 {
		return "overdrawn"
	} else if a.balance < 100 {
		return "low_balance"
	} else {
		return "standard"
	}
}

func main() {
	fmt.Println("Refactoring Methods Examples in Go:")
	fmt.Println("01 - Extract Method: Break down large methods into smaller ones")
	fmt.Println("02 - Variable Refactoring: Use descriptive variable names")
	fmt.Println("03 - Moving Features: Move methods between objects for better cohesion")
	fmt.Println("04 - Data Organization: Group related data into objects")
	fmt.Println("05 - Conditional Expressions: Simplify complex conditionals")

	// Example usage
	processor := OrderProcessor{}
	items := []map[string]interface{}{
		{"name": "Widget", "quantity": 2, "price": 10.99},
		{"name": "Gadget", "quantity": 1, "price": 25.50},
	}
	processor.ProcessOrder(123, items)

	user := NewUser(1, "John Doe", "john@example.com")
	userService := UserService{}
	fmt.Println(userService.GetUserDisplayInfo(user))
}
