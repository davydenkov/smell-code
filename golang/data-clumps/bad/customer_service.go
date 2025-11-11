package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type CustomerService struct{}

func (cs CustomerService) CreateCustomer(
	firstName, lastName, email,
	street, city, state, zipCode,
	phone, dateOfBirth string,
) (map[string]string, error) {
	// Validate data
	if firstName == "" || lastName == "" {
		return nil, errors.New("name is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return nil, errors.New("invalid email")
	}

	// Create customer record
	customerData := map[string]string{
		"first_name":   firstName,
		"last_name":    lastName,
		"email":        email,
		"street":       street,
		"city":         city,
		"state":        state,
		"zip_code":     zipCode,
		"phone":        phone,
		"date_of_birth": dateOfBirth,
	}

	// Save to database (simulated)
	return customerData, nil
}

func (cs CustomerService) UpdateCustomerAddress(
	customerId int,
	street, city, state, zipCode string,
) map[string]string {
	// Update address
	addressData := map[string]string{
		"customer_id": strconv.Itoa(customerId),
		"street":      street,
		"city":        city,
		"state":       state,
		"zip_code":    zipCode,
	}

	// Save to database (simulated)
	return addressData
}

func (cs CustomerService) SendWelcomeEmail(
	firstName, lastName, email,
	street, city, state, zipCode string,
) map[string]string {
	fullName := firstName + " " + lastName
	fullAddress := street + ", " + city + ", " + state + " " + zipCode

	message := fmt.Sprintf("Welcome %s!\n\n", fullName)
	message += fmt.Sprintf("Your address: %s\n", fullAddress)

	// Send email (simulated)
	return map[string]string{
		"to":      email,
		"message": message,
	}
}

func (cs CustomerService) ValidateShippingAddress(street, city, state, zipCode string) bool {
	if street == "" || city == "" || state == "" || zipCode == "" {
		return false
	}

	// Additional validation logic
	if len(zipCode) != 5 {
		return false
	}

	return true
}

func (cs CustomerService) FormatAddressLabel(
	firstName, lastName,
	street, city, state, zipCode string,
) string {
	fullName := firstName + " " + lastName
	fullAddress := street + "\n" + city + ", " + state + " " + zipCode

	return fullName + "\n" + fullAddress
}

func main() {
	cs := CustomerService{}

	// Example usage
	customer, err := cs.CreateCustomer(
		"John", "Doe", "john@example.com",
		"123 Main St", "Anytown", "CA", "12345",
		"555-1234", "1990-01-01",
	)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("Created customer: %+v\n", customer)

	isValid := cs.ValidateShippingAddress("123 Main St", "Anytown", "CA", "12345")
	fmt.Printf("Address valid: %t\n", isValid)
}
