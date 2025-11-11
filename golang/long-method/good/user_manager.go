package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

type UserValidator interface {
	ValidateRegistrationData(userData map[string]string) error
}

type UserRepository interface {
	UserExists(email string) bool
	CreateUser(userData map[string]string) (int64, error)
	CreateUserProfile(userId int64, userData map[string]string) error
	CreateUserSettings(userId int64) error
}

type EmailService interface {
	SendVerificationEmail(email, firstName, verificationToken string) error
}

type NotificationService interface {
	SendWelcomeNotification(userId int64) error
}

type Logger struct{}

func (l Logger) LogRegistration(email string) {
	logMessage := fmt.Sprintf("User registered: %s at %s", email, time.Now().Format("2006-01-02 15:04:05"))
	log.Println(logMessage)
}

type UserManager struct {
	validator           UserValidator
	repository          UserRepository
	emailService        EmailService
	notificationService NotificationService
	logger              Logger
}

func NewUserManager(
	validator UserValidator,
	repository UserRepository,
	emailService EmailService,
	notificationService NotificationService,
) *UserManager {
	return &UserManager{
		validator:           validator,
		repository:          repository,
		emailService:        emailService,
		notificationService: notificationService,
		logger:              Logger{},
	}
}

func (um *UserManager) RegisterUser(userData map[string]string) (int64, error) {
	err := um.validator.ValidateRegistrationData(userData)
	if err != nil {
		return 0, err
	}

	if um.repository.UserExists(userData["email"]) {
		return 0, fmt.Errorf("user already exists")
	}

	userData = um.prepareUserData(userData)

	userId, err := um.repository.CreateUser(userData)
	if err != nil {
		return 0, err
	}

	err = um.repository.CreateUserProfile(userId, userData)
	if err != nil {
		return 0, err
	}

	err = um.repository.CreateUserSettings(userId)
	if err != nil {
		return 0, err
	}

	err = um.emailService.SendVerificationEmail(
		userData["email"],
		userData["firstName"],
		userData["verificationToken"],
	)
	if err != nil {
		return 0, err
	}

	err = um.notificationService.SendWelcomeNotification(userId)
	if err != nil {
		return 0, err
	}

	um.logger.LogRegistration(userData["email"])

	return userId, nil
}

func (um *UserManager) prepareUserData(userData map[string]string) map[string]string {
	userData["password"] = um.hashPassword(userData["password"])
	userData["verificationToken"] = um.generateToken()
	return userData
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
	// This would normally use real implementations of the interfaces
	fmt.Println("Long method has been refactored into smaller, focused methods and separate classes:")
	fmt.Println("- UserManager.RegisterUser() now orchestrates the process")
	fmt.Println("- Validation is handled by UserValidator")
	fmt.Println("- Database operations by UserRepository")
	fmt.Println("- Emails by EmailService")
	fmt.Println("- Notifications by NotificationService")
	fmt.Println("- Logging by Logger")
}
