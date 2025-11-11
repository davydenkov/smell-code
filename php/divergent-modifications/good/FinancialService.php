<?php

class FinancialService {
    private FinancialCalculator $calculator;
    private TransactionRepository $transactionRepository;
    private UserRepository $userRepository;
    private ReportGenerator $reportGenerator;
    private EmailService $emailService;

    public function __construct($db) {
        $this->calculator = new FinancialCalculator();
        $this->transactionRepository = new TransactionRepository($db);
        $this->userRepository = new UserRepository($db);
        $this->reportGenerator = new ReportGenerator($this->transactionRepository, $this->calculator);
        $this->emailService = new EmailService();
    }

    // Financial calculations - delegated to FinancialCalculator
    public function calculateInterest(float $principal, float $rate, int $time): float {
        return $this->calculator->calculateInterest($principal, $rate, $time);
    }

    public function calculateTax(float $income, float $deductions): float {
        return $this->calculator->calculateTax($income, $deductions);
    }

    // Database operations - delegated to repositories
    public function saveTransaction(int $userId, float $amount, string $type): int {
        return $this->transactionRepository->saveTransaction($userId, $amount, $type);
    }

    public function getUserBalance(int $userId): float {
        return $this->transactionRepository->getUserBalance($userId);
    }

    public function updateUserProfile(int $userId, string $name, string $email): void {
        $this->userRepository->updateUserProfile($userId, $name, $email);
    }

    // Reporting - delegated to ReportGenerator
    public function generateMonthlyReport(int $userId, int $month, int $year): array {
        return $this->reportGenerator->generateMonthlyReport($userId, $month, $year);
    }

    public function generateTaxReport(int $userId, int $year): array {
        return $this->reportGenerator->generateTaxReport($userId, $year);
    }

    // Email notifications - delegated to EmailService
    public function sendMonthlyStatement(int $userId): void {
        $email = $this->userRepository->getUserEmail($userId);
        if (!$email) {
            return;
        }

        $balance = $this->getUserBalance($userId);
        $report = $this->generateMonthlyReport($userId, (int)date('m'), (int)date('Y'));

        $this->emailService->sendMonthlyStatement($email, $balance, $report);
    }
}
