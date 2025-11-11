package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

type UserService struct {
	db            Database
	emailConfig   map[string]string
	paymentConfig map[string]string

	// User properties
	userId      int
	userEmail   string
	userName    string
	userBalance float64

	// Email properties
	smtpHost     string
	smtpPort     string
	smtpUsername string
	smtpPassword string

	// Payment properties
	stripeSecretKey string
	paypalClientId  string
	paypalSecret    string
}

func NewUserService(db Database, emailConfig, paymentConfig map[string]string) *UserService {
	return &UserService{
		db:            db,
		emailConfig:   emailConfig,
		paymentConfig: paymentConfig,
		smtpHost:      emailConfig["host"],
		smtpPort:      emailConfig["port"],
		smtpUsername:  emailConfig["username"],
		smtpPassword:  emailConfig["password"],
		stripeSecretKey: paymentConfig["stripe_secret"],
		paypalClientId:  paymentConfig["paypal_client_id"],
		paypalSecret:    paymentConfig["paypal_secret"],
	}
}

// User management methods
func (us *UserService) CreateUser(email, name, password string) (int64, error) {
	hashedPassword := us.hashPassword(password)
	result, err := us.db.Exec("INSERT INTO users (email, name, password) VALUES (?, ?, ?)",
		email, name, hashedPassword)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (us *UserService) AuthenticateUser(email, password string) (int, error) {
	rows, err := us.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		var id int
		var hashedPassword string
		rows.Scan(&id, &hashedPassword)

		if us.verifyPassword(password, hashedPassword) {
			us.userId = id
			us.userEmail = email
			return id, nil
		}
	}
	return 0, fmt.Errorf("invalid credentials")
}

func (us *UserService) UpdateUserProfile(userId int, name, email string) error {
	_, err := us.db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", name, email, userId)
	return err
}

func (us *UserService) GetUserBalance(userId int) float64 {
	rows, err := us.db.Query("SELECT balance FROM users WHERE id = ?", userId)
	if err != nil {
		log.Printf("Error getting balance: %v", err)
		return 0
	}
	defer rows.Close()

	if rows.Next() {
		var balance *float64
		rows.Scan(&balance)
		if balance != nil {
			return *balance
		}
	}
	return 0
}

// Email methods
func (us *UserService) SendWelcomeEmail(email, name string) {
	subject := "Welcome to our platform!"
	message := fmt.Sprintf("Hello %s,\n\nWelcome to our platform!", name)

	fmt.Printf("Sending email from %s to %s: %s\n", us.smtpUsername, email, subject)
	// In real implementation, would send actual email
}

func (us *UserService) SendPasswordResetEmail(email, resetToken string) {
	subject := "Password Reset"
	message := fmt.Sprintf("Click here to reset your password: http://example.com/reset?token=%s", resetToken)

	fmt.Printf("Sending email from %s to %s: %s\n", us.smtpUsername, email, subject)
	// In real implementation, would send actual email
}

func (us *UserService) SendNotificationEmail(email, subject, message string) {
	fmt.Printf("Sending email from %s to %s: %s\n", us.smtpUsername, email, subject)
	// In real implementation, would send actual email
}

// Payment methods
func (us *UserService) ProcessStripePayment(amount float64, token string) map[string]interface{} {
	if token != "" && amount > 0 {
		return map[string]interface{}{
			"success":       true,
			"transaction_id": "stripe_" + us.generateToken(),
		}
	}
	return map[string]interface{}{
		"success": false,
		"error":   "Invalid payment data",
	}
}

func (us *UserService) ProcessPayPalPayment(amount float64, paypalToken string) map[string]interface{} {
	if paypalToken != "" && amount > 0 {
		return map[string]interface{}{
			"success":       true,
			"transaction_id": "paypal_" + us.generateToken(),
		}
	}
	return map[string]interface{}{
		"success": false,
		"error":   "Invalid payment data",
	}
}

func (us *UserService) RefundPayment(transactionId string, amount float64) map[string]interface{} {
	if len(transactionId) > 7 && transactionId[:7] == "stripe_" {
		return map[string]interface{}{
			"success":  true,
			"refund_id": "refund_" + us.generateToken(),
		}
	} else if len(transactionId) > 7 && transactionId[:7] == "paypal_" {
		return map[string]interface{}{
			"success":  true,
			"refund_id": "refund_" + us.generateToken(),
		}
	}
	return map[string]interface{}{
		"success": false,
		"error":   "Unknown transaction type",
	}
}

// Reporting methods
func (us *UserService) GenerateUserReport(userId int) map[string]interface{} {
	rows, err := us.db.Query(`
		SELECT u.name, u.email, u.created_at,
			   COUNT(o.id) as order_count,
			   SUM(o.total) as total_spent
		FROM users u
		LEFT JOIN orders o ON u.id = o.user_id
		WHERE u.id = ?
		GROUP BY u.id
	`, userId)

	if err != nil {
		log.Printf("Error generating user report: %v", err)
		return nil
	}
	defer rows.Close()

	if rows.Next() {
		var name, email string
		var createdAt string
		var orderCount *int
		var totalSpent *float64

		rows.Scan(&name, &email, &createdAt, &orderCount, &totalSpent)

		return map[string]interface{}{
			"name":        name,
			"email":       email,
			"created_at":  createdAt,
			"order_count": orderCount,
			"total_spent": totalSpent,
		}
	}
	return nil
}

func (us *UserService) GenerateSalesReport(startDate, endDate string) []map[string]interface{} {
	rows, err := us.db.Query(`
		SELECT DATE(created_at) as date,
			   COUNT(*) as order_count,
			   SUM(total) as total_sales
		FROM orders
		WHERE created_at BETWEEN ? AND ?
		GROUP BY DATE(created_at)
	`, startDate, endDate)

	if err != nil {
		log.Printf("Error generating sales report: %v", err)
		return nil
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var date string
		var orderCount int
		var totalSales *float64

		rows.Scan(&date, &orderCount, &totalSales)

		results = append(results, map[string]interface{}{
			"date":        date,
			"order_count": orderCount,
			"total_sales": totalSales,
		})
	}
	return results
}

// Utility methods
func (us *UserService) ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func (us *UserService) GenerateToken() string {
	return us.generateToken()
}

func (us *UserService) hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func (us *UserService) verifyPassword(password, hash string) bool {
	passwordHash := us.hashPassword(password)
	return passwordHash == hash
}

func (us *UserService) generateToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (us *UserService) LogActivity(userId int, action string) error {
	_, err := us.db.Exec("INSERT INTO activity_log (user_id, action, created_at) VALUES (?, ?, ?)",
		userId, action, time.Now())
	return err
}

func main() {
	// This would normally use a real database connection
	var db Database

	emailConfig := map[string]string{
		"host":     "smtp.example.com",
		"port":     "587",
		"username": "noreply@example.com",
		"password": "password",
	}

	paymentConfig := map[string]string{
		"stripe_secret":    "sk_test_...",
		"paypal_client_id": "client_id",
		"paypal_secret":    "secret",
	}

	us := NewUserService(db, emailConfig, paymentConfig)

	// Example usage
	token := us.GenerateToken()
	fmt.Printf("Generated token: %s\n", token)

	isValid := us.ValidateEmail("test@example.com")
	fmt.Printf("Email valid: %t\n", isValid)
}
