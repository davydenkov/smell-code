package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
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

func (a Address) GetStreet() string {
	return a.street
}

func (a Address) GetCity() string {
	return a.city
}

func (a Address) GetState() string {
	return a.state
}

func (a Address) GetZipCode() string {
	return a.zipCode
}

func (a Address) IsValid() bool {
	if a.street == "" || a.city == "" || a.state == "" || a.zipCode == "" {
		return false
	}

	if len(a.zipCode) != 5 {
		return false
	}

	return true
}

func (a Address) ToString() string {
	return fmt.Sprintf("%s, %s, %s %s", a.street, a.city, a.state, a.zipCode)
}

func (a Address) ToLabelFormat() string {
	return fmt.Sprintf("%s\n%s, %s %s", a.street, a.city, a.state, a.zipCode)
}

type Person struct {
	firstName   string
	lastName    string
	email       string
	phone       *string
	dateOfBirth *string
}

func NewPerson(firstName, lastName, email string, phone, dateOfBirth *string) *Person {
	return &Person{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		phone:       phone,
		dateOfBirth: dateOfBirth,
	}
}

func (p Person) GetFirstName() string {
	return p.firstName
}

func (p Person) GetLastName() string {
	return p.lastName
}

func (p Person) GetFullName() string {
	return p.firstName + " " + p.lastName
}

func (p Person) GetEmail() string {
	return p.email
}

func (p Person) GetPhone() *string {
	return p.phone
}

func (p Person) GetDateOfBirth() *string {
	return p.dateOfBirth
}

func (p Person) IsValid() bool {
	if p.firstName == "" || p.lastName == "" {
		return false
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(p.email) {
		return false
	}

	return true
}

type CustomerService struct{}

func (cs CustomerService) CreateCustomer(person *Person, address *Address) (map[string]string, error) {
	if !person.IsValid() {
		return nil, errors.New("invalid person data")
	}

	if !address.IsValid() {
		return nil, errors.New("invalid address data")
	}

	// Create customer record
	customerData := map[string]string{
		"first_name":    person.GetFirstName(),
		"last_name":     person.GetLastName(),
		"email":         person.GetEmail(),
		"phone":         *person.GetPhone(),
		"date_of_birth": *person.GetDateOfBirth(),
		"street":        address.GetStreet(),
		"city":          address.GetCity(),
		"state":         address.GetState(),
		"zip_code":      address.GetZipCode(),
	}

	// Save to database (simulated)
	return customerData, nil
}

func (cs CustomerService) UpdateCustomerAddress(customerId int, address *Address) (map[string]string, error) {
	if !address.IsValid() {
		return nil, errors.New("invalid address data")
	}

	// Update address
	addressData := map[string]string{
		"customer_id": strconv.Itoa(customerId),
		"street":      address.GetStreet(),
		"city":        address.GetCity(),
		"state":       address.GetState(),
		"zip_code":    address.GetZipCode(),
	}

	// Save to database (simulated)
	return addressData, nil
}

func (cs CustomerService) SendWelcomeEmail(person *Person, address *Address) map[string]string {
	message := fmt.Sprintf("Welcome %s!\n\n", person.GetFullName())
	message += fmt.Sprintf("Your address: %s\n", address.ToString())

	// Send email (simulated)
	return map[string]string{
		"to":      person.GetEmail(),
		"message": message,
	}
}

func (cs CustomerService) ValidateShippingAddress(address *Address) bool {
	return address.IsValid()
}

func (cs CustomerService) FormatAddressLabel(person *Person, address *Address) string {
	return person.GetFullName() + "\n" + address.ToLabelFormat()
}

func main() {
	cs := CustomerService{}

	// Create person and address objects
	phone := "555-1234"
	dob := "1990-01-01"
	person := NewPerson("John", "Doe", "john@example.com", &phone, &dob)
	address := NewAddress("123 Main St", "Anytown", "CA", "12345")

	// Example usage
	customer, err := cs.CreateCustomer(person, address)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("Created customer: %+v\n", customer)

	isValid := cs.ValidateShippingAddress(address)
	fmt.Printf("Address valid: %t\n", isValid)

	welcomeEmail := cs.SendWelcomeEmail(person, address)
	fmt.Printf("Welcome email: %+v\n", welcomeEmail)

	label := cs.FormatAddressLabel(person, address)
	fmt.Printf("Address label:\n%s\n", label)
}
