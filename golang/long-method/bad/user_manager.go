package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
	"time"
)

type UserManager struct {
	db Database
}

func NewUserManager(db Database) *UserManager {
	return &UserManager{db: db}
}

func (um *UserManager) RegisterUser(userData map[string]string) (int64, error) {
	// Validate input data
	if userData["email"] == "" {
		return 0, fmt.Errorf("email is required")
	}
	if userData["password"] == "" {
		return 0, fmt.Errorf("password is required")
	}
	if len(userData["password"]) < 8 {
		return 0, fmt.Errorf("password must be at least 8 characters")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(userData["email"]) {
		return 0, fmt.Errorf("invalid email format")
	}

	// Check if user already exists
	rows, err := um.db.Query("SELECT id FROM users WHERE email = ?", userData["email"])
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		return 0, fmt.Errorf("user already exists")
	}

	// Hash password
	hashedPassword := um.hashPassword(userData["password"])

	// Generate verification token
	verificationToken := um.generateToken()

	// Insert user into database
	result, err := um.db.Exec(`
		INSERT INTO users (email, password, first_name, last_name, verification_token, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`, userData["email"], hashedPassword, userData["firstName"], userData["lastName"], verificationToken, time.Now())
	if err != nil {
		return 0, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Create user profile
	_, err = um.db.Exec(`
		INSERT INTO user_profiles (user_id, phone, address, city, state, zip_code)
		VALUES (?, ?, ?, ?, ?, ?)
	`, userId, userData["phone"], userData["address"], userData["city"], userData["state"], userData["zipCode"])
	if err != nil {
		return 0, err
	}

	// Send verification email
	subject := "Please verify your email address"
	message := fmt.Sprintf("Hello %s,\n\n", userData["firstName"])
	message += "Thank you for registering. Please click the link below to verify your email:\n\n"
	message += fmt.Sprintf("http://example.com/verify?token=%s\n\n", verificationToken)
	message += "Best regards,\nThe Team"

	fmt.Printf("Sending email to %s: %s\n", userData["email"], subject)
	// In real implementation, would send actual email

	// Log registration
	logMessage := fmt.Sprintf("User registered: %s at %s", userData["email"], time.Now().Format("2006-01-02 15:04:05"))
	log.Println(logMessage)

	// Create default settings
	_, err = um.db.Exec("INSERT INTO user_settings (user_id, theme, notifications_enabled) VALUES (?, 'light', true)", userId)
	if err != nil {
		return 0, err
	}

	// Send welcome notification
	_, err = um.db.Exec("INSERT INTO notifications (user_id, type, message, created_at) VALUES (?, 'welcome', 'Welcome to our platform!', ?)", userId, time.Now())
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (um *UserManager) hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func (um *UserManager) generateToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func main() {
	// This would normally use a real database connection
	var db Database

	um := NewUserManager(db)

	userData := map[string]string{
		"email":     "john@example.com",
		"password":  "password123",
		"firstName": "John",
		"lastName":  "Doe",
		"phone":     "555-1234",
		"address":   "123 Main St",
		"city":      "Anytown",
		"state":     "CA",
		"zipCode":   "12345",
	}

	userId, err := um.RegisterUser(userData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User registered with ID: %d\n", userId)
}
