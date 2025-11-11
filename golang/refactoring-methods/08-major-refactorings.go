package main

import (
	"fmt"
	"strings"
)

// 69. Separation of inheritance (Tease Apart Inheritance)
//
// BEFORE: Mixed responsibilities in hierarchy
type EmployeeTeaseBefore struct {
	name string
	rate float64
}

func NewEmployeeTeaseBefore(name string, rate float64) *EmployeeTeaseBefore {
	return &EmployeeTeaseBefore{name: name, rate: rate}
}

func (e EmployeeTeaseBefore) GetName() string {
	return e.name
}

type SalariedEmployeeTeaseBefore struct {
	*EmployeeTeaseBefore
}

func (se SalariedEmployeeTeaseBefore) GetPay() float64 {
	return se.rate // Monthly salary
}

type CommissionedEmployeeTeaseBefore struct {
	*EmployeeTeaseBefore
	commission float64
}

func NewCommissionedEmployeeTeaseBefore(name string, rate, commission float64) *CommissionedEmployeeTeaseBefore {
	return &CommissionedEmployeeTeaseBefore{
		EmployeeTeaseBefore: NewEmployeeTeaseBefore(name, rate),
		commission:          commission,
	}
}

func (ce CommissionedEmployeeTeaseBefore) GetPay() float64 {
	return ce.rate + ce.commission
}

// AFTER: Tease apart inheritance using interfaces
type Payable interface {
	GetPay() float64
}

type EmployeeTeaseAfter struct {
	name string
}

func NewEmployeeTeaseAfter(name string) *EmployeeTeaseAfter {
	return &EmployeeTeaseAfter{name: name}
}

func (e EmployeeTeaseAfter) GetName() string {
	return e.name
}

type SalariedEmployeeTeaseAfter struct {
	*EmployeeTeaseAfter
	salary float64
}

func NewSalariedEmployeeTeaseAfter(name string, salary float64) *SalariedEmployeeTeaseAfter {
	return &SalariedEmployeeTeaseAfter{
		EmployeeTeaseAfter: NewEmployeeTeaseAfter(name),
		salary:             salary,
	}
}

func (se SalariedEmployeeTeaseAfter) GetPay() float64 {
	return se.salary
}

type CommissionedEmployeeTeaseAfter struct {
	*EmployeeTeaseAfter
	baseSalary float64
	commission float64
}

func NewCommissionedEmployeeTeaseAfter(name string, baseSalary, commission float64) *CommissionedEmployeeTeaseAfter {
	return &CommissionedEmployeeTeaseAfter{
		EmployeeTeaseAfter: NewEmployeeTeaseAfter(name),
		baseSalary:         baseSalary,
		commission:         commission,
	}
}

func (ce CommissionedEmployeeTeaseAfter) GetPay() float64 {
	return ce.baseSalary + ce.commission
}

// 70. Converting a procedural project into objects (Convert Procedural Design to Objects)
//
// BEFORE: Procedural functions with global state
var accountsBefore = make(map[string]float64)

func CreateAccountBefore(id string, balance float64) {
	accountsBefore[id] = balance
}

func GetBalanceBefore(id string) float64 {
	return accountsBefore[id]
}

func DepositBefore(id string, amount float64) {
	if balance, exists := accountsBefore[id]; exists {
		accountsBefore[id] = balance + amount
	}
}

func WithdrawBefore(id string, amount float64) bool {
	if balance, exists := accountsBefore[id]; exists && balance >= amount {
		accountsBefore[id] = balance - amount
		return true
	}
	return false
}

// AFTER: Convert to object-oriented design
type AccountObjectsAfter struct {
	id      string
	balance float64
}

func NewAccountObjectsAfter(id string, initialBalance float64) *AccountObjectsAfter {
	return &AccountObjectsAfter{
		id:      id,
		balance: initialBalance,
	}
}

func (a AccountObjectsAfter) GetID() string {
	return a.id
}

func (a AccountObjectsAfter) GetBalance() float64 {
	return a.balance
}

func (a *AccountObjectsAfter) Deposit(amount float64) {
	a.balance += amount
}

func (a *AccountObjectsAfter) Withdraw(amount float64) bool {
	if a.balance >= amount {
		a.balance -= amount
		return true
	}
	return false
}

type BankObjectsAfter struct {
	accounts map[string]*AccountObjectsAfter
}

func NewBankObjectsAfter() *BankObjectsAfter {
	return &BankObjectsAfter{
		accounts: make(map[string]*AccountObjectsAfter),
	}
}

func (b *BankObjectsAfter) CreateAccount(id string, initialBalance float64) {
	b.accounts[id] = NewAccountObjectsAfter(id, initialBalance)
}

func (b BankObjectsAfter) GetAccount(id string) *AccountObjectsAfter {
	return b.accounts[id]
}

// 71. Separating the domain from the representation (Separate Domain from Presentation)
//
// BEFORE: Domain logic mixed with presentation
type ProductDomainBefore struct {
	name  string
	price float64
}

func (p ProductDomainBefore) Display() string {
	return fmt.Sprintf("Product: %s - $%.2f", p.name, p.price) // Presentation logic in domain
}

func (p ProductDomainBefore) ApplyDiscount(discount float64) {
	p.price = p.price * (1 - discount) // Side effect on domain object
}

// AFTER: Separate domain from presentation
type ProductDomainAfter struct {
	name  string
	price float64
}

func NewProductDomainAfter(name string, price float64) *ProductDomainAfter {
	return &ProductDomainAfter{name: name, price: price}
}

func (p ProductDomainAfter) GetName() string {
	return p.name
}

func (p ProductDomainAfter) GetPrice() float64 {
	return p.price
}

func (p *ProductDomainAfter) SetPrice(price float64) {
	p.price = price
}

func (p *ProductDomainAfter) ApplyDiscount(discount float64) {
	p.price = p.price * (1 - discount)
}

// Presentation layer
type ProductPresenter struct{}

func (pp ProductPresenter) Display(product *ProductDomainAfter) string {
	return fmt.Sprintf("Product: %s - $%.2f", product.GetName(), product.GetPrice())
}

func (pp ProductPresenter) DisplayWithDiscount(product *ProductDomainAfter, discount float64) string {
	discountedPrice := product.GetPrice() * (1 - discount)
	return fmt.Sprintf("Product: %s - Regular: $%.2f, Sale: $%.2f",
		product.GetName(), product.GetPrice(), discountedPrice)
}

// 72. Hierarchy Extraction (Extract Hierarchy)
//
// BEFORE: Flat structure with conditional logic
type EmployeeHierarchyBefore struct {
	name     string
	empType  string
	salary   float64
	commission float64
	bonus    float64
}

func (e EmployeeHierarchyBefore) CalculatePay() float64 {
	switch e.empType {
	case "salaried":
		return e.salary
	case "commissioned":
		return e.salary + e.commission
	case "manager":
		return e.salary + e.bonus
	default:
		return 0
	}
}

// AFTER: Extract hierarchy
type EmployeeHierarchyAfter interface {
	CalculatePay() float64
	GetName() string
}

type BaseEmployee struct {
	name string
}

func (be BaseEmployee) GetName() string {
	return be.name
}

type SalariedEmployeeHierarchy struct {
	BaseEmployee
	salary float64
}

func NewSalariedEmployeeHierarchy(name string, salary float64) *SalariedEmployeeHierarchy {
	return &SalariedEmployeeHierarchy{
		BaseEmployee: BaseEmployee{name: name},
		salary:       salary,
	}
}

func (se SalariedEmployeeHierarchy) CalculatePay() float64 {
	return se.salary
}

type CommissionedEmployeeHierarchy struct {
	BaseEmployee
	salary     float64
	commission float64
}

func NewCommissionedEmployeeHierarchy(name string, salary, commission float64) *CommissionedEmployeeHierarchy {
	return &CommissionedEmployeeHierarchy{
		BaseEmployee: BaseEmployee{name: name},
		salary:       salary,
		commission:   commission,
	}
}

func (ce CommissionedEmployeeHierarchy) CalculatePay() float64 {
	return ce.salary + ce.commission
}

type ManagerEmployeeHierarchy struct {
	BaseEmployee
	salary float64
	bonus  float64
}

func NewManagerEmployeeHierarchy(name string, salary, bonus float64) *ManagerEmployeeHierarchy {
	return &ManagerEmployeeHierarchy{
		BaseEmployee: BaseEmployee{name: name},
		salary:       salary,
		bonus:        bonus,
	}
}

func (me ManagerEmployeeHierarchy) CalculatePay() float64 {
	return me.salary + me.bonus
}

// Example usage for major refactorings
func demonstrateMajorRefactorings() {
	fmt.Println("=== 69. Tease Apart Inheritance ===")

	// Before - mixed responsibilities
	salariedBefore := SalariedEmployeeTeaseBefore{
		EmployeeTeaseBefore: NewEmployeeTeaseBefore("John", 5000),
	}
	commissionedBefore := NewCommissionedEmployeeTeaseBefore("Jane", 3000, 500)

	fmt.Printf("Salaried pay before: $%.2f\n", salariedBefore.GetPay())
	fmt.Printf("Commissioned pay before: $%.2f\n", commissionedBefore.GetPay())

	// After - separated responsibilities
	var employeesAfter []Payable
	employeesAfter = append(employeesAfter, NewSalariedEmployeeTeaseAfter("John", 5000))
	employeesAfter = append(employeesAfter, NewCommissionedEmployeeTeaseAfter("Jane", 3000, 500))

	for i, emp := range employeesAfter {
		fmt.Printf("Employee %d pay after: $%.2f\n", i+1, emp.GetPay())
	}

	fmt.Println("\n=== 70. Convert Procedural Design to Objects ===")

	// Before - procedural
	CreateAccountBefore("ACC001", 1000)
	DepositBefore("ACC001", 500)
	fmt.Printf("Balance before: $%.2f\n", GetBalanceBefore("ACC001"))

	// After - object-oriented
	bank := NewBankObjectsAfter()
	bank.CreateAccount("ACC001", 1000)
	account := bank.GetAccount("ACC001")
	account.Deposit(500)
	fmt.Printf("Balance after: $%.2f\n", account.GetBalance())

	fmt.Println("\n=== 71. Separate Domain from Presentation ===")

	// Before - mixed concerns
	productBefore := ProductDomainBefore{name: "Widget", price: 100}
	fmt.Printf("Display before: %s\n", productBefore.Display())

	// After - separated concerns
	productAfter := NewProductDomainAfter("Widget", 100)
	presenter := ProductPresenter{}
	fmt.Printf("Display after: %s\n", presenter.Display(productAfter))
	fmt.Printf("Display with discount: %s\n", presenter.DisplayWithDiscount(productAfter, 0.1))

	fmt.Println("\n=== 72. Extract Hierarchy ===")

	// Before - conditional logic
	empBefore := EmployeeHierarchyBefore{name: "Bob", empType: "commissioned", salary: 4000, commission: 600}
	fmt.Printf("Pay before: $%.2f\n", empBefore.CalculatePay())

	// After - proper hierarchy
	var employeeAfter EmployeeHierarchyAfter = NewCommissionedEmployeeHierarchy("Bob", 4000, 600)
	fmt.Printf("Pay after: $%.2f\n", employeeAfter.CalculatePay())
}
