<?php

/**
 * 3. Embedding a temporary variable (Inline Temp)
 *
 * BEFORE: Unnecessary temporary variable
 */
class PriceCalculatorBefore {
    public function getPrice() {
        $basePrice = $this->quantity * $this->itemPrice;
        if ($basePrice > 1000) {
            return $basePrice * 0.95;
        } else {
            return $basePrice * 0.98;
        }
    }
}

/**
 * AFTER: Inline the temporary variable
 */
class PriceCalculatorAfter {
    public function getPrice() {
        if ($this->quantity * $this->itemPrice > 1000) {
            return $this->quantity * $this->itemPrice * 0.95;
        } else {
            return $this->quantity * $this->itemPrice * 0.98;
        }
    }
}

/**
 * 4. Replacing a temporary variable with a method call (Replace Temp with Query)
 *
 * BEFORE: Temporary variable used multiple times
 */
class OrderBefore {
    private $quantity;
    private $itemPrice;

    public function getPrice() {
        $basePrice = $this->quantity * $this->itemPrice;
        return $basePrice - $this->getDiscount($basePrice);
    }

    private function getDiscount($basePrice) {
        return max(0, $basePrice - 500) * 0.05;
    }
}

/**
 * AFTER: Replace temp with query
 */
class OrderAfter {
    private $quantity;
    private $itemPrice;

    public function getPrice() {
        return $this->getBasePrice() - $this->getDiscount();
    }

    private function getBasePrice() {
        return $this->quantity * $this->itemPrice;
    }

    private function getDiscount() {
        return max(0, $this->getBasePrice() - 500) * 0.05;
    }
}

/**
 * 5. Introduction of an explanatory variable (Introduce Explaining Variable)
 *
 * BEFORE: Complex expression hard to understand
 */
class PerformanceCalculatorBefore {
    public function getPerformance() {
        return ($this->goals * 2) + ($this->assists * 1.5) + ($this->minutesPlayed / 60) * 0.1;
    }
}

/**
 * AFTER: Introduce explaining variables for clarity
 */
class PerformanceCalculatorAfter {
    public function getPerformance() {
        $goalPoints = $this->goals * 2;
        $assistPoints = $this->assists * 1.5;
        $playingTimeBonus = ($this->minutesPlayed / 60) * 0.1;

        return $goalPoints + $assistPoints + $playingTimeBonus;
    }
}

/**
 * 6. Splitting a Temporary Variable
 *
 * BEFORE: Same variable used for different purposes
 */
class TemperatureMonitorBefore {
    public function getReading() {
        $temp = $this->getCurrentTemperature();

        // First use: get initial reading
        $initialTemp = $temp;

        // Later: temp is reused for different calculation
        $temp = $temp + $this->getAdjustment();
        $adjustedTemp = $temp;

        return ['initial' => $initialTemp, 'adjusted' => $adjustedTemp];
    }
}

/**
 * AFTER: Split the temporary variable
 */
class TemperatureMonitorAfter {
    public function getReading() {
        $temp = $this->getCurrentTemperature();
        $initialTemp = $temp;

        $adjustedTemp = $temp + $this->getAdjustment();

        return ['initial' => $initialTemp, 'adjusted' => $adjustedTemp];
    }
}

/**
 * 7. Removing parameter Assignments (Remove Assignments to Parameters)
 *
 * BEFORE: Parameter is modified inside method
 */
class DiscountCalculatorBefore {
    public function applyDiscount($price) {
        if ($price > 100) {
            $price = $price * 0.9; // Modifying parameter
        }
        return $price;
    }
}

/**
 * AFTER: Use a local variable instead
 */
class DiscountCalculatorAfter {
    public function applyDiscount($price) {
        $result = $price;
        if ($price > 100) {
            $result = $price * 0.9;
        }
        return $result;
    }
}

/**
 * 8. Replacing a method with a method Object (Replace Method with Method Object)
 *
 * BEFORE: Method with many parameters and local variables
 */
class AccountBefore {
    public function calculateInterest($principal, $rate, $time, $compoundingFrequency) {
        $amount = $principal * pow(1 + ($rate / $compoundingFrequency), $compoundingFrequency * $time);
        $interest = $amount - $principal;
        return $interest;
    }
}

/**
 * AFTER: Extract to a method object
 */
class InterestCalculation {
    private $principal;
    private $rate;
    private $time;
    private $compoundingFrequency;

    public function __construct($principal, $rate, $time, $compoundingFrequency) {
        $this->principal = $principal;
        $this->rate = $rate;
        $this->time = $time;
        $this->compoundingFrequency = $compoundingFrequency;
    }

    public function calculate() {
        $amount = $this->principal * pow(1 + ($this->rate / $this->compoundingFrequency),
                                        $this->compoundingFrequency * $this->time);
        return $amount - $this->principal;
    }
}

class AccountAfter {
    public function calculateInterest($principal, $rate, $time, $compoundingFrequency) {
        $calculation = new InterestCalculation($principal, $rate, $time, $compoundingFrequency);
        return $calculation->calculate();
    }
}
