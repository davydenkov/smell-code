<?php

class FinancialService {
    private $db;

    public function __construct($db) {
        $this->db = $db;
    }

    // Financial calculations - changes when business rules change
    public function calculateInterest($principal, $rate, $time) {
        return $principal * $rate * $time;
    }

    public function calculateTax($income, $deductions) {
        $taxableIncome = $income - $deductions;
        if ($taxableIncome <= 50000) {
            return $taxableIncome * 0.1;
        } elseif ($taxableIncome <= 100000) {
            return 5000 + ($taxableIncome - 50000) * 0.2;
        } else {
            return 15000 + ($taxableIncome - 100000) * 0.3;
        }
    }

    // Database operations - changes when data schema changes
    public function saveTransaction($userId, $amount, $type) {
        $stmt = $this->db->prepare("INSERT INTO transactions (user_id, amount, type, created_at) VALUES (?, ?, ?, NOW())");
        $stmt->execute([$userId, $amount, $type]);
        return $this->db->lastInsertId();
    }

    public function getUserBalance($userId) {
        $stmt = $this->db->prepare("SELECT SUM(amount) as balance FROM transactions WHERE user_id = ?");
        $stmt->execute([$userId]);
        $result = $stmt->fetch();
        return $result['balance'] ?? 0;
    }

    public function updateUserProfile($userId, $name, $email) {
        $stmt = $this->db->prepare("UPDATE users SET name = ?, email = ? WHERE id = ?");
        $stmt->execute([$name, $email, $userId]);
    }

    // Reporting - changes when reporting requirements change
    public function generateMonthlyReport($userId, $month, $year) {
        $stmt = $this->db->prepare("
            SELECT type, SUM(amount) as total, COUNT(*) as count
            FROM transactions
            WHERE user_id = ? AND MONTH(created_at) = ? AND YEAR(created_at) = ?
            GROUP BY type
        ");
        $stmt->execute([$userId, $month, $year]);
        return $stmt->fetchAll();
    }

    public function generateTaxReport($userId, $year) {
        $stmt = $this->db->prepare("
            SELECT SUM(amount) as total_income
            FROM transactions
            WHERE user_id = ? AND type = 'income' AND YEAR(created_at) = ?
        ");
        $stmt->execute([$userId, $year]);
        $income = $stmt->fetch()['total_income'] ?? 0;

        $stmt = $this->db->prepare("
            SELECT SUM(amount) as total_deductions
            FROM transactions
            WHERE user_id = ? AND type = 'deduction' AND YEAR(created_at) = ?
        ");
        $stmt->execute([$userId, $year]);
        $deductions = $stmt->fetch()['total_deductions'] ?? 0;

        $taxOwed = $this->calculateTax($income, $deductions);

        return [
            'year' => $year,
            'total_income' => $income,
            'total_deductions' => $deductions,
            'tax_owed' => $taxOwed
        ];
    }

    // Email notifications - changes when communication requirements change
    public function sendMonthlyStatement($userId, $email) {
        $balance = $this->getUserBalance($userId);
        $report = $this->generateMonthlyReport($userId, date('m'), date('Y'));

        $subject = 'Monthly Financial Statement';
        $message = "Your current balance: $" . number_format($balance, 2) . "\n\n";
        $message .= "Transactions this month:\n";

        foreach ($report as $row) {
            $message .= "- {$row['type']}: {$row['count']} transactions, total: $" . number_format($row['total'], 2) . "\n";
        }

        $headers = 'From: finance@company.com';
        mail($email, $subject, $message, $headers);
    }
}
