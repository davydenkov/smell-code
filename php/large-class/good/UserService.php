<?php

class UserService {
    private UserRepository $userRepository;
    private EmailService $emailService;
    private PaymentService $paymentService;
    private ReportService $reportService;
    private ActivityLogger $activityLogger;

    public function __construct($db, array $emailConfig, array $paymentConfig) {
        $this->userRepository = new UserRepository($db);
        $this->emailService = new EmailService($emailConfig);
        $this->paymentService = new PaymentService($paymentConfig);
        $this->reportService = new ReportService($db);
        $this->activityLogger = new ActivityLogger($db);
    }

    // User management methods
    public function createUser(string $email, string $name, string $password): int {
        $hashedPassword = Utility::hashPassword($password);
        $userId = $this->userRepository->create($email, $name, $hashedPassword);

        $this->emailService->sendWelcomeEmail($email, $name);
        $this->activityLogger->logActivity($userId, 'user_created');

        return $userId;
    }

    public function authenticateUser(string $email, string $password): ?User {
        $user = $this->userRepository->authenticate($email, $password);

        if ($user) {
            $this->activityLogger->logActivity($user->getId(), 'user_login');
        }

        return $user;
    }

    public function updateUserProfile(int $userId, string $name, string $email): void {
        $user = $this->userRepository->findById($userId);
        if ($user) {
            $user->updateProfile($name, $email);
            $this->userRepository->update($user);
            $this->activityLogger->logActivity($userId, 'profile_updated');
        }
    }

    public function getUserBalance(int $userId): float {
        $user = $this->userRepository->findById($userId);
        return $user ? $user->getBalance() : 0.0;
    }

    // Email methods
    public function sendWelcomeEmail(string $email, string $name): bool {
        return $this->emailService->sendWelcomeEmail($email, $name);
    }

    public function sendPasswordResetEmail(string $email, string $resetToken): bool {
        return $this->emailService->sendPasswordResetEmail($email, $resetToken);
    }

    public function sendNotificationEmail(string $email, string $subject, string $message): bool {
        return $this->emailService->sendNotificationEmail($email, $subject, $message);
    }

    // Payment methods
    public function processStripePayment(float $amount, string $token): array {
        return $this->paymentService->processStripePayment($amount, $token);
    }

    public function processPayPalPayment(float $amount, string $paypalToken): array {
        return $this->paymentService->processPayPalPayment($amount, $paypalToken);
    }

    public function refundPayment(string $transactionId, float $amount): array {
        return $this->paymentService->refundPayment($transactionId, $amount);
    }

    // Reporting methods
    public function generateUserReport(int $userId): array {
        return $this->reportService->generateUserReport($userId);
    }

    public function generateSalesReport(string $startDate, string $endDate): array {
        return $this->reportService->generateSalesReport($startDate, $endDate);
    }
}
