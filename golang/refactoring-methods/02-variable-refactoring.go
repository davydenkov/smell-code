package main

import (
	"fmt"
	"math"
)

// 3. Embedding a temporary variable (Inline Temp)
//
// BEFORE: Unnecessary temporary variable
type PriceCalculatorBefore struct {
	quantity  int
	itemPrice float64
}

func (pc PriceCalculatorBefore) GetPrice() float64 {
	basePrice := float64(pc.quantity) * pc.itemPrice
	if basePrice > 1000 {
		return basePrice * 0.95
	} else {
		return basePrice * 0.98
	}
}

// AFTER: Inline the temporary variable
type PriceCalculatorAfter struct {
	quantity  int
	itemPrice float64
}

func (pc PriceCalculatorAfter) GetPrice() float64 {
	if float64(pc.quantity)*pc.itemPrice > 1000 {
		return float64(pc.quantity) * pc.itemPrice * 0.95
	} else {
		return float64(pc.quantity) * pc.itemPrice * 0.98
	}
}

// 4. Replacing a temporary variable with a method call (Replace Temp with Query)
//
// BEFORE: Temporary variable used multiple times
type OrderBefore struct {
	quantity  int
	itemPrice float64
}

func (o OrderBefore) GetPrice() float64 {
	basePrice := float64(o.quantity) * o.itemPrice
	return basePrice - o.getDiscount(basePrice)
}

func (o OrderBefore) getDiscount(basePrice float64) float64 {
	return math.Max(0, basePrice-500) * 0.05
}

// AFTER: Replace temp with query
type OrderAfter struct {
	quantity  int
	itemPrice float64
}

func (o OrderAfter) GetPrice() float64 {
	return o.getBasePrice() - o.getDiscount()
}

func (o OrderAfter) getBasePrice() float64 {
	return float64(o.quantity) * o.itemPrice
}

func (o OrderAfter) getDiscount() float64 {
	return math.Max(0, o.getBasePrice()-500) * 0.05
}

// 5. Introduction of an explanatory variable (Introduce Explaining Variable)
//
// BEFORE: Complex expression hard to understand
type PerformanceCalculatorBefore struct {
	goals         int
	assists       int
	minutesPlayed int
}

func (pc PerformanceCalculatorBefore) GetPerformance() float64 {
	return float64(pc.goals*2) + float64(pc.assists)*1.5 + float64(pc.minutesPlayed)/60*0.1
}

// AFTER: Introduce explaining variables for clarity
type PerformanceCalculatorAfter struct {
	goals         int
	assists       int
	minutesPlayed int
}

func (pc PerformanceCalculatorAfter) GetPerformance() float64 {
	goalPoints := pc.goals * 2
	assistPoints := float64(pc.assists) * 1.5
	playingTimeBonus := float64(pc.minutesPlayed) / 60 * 0.1

	return float64(goalPoints) + assistPoints + playingTimeBonus
}

// 6. Splitting a Temporary Variable
//
// BEFORE: Same variable used for different purposes
type TemperatureMonitorBefore struct{}

func (tm TemperatureMonitorBefore) GetReading() map[string]float64 {
	temp := tm.getCurrentTemperature()

	// First use: get initial reading
	initialTemp := temp

	// Later: temp is reused for different calculation
	temp = temp + tm.getAdjustment()
	adjustedTemp := temp

	return map[string]float64{
		"initial":  initialTemp,
		"adjusted": adjustedTemp,
	}
}

// AFTER: Split the temporary variable
type TemperatureMonitorAfter struct{}

func (tm TemperatureMonitorAfter) GetReading() map[string]float64 {
	temp := tm.getCurrentTemperature()
	initialTemp := temp

	adjustedTemp := temp + tm.getAdjustment()

	return map[string]float64{
		"initial":  initialTemp,
		"adjusted": adjustedTemp,
	}
}

// Helper methods for temperature monitoring
func (tm TemperatureMonitorBefore) getCurrentTemperature() float64 {
	return 25.0 // Mock temperature
}

func (tm TemperatureMonitorBefore) getAdjustment() float64 {
	return 2.5 // Mock adjustment
}

func (tm TemperatureMonitorAfter) getCurrentTemperature() float64 {
	return 25.0 // Mock temperature
}

func (tm TemperatureMonitorAfter) getAdjustment() float64 {
	return 2.5 // Mock adjustment
}

// 7. Removing parameter Assignments (Remove Assignments to Parameters)
//
// BEFORE: Parameter is modified inside method
type DiscountCalculatorBefore struct{}

func (dc DiscountCalculatorBefore) ApplyDiscount(price float64) float64 {
	if price > 100 {
		price = price * 0.9 // Modifying parameter - not idiomatic in Go
	}
	return price
}

// AFTER: Use a local variable instead
type DiscountCalculatorAfter struct{}

func (dc DiscountCalculatorAfter) ApplyDiscount(price float64) float64 {
	result := price
	if price > 100 {
		result = price * 0.9
	}
	return result
}

// 8. Replacing a method with a method Object (Replace Method with Method Object)
//
// BEFORE: Method with many parameters and local variables
type AccountBefore struct{}

func (a AccountBefore) CalculateInterest(principal, rate, time, compoundingFrequency float64) float64 {
	amount := principal * math.Pow(1+(rate/compoundingFrequency), compoundingFrequency*time)
	interest := amount - principal
	return interest
}

// AFTER: Extract to a method object
type InterestCalculation struct {
	principal            float64
	rate                 float64
	time                 float64
	compoundingFrequency float64
}

func NewInterestCalculation(principal, rate, time, compoundingFrequency float64) *InterestCalculation {
	return &InterestCalculation{
		principal:            principal,
		rate:                 rate,
		time:                 time,
		compoundingFrequency: compoundingFrequency,
	}
}

func (ic InterestCalculation) Calculate() float64 {
	amount := ic.principal * math.Pow(1+(ic.rate/ic.compoundingFrequency),
		ic.compoundingFrequency*ic.time)
	return amount - ic.principal
}

type AccountAfter struct{}

func (a AccountAfter) CalculateInterest(principal, rate, time, compoundingFrequency float64) float64 {
	calculation := NewInterestCalculation(principal, rate, time, compoundingFrequency)
	return calculation.Calculate()
}

// Example usage for variable refactoring
func demonstrateVariableRefactoring() {
	fmt.Println("=== 3. Inline Temp Refactoring ===")

	calcBefore := PriceCalculatorBefore{quantity: 20, itemPrice: 60.0} // total = 1200
	calcAfter := PriceCalculatorAfter{quantity: 20, itemPrice: 60.0}

	fmt.Printf("Price before (with temp): $%.2f\n", calcBefore.GetPrice())
	fmt.Printf("Price after (inlined): $%.2f\n", calcAfter.GetPrice())

	fmt.Println("\n=== 4. Replace Temp with Query ===")

	orderBefore := OrderBefore{quantity: 15, itemPrice: 50.0}
	orderAfter := OrderAfter{quantity: 15, itemPrice: 50.0}

	fmt.Printf("Order price before: $%.2f\n", orderBefore.GetPrice())
	fmt.Printf("Order price after: $%.2f\n", orderAfter.GetPrice())

	fmt.Println("\n=== 5. Introduce Explaining Variable ===")

	perfBefore := PerformanceCalculatorBefore{goals: 10, assists: 5, minutesPlayed: 180}
	perfAfter := PerformanceCalculatorAfter{goals: 10, assists: 5, minutesPlayed: 180}

	fmt.Printf("Performance before: %.2f\n", perfBefore.GetPerformance())
	fmt.Printf("Performance after: %.2f\n", perfAfter.GetPerformance())

	fmt.Println("\n=== 6. Split Temporary Variable ===")

	monitorBefore := TemperatureMonitorBefore{}
	monitorAfter := TemperatureMonitorAfter{}

	readingBefore := monitorBefore.GetReading()
	readingAfter := monitorAfter.GetReading()

	fmt.Printf("Reading before - Initial: %.1f, Adjusted: %.1f\n",
		readingBefore["initial"], readingBefore["adjusted"])
	fmt.Printf("Reading after - Initial: %.1f, Adjusted: %.1f\n",
		readingAfter["initial"], readingAfter["adjusted"])

	fmt.Println("\n=== 7. Remove Assignments to Parameters ===")

	discountBefore := DiscountCalculatorBefore{}
	discountAfter := DiscountCalculatorAfter{}

	price := 150.0
	fmt.Printf("Discount before: $%.2f -> $%.2f\n", price, discountBefore.ApplyDiscount(price))
	fmt.Printf("Discount after: $%.2f -> $%.2f\n", price, discountAfter.ApplyDiscount(price))

	fmt.Println("\n=== 8. Replace Method with Method Object ===")

	accountBefore := AccountBefore{}
	accountAfter := AccountAfter{}

	principal, rate, time, freq := 1000.0, 0.05, 2.0, 12.0
	fmt.Printf("Interest before: $%.2f\n", accountBefore.CalculateInterest(principal, rate, time, freq))
	fmt.Printf("Interest after: $%.2f\n", accountAfter.CalculateInterest(principal, rate, time, freq))
}
