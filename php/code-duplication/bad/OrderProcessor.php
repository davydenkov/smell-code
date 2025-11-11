<?php

class OrderProcessor {
    public function calculateTax($price, $state) {
        $taxRate = 0.0;

        if ($state === 'CA') {
            $taxRate = 0.0825;
        } elseif ($state === 'NY') {
            $taxRate = 0.04;
        } elseif ($state === 'TX') {
            $taxRate = 0.0625;
        }

        return $price * $taxRate;
    }

    public function calculateShipping($weight, $distance) {
        $baseRate = 5.0;
        $weightRate = $weight * 0.5;
        $distanceRate = $distance * 0.1;

        return $baseRate + $weightRate + $distanceRate;
    }
}

class InvoiceGenerator {
    public function calculateTax($price, $state) {
        $taxRate = 0.0;

        if ($state === 'CA') {
            $taxRate = 0.0825;
        } elseif ($state === 'NY') {
            $taxRate = 0.04;
        } elseif ($state === 'TX') {
            $taxRate = 0.0625;
        }

        return $price * $taxRate;
    }

    public function calculateShipping($weight, $distance) {
        $baseRate = 5.0;
        $weightRate = $weight * 0.5;
        $distanceRate = $distance * 0.1;

        return $baseRate + $weightRate + $distanceRate;
    }
}

class QuoteGenerator {
    public function calculateTax($price, $state) {
        $taxRate = 0.0;

        if ($state === 'CA') {
            $taxRate = 0.0825;
        } elseif ($state === 'NY') {
            $taxRate = 0.04;
        } elseif ($state === 'TX') {
            $taxRate = 0.0625;
        }

        return $price * $taxRate;
    }

    public function calculateShipping($weight, $distance) {
        $baseRate = 5.0;
        $weightRate = $weight * 0.5;
        $distanceRate = $distance * 0.1;

        return $baseRate + $weightRate + $distanceRate;
    }
}
