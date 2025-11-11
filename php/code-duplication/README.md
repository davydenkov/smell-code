# Code Duplication

## Problem

Code duplication occurs when the same code logic appears in multiple places. This violates the **DRY (Don't Repeat Yourself)** principle and leads to maintenance nightmares.

## Why It's Bad

- **Maintenance Burden**: Changes need to be made in multiple places
- **Bug Risk**: Fixing a bug in one location but forgetting others
- **Inconsistency**: Different implementations of the same logic
- **Code Bloat**: Unnecessarily large codebase

## Example Scenario

In an e-commerce system, tax calculation logic might be duplicated across:
- Order processing
- Invoice generation
- Report calculations
- Shopping cart totals

## Before Refactoring

```php
<?php
class OrderProcessor {
    public function calculateTotal($subtotal) {
        $tax = $subtotal * 0.08;  // Tax calculation
        $shipping = $subtotal > 100 ? 5.00 : 10.00;  // Shipping logic
        return $subtotal + $tax + $shipping;
    }
}

class InvoiceProcessor {
    public function calculateTotal($subtotal) {
        $tax = $subtotal * 0.08;  // Same tax calculation - DUPLICATED
        $shipping = $subtotal > 100 ? 5.00 : 10.00;  // Same shipping logic - DUPLICATED
        return $subtotal + $tax + $shipping;
    }
}
```

## After Refactoring

```php
<?php
class TaxCalculator {
    public function calculateTax($amount, $rate = 0.08) {
        return $amount * $rate;
    }
}

class ShippingCalculator {
    public function calculateShipping($subtotal) {
        return $subtotal > 100 ? 5.00 : 10.00;
    }
}

class OrderProcessor {
    private $taxCalculator;
    private $shippingCalculator;

    public function __construct(TaxCalculator $taxCalc, ShippingCalculator $shippingCalc) {
        $this->taxCalculator = $taxCalc;
        $this->shippingCalculator = $shippingCalc;
    }

    public function calculateTotal($subtotal) {
        $tax = $this->taxCalculator->calculateTax($subtotal);
        $shipping = $this->shippingCalculator->calculateShipping($subtotal);
        return $subtotal + $tax + $shipping;
    }
}
```

## Refactoring Techniques Used

1. **Extract Method**: Move duplicated logic into separate methods
2. **Extract Class**: Create dedicated classes for specific responsibilities
3. **Dependency Injection**: Pass calculators as dependencies

## Files in This Example

- `bad/OrderProcessor.php` - Shows the problematic duplicated code
- `good/OrderProcessor.php` - Refactored OrderProcessor using composition
- `good/TaxCalculator.php` - Extracted tax calculation logic

## Running the Examples

```bash
# Run the bad example
php bad/OrderProcessor.php

# Run the good examples
php good/OrderProcessor.php
php good/TaxCalculator.php
```

## Key Takeaways

- Always look for repeated code patterns
- Extract common logic into reusable components
- Use composition over duplication
- Test refactored code thoroughly to ensure behavior is preserved
