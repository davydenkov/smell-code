package main

import (
	"fmt"
	"math"
)

// 57. Lifting the field (Pull Up Field)
//
// BEFORE: Duplicate fields in related types (Go equivalent using composition)
type EmployeePullBefore struct {
	name string
}

type ManagerPullBefore struct {
	employee EmployeePullBefore // Embedded but field is duplicated in concept
	budget   float64
}

type EngineerPullBefore struct {
	employee EmployeePullBefore // Embedded but field is duplicated in concept
	skills   []string
}

// AFTER: Pull up field to common embedded struct
type EmployeePullAfter struct {
	name string
}

type ManagerPullAfter struct {
	EmployeePullAfter
	budget float64
}

type EngineerPullAfter struct {
	EmployeePullAfter
	skills []string
}

// 58. Lifting the method (Pull Up Method)
//
// BEFORE: Duplicate methods in related types
type ShapeMethodBefore interface {
	Area() float64
}

type CircleMethodBefore struct {
	radius float64
}

func (c CircleMethodBefore) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c CircleMethodBefore) Circumference() float64 {
	return 2 * math.Pi * c.radius
}

type SquareMethodBefore struct {
	side float64
}

func (s SquareMethodBefore) Area() float64 {
	return s.side * s.side
}

func (s SquareMethodBefore) Circumference() float64 { // Duplicate logic concept
	return 4 * s.side
}

// AFTER: Pull up method to interface (Go equivalent)
type ShapeMethodAfter interface {
	Area() float64
	Circumference() float64
}

type CircleMethodAfter struct {
	radius float64
}

func (c CircleMethodAfter) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c CircleMethodAfter) Circumference() float64 {
	return 2 * math.Pi * c.radius
}

type SquareMethodAfter struct {
	side float64
}

func (s SquareMethodAfter) Area() float64 {
	return s.side * s.side
}

func (s SquareMethodAfter) Circumference() float64 {
	return 4 * s.side
}

// 59. Lifting the constructor Body (Pull Up Constructor Body)
//
// BEFORE: Duplicate constructor logic
type VehicleBefore struct {
	make  string
	model string
	year  int
}

func NewVehicleBefore(make, model string, year int) *VehicleBefore {
	return &VehicleBefore{
		make:  make,
		model: model,
		year:  year,
	}
}

type CarBefore struct {
	VehicleBefore
	doors int
}

func NewCarBefore(make, model string, year, doors int) *CarBefore {
	car := &CarBefore{
		doors: doors,
	}
	car.make = make   // Duplicate initialization
	car.model = model // Duplicate initialization
	car.year = year   // Duplicate initialization
	return car
}

type TruckBefore struct {
	VehicleBefore
	payload float64
}

func NewTruckBefore(make, model string, year int, payload float64) *TruckBefore {
	truck := &TruckBefore{
		payload: payload,
	}
	truck.make = make   // Duplicate initialization
	truck.model = model // Duplicate initialization
	truck.year = year   // Duplicate initialization
	return truck
}

// AFTER: Pull up constructor body
type VehicleAfter struct {
	make  string
	model string
	year  int
}

func NewVehicleAfter(make, model string, year int) *VehicleAfter {
	return &VehicleAfter{
		make:  make,
		model: model,
		year:  year,
	}
}

type CarAfter struct {
	*VehicleAfter
	doors int
}

func NewCarAfter(make, model string, year, doors int) *CarAfter {
	return &CarAfter{
		VehicleAfter: NewVehicleAfter(make, model, year),
		doors:        doors,
	}
}

type TruckAfter struct {
	*VehicleAfter
	payload float64
}

func NewTruckAfter(make, model string, year int, payload float64) *TruckAfter {
	return &TruckAfter{
		VehicleAfter: NewVehicleAfter(make, model, year),
		payload:      payload,
	}
}

// 60. Method Descent (Push Down Method)
//
// BEFORE: Method in superclass not used by all subclasses
type EmployeePushBefore struct {
	name string
}

func (e EmployeePushBefore) GetName() string {
	return e.name
}

func (e EmployeePushBefore) CalculateBonus() float64 {
	return 1000.0 // Generic bonus - not appropriate for all
}

type ManagerPushBefore struct {
	EmployeePushBefore
	department string
}

type EngineerPushBefore struct {
	EmployeePushBefore
	skills []string
}

// AFTER: Push down method to specific subclass
type EmployeePushAfter struct {
	name string
}

func (e EmployeePushAfter) GetName() string {
	return e.name
}

type ManagerPushAfter struct {
	EmployeePushAfter
	department string
}

func (m ManagerPushAfter) CalculateBonus() float64 {
	return 2000.0 // Manager-specific bonus
}

type EngineerPushAfter struct {
	EmployeePushAfter
	skills []string
}

func (e EngineerPushAfter) CalculateBonus() float64 {
	return 1500.0 // Engineer-specific bonus
}

// 61. Field Descent (Push Down Field)
//
// BEFORE: Field in superclass not used by all subclasses
type ProductPushBefore struct {
	name  string
	price float64
	category string // Not used by all subclasses
}

type BookPushBefore struct {
	ProductPushBefore
	author string
}

type ElectronicsPushBefore struct {
	ProductPushBefore
	warranty int
}

// AFTER: Push down field to specific subclass
type ProductPushAfter struct {
	name  string
	price float64
}

type BookPushAfter struct {
	ProductPushAfter
	author   string
	category string // Only books have categories
}

type ElectronicsPushAfter struct {
	ProductPushAfter
	warranty int
}

// 62. Subclass extraction (Extract Subclass)
//
// BEFORE: Class with conditional behavior
type EmployeeExtractBefore struct {
	name     string
	isManager bool
	bonus    float64
}

func (e EmployeeExtractBefore) CalculateSalary(baseSalary float64) float64 {
	salary := baseSalary
	if e.isManager {
		salary += e.bonus
	}
	return salary
}

// AFTER: Extract subclass
type EmployeeExtractAfter struct {
	name string
}

func (e EmployeeExtractAfter) CalculateSalary(baseSalary float64) float64 {
	return baseSalary
}

type ManagerExtractAfter struct {
	EmployeeExtractAfter
	bonus float64
}

func (m ManagerExtractAfter) CalculateSalary(baseSalary float64) float64 {
	return baseSalary + m.bonus
}

// 63. Allocation of the parent class (Extract Superclass)
//
// BEFORE: Two classes with common behavior
type SavingsAccountBefore struct {
	balance float64
}

func (s SavingsAccountBefore) GetBalance() float64 {
	return s.balance
}

func (s *SavingsAccountBefore) Deposit(amount float64) {
	s.balance += amount
}

type CheckingAccountBefore struct {
	balance float64
}

func (c CheckingAccountBefore) GetBalance() float64 {
	return c.balance
}

func (c *CheckingAccountBefore) Deposit(amount float64) {
	c.balance += amount
}

// AFTER: Extract superclass
type AccountAfter struct {
	balance float64
}

func (a AccountAfter) GetBalance() float64 {
	return a.balance
}

func (a *AccountAfter) Deposit(amount float64) {
	a.balance += amount
}

type SavingsAccountAfter struct {
	AccountAfter
	interestRate float64
}

type CheckingAccountAfter struct {
	AccountAfter
	overdraftLimit float64
}

// 64. Interface extraction (Extract Interface)
//
// BEFORE: Class with multiple responsibilities
type DataProcessorBefore struct{}

func (dp DataProcessorBefore) ReadData() string {
	return "data from source"
}

func (dp DataProcessorBefore) ProcessData(data string) string {
	return "processed: " + data
}

func (dp DataProcessorBefore) SaveData(data string) error {
	fmt.Printf("Saving: %s\n", data)
	return nil
}

func (dp DataProcessorBefore) ValidateData(data string) bool {
	return len(data) > 0
}

// AFTER: Extract interface
type DataReader interface {
	ReadData() string
}

type DataProcessor interface {
	ProcessData(data string) string
}

type DataSaver interface {
	SaveData(data string) error
}

type DataValidator interface {
	ValidateData(data string) bool
}

type DataProcessorAfter struct{}

func (dp DataProcessorAfter) ReadData() string {
	return "data from source"
}

func (dp DataProcessorAfter) ProcessData(data string) string {
	return "processed: " + data
}

func (dp DataProcessorAfter) SaveData(data string) error {
	fmt.Printf("Saving: %s\n", data)
	return nil
}

func (dp DataProcessorAfter) ValidateData(data string) bool {
	return len(data) > 0
}

// 65. Collapse Hierarchy
//
// BEFORE: Unnecessary inheritance hierarchy
type VehicleHierarchyBefore struct {
	make  string
	model string
}

type CarHierarchyBefore struct {
	VehicleHierarchyBefore
	doors int
}

type SedanHierarchyBefore struct {
	CarHierarchyBefore
	trunkSize float64
}

// AFTER: Collapse hierarchy
type VehicleHierarchyAfter struct {
	make      string
	model     string
	doors     int
	trunkSize float64
}

// Example usage for generalization problems
func demonstrateGeneralizationProblems() {
	fmt.Println("=== 57. Pull Up Field ===")

	managerBefore := ManagerPullBefore{employee: EmployeePullBefore{name: "John"}}
	engineerBefore := EngineerPullBefore{employee: EmployeePullBefore{name: "Jane"}}

	managerAfter := ManagerPullAfter{EmployeePullAfter: EmployeePullAfter{name: "John"}}
	engineerAfter := EngineerPullAfter{EmployeePullAfter: EmployeePullAfter{name: "Jane"}}

	fmt.Printf("Manager name before: %s\n", managerBefore.employee.name)
	fmt.Printf("Engineer name before: %s\n", engineerBefore.employee.name)
	fmt.Printf("Manager name after: %s\n", managerAfter.name)
	fmt.Printf("Engineer name after: %s\n", engineerAfter.name)

	fmt.Println("\n=== 58. Pull Up Method ===")

	var shapesBefore []ShapeMethodBefore = []ShapeMethodBefore{
		CircleMethodBefore{radius: 5},
		SquareMethodBefore{side: 4},
	}

	for _, shape := range shapesBefore {
		fmt.Printf("Area: %.2f, Circumference: %.2f\n", shape.Area(), shape.(interface{ Circumference() float64 }).Circumference())
	}

	var shapesAfter []ShapeMethodAfter = []ShapeMethodAfter{
		CircleMethodAfter{radius: 5},
		SquareMethodAfter{side: 4},
	}

	for _, shape := range shapesAfter {
		fmt.Printf("Area: %.2f, Circumference: %.2f\n", shape.Area(), shape.Circumference())
	}

	fmt.Println("\n=== 59. Pull Up Constructor Body ===")

	carBefore := NewCarBefore("Toyota", "Camry", 2020, 4)
	truckBefore := NewTruckBefore("Ford", "F150", 2020, 2000)

	carAfter := NewCarAfter("Toyota", "Camry", 2020, 4)
	truckAfter := NewTruckAfter("Ford", "F150", 2020, 2000)

	fmt.Printf("Car before: %s %s %d, doors: %d\n", carBefore.make, carBefore.model, carBefore.year, carBefore.doors)
	fmt.Printf("Truck before: %s %s %d, payload: %.0f\n", truckBefore.make, truckBefore.model, truckBefore.year, truckBefore.payload)
	fmt.Printf("Car after: %s %s %d, doors: %d\n", carAfter.make, carAfter.model, carAfter.year, carAfter.doors)
	fmt.Printf("Truck after: %s %s %d, payload: %.0f\n", truckAfter.make, truckAfter.model, truckAfter.year, truckAfter.payload)

	fmt.Println("\n=== 63. Extract Superclass ===")

	savingsBefore := SavingsAccountBefore{balance: 1000}
	checkingBefore := CheckingAccountBefore{balance: 500}

	savingsAfter := SavingsAccountAfter{AccountAfter: AccountAfter{balance: 1000}}
	checkingAfter := CheckingAccountAfter{AccountAfter: AccountAfter{balance: 500}}

	fmt.Printf("Savings balance before: $%.2f\n", savingsBefore.GetBalance())
	fmt.Printf("Checking balance before: $%.2f\n", checkingBefore.GetBalance())
	fmt.Printf("Savings balance after: $%.2f\n", savingsAfter.GetBalance())
	fmt.Printf("Checking balance after: $%.2f\n", checkingAfter.GetBalance())

	fmt.Println("\n=== 64. Extract Interface ===")

	processor := DataProcessorAfter{}

	// Using specific interfaces
	var reader DataReader = processor
	var dataProcessor DataProcessor = processor
	var saver DataSaver = processor
	var validator DataValidator = processor

	data := reader.ReadData()
	processed := dataProcessor.ProcessData(data)
	isValid := validator.ValidateData(processed)
	if isValid {
		saver.SaveData(processed)
	}
}
