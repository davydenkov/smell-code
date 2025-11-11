#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <memory>

/**
 * 1. Method Extraction (Extract Method)
 *
 * BEFORE: A method contains too much logic, making it hard to understand
 */
class OrderProcessorBefore {
private:
    void saveOrder(const std::unordered_map<std::string, double>& order, double total) {
        // Database save logic
        std::cout << "Saving order with total: " << total << std::endl;
    }

public:
    double processOrder(const std::unordered_map<std::string, double>& order) {
        // Validate order
        if (order.at("total") <= 0) {
            throw std::runtime_error("Invalid order total");
        }

        // Calculate tax
        double tax = order.at("subtotal") * 0.08;

        // Calculate shipping
        double shipping = order.at("weight") > 10 ? 15.00 : 5.00;

        // Calculate total
        double total = order.at("subtotal") + tax + shipping;

        // Save to database
        saveOrder(order, total);

        return total;
    }
};

/**
 * AFTER: Extract methods to separate concerns
 */
class OrderProcessorAfter {
private:
    void validateOrder(const std::unordered_map<std::string, double>& order) {
        if (order.at("total") <= 0) {
            throw std::runtime_error("Invalid order total");
        }
    }

    double calculateTax(const std::unordered_map<std::string, double>& order) {
        return order.at("subtotal") * 0.08;
    }

    double calculateShipping(const std::unordered_map<std::string, double>& order) {
        return order.at("weight") > 10 ? 15.00 : 5.00;
    }

    double calculateTotal(const std::unordered_map<std::string, double>& order, double tax, double shipping) {
        return order.at("subtotal") + tax + shipping;
    }

    void saveOrder(const std::unordered_map<std::string, double>& order, double total) {
        // Database save logic
        std::cout << "Saving order with total: " << total << std::endl;
    }

public:
    double processOrder(const std::unordered_map<std::string, double>& order) {
        validateOrder(order);

        double tax = calculateTax(order);
        double shipping = calculateShipping(order);
        double total = calculateTotal(order, tax, shipping);

        saveOrder(order, total);

        return total;
    }
};

/**
 * 2. Embedding a method (Inline Method)
 *
 * BEFORE: A method is too simple and adds no value
 */
class UserBefore {
private:
    std::string firstName;
    std::string lastName;

public:
    std::string getFullName() {
        return getFirstName() + " " + getLastName();
    }

    std::string getFirstName() {
        return firstName;
    }

    std::string getLastName() {
        return lastName;
    }
};

/**
 * AFTER: Inline the simple method
 */
class UserAfter {
private:
    std::string firstName;
    std::string lastName;

public:
    std::string getFullName() {
        return firstName + " " + lastName;
    }

    // getFirstName() and getLastName() methods removed
};
