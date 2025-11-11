#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <memory>
#include <stdexcept>

/**
 * 9. Substitution Algorithm
 *
 * BEFORE: Complex algorithm that can be simplified
 */
class PricingServiceBefore {
public:
    double calculatePrice(const std::vector<std::unordered_map<std::string, std::string>>& items) {
        double total = 0;
        for (const auto& item : items) {
            if (item.at("type") == "book") {
                total += std::stod(item.at("price")) * 0.9; // 10% discount for books
            } else if (item.at("type") == "electronics") {
                total += std::stod(item.at("price")) * 1.1; // 10% markup for electronics
            } else {
                total += std::stod(item.at("price"));
            }
        }
        return total;
    }
};

/**
 * AFTER: Substitute with a simpler algorithm
 */
class PricingServiceAfter {
private:
    std::unordered_map<std::string, double> discounts = {
        {"book", 0.9},
        {"electronics", 1.1},
        {"default", 1.0}
    };

public:
    double calculatePrice(const std::vector<std::unordered_map<std::string, std::string>>& items) {
        double total = 0;
        for (const auto& item : items) {
            double multiplier = discounts.count(item.at("type")) ?
                               discounts.at(item.at("type")) : discounts.at("default");
            total += std::stod(item.at("price")) * multiplier;
        }
        return total;
    }
};

/**
 * 10. Moving functions between objects (Move Method)
 *
 * BEFORE: Method in wrong class
 */
class AccountBefore {
private:
    double balance;

public:
    AccountBefore(double balance) : balance(balance) {}

    double getBalance() const {
        return balance;
    }

    // This method belongs in Bank class, not Account
    bool transferTo(AccountBefore& target, double amount) {
        if (balance >= amount) {
            balance -= amount;
            target.balance += amount;
            return true;
        }
        return false;
    }
};

/**
 * AFTER: Move method to appropriate class
 */
class AccountAfter {
private:
    double balance;

public:
    AccountAfter(double balance) : balance(balance) {}

    double getBalance() const {
        return balance;
    }

    void decreaseBalance(double amount) {
        balance -= amount;
    }

    void increaseBalance(double amount) {
        balance += amount;
    }
};

class Bank {
public:
    bool transfer(AccountAfter& from, AccountAfter& to, double amount) {
        if (from.getBalance() >= amount) {
            from.decreaseBalance(amount);
            to.increaseBalance(amount);
            return true;
        }
        return false;
    }
};

/**
 * 11. Moving the field (Move Field)
 *
 * BEFORE: Field in wrong class
 */
class CustomerBefore {
private:
    std::string name;
    std::unordered_map<std::string, std::string> address; // This should be in Address class

public:
    CustomerBefore(const std::string& name, const std::string& street, const std::string& city, const std::string& zipCode) {
        this->name = name;
        address = {{"street", street}, {"city", city}, {"zipCode", zipCode}};
    }

    std::string getAddress() const {
        return address.at("street") + ", " + address.at("city") + " " + address.at("zipCode");
    }
};

/**
 * AFTER: Move field to dedicated class
 */
class Address {
private:
    std::string street;
    std::string city;
    std::string zipCode;

public:
    Address(const std::string& street, const std::string& city, const std::string& zipCode)
        : street(street), city(city), zipCode(zipCode) {}

    std::string getFullAddress() const {
        return street + ", " + city + " " + zipCode;
    }
};

class CustomerAfter {
private:
    std::string name;
    Address address;

public:
    CustomerAfter(const std::string& name, const Address& address)
        : name(name), address(address) {}

    std::string getAddress() const {
        return address.getFullAddress();
    }
};

/**
 * 12. Class Allocation (Extract Class)
 *
 * BEFORE: Class has too many responsibilities
 */
class PersonBefore {
private:
    std::string name;
    std::string phoneNumber;
    std::string officeAreaCode;
    std::string officeNumber;

public:
    std::string getTelephoneNumber() const {
        return "(" + officeAreaCode + ") " + officeNumber;
    }
};

/**
 * AFTER: Extract telephone number to separate class
 */
class TelephoneNumber {
private:
    std::string areaCode;
    std::string number;

public:
    TelephoneNumber(const std::string& areaCode, const std::string& number)
        : areaCode(areaCode), number(number) {}

    std::string getTelephoneNumber() const {
        return "(" + areaCode + ") " + number;
    }
};

class PersonAfter {
private:
    std::string name;
    std::string phoneNumber;
    TelephoneNumber officeTelephone;

public:
    PersonAfter(const std::string& name) : name(name) {}

    std::string getOfficeTelephone() const {
        return officeTelephone.getTelephoneNumber();
    }

    void setOfficeTelephone(const TelephoneNumber& telephone) {
        officeTelephone = telephone;
    }
};

/**
 * 13. Embedding a class (Inline Class)
 *
 * BEFORE: Unnecessary class with single responsibility
 */
class OrderProcessorBefore {
private:
    std::unique_ptr<class OrderValidator> validator;

public:
    OrderProcessorBefore() {
        validator = std::make_unique<OrderValidator>();
    }

    void process(const std::unordered_map<std::string, double>& order) {
        if (validator->isValid(order)) {
            // Process order
            std::cout << "Processing valid order" << std::endl;
        }
    }
};

class OrderValidator {
public:
    bool isValid(const std::unordered_map<std::string, double>& order) {
        return order.at("total") > 0;
    }
};

/**
 * AFTER: Inline the class
 */
class OrderProcessorAfter {
public:
    void process(const std::unordered_map<std::string, double>& order) {
        if (isValidOrder(order)) {
            // Process order
            std::cout << "Processing valid order" << std::endl;
        }
    }

private:
    bool isValidOrder(const std::unordered_map<std::string, double>& order) {
        return order.at("total") > 0;
    }
};

/**
 * 14. Hiding delegation (Hide Delegate)
 *
 * BEFORE: Client has to know about delegation
 */
class DepartmentBefore {
private:
    std::string manager;

public:
    DepartmentBefore(const std::string& manager) : manager(manager) {}

    std::string getManager() const {
        return manager;
    }
};

class PersonDelegate {
private:
    DepartmentBefore department;

public:
    PersonDelegate(const DepartmentBefore& department) : department(department) {}

    const DepartmentBefore& getDepartment() const {
        return department;
    }
};

// Client code
// std::string manager = person.getDepartment().getManager();

/**
 * AFTER: Hide the delegation
 */
class DepartmentAfter {
private:
    std::string manager;

public:
    DepartmentAfter(const std::string& manager) : manager(manager) {}

    std::string getManager() const {
        return manager;
    }
};

class PersonAfter {
private:
    DepartmentAfter department;

public:
    PersonAfter(const DepartmentAfter& department) : department(department) {}

    const DepartmentAfter& getDepartment() const {
        return department;
    }

    std::string getManager() const {
        return department.getManager();
    }
};

// Client code - much cleaner
// std::string manager = person.getManager();

/**
 * 15. Removing the intermediary (Remove Middle Man)
 *
 * BEFORE: Too much delegation
 */
class PersonWithMiddleMan {
private:
    DepartmentAfter department;

public:
    PersonWithMiddleMan(const DepartmentAfter& department) : department(department) {}

    const DepartmentAfter& getDepartment() const {
        return department;
    }

    std::string getManager() const {
        return department.getManager();
    }

    std::string getDepartmentName() const {
        return "Engineering"; // Simplified for example
    }
};

/**
 * AFTER: Remove middle man if delegation is too heavy
 */
class PersonDirect {
private:
    DepartmentAfter department;
    std::string manager; // Direct reference

public:
    PersonDirect(const std::string& manager) : manager(manager) {}

    std::string getManager() const {
        return manager;
    }

    const DepartmentAfter& getDepartment() const {
        return department;
    }
};

/**
 * 16. Introduction of an external method (Introduce Foreign Method)
 *
 * BEFORE: Using external class method in wrong place
 */
class ReportGeneratorBefore {
public:
    void generateReport() {
        // Simplified for example - would use DateTime in real scenario
        std::string nextMonth = "2024-02"; // Foreign method usage

        // Generate report for next month
        std::cout << "Generating report for: " << nextMonth << std::endl;
    }
};

/**
 * AFTER: Introduce foreign method
 */
class ReportGeneratorAfter {
public:
    void generateReport() {
        std::string nextMonth = nextMonthStr();
        // Generate report for next month
        std::cout << "Generating report for: " << nextMonth << std::endl;
    }

private:
    std::string nextMonthStr() {
        // Simplified implementation
        return "2024-02";
    }
};

/**
 * 17. The introduction of local extension (Introduce Local Extension)
 *
 * BEFORE: Adding methods to external class (not possible)
 */
class DateUtil {
public:
    static std::string nextMonth() {
        return "2024-02";
    }

    static std::string previousMonth() {
        return "2024-01";
    }
};

/**
 * AFTER: Create local extension class
 */
class DateTimeExtension {
private:
    std::string date;

public:
    DateTimeExtension(const std::string& date) : date(date) {}

    std::string nextMonth() {
        // Simplified implementation
        return "2024-02";
    }

    std::string previousMonth() {
        // Simplified implementation
        return "2024-01";
    }

    std::string getDate() const {
        return date;
    }
};
