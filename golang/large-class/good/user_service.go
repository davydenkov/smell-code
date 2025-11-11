package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type UserRepository interface {
	Create(email, name, hashedPassword string) (int, error)
	Authenticate(email, password string) (*User, error)
	FindById(id int) (*User, error)
	Update(user *User) error
}

type EmailService interface {
	SendWelcomeEmail(email, name string) error
	SendPasswordResetEmail(email, resetToken string) error
	SendNotificationEmail(email, subject, message string) error
}

type PaymentService interface {
	ProcessStripePayment(amount float64, token string) map[string]interface{}
	ProcessPayPalPayment(amount float64, paypalToken string) map[string]interface{}
	RefundPayment(transactionId string, amount float64) map[string]interface{}
}

type ReportService interface {
	GenerateUserReport(userId int) map[string]interface{}
	GenerateSalesReport(startDate, endDate string) []map[string]interface{}
}

type ActivityLogger interface {
	LogActivity(userId int, action string) error
}

type UserService struct {
	userRepository UserRepository
	emailService   EmailService
	paymentService PaymentService
	reportService  ReportService
	activityLogger ActivityLogger
}

func NewUserService(
	userRepo UserRepository,
	emailSvc EmailService,
	paymentSvc PaymentService,
	reportSvc ReportService,
	activityLog ActivityLogger,
) *UserService {
	return &UserService{
		userRepository: userRepo,
		emailService:   emailSvc,
		paymentService: paymentSvc,
		reportService:  reportSvc,
		activityLogger: activityLog,
	}
}

// User management methods
func (us *UserService) CreateUser(email, name, password string) (int, error) {
	hashedPassword := hashPassword(password)
	userId, err := us.userRepository.Create(email, name, hashedPassword)
	if err != nil {
		return 0, err
	}

	us.emailService.SendWelcomeEmail(email, name)
	us.activityLogger.LogActivity(userId, "user_created")

	return userId, nil
}

func (us *UserService) AuthenticateUser(email, password string) (*User, error) {
	user, err := us.userRepository.Authenticate(email, password)
	if err != nil {
		return nil, err
	}

	if user != nil {
		us.activityLogger.LogActivity(user.GetId(), "user_login")
	}

	return user, nil
}

func (us *UserService) UpdateUserProfile(userId int, name, email string) error {
	user, err := us.userRepository.FindById(userId)
	if err != nil {
		return err
	}

	if user != nil {
		user.UpdateProfile(name, email)
		err = us.userRepository.Update(user)
		if err != nil {
			return err
		}
		us.activityLogger.LogActivity(userId, "profile_updated")
	}

	return nil
}

func (us *UserService) GetUserBalance(userId int) float64 {
	user, err := us.userRepository.FindById(userId)
	if err != nil {
		return 0.0
	}
	if user != nil {
		return user.GetBalance()
	}
	return 0.0
}

// Email methods - delegated to EmailService
func (us *UserService) SendWelcomeEmail(email, name string) error {
	return us.emailService.SendWelcomeEmail(email, name)
}

func (us *UserService) SendPasswordResetEmail(email, resetToken string) error {
	return us.emailService.SendPasswordResetEmail(email, resetToken)
}

func (us *UserService) SendNotificationEmail(email, subject, message string) error {
	return us.emailService.SendNotificationEmail(email, subject, message)
}

// Payment methods - delegated to PaymentService
func (us *UserService) ProcessStripePayment(amount float64, token string) map[string]interface{} {
	return us.paymentService.ProcessStripePayment(amount, token)
}

func (us *UserService) ProcessPayPalPayment(amount float64, paypalToken string) map[string]interface{} {
	return us.paymentService.ProcessPayPalPayment(amount, paypalToken)
}

func (us *UserService) RefundPayment(transactionId string, amount float64) map[string]interface{} {
	return us.paymentService.RefundPayment(transactionId, amount)
}

// Reporting methods - delegated to ReportService
func (us *UserService) GenerateUserReport(userId int) map[string]interface{} {
	return us.reportService.GenerateUserReport(userId)
}

func (us *UserService) GenerateSalesReport(startDate, endDate string) []map[string]interface{} {
	return us.reportService.GenerateSalesReport(startDate, endDate)
}

// Utility function
func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
