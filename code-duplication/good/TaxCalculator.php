<?php

class TaxCalculator {
    private const TAX_RATES = [
        'CA' => 0.0825,
        'NY' => 0.04,
        'TX' => 0.0625,
    ];

    public function calculateTax(float $price, string $state): float {
        $taxRate = self::TAX_RATES[$state] ?? 0.0;
        return $price * $taxRate;
    }
}

class ShippingCalculator {
    private const BASE_RATE = 5.0;
    private const WEIGHT_RATE_PER_UNIT = 0.5;
    private const DISTANCE_RATE_PER_UNIT = 0.1;

    public function calculateShipping(float $weight, float $distance): float {
        $weightRate = $weight * self::WEIGHT_RATE_PER_UNIT;
        $distanceRate = $distance * self::DISTANCE_RATE_PER_UNIT;

        return self::BASE_RATE + $weightRate + $distanceRate;
    }
}
