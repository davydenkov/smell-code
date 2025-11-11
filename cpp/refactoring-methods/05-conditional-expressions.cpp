#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <cassert>
#include <stdexcept>

/**
 * 34. Decomposition of a conditional operator (Decompose Conditional)
 *
 * BEFORE: Complex conditional logic
 */
class PaymentProcessorBefore {
public:
    double calculateFee(double amount, bool isInternational, bool isPremium) {
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
};

/**
 * AFTER: Decompose conditional
 */
class PaymentProcessorAfter {
public:
    double calculateFee(double amount, bool isInternational, bool isPremium) {
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

private:
    bool isHighValueInternationalPremium(double amount, bool isInternational, bool isPremium) {
        return amount > 100 && isInternational && isPremium;
    }

    bool isHighValueInternationalStandard(double amount, bool isInternational, bool isPremium) {
        return amount > 100 && isInternational && !isPremium;
    }

    bool isLowValueInternational(double amount, bool isInternational) {
        return amount <= 100 && isInternational;
    }

    double calculateHighValueInternationalPremiumFee(double amount) {
        return amount * 0.05 + 10;
    }

    double calculateHighValueInternationalStandardFee(double amount) {
        return amount * 0.05 + 15;
    }

    double calculateLowValueInternationalFee(double amount) {
        return amount * 0.03 + 5;
    }

    double calculateDomesticFee(double amount) {
        return amount * 0.02;
    }
};

/**
 * 35. Consolidation of a conditional expression (Consolidate Conditional Expression)
 *
 * BEFORE: Multiple conditionals with same result
 */
class InsuranceCalculatorBefore {
public:
    bool isEligibleForDiscount(int age, bool isStudent, bool hasGoodRecord) {
        if (age < 25) return false;
        if (isStudent) return true;
        if (hasGoodRecord) return true;
        return false;
    }
};

/**
 * AFTER: Consolidate conditionals
 */
class InsuranceCalculatorAfter {
public:
    bool isEligibleForDiscount(int age, bool isStudent, bool hasGoodRecord) {
        return age >= 25 && (isStudent || hasGoodRecord);
    }
};

/**
 * 36. Consolidation of duplicate conditional fragments
 * (Consolidate Duplicate Conditional Fragments)
 *
 * BEFORE: Duplicate code in conditional branches
 */
class FileProcessorBefore {
public:
    void processFile(const std::string& file) {
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

private:
    bool isValidFile(const std::string& file) {
        return !file.empty();
    }

    void logProcessing(const std::string& file) {
        std::cout << "Processing: " << file << std::endl;
    }

    void validateContent(const std::string& file) {
        std::cout << "Validating: " << file << std::endl;
    }

    void saveToDatabase(const std::string& file) {
        std::cout << "Saving: " << file << std::endl;
    }

    void logError(const std::string& file) {
        std::cout << "Error processing: " << file << std::endl;
    }

    void sendNotification(const std::string& file) {
        std::cout << "Notification sent for: " << file << std::endl;
    }
};

/**
 * AFTER: Consolidate duplicate fragments
 */
class FileProcessorAfter {
public:
    void processFile(const std::string& file) {
        sendNotification(file); // Moved outside conditional

        if (isValidFile(file)) {
            logProcessing(file);
            validateContent(file);
            saveToDatabase(file);
        } else {
            logError(file);
        }
    }

private:
    bool isValidFile(const std::string& file) {
        return !file.empty();
    }

    void logProcessing(const std::string& file) {
        std::cout << "Processing: " << file << std::endl;
    }

    void validateContent(const std::string& file) {
        std::cout << "Validating: " << file << std::endl;
    }

    void saveToDatabase(const std::string& file) {
        std::cout << "Saving: " << file << std::endl;
    }

    void logError(const std::string& file) {
        std::cout << "Error processing: " << file << std::endl;
    }

    void sendNotification(const std::string& file) {
        std::cout << "Notification sent for: " << file << std::endl;
    }
};

/**
 * 37. Remove Control Flag
 *
 * BEFORE: Control flag to break out of loop
 */
class DataProcessorBefore {
public:
    std::string findPerson(const std::vector<std::unordered_map<std::string, std::string>>& people, const std::string& name) {
        bool found = false;
        std::string result;

        for (const auto& person : people) {
            if (person.at("name") == name) {
                result = person.at("name");
                found = true;
                break; // Control flag usage
            }
        }

        return found ? result : "";
    }
};

/**
 * AFTER: Remove control flag
 */
class DataProcessorAfter {
public:
    std::string findPerson(const std::vector<std::unordered_map<std::string, std::string>>& people, const std::string& name) {
        for (const auto& person : people) {
            if (person.at("name") == name) {
                return person.at("name"); // Direct return
            }
        }
        return "";
    }
};

/**
 * 38. Replacing Nested Conditional statements with a boundary operator
 * (Replace Nested Conditional with Guard Clauses)
 *
 * BEFORE: Nested conditionals
 */
class PaymentValidatorBefore {
public:
    bool isValidPayment(const std::unordered_map<std::string, std::string>& payment) {
        if (std::stod(payment.at("amount")) > 0) {
            if (!payment.at("cardNumber").empty()) {
                if (payment.at("cardNumber").length() == 16) {
                    if (isValidExpiry(payment.at("expiry"))) {
                        return true;
                    }
                }
            }
        }
        return false;
    }

private:
    bool isValidExpiry(const std::string& expiry) {
        // Simplified check
        return expiry.length() == 5; // MM/YY format
    }
};

/**
 * AFTER: Replace with guard clauses
 */
class PaymentValidatorAfter {
public:
    bool isValidPayment(const std::unordered_map<std::string, std::string>& payment) {
        if (std::stod(payment.at("amount")) <= 0) {
            return false;
        }

        if (payment.at("cardNumber").empty()) {
            return false;
        }

        if (payment.at("cardNumber").length() != 16) {
            return false;
        }

        if (!isValidExpiry(payment.at("expiry"))) {
            return false;
        }

        return true;
    }

private:
    bool isValidExpiry(const std::string& expiry) {
        // Simplified check
        return expiry.length() == 5; // MM/YY format
    }
};

/**
 * 39. Replacing a conditional operator with polymorphism (Replace Conditional with Polymorphism)
 *
 * BEFORE: Type checking with conditionals
 */
class BirdBefore {
public:
    static const int EUROPEAN = 0;
    static const int AFRICAN = 1;
    static const int NORWEGIAN_BLUE = 2;

private:
    int type;
    double voltage;
    bool isNailed;

public:
    BirdBefore(int type) : type(type), voltage(0), isNailed(false) {}

    double getSpeed() {
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

private:
    double getBaseSpeed() {
        return 10.0;
    }
};

/**
 * AFTER: Replace conditional with polymorphism
 */
class BirdAfter {
public:
    virtual ~BirdAfter() = default;
    virtual double getSpeed() = 0;

protected:
    double getBaseSpeed() {
        return 10.0;
    }
};

class EuropeanSwallow : public BirdAfter {
public:
    double getSpeed() override {
        return getBaseSpeed();
    }
};

class AfricanSwallow : public BirdAfter {
private:
    double voltage;

public:
    AfricanSwallow(double voltage) : voltage(voltage) {}

    double getSpeed() override {
        return getBaseSpeed() - voltage * 2;
    }
};

class NorwegianBlueParrot : public BirdAfter {
private:
    bool isNailed;

public:
    NorwegianBlueParrot(bool isNailed) : isNailed(isNailed) {}

    double getSpeed() override {
        return isNailed ? 0 : getBaseSpeed();
    }
};

/**
 * 40. Introduction of the object (Introduce Object)
 *
 * BEFORE: Primitive obsession with conditionals
 */
class UserValidatorBefore {
public:
    std::string validateUser(const std::unordered_map<std::string, std::string>& user) {
        if (user.at("name").empty()) {
            return "Name is required";
        }

        if (user.at("name").length() < 2) {
            return "Name must be at least 2 characters";
        }

        if (user.at("email").find('@') == std::string::npos) {
            return "Invalid email format";
        }

        return "valid";
    }
};

/**
 * AFTER: Introduce validation result object
 */
class ValidationResult {
private:
    bool isValid;
    std::vector<std::string> errors;

public:
    ValidationResult(bool isValid = true, const std::vector<std::string>& errors = {})
        : isValid(isValid), errors(errors) {}

    bool isValidResult() const {
        return isValid;
    }

    const std::vector<std::string>& getErrors() const {
        return errors;
    }

    ValidationResult& addError(const std::string& error) {
        isValid = false;
        errors.push_back(error);
        return *this;
    }
};

class UserValidatorAfter {
public:
    ValidationResult validateUser(const std::unordered_map<std::string, std::string>& user) {
        ValidationResult result;

        if (user.at("name").empty()) {
            result.addError("Name is required");
        }

        if (user.at("name").length() < 2) {
            result.addError("Name must be at least 2 characters");
        }

        if (user.at("email").find('@') == std::string::npos) {
            result.addError("Invalid email format");
        }

        return result;
    }
};

/**
 * 41. Introduction of the statement (Introduction Statement)
 *
 * BEFORE: Magic assertion
 */
class AccountAssertion {
private:
    double balance;

public:
    void withdraw(double amount) {
        assert(amount > 0 && amount <= balance);
        balance -= amount;
    }
};

/**
 * AFTER: Introduce assertion method
 */
class AccountAssertionAfter {
private:
    double balance;

public:
    void withdraw(double amount) {
        assertValidWithdrawal(amount);
        balance -= amount;
    }

private:
    void assertValidWithdrawal(double amount) {
        assert(amount > 0 && amount <= balance && "Invalid withdrawal amount");
    }
};
