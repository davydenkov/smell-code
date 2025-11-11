/**
 * 34. Decomposition of a conditional operator (Decompose Conditional)
 *
 * BEFORE: Complex conditional logic
 */
class PaymentProcessorBefore {
    calculateFee(amount, isInternational, isPremium) {
        let fee = 0;
        if (amount > 100 && isInternational && isPremium) {
            fee = amount * 0.05 + 10;
        } else if (amount > 100 && isInternational && !isPremium) {
            fee = amount * 0.05 + 15;
        } else if (amount <= 100 && isInternational) {
            fee = amount * 0.03 + 5;
        } else {
            fee = amount * 0.02;
        }
        return fee;
    }
}

/**
 * AFTER: Decompose conditional
 */
class PaymentProcessorAfter {
    calculateFee(amount, isInternational, isPremium) {
        if (this.isHighValueInternationalPremium(amount, isInternational, isPremium)) {
            return this.calculateHighValueInternationalPremiumFee(amount);
        } else if (this.isHighValueInternationalStandard(amount, isInternational, isPremium)) {
            return this.calculateHighValueInternationalStandardFee(amount);
        } else if (this.isLowValueInternational(amount, isInternational)) {
            return this.calculateLowValueInternationalFee(amount);
        } else {
            return this.calculateDomesticFee(amount);
        }
    }

    isHighValueInternationalPremium(amount, isInternational, isPremium) {
        return amount > 100 && isInternational && isPremium;
    }

    isHighValueInternationalStandard(amount, isInternational, isPremium) {
        return amount > 100 && isInternational && !isPremium;
    }

    isLowValueInternational(amount, isInternational) {
        return amount <= 100 && isInternational;
    }

    calculateHighValueInternationalPremiumFee(amount) {
        return amount * 0.05 + 10;
    }

    calculateHighValueInternationalStandardFee(amount) {
        return amount * 0.05 + 15;
    }

    calculateLowValueInternationalFee(amount) {
        return amount * 0.03 + 5;
    }

    calculateDomesticFee(amount) {
        return amount * 0.02;
    }
}

/**
 * 35. Consolidation of a conditional expression (Consolidate Conditional Expression)
 *
 * BEFORE: Multiple conditionals with same result
 */
class InsuranceCalculatorBefore {
    isEligibleForDiscount(age, isStudent, hasGoodRecord) {
        if (age < 25) return false;
        if (isStudent) return true;
        if (hasGoodRecord) return true;
        return false;
    }
}

/**
 * AFTER: Consolidate conditionals
 */
class InsuranceCalculatorAfter {
    isEligibleForDiscount(age, isStudent, hasGoodRecord) {
        return age >= 25 && (isStudent || hasGoodRecord);
    }
}

/**
 * 36. Consolidation of duplicate conditional fragments
 * (Consolidate Duplicate Conditional Fragments)
 *
 * BEFORE: Duplicate code in conditional branches
 */
class FileProcessorBefore {
    processFile(file) {
        if (this.isValidFile(file)) {
            this.logProcessing(file);
            this.validateContent(file);
            this.saveToDatabase(file);
            this.sendNotification(file);
        } else {
            this.logError(file);
            this.sendNotification(file); // Duplicate
        }
    }

    sendNotification(file) {
        // Send notification logic
    }
}

/**
 * AFTER: Consolidate duplicate fragments
 */
class FileProcessorAfter {
    processFile(file) {
        this.sendNotification(file); // Moved outside conditional

        if (this.isValidFile(file)) {
            this.logProcessing(file);
            this.validateContent(file);
            this.saveToDatabase(file);
        } else {
            this.logError(file);
        }
    }

    sendNotification(file) {
        // Send notification logic
    }
}

/**
 * 37. Remove Control Flag
 *
 * BEFORE: Control flag to break out of loop
 */
class DataProcessorBefore {
    findPerson(people, name) {
        let found = false;
        for (const person of people) {
            if (person.name === name) {
                found = person;
                break; // Control flag usage
            }
        }
        return found;
    }
}

/**
 * AFTER: Remove control flag
 */
class DataProcessorAfter {
    findPerson(people, name) {
        for (const person of people) {
            if (person.name === name) {
                return person; // Direct return
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
    isValidPayment(payment) {
        if (payment.amount > 0) {
            if (payment.cardNumber !== null) {
                if (payment.cardNumber.length === 16) {
                    if (this.isValidExpiry(payment.expiry)) {
                        return true;
                    }
                }
            }
        }
        return false;
    }

    isValidExpiry(expiry) {
        return new Date(expiry) > new Date();
    }
}

/**
 * AFTER: Replace with guard clauses
 */
class PaymentValidatorAfter {
    isValidPayment(payment) {
        if (payment.amount <= 0) {
            return false;
        }

        if (payment.cardNumber === null) {
            return false;
        }

        if (payment.cardNumber.length !== 16) {
            return false;
        }

        if (!this.isValidExpiry(payment.expiry)) {
            return false;
        }

        return true;
    }

    isValidExpiry(expiry) {
        return new Date(expiry) > new Date();
    }
}

/**
 * 39. Replacing a conditional operator with polymorphism (Replace Conditional with Polymorphism)
 *
 * BEFORE: Type checking with conditionals
 */
class BirdBefore {
    static EUROPEAN = 'european';
    static AFRICAN = 'african';
    static NORWEGIAN_BLUE = 'norwegian_blue';

    constructor(type) {
        this.type = type;
        this.voltage = null;
        this.isNailed = false;
    }

    getSpeed() {
        switch (this.type) {
            case BirdBefore.EUROPEAN:
                return this.getBaseSpeed();
            case BirdBefore.AFRICAN:
                return this.getBaseSpeed() - this.voltage * 2;
            case BirdBefore.NORWEGIAN_BLUE:
                return this.isNailed ? 0 : this.getBaseSpeed();
            default:
                return this.getBaseSpeed();
        }
    }

    getBaseSpeed() {
        return 10;
    }
}

/**
 * AFTER: Replace conditional with polymorphism
 */
class BirdAfter {
    getSpeed() {
        throw new Error('Abstract method must be implemented');
    }

    getBaseSpeed() {
        return 10;
    }
}

class EuropeanSwallow extends BirdAfter {
    getSpeed() {
        return this.getBaseSpeed();
    }
}

class AfricanSwallow extends BirdAfter {
    constructor(voltage) {
        super();
        this.voltage = voltage;
    }

    getSpeed() {
        return this.getBaseSpeed() - this.voltage * 2;
    }
}

class NorwegianBlueParrot extends BirdAfter {
    constructor(isNailed) {
        super();
        this.isNailed = isNailed;
    }

    getSpeed() {
        return this.isNailed ? 0 : this.getBaseSpeed();
    }
}

/**
 * 40. Introduction of the object (Introduce Object)
 *
 * BEFORE: Primitive obsession with conditionals
 */
class UserValidatorBefore {
    validateUser(user) {
        if (!user.name) {
            return 'Name is required';
        }

        if (user.name.length < 2) {
            return 'Name must be at least 2 characters';
        }

        if (!this.isValidEmail(user.email)) {
            return 'Invalid email format';
        }

        return true;
    }

    isValidEmail(email) {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
    }
}

/**
 * AFTER: Introduce validation result object
 */
class ValidationResult {
    constructor(isValid = true, errors = []) {
        this.isValid = isValid;
        this.errors = errors;
    }

    isValid() {
        return this.isValid;
    }

    getErrors() {
        return this.errors;
    }

    addError(error) {
        this.isValid = false;
        this.errors.push(error);
        return this;
    }
}

class UserValidatorAfter {
    validateUser(user) {
        const result = new ValidationResult();

        if (!user.name) {
            result.addError('Name is required');
        }

        if (user.name && user.name.length < 2) {
            result.addError('Name must be at least 2 characters');
        }

        if (!this.isValidEmail(user.email)) {
            result.addError('Invalid email format');
        }

        return result;
    }

    isValidEmail(email) {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
    }
}

/**
 * 41. Introduction of the statement (Introduction Statement)
 *
 * BEFORE: Magic assertion
 */
class AccountAssertion {
    constructor() {
        this.balance = 0;
    }

    withdraw(amount) {
        console.assert(amount > 0 && amount <= this.balance);
        this.balance -= amount;
    }
}

/**
 * AFTER: Introduce assertion method
 */
class AccountAssertionAfter {
    constructor() {
        this.balance = 0;
    }

    withdraw(amount) {
        this.assertValidWithdrawal(amount);
        this.balance -= amount;
    }

    assertValidWithdrawal(amount) {
        console.assert(amount > 0 && amount <= this.balance, 'Invalid withdrawal amount');
    }
}

module.exports = {
    PaymentProcessorBefore,
    PaymentProcessorAfter,
    InsuranceCalculatorBefore,
    InsuranceCalculatorAfter,
    FileProcessorBefore,
    FileProcessorAfter,
    DataProcessorBefore,
    DataProcessorAfter,
    PaymentValidatorBefore,
    PaymentValidatorAfter,
    BirdBefore,
    BirdAfter,
    EuropeanSwallow,
    AfricanSwallow,
    NorwegianBlueParrot,
    UserValidatorBefore,
    ValidationResult,
    UserValidatorAfter,
    AccountAssertion,
    AccountAssertionAfter
};
