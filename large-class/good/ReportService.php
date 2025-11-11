<?php

class ReportService {
    private $db;

    public function __construct($db) {
        $this->db = $db;
    }

    public function generateUserReport(int $userId): array {
        $stmt = $this->db->prepare("
            SELECT u.name, u.email, u.created_at,
                   COUNT(o.id) as order_count,
                   SUM(o.total) as total_spent
            FROM users u
            LEFT JOIN orders o ON u.id = o.user_id
            WHERE u.id = ?
            GROUP BY u.id
        ");
        $stmt->execute([$userId]);
        return $stmt->fetch();
    }

    public function generateSalesReport(string $startDate, string $endDate): array {
        $stmt = $this->db->prepare("
            SELECT DATE(created_at) as date,
                   COUNT(*) as order_count,
                   SUM(total) as total_sales
            FROM orders
            WHERE created_at BETWEEN ? AND ?
            GROUP BY DATE(created_at)
        ");
        $stmt->execute([$startDate, $endDate]);
        return $stmt->fetchAll();
    }
}
