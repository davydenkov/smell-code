package main

import (
	"fmt"
	"sync"
)

// 18. Self-Encapsulate Field
//
// BEFORE: Direct field access (Go equivalent - exported field)
type PersonBefore struct {
	Name string // Direct access - not encapsulated
}

func (p PersonBefore) GetName() string {
	return p.Name
}

func (p *PersonBefore) SetName(name string) {
	p.Name = name
}

// AFTER: Self-encapsulate field
type PersonAfter struct {
	name string // Unexported field
}

func (p PersonAfter) GetName() string {
	return p.name
}

func (p *PersonAfter) SetName(name string) {
	p.name = name
}

// 19. Replacing the data value with an object (Replace Data Value with Object)
//
// BEFORE: Primitive data type that should be an object
type OrderBefore struct {
	customer string // Just a string
}

func (o OrderBefore) GetCustomerName() string {
	return o.customer
}

func (o *OrderBefore) SetCustomer(customer string) {
	o.customer = customer
}

// AFTER: Replace with object
type Customer struct {
	name string
}

func NewCustomer(name string) *Customer {
	return &Customer{name: name}
}

func (c Customer) GetName() string {
	return c.name
}

type OrderAfter struct {
	customer *Customer
}

func (o OrderAfter) GetCustomer() *Customer {
	return o.customer
}

func (o *OrderAfter) SetCustomer(customer *Customer) {
	o.customer = customer
}

func (o OrderAfter) GetCustomerName() string {
	if o.customer != nil {
		return o.customer.GetName()
	}
	return ""
}

// 20. Replacing the value with a reference (Change Value to Reference)
//
// BEFORE: Multiple instances of same object
type CustomerValue struct {
	name string
}

func NewCustomerValue(name string) *CustomerValue {
	return &CustomerValue{name: name}
}

func (c CustomerValue) GetName() string {
	return c.name
}

type OrderValue struct {
	customer *CustomerValue
}

func (o OrderValue) GetCustomer() *CustomerValue {
	return o.customer
}

// AFTER: Use shared reference
var (
	customerRegistry = make(map[string]*CustomerReference)
	registryMutex   sync.RWMutex
)

type CustomerReference struct {
	name string
}

func GetCustomerReference(name string) *CustomerReference {
	registryMutex.RLock()
	if customer, exists := customerRegistry[name]; exists {
		registryMutex.RUnlock()
		return customer
	}
	registryMutex.RUnlock()

	registryMutex.Lock()
	defer registryMutex.Unlock()

	// Double-check pattern
	if customer, exists := customerRegistry[name]; exists {
		return customer
	}

	customer := &CustomerReference{name: name}
	customerRegistry[name] = customer
	return customer
}

func (c CustomerReference) GetName() string {
	return c.name
}

type OrderReference struct {
	customer *CustomerReference
}

func (o OrderReference) GetCustomer() *CustomerReference {
	return o.customer
}

// 21. Replacing a reference with a value (Change Reference to Value)
//
// BEFORE: Using shared reference when independent copies are needed
type OrderRefBefore struct {
	customer *Customer
}

// AFTER: Use independent copies
type CustomerValueCopy struct {
	name string
}

func NewCustomerValueCopy(name string) CustomerValueCopy {
	return CustomerValueCopy{name: name}
}

func (c CustomerValueCopy) GetName() string {
	return c.name
}

type OrderRefAfter struct {
	customer CustomerValueCopy // Value type, not reference
}

func (o OrderRefAfter) GetCustomer() CustomerValueCopy {
	return o.customer
}

// 22. Replacing an array with an object (Replace Array with Object)
//
// BEFORE: Using map/array for structured data
type PerformanceDataBefore struct {
	data []interface{} // Array with mixed types
}

func (pd PerformanceDataBefore) GetGoals() int {
	return pd.data[0].(int)
}

func (pd PerformanceDataBefore) GetAssists() int {
	return pd.data[1].(int)
}

func (pd PerformanceDataBefore) GetMinutesPlayed() int {
	return pd.data[2].(int)
}

// AFTER: Replace with proper object
type PerformanceDataAfter struct {
	goals         int
	assists       int
	minutesPlayed int
}

func NewPerformanceDataAfter(goals, assists, minutesPlayed int) *PerformanceDataAfter {
	return &PerformanceDataAfter{
		goals:         goals,
		assists:       assists,
		minutesPlayed: minutesPlayed,
	}
}

func (pd PerformanceDataAfter) GetGoals() int {
	return pd.goals
}

func (pd PerformanceDataAfter) GetAssists() int {
	return pd.assists
}

func (pd PerformanceDataAfter) GetMinutesPlayed() int {
	return pd.minutesPlayed
}

// 23. Duplication of visible data (Duplicate Observed Data)
//
// BEFORE: Separating domain and presentation data
type ProductDomainBefore struct {
	name  string
	price float64
}

type ProductPresentationBefore struct {
	name  string
	price string // String representation for display
}

// AFTER: Duplicate data for separation of concerns
type ProductDomainAfter struct {
	name  string
	price float64
}

func (p ProductDomainAfter) GetDisplayPrice() string {
	return fmt.Sprintf("$%.2f", p.price)
}

type ProductPresentationAfter struct {
	name        string
	displayPrice string // Separate from domain
}

// 24. Replacing Unidirectional communication with Bidirectional
//
// BEFORE: Unidirectional relationship
type DepartmentUniBefore struct {
	name string
}

type EmployeeUniBefore struct {
	name       string
	department *DepartmentUniBefore
}

// AFTER: Bidirectional relationship
type DepartmentBiAfter struct {
	name      string
	employees []*EmployeeBiAfter
}

func (d *DepartmentBiAfter) AddEmployee(employee *EmployeeBiAfter) {
	d.employees = append(d.employees, employee)
	employee.department = d
}

type EmployeeBiAfter struct {
	name       string
	department *DepartmentBiAfter
}

// 25. Replacing Bidirectional communication with Unidirectional
//
// BEFORE: Unidirectional relationship (already unidirectional)

// AFTER: Keep unidirectional for simplicity (same as before)

// 26. Replacing the magic number with a symbolic constant
//
// BEFORE: Magic numbers
type CircleBefore struct {
	radius float64
}

func (c CircleBefore) GetArea() float64 {
	return 3.14159 * c.radius * c.radius // Magic number PI
}

func (c CircleBefore) GetCircumference() float64 {
	return 2 * 3.14159 * c.radius // Magic number PI again
}

// AFTER: Use symbolic constants
const PI = 3.14159

type CircleAfter struct {
	radius float64
}

func (c CircleAfter) GetArea() float64 {
	return PI * c.radius * c.radius
}

func (c CircleAfter) GetCircumference() float64 {
	return 2 * PI * c.radius
}

// Example usage for data organization
func demonstrateDataOrganization() {
	fmt.Println("=== 18. Self-Encapsulate Field ===")

	personBefore := PersonBefore{Name: "John"}
	personAfter := PersonAfter{}
	personAfter.SetName("John")

	fmt.Printf("Before: Name field is %s (exported)\n", personBefore.Name)
	fmt.Printf("After: Name is encapsulated, value: %s\n", personAfter.GetName())

	fmt.Println("\n=== 19. Replace Data Value with Object ===")

	orderBefore := OrderBefore{}
	orderBefore.SetCustomer("John Doe")

	customer := NewCustomer("John Doe")
	orderAfter := OrderAfter{}
	orderAfter.SetCustomer(customer)

	fmt.Printf("Before: Customer as string: %s\n", orderBefore.GetCustomerName())
	fmt.Printf("After: Customer as object: %s\n", orderAfter.GetCustomerName())

	fmt.Println("\n=== 20. Change Value to Reference ===")

	// Value objects - each order gets its own copy
	customerValue1 := NewCustomerValue("Jane Smith")
	customerValue2 := NewCustomerValue("Jane Smith") // Different instances

	orderValue1 := OrderValue{customer: customerValue1}
	orderValue2 := OrderValue{customer: customerValue2}

	fmt.Printf("Value objects: Different instances - %p vs %p\n",
		orderValue1.GetCustomer(), orderValue2.GetCustomer())

	// Reference objects - shared instances
	customerRef1 := GetCustomerReference("Jane Smith")
	customerRef2 := GetCustomerReference("Jane Smith") // Same instance

	orderRef1 := OrderReference{customer: customerRef1}
	orderRef2 := OrderReference{customer: customerRef2}

	fmt.Printf("Reference objects: Same instance - %p vs %p\n",
		orderRef1.GetCustomer(), orderRef2.GetCustomer())

	fmt.Println("\n=== 22. Replace Array with Object ===")

	// Before - using array
	dataBefore := PerformanceDataBefore{data: []interface{}{10, 5, 180}}
	fmt.Printf("Before: Goals=%d, Assists=%d, Minutes=%d\n",
		dataBefore.GetGoals(), dataBefore.GetAssists(), dataBefore.GetMinutesPlayed())

	// After - using struct
	dataAfter := NewPerformanceDataAfter(10, 5, 180)
	fmt.Printf("After: Goals=%d, Assists=%d, Minutes=%d\n",
		dataAfter.GetGoals(), dataAfter.GetAssists(), dataAfter.GetMinutesPlayed())

	fmt.Println("\n=== 26. Replace Magic Number with Symbolic Constant ===")

	circleBefore := CircleBefore{radius: 5.0}
	circleAfter := CircleAfter{radius: 5.0}

	fmt.Printf("Before: Area=%.2f, Circumference=%.2f (magic numbers)\n",
		circleBefore.GetArea(), circleBefore.GetCircumference())
	fmt.Printf("After: Area=%.2f, Circumference=%.2f (symbolic constant)\n",
		circleAfter.GetArea(), circleAfter.GetCircumference())
}
