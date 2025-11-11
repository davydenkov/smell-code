<?php

class FinancialCalculator {
    public function calculateInterest(float $principal, float $rate, int $time): float {
        return $principal * $rate * $time;
    }

    public function calculateTax(float $income, float $deductions): float {
        $taxableIncome = $income - $deductions;
        if ($taxableIncome <= 50000) {
            return $taxableIncome * 0.1;
        } elseif ($taxableIncome <= 100000) {
            return 5000 + ($taxableIncome - 50000) * 0.2;
        } else {
            return 15000 + ($taxableIncome - 100000) * 0.3;
        }
    }
}
