<?php

class TransactionRepository {
    private $db;

    public function __construct($db) {
        $this->db = $db;
    }

    public function saveTransaction(int $userId, float $amount, string $type): int {
        $stmt = $this->db->prepare("INSERT INTO transactions (user_id, amount, type, created_at) VALUES (?, ?, ?, NOW())");
        $stmt->execute([$userId, $amount, $type]);
        return $this->db->lastInsertId();
    }

    public function getUserBalance(int $userId): float {
        $stmt = $this->db->prepare("SELECT SUM(amount) as balance FROM transactions WHERE user_id = ?");
        $stmt->execute([$userId]);
        $result = $stmt->fetch();
        return $result['balance'] ?? 0;
    }

    public function getMonthlyTransactions(int $userId, int $month, int $year): array {
        $stmt = $this->db->prepare("
            SELECT type, SUM(amount) as total, COUNT(*) as count
            FROM transactions
            WHERE user_id = ? AND MONTH(created_at) = ? AND YEAR(created_at) = ?
            GROUP BY type
        ");
        $stmt->execute([$userId, $month, $year]);
        return $stmt->fetchAll();
    }

    public function getYearlyIncome(int $userId, int $year): float {
        $stmt = $this->db->prepare("
            SELECT SUM(amount) as total_income
            FROM transactions
            WHERE user_id = ? AND type = 'income' AND YEAR(created_at) = ?
        ");
        $stmt->execute([$userId, $year]);
        return $stmt->fetch()['total_income'] ?? 0;
    }

    public function getYearlyDeductions(int $userId, int $year): float {
        $stmt = $this->db->prepare("
            SELECT SUM(amount) as total_deductions
            FROM transactions
            WHERE user_id = ? AND type = 'deduction' AND YEAR(created_at) = ?
        ");
        $stmt->execute([$userId, $year]);
        return $stmt->fetch()['total_deductions'] ?? 0;
    }
}
