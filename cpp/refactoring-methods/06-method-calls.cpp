#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <memory>
#include <stdexcept>

/**
 * 42. Renaming a method (Rename Method)
 *
 * BEFORE: Poorly named method
 */
class CalculatorBefore {
public:
    double calc(double a, double b) { // Unclear name
        return a + b;
    }
};

/**
 * AFTER: Rename method to be more descriptive
 */
class CalculatorAfter {
public:
    double add(double a, double b) {
        return a + b;
    }
};

/**
 * 43. Adding a parameter (Add Parameter)
 *
 * BEFORE: Method missing required parameter
 */
class EmailSenderBefore {
public:
    void sendEmail(const std::string& to, const std::string& subject,
                  const std::string& body) {
        std::string priority = "normal";
        // Send logic
        std::cout << "Sending email with priority: " << priority << std::endl;
    }
};

/**
 * AFTER: Add parameter
 */
class EmailSenderAfter {
public:
    void sendEmail(const std::string& to, const std::string& subject,
                  const std::string& body, const std::string& priority = "normal") {
        // Send logic with priority
        std::cout << "Sending email with priority: " << priority << std::endl;
    }
};

/**
 * 44. Deleting a parameter (Remove Parameter)
 *
 * BEFORE: Unnecessary parameter
 */
class ReportGeneratorBefore {
public:
    void generateReport(const std::vector<std::string>& data, const std::string& format, bool includeHeader = true) {
        bool shouldIncludeHeader = (format == "html") ? true : includeHeader;
        // Generate report
        std::cout << "Generating report with header: " << (shouldIncludeHeader ? "yes" : "no") << std::endl;
    }
};

/**
 * AFTER: Remove unnecessary parameter
 */
class ReportGeneratorAfter {
public:
    void generateReport(const std::vector<std::string>& data, const std::string& format) {
        bool includeHeader = (format == "html");
        // Generate report
        std::cout << "Generating report with header: " << (includeHeader ? "yes" : "no") << std::endl;
    }
};

/**
 * 45. Separation of Query and Modifier (Separate Query from Modifier)
 *
 * BEFORE: Method that both queries and modifies
 */
class BankAccountBefore {
private:
    double balance = 0;

public:
    bool withdraw(double amount) {
        if (balance >= amount) {
            balance -= amount;
            return true;
        }
        return false;
    }
};

/**
 * AFTER: Separate query from modifier
 */
class BankAccountAfter {
private:
    double balance = 0;

public:
    bool canWithdraw(double amount) const {
        return balance >= amount;
    }

    bool withdraw(double amount) {
        if (canWithdraw(amount)) {
            balance -= amount;
            return true;
        }
        return false;
    }
};

/**
 * 46. Parameterization of the method (Parameterize Method)
 *
 * BEFORE: Similar methods with different values
 */
class ReportGeneratorParamBefore {
public:
    void generateWeeklyReport() {
        generateReport(7);
    }

    void generateMonthlyReport() {
        generateReport(30);
    }

    void generateQuarterlyReport() {
        generateReport(90);
    }

private:
    void generateReport(int days) {
        // Generate report for specified days
        std::cout << "Generating report for " << days << " days" << std::endl;
    }
};

/**
 * AFTER: Parameterize method
 */
class ReportGeneratorParamAfter {
public:
    void generateReport(int days) {
        // Generate report for specified days
        std::cout << "Generating report for " << days << " days" << std::endl;
    }

    void generateWeeklyReport() {
        generateReport(7);
    }

    void generateMonthlyReport() {
        generateReport(30);
    }

    void generateQuarterlyReport() {
        generateReport(90);
    }
};

/**
 * 47. Replacing a parameter with explicit methods (Replace Parameter with Explicit Methods)
 *
 * BEFORE: Parameter determines behavior
 */
class EmployeeExplicitBefore {
public:
    static const int ENGINEER = 0;
    static const int SALESMAN = 1;
    static const int MANAGER = 2;

private:
    int type;

public:
    EmployeeExplicitBefore(int type) : type(type) {}

    double getSalary(double baseSalary) {
        switch (type) {
            case ENGINEER: return baseSalary * 1.0;
            case SALESMAN: return baseSalary * 1.1;
            case MANAGER: return baseSalary * 1.2;
            default: return baseSalary;
        }
    }
};

/**
 * AFTER: Replace parameter with explicit methods
 */
class EmployeeExplicitAfter {
public:
    double getEngineerSalary(double baseSalary) {
        return baseSalary * 1.0;
    }

    double getSalesmanSalary(double baseSalary) {
        return baseSalary * 1.1;
    }

    double getManagerSalary(double baseSalary) {
        return baseSalary * 1.2;
    }
};

/**
 * 48. Save the Whole Object
 *
 * BEFORE: Passing individual fields
 */
class OrderWholeBefore {
private:
    std::unordered_map<std::string, std::string> customer;

public:
    OrderWholeBefore(const std::string& customerName, const std::string& customerAddress) {
        customer = {{"name", customerName}, {"address", customerAddress}};
    }

    double calculateShipping() {
        return getShippingCost(customer.at("name"), customer.at("address"));
    }

private:
    double getShippingCost(const std::string& name, const std::string& address) {
        // Calculate based on name and address
        return 10.0;
    }
};

/**
 * AFTER: Pass whole object
 */
class Customer {
private:
    std::string name;
    std::string address;

public:
    Customer(const std::string& name, const std::string& address)
        : name(name), address(address) {}

    std::string getName() const {
        return name;
    }

    std::string getAddress() const {
        return address;
    }
};

class OrderWholeAfter {
private:
    Customer customer;

public:
    OrderWholeAfter(const Customer& customer) : customer(customer) {}

    double calculateShipping() {
        return getShippingCost(customer);
    }

private:
    double getShippingCost(const Customer& customer) {
        // Calculate based on customer object
        return 10.0;
    }
};

/**
 * 49. Replacing a parameter with a method call (Replace Parameter with Method)
 *
 * BEFORE: Parameter calculated outside method
 */
class DiscountCalculatorParamBefore {
public:
    double calculateDiscount(double price, const std::string& customerType) {
        // customerType passed in
        return price * getDiscountRate(customerType);
    }

private:
    double getDiscountRate(const std::string& customerType) {
        if (customerType == "premium") return 0.1;
        if (customerType == "regular") return 0.05;
        return 0.0;
    }
};

class OrderParamBefore {
private:
    std::string customerType;

public:
    double getDiscountedPrice(double price) {
        DiscountCalculatorParamBefore calculator;
        return calculator.calculateDiscount(price, customerType);
    }
};

/**
 * AFTER: Replace parameter with method call
 */
class DiscountCalculatorParamAfter {
public:
    double calculateDiscount(double price, const std::string& customerType) {
        return price * getDiscountRate(customerType);
    }

private:
    double getDiscountRate(const std::string& customerType) {
        if (customerType == "premium") return 0.1;
        if (customerType == "regular") return 0.05;
        return 0.0;
    }
};

class OrderParamAfter {
private:
    std::string customerType;

public:
    double getDiscountedPrice(double price) {
        DiscountCalculatorParamAfter calculator;
        return calculator.calculateDiscount(price, customerType);
    }
};

/**
 * 50. Introduction of the boundary object (Introduce Parameter Object)
 *
 * BEFORE: Multiple parameters
 */
class TemperatureRangeBefore {
public:
    bool withinRange(double minTemp, double maxTemp, double currentTemp) {
        return currentTemp >= minTemp && currentTemp <= maxTemp;
    }

    double getAverageTemp(double minTemp, double maxTemp) {
        return (minTemp + maxTemp) / 2;
    }
};

/**
 * AFTER: Introduce parameter object
 */
class TemperatureRange {
private:
    double minTemp;
    double maxTemp;

public:
    TemperatureRange(double minTemp, double maxTemp)
        : minTemp(minTemp), maxTemp(maxTemp) {}

    double getMinTemp() const {
        return minTemp;
    }

    double getMaxTemp() const {
        return maxTemp;
    }

    bool withinRange(double currentTemp) const {
        return currentTemp >= minTemp && currentTemp <= maxTemp;
    }

    double getAverageTemp() const {
        return (minTemp + maxTemp) / 2;
    }
};

/**
 * 51. Removing the Value Setting Method
 *
 * BEFORE: Setter that's not needed
 */
class SensorBefore {
private:
    double temperature;

public:
    SensorBefore(double temperature) : temperature(temperature) {}

    double getTemperature() const {
        return temperature;
    }

    void setTemperature(double temperature) { // Not needed if immutable
        this->temperature = temperature;
    }
};

/**
 * AFTER: Remove setter for immutable object
 */
class SensorAfter {
private:
    double temperature;

public:
    SensorAfter(double temperature) : temperature(temperature) {}

    double getTemperature() const {
        return temperature;
    }

    // setTemperature removed
};

/**
 * 52. Hiding a method (Hide Method)
 *
 * BEFORE: Public method that should be private
 */
class DataProcessorHideBefore {
public:
    bool validateData(const std::vector<std::string>& data) { // Should be private
        return !data.empty();
    }

    void processData(const std::vector<std::string>& data) {
        if (validateData(data)) {
            // Process data
            std::cout << "Processing data" << std::endl;
        }
    }
};

/**
 * AFTER: Hide method
 */
class DataProcessorHideAfter {
private:
    bool validateData(const std::vector<std::string>& data) {
        return !data.empty();
    }

public:
    void processData(const std::vector<std::string>& data) {
        if (validateData(data)) {
            // Process data
            std::cout << "Processing data" << std::endl;
        }
    }
};

/**
 * 53. Replacing the constructor with the factory method (Replace Constructor with Factory Method)
 *
 * BEFORE: Complex constructor
 */
class ComplexObjectBefore {
private:
    std::string type;
    std::unordered_map<std::string, std::string> config;

public:
    ComplexObjectBefore(const std::string& type, const std::unordered_map<std::string, std::string>& config = {}) {
        this->type = type;
        this->config = config;

        if (type == "database") {
            this->config["host"] = config.count("host") ? config.at("host") : "localhost";
            this->config["port"] = config.count("port") ? config.at("port") : "3306";
        } else if (type == "file") {
            this->config["path"] = config.count("path") ? config.at("path") : "/tmp";
            this->config["format"] = config.count("format") ? config.at("format") : "json";
        }
    }
};

/**
 * AFTER: Replace constructor with factory method
 */
class ComplexObjectAfter {
private:
    std::string type;
    std::unordered_map<std::string, std::string> config;

    ComplexObjectAfter(const std::string& type, const std::unordered_map<std::string, std::string>& config)
        : type(type), config(config) {}

public:
    static std::unique_ptr<ComplexObjectAfter> createDatabaseConnection(
        const std::unordered_map<std::string, std::string>& config = {}) {
        std::unordered_map<std::string, std::string> finalConfig = config;
        finalConfig["host"] = config.count("host") ? config.at("host") : "localhost";
        finalConfig["port"] = config.count("port") ? config.at("port") : "3306";
        return std::make_unique<ComplexObjectAfter>("database", finalConfig);
    }

    static std::unique_ptr<ComplexObjectAfter> createFileHandler(
        const std::unordered_map<std::string, std::string>& config = {}) {
        std::unordered_map<std::string, std::string> finalConfig = config;
        finalConfig["path"] = config.count("path") ? config.at("path") : "/tmp";
        finalConfig["format"] = config.count("format") ? config.at("format") : "json";
        return std::make_unique<ComplexObjectAfter>("file", finalConfig);
    }
};

/**
 * 54. Encapsulation of top-down type conversion (Encapsulate Downcast)
 *
 * BEFORE: Downcast in client code
 */
class Shape {
public:
    virtual ~Shape() = default;
    virtual std::string getType() const = 0;
};

class Circle : public Shape {
public:
    std::string getType() const override { return "circle"; }
};

class Square : public Shape {
public:
    std::string getType() const override { return "square"; }
};

class ShapeCollectionBefore {
private:
    std::vector<std::shared_ptr<Shape>> shapes;

public:
    void addShape(std::shared_ptr<Shape> shape) {
        shapes.push_back(shape);
    }

    const std::vector<std::shared_ptr<Shape>>& getShapes() const {
        return shapes;
    }
};

// Client code
// auto circles = std::vector<std::shared_ptr<Circle>>();
// for (auto& shape : collection.getShapes()) {
//     if (auto circle = std::dynamic_pointer_cast<Circle>(shape)) {
//         circles.push_back(circle); // Downcast
//     }
// }

/**
 * AFTER: Encapsulate downcast
 */
class ShapeCollectionAfter {
private:
    std::vector<std::shared_ptr<Shape>> shapes;

public:
    void addShape(std::shared_ptr<Shape> shape) {
        shapes.push_back(shape);
    }

    std::vector<std::shared_ptr<Circle>> getCircles() const {
        std::vector<std::shared_ptr<Circle>> circles;
        for (auto& shape : shapes) {
            if (auto circle = std::dynamic_pointer_cast<Circle>(shape)) {
                circles.push_back(circle);
            }
        }
        return circles;
    }

    std::vector<std::shared_ptr<Square>> getSquares() const {
        std::vector<std::shared_ptr<Square>> squares;
        for (auto& shape : shapes) {
            if (auto square = std::dynamic_pointer_cast<Square>(shape)) {
                squares.push_back(square);
            }
        }
        return squares;
    }
};

/**
 * 55. Replacing the error code with an exceptional situation (Replace Error Code with Exception)
 *
 * BEFORE: Error codes
 */
class FileReaderErrorBefore {
public:
    static const int FILE_NOT_FOUND = 1;
    static const int PERMISSION_DENIED = 2;

    std::string readFile(const std::string& filename) {
        if (filename == "nonexistent") {
            return std::to_string(FILE_NOT_FOUND);
        }

        if (filename == "denied") {
            return std::to_string(PERMISSION_DENIED);
        }

        return "file content";
    }
};

// Client code
// FileReaderErrorBefore reader;
// std::string result = reader.readFile("test.txt");
// if (result == std::to_string(FileReaderErrorBefore::FILE_NOT_FOUND)) {
//     // Handle error
// } else if (result == std::to_string(FileReaderErrorBefore::PERMISSION_DENIED)) {
//     // Handle error
// } else {
//     // Use content
// }

/**
 * AFTER: Replace error codes with exceptions
 */
class FileNotFoundException : public std::runtime_error {
public:
    FileNotFoundException(const std::string& filename)
        : std::runtime_error("File not found: " + filename) {}
};

class PermissionDeniedException : public std::runtime_error {
public:
    PermissionDeniedException(const std::string& filename)
        : std::runtime_error("Permission denied: " + filename) {}
};

class FileReaderExceptionAfter {
public:
    std::string readFile(const std::string& filename) {
        if (filename == "nonexistent") {
            throw FileNotFoundException(filename);
        }

        if (filename == "denied") {
            throw PermissionDeniedException(filename);
        }

        return "file content";
    }
};

// Client code
// try {
//     FileReaderExceptionAfter reader;
//     std::string content = reader.readFile("test.txt");
//     // Use content
// } catch (const FileNotFoundException& e) {
//     // Handle file not found
// } catch (const PermissionDeniedException& e) {
//     // Handle permission denied
// }

/**
 * 56. Replacing an exceptional situation with a check (Replace Exception with Test)
 *
 * BEFORE: Using exception for control flow
 */
class StackExceptionBefore {
private:
    std::vector<int> items;

public:
    int pop() {
        if (items.empty()) {
            throw std::runtime_error("Stack is empty");
        }
        int item = items.back();
        items.pop_back();
        return item;
    }
};

class EmptyStackException : public std::runtime_error {
public:
    EmptyStackException() : std::runtime_error("Stack is empty") {}
};

// Client code
// StackExceptionBefore stack;
// try {
//     int item = stack.pop();
// } catch (const EmptyStackException& e) {
//     item = 0; // Default value
// }

/**
 * AFTER: Replace exception with test
 */
class StackTestAfter {
private:
    std::vector<int> items;

public:
    bool isEmpty() const {
        return items.empty();
    }

    int pop() {
        int item = items.back();
        items.pop_back();
        return item;
    }
};

// Client code
// StackTestAfter stack;
// int item = stack.isEmpty() ? 0 : stack.pop();
