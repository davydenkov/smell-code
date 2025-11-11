<?php

class ReportGenerator {
    private TransactionRepository $transactionRepository;
    private FinancialCalculator $financialCalculator;

    public function __construct(TransactionRepository $transactionRepository, FinancialCalculator $financialCalculator) {
        $this->transactionRepository = $transactionRepository;
        $this->financialCalculator = $financialCalculator;
    }

    public function generateMonthlyReport(int $userId, int $month, int $year): array {
        return $this->transactionRepository->getMonthlyTransactions($userId, $month, $year);
    }

    public function generateTaxReport(int $userId, int $year): array {
        $income = $this->transactionRepository->getYearlyIncome($userId, $year);
        $deductions = $this->transactionRepository->getYearlyDeductions($userId, $year);
        $taxOwed = $this->financialCalculator->calculateTax($income, $deductions);

        return [
            'year' => $year,
            'total_income' => $income,
            'total_deductions' => $deductions,
            'tax_owed' => $taxOwed
        ];
    }
}
