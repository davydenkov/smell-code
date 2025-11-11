<?php

class EmailService {
    public function sendMonthlyStatement(string $email, float $balance, array $monthlyReport): void {
        $subject = 'Monthly Financial Statement';
        $message = "Your current balance: $" . number_format($balance, 2) . "\n\n";
        $message .= "Transactions this month:\n";

        foreach ($monthlyReport as $row) {
            $message .= "- {$row['type']}: {$row['count']} transactions, total: $" . number_format($row['total'], 2) . "\n";
        }

        $headers = 'From: finance@company.com';
        mail($email, $subject, $message, $headers);
    }
}
