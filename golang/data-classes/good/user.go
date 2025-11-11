package main

import (
	"errors"
	"fmt"
	"regexp"
)

type EmailValidator struct{}

func NewEmailValidator() *EmailValidator {
	return &EmailValidator{}
}

func (ev EmailValidator) IsValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

type User struct {
	id             int
	name           string
	email          string
	age            int
	emailValidator *EmailValidator
}

func NewUser(id int, name string, email string, age int, emailValidator *EmailValidator) (*User, error) {
	if emailValidator == nil {
		emailValidator = NewEmailValidator()
	}

	user := &User{
		id:             id,
		name:           name,
		age:            age,
		emailValidator: emailValidator,
	}

	err := user.SetEmail(email) // Use setter for validation
	if err != nil {
		return nil, err
	}

	err = user.SetAge(age)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetEmail(email string) error {
	if !u.emailValidator.IsValid(email) {
		return errors.New("invalid email address")
	}
	u.email = email
	return nil
}

func (u *User) GetAge() int {
	return u.age
}

func (u *User) SetAge(age int) error {
	if age < 0 || age > 150 {
		return errors.New("age must be between 0 and 150")
	}
	u.age = age
	return nil
}

func (u *User) GetDisplayName() string {
	return fmt.Sprintf("%s (%d years old)", u.name, u.age)
}

func (u *User) CanVote() bool {
	return u.age >= 18
}

func (u *User) IsAdult() bool {
	return u.age >= 18
}

func (u *User) GetAgeCategory() string {
	if u.age < 13 {
		return "child"
	}
	if u.age < 20 {
		return "teenager"
	}
	if u.age < 65 {
		return "adult"
	}
	return "senior"
}

func main() {
	user, err := NewUser(1, "John Doe", "john@example.com", 30, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User ID: %d\n", user.GetId())
	fmt.Printf("User Name: %s\n", user.GetName())
	fmt.Printf("User Email: %s\n", user.GetEmail())
	fmt.Printf("User Age: %d\n", user.GetAge())
	fmt.Printf("Display Name: %s\n", user.GetDisplayName())
	fmt.Printf("Can Vote: %t\n", user.CanVote())
	fmt.Printf("Is Adult: %t\n", user.IsAdult())
	fmt.Printf("Age Category: %s\n", user.GetAgeCategory())

	// Test validation
	err = user.SetEmail("invalid-email")
	if err != nil {
		fmt.Printf("Email validation error: %s\n", err.Error())
	}

	err = user.SetAge(-5)
	if err != nil {
		fmt.Printf("Age validation error: %s\n", err.Error())
	}
}
