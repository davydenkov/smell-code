<?php

class OrderProcessor {
    private TaxCalculator $taxCalculator;
    private ShippingCalculator $shippingCalculator;

    public function __construct() {
        $this->taxCalculator = new TaxCalculator();
        $this->shippingCalculator = new ShippingCalculator();
    }

    public function calculateTax(float $price, string $state): float {
        return $this->taxCalculator->calculateTax($price, $state);
    }

    public function calculateShipping(float $weight, float $distance): float {
        return $this->shippingCalculator->calculateShipping($weight, $distance);
    }
}

class InvoiceGenerator {
    private TaxCalculator $taxCalculator;
    private ShippingCalculator $shippingCalculator;

    public function __construct() {
        $this->taxCalculator = new TaxCalculator();
        $this->shippingCalculator = new ShippingCalculator();
    }

    public function calculateTax(float $price, string $state): float {
        return $this->taxCalculator->calculateTax($price, $state);
    }

    public function calculateShipping(float $weight, float $distance): float {
        return $this->shippingCalculator->calculateShipping($weight, $distance);
    }
}

class QuoteGenerator {
    private TaxCalculator $taxCalculator;
    private ShippingCalculator $shippingCalculator;

    public function __construct() {
        $this->taxCalculator = new TaxCalculator();
        $this->shippingCalculator = new ShippingCalculator();
    }

    public function calculateTax(float $price, string $state): float {
        return $this->taxCalculator->calculateTax($price, $state);
    }

    public function calculateShipping(float $weight, float $distance): float {
        return $this->shippingCalculator->calculateShipping($weight, $distance);
    }
}
