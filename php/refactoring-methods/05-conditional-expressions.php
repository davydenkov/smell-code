<?php

/**
 * 34. Decomposition of a conditional operator (Decompose Conditional)
 *
 * BEFORE: Complex conditional logic
 */
class PaymentProcessorBefore {
    public function calculateFee($amount, $isInternational, $isPremium) {
        $fee = 0;
        if ($amount > 100 && $isInternational && $isPremium) {
            $fee = $amount * 0.05 + 10;
        } elseif ($amount > 100 && $isInternational && !$isPremium) {
            $fee = $amount * 0.05 + 15;
        } elseif ($amount <= 100 && $isInternational) {
            $fee = $amount * 0.03 + 5;
        } else {
            $fee = $amount * 0.02;
        }
        return $fee;
    }
}

/**
 * AFTER: Decompose conditional
 */
class PaymentProcessorAfter {
    public function calculateFee($amount, $isInternational, $isPremium) {
        if ($this->isHighValueInternationalPremium($amount, $isInternational, $isPremium)) {
            return $this->calculateHighValueInternationalPremiumFee($amount);
        } elseif ($this->isHighValueInternationalStandard($amount, $isInternational, $isPremium)) {
            return $this->calculateHighValueInternationalStandardFee($amount);
        } elseif ($this->isLowValueInternational($amount, $isInternational)) {
            return $this->calculateLowValueInternationalFee($amount);
        } else {
            return $this->calculateDomesticFee($amount);
        }
    }

    private function isHighValueInternationalPremium($amount, $isInternational, $isPremium) {
        return $amount > 100 && $isInternational && $isPremium;
    }

    private function isHighValueInternationalStandard($amount, $isInternational, $isPremium) {
        return $amount > 100 && $isInternational && !$isPremium;
    }

    private function isLowValueInternational($amount, $isInternational) {
        return $amount <= 100 && $isInternational;
    }

    private function calculateHighValueInternationalPremiumFee($amount) {
        return $amount * 0.05 + 10;
    }

    private function calculateHighValueInternationalStandardFee($amount) {
        return $amount * 0.05 + 15;
    }

    private function calculateLowValueInternationalFee($amount) {
        return $amount * 0.03 + 5;
    }

    private function calculateDomesticFee($amount) {
        return $amount * 0.02;
    }
}

/**
 * 35. Consolidation of a conditional expression (Consolidate Conditional Expression)
 *
 * BEFORE: Multiple conditionals with same result
 */
class InsuranceCalculatorBefore {
    public function isEligibleForDiscount($age, $isStudent, $hasGoodRecord) {
        if ($age < 25) return false;
        if ($isStudent) return true;
        if ($hasGoodRecord) return true;
        return false;
    }
}

/**
 * AFTER: Consolidate conditionals
 */
class InsuranceCalculatorAfter {
    public function isEligibleForDiscount($age, $isStudent, $hasGoodRecord) {
        return $age >= 25 && ($isStudent || $hasGoodRecord);
    }
}

/**
 * 36. Consolidation of duplicate conditional fragments
 * (Consolidate Duplicate Conditional Fragments)
 *
 * BEFORE: Duplicate code in conditional branches
 */
class FileProcessorBefore {
    public function processFile($file) {
        if ($this->isValidFile($file)) {
            $this->logProcessing($file);
            $this->validateContent($file);
            $this->saveToDatabase($file);
            $this->sendNotification($file);
        } else {
            $this->logError($file);
            $this->sendNotification($file); // Duplicate
        }
    }

    private function sendNotification($file) {
        // Send notification logic
    }
}

/**
 * AFTER: Consolidate duplicate fragments
 */
class FileProcessorAfter {
    public function processFile($file) {
        $this->sendNotification($file); // Moved outside conditional

        if ($this->isValidFile($file)) {
            $this->logProcessing($file);
            $this->validateContent($file);
            $this->saveToDatabase($file);
        } else {
            $this->logError($file);
        }
    }

    private function sendNotification($file) {
        // Send notification logic
    }
}

/**
 * 37. Remove Control Flag
 *
 * BEFORE: Control flag to break out of loop
 */
class DataProcessorBefore {
    public function findPerson($people, $name) {
        $found = false;
        foreach ($people as $person) {
            if ($person['name'] === $name) {
                $found = $person;
                break; // Control flag usage
            }
        }
        return $found;
    }
}

/**
 * AFTER: Remove control flag
 */
class DataProcessorAfter {
    public function findPerson($people, $name) {
        foreach ($people as $person) {
            if ($person['name'] === $name) {
                return $person; // Direct return
            }
        }
        return false;
    }
}

/**
 * 38. Replacing Nested Conditional statements with a boundary operator
 * (Replace Nested Conditional with Guard Clauses)
 *
 * BEFORE: Nested conditionals
 */
class PaymentValidatorBefore {
    public function isValidPayment($payment) {
        if ($payment['amount'] > 0) {
            if ($payment['cardNumber'] !== null) {
                if (strlen($payment['cardNumber']) === 16) {
                    if ($this->isValidExpiry($payment['expiry'])) {
                        return true;
                    }
                }
            }
        }
        return false;
    }

    private function isValidExpiry($expiry) {
        return strtotime($expiry) > time();
    }
}

/**
 * AFTER: Replace with guard clauses
 */
class PaymentValidatorAfter {
    public function isValidPayment($payment) {
        if ($payment['amount'] <= 0) {
            return false;
        }

        if ($payment['cardNumber'] === null) {
            return false;
        }

        if (strlen($payment['cardNumber']) !== 16) {
            return false;
        }

        if (!$this->isValidExpiry($payment['expiry'])) {
            return false;
        }

        return true;
    }

    private function isValidExpiry($expiry) {
        return strtotime($expiry) > time();
    }
}

/**
 * 39. Replacing a conditional operator with polymorphism (Replace Conditional with Polymorphism)
 *
 * BEFORE: Type checking with conditionals
 */
class BirdBefore {
    const EUROPEAN = 'european';
    const AFRICAN = 'african';
    const NORWEGIAN_BLUE = 'norwegian_blue';

    private $type;
    private $voltage;
    private $isNailed;

    public function __construct($type) {
        $this->type = $type;
    }

    public function getSpeed() {
        switch ($this->type) {
            case self::EUROPEAN:
                return $this->getBaseSpeed();
            case self::AFRICAN:
                return $this->getBaseSpeed() - $this->voltage * 2;
            case self::NORWEGIAN_BLUE:
                return $this->isNailed ? 0 : $this->getBaseSpeed();
            default:
                return $this->getBaseSpeed();
        }
    }

    private function getBaseSpeed() {
        return 10;
    }
}

/**
 * AFTER: Replace conditional with polymorphism
 */
abstract class BirdAfter {
    abstract public function getSpeed();

    protected function getBaseSpeed() {
        return 10;
    }
}

class EuropeanSwallow extends BirdAfter {
    public function getSpeed() {
        return $this->getBaseSpeed();
    }
}

class AfricanSwallow extends BirdAfter {
    private $voltage;

    public function __construct($voltage) {
        $this->voltage = $voltage;
    }

    public function getSpeed() {
        return $this->getBaseSpeed() - $this->voltage * 2;
    }
}

class NorwegianBlueParrot extends BirdAfter {
    private $isNailed;

    public function __construct($isNailed) {
        $this->isNailed = $isNailed;
    }

    public function getSpeed() {
        return $this->isNailed ? 0 : $this->getBaseSpeed();
    }
}

/**
 * 40. Introduction of the object (Introduce Object)
 *
 * BEFORE: Primitive obsession with conditionals
 */
class UserValidatorBefore {
    public function validateUser($user) {
        if (empty($user['name'])) {
            return 'Name is required';
        }

        if (strlen($user['name']) < 2) {
            return 'Name must be at least 2 characters';
        }

        if (!filter_var($user['email'], FILTER_VALIDATE_EMAIL)) {
            return 'Invalid email format';
        }

        return true;
    }
}

/**
 * AFTER: Introduce validation result object
 */
class ValidationResult {
    private $isValid;
    private $errors;

    public function __construct($isValid = true, $errors = []) {
        $this->isValid = $isValid;
        $this->errors = $errors;
    }

    public function isValid() {
        return $this->isValid;
    }

    public function getErrors() {
        return $this->errors;
    }

    public function addError($error) {
        $this->isValid = false;
        $this->errors[] = $error;
        return $this;
    }
}

class UserValidatorAfter {
    public function validateUser($user) {
        $result = new ValidationResult();

        if (empty($user['name'])) {
            $result->addError('Name is required');
        }

        if (strlen($user['name']) < 2) {
            $result->addError('Name must be at least 2 characters');
        }

        if (!filter_var($user['email'], FILTER_VALIDATE_EMAIL)) {
            $result->addError('Invalid email format');
        }

        return $result;
    }
}

/**
 * 41. Introduction of the statement (Introduction Statement)
 *
 * BEFORE: Magic assertion
 */
class AccountAssertion {
    private $balance;

    public function withdraw($amount) {
        assert($amount > 0 && $amount <= $this->balance);
        $this->balance -= $amount;
    }
}

/**
 * AFTER: Introduce assertion method
 */
class AccountAssertionAfter {
    private $balance;

    public function withdraw($amount) {
        $this->assertValidWithdrawal($amount);
        $this->balance -= $amount;
    }

    private function assertValidWithdrawal($amount) {
        assert($amount > 0 && $amount <= $this->balance, 'Invalid withdrawal amount');
    }
}
