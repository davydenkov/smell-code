import java.util.*;

/**
 * 34. Decomposition of a conditional operator (Decompose Conditional)
 *
 * BEFORE: Complex conditional logic
 */
class PaymentProcessorBefore {
    public double calculateFee(double amount, boolean isInternational, boolean isPremium) {
        double fee = 0;
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
    public double calculateFee(double amount, boolean isInternational, boolean isPremium) {
        if (isHighValueInternationalPremium(amount, isInternational, isPremium)) {
            return calculateHighValueInternationalPremiumFee(amount);
        } else if (isHighValueInternationalStandard(amount, isInternational, isPremium)) {
            return calculateHighValueInternationalStandardFee(amount);
        } else if (isLowValueInternational(amount, isInternational)) {
            return calculateLowValueInternationalFee(amount);
        } else {
            return calculateDomesticFee(amount);
        }
    }

    private boolean isHighValueInternationalPremium(double amount, boolean isInternational, boolean isPremium) {
        return amount > 100 && isInternational && isPremium;
    }

    private boolean isHighValueInternationalStandard(double amount, boolean isInternational, boolean isPremium) {
        return amount > 100 && isInternational && !isPremium;
    }

    private boolean isLowValueInternational(double amount, boolean isInternational) {
        return amount <= 100 && isInternational;
    }

    private double calculateHighValueInternationalPremiumFee(double amount) {
        return amount * 0.05 + 10;
    }

    private double calculateHighValueInternationalStandardFee(double amount) {
        return amount * 0.05 + 15;
    }

    private double calculateLowValueInternationalFee(double amount) {
        return amount * 0.03 + 5;
    }

    private double calculateDomesticFee(double amount) {
        return amount * 0.02;
    }
}

/**
 * 35. Consolidation of a conditional expression (Consolidate Conditional Expression)
 *
 * BEFORE: Multiple conditionals with same result
 */
class InsuranceCalculatorBefore {
    public boolean isEligibleForDiscount(int age, boolean isStudent, boolean hasGoodRecord) {
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
    public boolean isEligibleForDiscount(int age, boolean isStudent, boolean hasGoodRecord) {
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
    public void processFile(String file) {
        if (isValidFile(file)) {
            logProcessing(file);
            validateContent(file);
            saveToDatabase(file);
            sendNotification(file);
        } else {
            logError(file);
            sendNotification(file); // Duplicate
        }
    }

    private void sendNotification(String file) {
        // Send notification logic
    }

    private boolean isValidFile(String file) { return true; }
    private void logProcessing(String file) {}
    private void validateContent(String file) {}
    private void saveToDatabase(String file) {}
    private void logError(String file) {}
}

/**
 * AFTER: Consolidate duplicate fragments
 */
class FileProcessorAfter {
    public void processFile(String file) {
        sendNotification(file); // Moved outside conditional

        if (isValidFile(file)) {
            logProcessing(file);
            validateContent(file);
            saveToDatabase(file);
        } else {
            logError(file);
        }
    }

    private void sendNotification(String file) {
        // Send notification logic
    }

    private boolean isValidFile(String file) { return true; }
    private void logProcessing(String file) {}
    private void validateContent(String file) {}
    private void saveToDatabase(String file) {}
    private void logError(String file) {}
}

/**
 * 37. Remove Control Flag
 *
 * BEFORE: Control flag to break out of loop
 */
class DataProcessorBefore {
    public Map<String, Object> findPerson(List<Map<String, Object>> people, String name) {
        Map<String, Object> found = null;
        for (Map<String, Object> person : people) {
            if (person.get("name").equals(name)) {
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
    public Map<String, Object> findPerson(List<Map<String, Object>> people, String name) {
        for (Map<String, Object> person : people) {
            if (person.get("name").equals(name)) {
                return person; // Direct return
            }
        }
        return null;
    }
}

/**
 * 38. Replacing Nested Conditional statements with a boundary operator
 * (Replace Nested Conditional with Guard Clauses)
 *
 * BEFORE: Nested conditionals
 */
class PaymentValidatorBefore {
    public boolean isValidPayment(Map<String, Object> payment) {
        if ((Double) payment.get("amount") > 0) {
            if (payment.get("cardNumber") != null) {
                if (((String) payment.get("cardNumber")).length() == 16) {
                    if (isValidExpiry((String) payment.get("expiry"))) {
                        return true;
                    }
                }
            }
        }
        return false;
    }

    private boolean isValidExpiry(String expiry) {
        // Simplified validation
        return expiry != null && !expiry.isEmpty();
    }
}

/**
 * AFTER: Replace with guard clauses
 */
class PaymentValidatorAfter {
    public boolean isValidPayment(Map<String, Object> payment) {
        if ((Double) payment.get("amount") <= 0) {
            return false;
        }

        if (payment.get("cardNumber") == null) {
            return false;
        }

        if (((String) payment.get("cardNumber")).length() != 16) {
            return false;
        }

        if (!isValidExpiry((String) payment.get("expiry"))) {
            return false;
        }

        return true;
    }

    private boolean isValidExpiry(String expiry) {
        // Simplified validation
        return expiry != null && !expiry.isEmpty();
    }
}

/**
 * 39. Replacing a conditional operator with polymorphism (Replace Conditional with Polymorphism)
 *
 * BEFORE: Type checking with conditionals
 */
class BirdBefore {
    public static final String EUROPEAN = "european";
    public static final String AFRICAN = "african";
    public static final String NORWEGIAN_BLUE = "norwegian_blue";

    private String type;
    private int voltage;
    private boolean isNailed;

    public BirdBefore(String type) {
        this.type = type;
    }

    public double getSpeed() {
        switch (type) {
            case EUROPEAN:
                return getBaseSpeed();
            case AFRICAN:
                return getBaseSpeed() - voltage * 2;
            case NORWEGIAN_BLUE:
                return isNailed ? 0 : getBaseSpeed();
            default:
                return getBaseSpeed();
        }
    }

    private double getBaseSpeed() {
        return 10;
    }
}

/**
 * AFTER: Replace conditional with polymorphism
 */
abstract class BirdAfter {
    public abstract double getSpeed();

    protected double getBaseSpeed() {
        return 10;
    }
}

class EuropeanSwallow extends BirdAfter {
    public double getSpeed() {
        return getBaseSpeed();
    }
}

class AfricanSwallow extends BirdAfter {
    private int voltage;

    public AfricanSwallow(int voltage) {
        this.voltage = voltage;
    }

    public double getSpeed() {
        return getBaseSpeed() - voltage * 2;
    }
}

class NorwegianBlueParrot extends BirdAfter {
    private boolean isNailed;

    public NorwegianBlueParrot(boolean isNailed) {
        this.isNailed = isNailed;
    }

    public double getSpeed() {
        return isNailed ? 0 : getBaseSpeed();
    }
}

/**
 * 40. Introduction of the object (Introduce Object)
 *
 * BEFORE: Primitive obsession with conditionals
 */
class UserValidatorBefore {
    public Object validateUser(Map<String, Object> user) {
        if (user.get("name") == null || ((String) user.get("name")).isEmpty()) {
            return "Name is required";
        }

        if (((String) user.get("name")).length() < 2) {
            return "Name must be at least 2 characters";
        }

        if (!isValidEmail((String) user.get("email"))) {
            return "Invalid email format";
        }

        return true;
    }

    private boolean isValidEmail(String email) {
        // Simplified email validation
        return email != null && email.contains("@");
    }
}

/**
 * AFTER: Introduce validation result object
 */
class ValidationResult {
    private boolean isValid;
    private List<String> errors;

    public ValidationResult(boolean isValid, List<String> errors) {
        this.isValid = isValid;
        this.errors = errors != null ? errors : new ArrayList<>();
    }

    public ValidationResult() {
        this(true, new ArrayList<>());
    }

    public boolean isValid() {
        return isValid;
    }

    public List<String> getErrors() {
        return new ArrayList<>(errors);
    }

    public ValidationResult addError(String error) {
        this.isValid = false;
        this.errors.add(error);
        return this;
    }
}

class UserValidatorAfter {
    public ValidationResult validateUser(Map<String, Object> user) {
        ValidationResult result = new ValidationResult();

        if (user.get("name") == null || ((String) user.get("name")).isEmpty()) {
            result.addError("Name is required");
        }

        if (user.get("name") != null && ((String) user.get("name")).length() < 2) {
            result.addError("Name must be at least 2 characters");
        }

        if (!isValidEmail((String) user.get("email"))) {
            result.addError("Invalid email format");
        }

        return result;
    }

    private boolean isValidEmail(String email) {
        // Simplified email validation
        return email != null && email.contains("@");
    }
}

/**
 * 41. Introduction of the statement (Introduction Statement)
 *
 * BEFORE: Magic assertion
 */
class AccountAssertion {
    private double balance;

    public void withdraw(double amount) {
        assert amount > 0 && amount <= balance;
        balance -= amount;
    }
}

/**
 * AFTER: Introduce assertion method
 */
class AccountAssertionAfter {
    private double balance;

    public void withdraw(double amount) {
        assertValidWithdrawal(amount);
        balance -= amount;
    }

    private void assertValidWithdrawal(double amount) {
        assert amount > 0 && amount <= balance : "Invalid withdrawal amount";
    }
}
