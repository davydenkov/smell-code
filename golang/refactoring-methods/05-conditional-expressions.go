package main

import (
	"errors"
	"fmt"
	"strings"
)

// 34. Decomposition of a conditional operator (Decompose Conditional)
//
// BEFORE: Complex conditional logic
type PaymentProcessorBefore struct{}

func (pp PaymentProcessorBefore) CalculateFee(amount float64, isInternational, isPremium bool) float64 {
	var fee float64
	if amount > 100 && isInternational && isPremium {
		fee = amount*0.05 + 10
	} else if amount > 100 && isInternational && !isPremium {
		fee = amount*0.05 + 15
	} else if amount <= 100 && isInternational {
		fee = amount*0.03 + 5
	} else {
		fee = amount * 0.02
	}
	return fee
}

// AFTER: Decompose conditional
type PaymentProcessorAfter struct{}

func (pp PaymentProcessorAfter) CalculateFee(amount float64, isInternational, isPremium bool) float64 {
	if pp.isHighValueInternationalPremium(amount, isInternational, isPremium) {
		return pp.calculateHighValueInternationalPremiumFee(amount)
	} else if pp.isHighValueInternationalStandard(amount, isInternational, isPremium) {
		return pp.calculateHighValueInternationalStandardFee(amount)
	} else if pp.isLowValueInternational(amount, isInternational) {
		return pp.calculateLowValueInternationalFee(amount)
	} else {
		return pp.calculateDomesticFee(amount)
	}
}

func (pp PaymentProcessorAfter) isHighValueInternationalPremium(amount float64, isInternational, isPremium bool) bool {
	return amount > 100 && isInternational && isPremium
}

func (pp PaymentProcessorAfter) isHighValueInternationalStandard(amount float64, isInternational, isPremium bool) bool {
	return amount > 100 && isInternational && !isPremium
}

func (pp PaymentProcessorAfter) isLowValueInternational(amount float64, isInternational bool) bool {
	return amount <= 100 && isInternational
}

func (pp PaymentProcessorAfter) calculateHighValueInternationalPremiumFee(amount float64) float64 {
	return amount*0.05 + 10
}

func (pp PaymentProcessorAfter) calculateHighValueInternationalStandardFee(amount float64) float64 {
	return amount*0.05 + 15
}

func (pp PaymentProcessorAfter) calculateLowValueInternationalFee(amount float64) float64 {
	return amount*0.03 + 5
}

func (pp PaymentProcessorAfter) calculateDomesticFee(amount float64) float64 {
	return amount * 0.02
}

// 35. Consolidation of a conditional expression (Consolidate Conditional Expression)
//
// BEFORE: Multiple conditionals with same result
type InsuranceCalculatorBefore struct{}

func (ic InsuranceCalculatorBefore) IsEligibleForDiscount(age int, isStudent, hasGoodRecord bool) bool {
	if age < 25 {
		return false
	}
	if isStudent {
		return true
	}
	if hasGoodRecord {
		return true
	}
	return false
}

// AFTER: Consolidate conditionals
type InsuranceCalculatorAfter struct{}

func (ic InsuranceCalculatorAfter) IsEligibleForDiscount(age int, isStudent, hasGoodRecord bool) bool {
	return age >= 25 && (isStudent || hasGoodRecord)
}

// 36. Consolidation of duplicate conditional fragments
//
// BEFORE: Duplicate code in conditional branches
type FileProcessorBefore struct{}

func (fp FileProcessorBefore) ProcessFile(filePath string) error {
	if fp.isValidFile(filePath) {
		content, err := fp.readFile(filePath)
		if err != nil {
			return err
		}
		fmt.Printf("Processing file: %s\n", filePath)
		return fp.processContent(content)
	} else {
		content, err := fp.readFile(filePath)
		if err != nil {
			return err
		}
		fmt.Printf("Processing invalid file: %s\n", filePath)
		return fp.processContent(content)
	}
}

// AFTER: Consolidate duplicate conditional fragments
type FileProcessorAfter struct{}

func (fp FileProcessorAfter) ProcessFile(filePath string) error {
	content, err := fp.readFile(filePath)
	if err != nil {
		return err
	}

	if fp.isValidFile(filePath) {
		fmt.Printf("Processing file: %s\n", filePath)
	} else {
		fmt.Printf("Processing invalid file: %s\n", filePath)
	}

	return fp.processContent(content)
}

// 37. Remove Control Flag
//
// BEFORE: Control flag variable
type SearchProcessorBefore struct{}

func (sp SearchProcessorBefore) FindUser(users []map[string]interface{}, targetName string) (map[string]interface{}, error) {
	found := false
	var result map[string]interface{}

	for _, user := range users {
		if name, ok := user["name"].(string); ok && name == targetName {
			result = user
			found = true
			break // Control flag used to break
		}
	}

	if !found {
		return nil, errors.New("user not found")
	}
	return result, nil
}

// AFTER: Remove control flag
type SearchProcessorAfter struct{}

func (sp SearchProcessorAfter) FindUser(users []map[string]interface{}, targetName string) (map[string]interface{}, error) {
	for _, user := range users {
		if name, ok := user["name"].(string); ok && name == targetName {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// 38. Replacing Nested Conditional statements with a boundary operator (Replace Nested Conditional with Guard Clauses)
//
// BEFORE: Nested conditionals
type AccountValidatorBefore struct{}

func (av AccountValidatorBefore) ValidateTransaction(account map[string]interface{}, amount float64) error {
	if balance, ok := account["balance"].(float64); ok {
		if balance >= amount {
			if amount > 0 {
				if status, ok := account["status"].(string); ok {
					if status == "active" {
						return nil // Valid transaction
					} else {
						return errors.New("account is not active")
					}
				}
			} else {
				return errors.New("amount must be positive")
			}
		} else {
			return errors.New("insufficient funds")
		}
	}
	return errors.New("invalid account balance")
}

// AFTER: Use guard clauses
type AccountValidatorAfter struct{}

func (av AccountValidatorAfter) ValidateTransaction(account map[string]interface{}, amount float64) error {
	balance, ok := account["balance"].(float64)
	if !ok {
		return errors.New("invalid account balance")
	}

	if balance < amount {
		return errors.New("insufficient funds")
	}

	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	status, ok := account["status"].(string)
	if !ok || status != "active" {
		return errors.New("account is not active")
	}

	return nil // Valid transaction
}

// 39. Replacing a conditional operator with polymorphism (Replace Conditional with Polymorphism)
//
// BEFORE: Conditional based on type
type ShapeBefore struct {
	shapeType string
	width     float64
	height    float64
}

func (s ShapeBefore) GetArea() float64 {
	if s.shapeType == "rectangle" {
		return s.width * s.height
	} else if s.shapeType == "triangle" {
		return s.width * s.height * 0.5
	} else if s.shapeType == "circle" {
		return 3.14159 * s.width * s.width // width used as radius
	}
	return 0
}

// AFTER: Use polymorphism
type Shape interface {
	GetArea() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) GetArea() float64 {
	return r.width * r.height
}

type Triangle struct {
	width, height float64
}

func (t Triangle) GetArea() float64 {
	return t.width * t.height * 0.5
}

type Circle struct {
	radius float64
}

func (c Circle) GetArea() float64 {
	return 3.14159 * c.radius * c.radius
}

// 40. Introduction of the object (Introduce Null Object)
//
// BEFORE: Checking for nil everywhere
type LoggerBefore struct{}

func (l LoggerBefore) ProcessData(data string, logger interface{}) {
	if logger != nil {
		// Type assertion and call would be done here
		fmt.Printf("Logging: %s\n", data)
	}
	// Process data
	fmt.Printf("Processing: %s\n", data)
}

// AFTER: Introduce null object
type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Printf("Console: %s\n", message)
}

type NullLogger struct{}

func (nl NullLogger) Log(message string) {
	// Do nothing
}

type DataProcessorAfter struct{}

func (dp DataProcessorAfter) ProcessData(data string, logger Logger) {
	logger.Log(fmt.Sprintf("Processing: %s", data))
	// Process data
	fmt.Printf("Processing: %s\n", data)
}

// 41. Introduction of the statement (Introduce Assertion)
//
// BEFORE: No validation
type CalculatorBefore struct{}

func (c CalculatorBefore) Divide(a, b float64) float64 {
	return a / b // Could divide by zero
}

// AFTER: Introduce assertion
type CalculatorAfter struct{}

func (c CalculatorAfter) Divide(a, b float64) float64 {
	c.assertDenominatorNotZero(b)
	return a / b
}

func (c CalculatorAfter) assertDenominatorNotZero(b float64) {
	if b == 0 {
		panic("denominator cannot be zero")
	}
}

// Helper methods for the examples
func (fp FileProcessorBefore) isValidFile(filePath string) bool {
	return strings.HasSuffix(filePath, ".txt")
}

func (fp FileProcessorBefore) readFile(filePath string) (string, error) {
	return "file content", nil
}

func (fp FileProcessorBefore) processContent(content string) error {
	fmt.Printf("Content processed: %s\n", content)
	return nil
}

func (fp FileProcessorAfter) isValidFile(filePath string) bool {
	return strings.HasSuffix(filePath, ".txt")
}

func (fp FileProcessorAfter) readFile(filePath string) (string, error) {
	return "file content", nil
}

func (fp FileProcessorAfter) processContent(content string) error {
	fmt.Printf("Content processed: %s\n", content)
	return nil
}

// Example usage for conditional expressions
func demonstrateConditionalExpressions() {
	fmt.Println("=== 34. Decompose Conditional ===")

	processorBefore := PaymentProcessorBefore{}
	processorAfter := PaymentProcessorAfter{}

	amount := 150.0
	fmt.Printf("Fee before: $%.2f\n", processorBefore.CalculateFee(amount, true, true))
	fmt.Printf("Fee after: $%.2f\n", processorAfter.CalculateFee(amount, true, true))

	fmt.Println("\n=== 35. Consolidate Conditional Expression ===")

	insuranceBefore := InsuranceCalculatorBefore{}
	insuranceAfter := InsuranceCalculatorAfter{}

	fmt.Printf("Discount before: %v\n", insuranceBefore.IsEligibleForDiscount(30, true, false))
	fmt.Printf("Discount after: %v\n", insuranceAfter.IsEligibleForDiscount(30, true, false))

	fmt.Println("\n=== 37. Remove Control Flag ===")

	users := []map[string]interface{}{
		{"name": "Alice", "age": 25},
		{"name": "Bob", "age": 30},
		{"name": "Charlie", "age": 35},
	}

	searchBefore := SearchProcessorBefore{}
	searchAfter := SearchProcessorAfter{}

	if user, err := searchBefore.FindUser(users, "Bob"); err == nil {
		fmt.Printf("Found user before: %s\n", user["name"])
	}

	if user, err := searchAfter.FindUser(users, "Bob"); err == nil {
		fmt.Printf("Found user after: %s\n", user["name"])
	}

	fmt.Println("\n=== 38. Replace Nested Conditional with Guard Clauses ===")

	account := map[string]interface{}{
		"balance": 1000.0,
		"status":  "active",
	}

	validatorBefore := AccountValidatorBefore{}
	validatorAfter := AccountValidatorAfter{}

	if err := validatorBefore.ValidateTransaction(account, 100.0); err == nil {
		fmt.Println("Transaction valid (before)")
	}

	if err := validatorAfter.ValidateTransaction(account, 100.0); err == nil {
		fmt.Println("Transaction valid (after)")
	}

	fmt.Println("\n=== 39. Replace Conditional with Polymorphism ===")

	// Before - using conditional
	shapeBefore := ShapeBefore{shapeType: "rectangle", width: 10, height: 5}
	fmt.Printf("Area before: %.2f\n", shapeBefore.GetArea())

	// After - using polymorphism
	var shape Shape = Rectangle{width: 10, height: 5}
	fmt.Printf("Area after: %.2f\n", shape.GetArea())

	fmt.Println("\n=== 40. Introduce Null Object ===")

	processor := DataProcessorAfter{}

	// With real logger
	consoleLogger := ConsoleLogger{}
	processor.ProcessData("test data", consoleLogger)

	// With null logger
	nullLogger := NullLogger{}
	processor.ProcessData("test data", nullLogger)

	fmt.Println("\n=== 41. Introduce Assertion ===")

	calc := CalculatorAfter{}
	result := calc.Divide(10, 2)
	fmt.Printf("Division result: %.2f\n", result)
}
