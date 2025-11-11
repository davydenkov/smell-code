#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <memory>

/**
 * 18. Self-Encapsulate Field
 *
 * BEFORE: Direct field access
 */
class PersonBefore {
public:
    std::string name; // Direct access

    std::string getName() const {
        return name;
    }

    void setName(const std::string& name) {
        this->name = name;
    }
};

/**
 * AFTER: Self-encapsulate field
 */
class PersonAfter {
private:
    std::string name;

public:
    std::string getName() const {
        return name;
    }

    void setName(const std::string& name) {
        this->name = name;
    }
};

/**
 * 19. Replacing the data value with an object (Replace Data Value with Object)
 *
 * BEFORE: Primitive data type that should be an object
 */
class OrderBefore {
private:
    std::string customer; // Just a string

public:
    std::string getCustomerName() const {
        return customer;
    }

    void setCustomer(const std::string& customer) {
        this->customer = customer;
    }
};

/**
 * AFTER: Replace with object
 */
class Customer {
private:
    std::string name;

public:
    Customer(const std::string& name) : name(name) {}

    std::string getName() const {
        return name;
    }
};

class OrderAfter {
private:
    Customer customer;

public:
    const Customer& getCustomer() const {
        return customer;
    }

    void setCustomer(const Customer& customer) {
        this->customer = customer;
    }

    std::string getCustomerName() const {
        return customer.getName();
    }
};

/**
 * 20. Replacing the value with a reference (Change Value to Reference)
 *
 * BEFORE: Multiple instances of same object
 */
class CustomerValue {
private:
    std::string name;

public:
    CustomerValue(const std::string& name) : name(name) {}

    std::string getName() const {
        return name;
    }
};

class OrderValue {
private:
    CustomerValue customer; // New instance for each order

public:
    OrderValue(const std::string& customerName) {
        customer = CustomerValue(customerName);
    }
};

/**
 * AFTER: Use reference to single instance
 */
class CustomerReference {
private:
    std::string name;
    static std::unordered_map<std::string, std::shared_ptr<CustomerReference>> instances;

    CustomerReference(const std::string& name) : name(name) {}

public:
    static std::shared_ptr<CustomerReference> create(const std::string& name) {
        if (instances.find(name) == instances.end()) {
            instances[name] = std::shared_ptr<CustomerReference>(new CustomerReference(name));
        }
        return instances[name];
    }

    std::string getName() const {
        return name;
    }
};

std::unordered_map<std::string, std::shared_ptr<CustomerReference>> CustomerReference::instances;

class OrderReference {
private:
    std::shared_ptr<CustomerReference> customer;

public:
    OrderReference(const std::string& customerName) {
        customer = CustomerReference::create(customerName);
    }
};

/**
 * 21. Replacing a reference with a value (Change Reference to Value)
 *
 * BEFORE: Unnecessary reference when value would suffice
 */
class CurrencyReference {
private:
    std::string code;

public:
    CurrencyReference(const std::string& code) : code(code) {}

    std::string getCode() const {
        return code;
    }
};

class ProductReference {
private:
    double price;
    CurrencyReference currency; // Reference object

public:
    ProductReference(double price, const CurrencyReference& currency)
        : price(price), currency(currency) {}
};

/**
 * AFTER: Use value object instead
 */
class CurrencyValue {
private:
    std::string code;

public:
    CurrencyValue(const std::string& code) : code(code) {}

    std::string getCode() const {
        return code;
    }
};

class ProductValue {
private:
    double price;
    std::string currencyCode; // Just the value

public:
    ProductValue(double price, const std::string& currencyCode)
        : price(price), currencyCode(currencyCode) {}

    std::string getCurrencyCode() const {
        return currencyCode;
    }
};

/**
 * 22. Replacing an array with an object (Replace Array with Object)
 *
 * BEFORE: Using array for structured data
 */
class PerformanceArray {
public:
    std::unordered_map<std::string, int> getPerformanceData() {
        return {
            {"goals", 10},
            {"assists", 5},
            {"minutes", 120}
        };
    }

    double calculateScore(const std::unordered_map<std::string, int>& data) {
        return (data.at("goals") * 2) + (data.at("assists") * 1.5) + (data.at("minutes") / 60);
    }
};

/**
 * AFTER: Replace array with object
 */
class PerformanceData {
private:
    int goals;
    int assists;
    int minutes;

public:
    PerformanceData(int goals, int assists, int minutes)
        : goals(goals), assists(assists), minutes(minutes) {}

    int getGoals() const {
        return goals;
    }

    int getAssists() const {
        return assists;
    }

    int getMinutes() const {
        return minutes;
    }

    double calculateScore() const {
        return (goals * 2) + (assists * 1.5) + (minutes / 60);
    }
};

class PerformanceObject {
public:
    PerformanceData getPerformanceData() {
        return PerformanceData(10, 5, 120);
    }

    double calculateScore(const PerformanceData& data) {
        return data.calculateScore();
    }
};

/**
 * 23. Duplication of visible data (Duplicate Observed Data)
 *
 * BEFORE: Domain data mixed with presentation
 */
class OrderDomain {
private:
    double total = 0.0;

public:
    void addItem(double price) {
        total += price;
        // Have to update UI here too
        updateDisplay();
    }

private:
    void updateDisplay() {
        std::cout << "Total: $" << total << std::endl;
    }
};

/**
 * AFTER: Separate domain and presentation data
 */
class OrderObserver {
public:
    virtual ~OrderObserver() = default;
    virtual void update(double total) = 0;
};

class OrderDomainSeparated {
private:
    double total = 0.0;
    std::vector<OrderObserver*> observers;

public:
    void addItem(double price) {
        total += price;
        notifyObservers();
    }

    double getTotal() const {
        return total;
    }

    void addObserver(OrderObserver* observer) {
        observers.push_back(observer);
    }

private:
    void notifyObservers() {
        for (auto observer : observers) {
            observer->update(total);
        }
    }
};

class OrderDisplay : public OrderObserver {
private:
    OrderDomainSeparated& order;

public:
    OrderDisplay(OrderDomainSeparated& order) : order(order) {
        order.addObserver(this);
    }

    void update(double total) override {
        std::cout << "Total: $" << total << std::endl;
    }
};

/**
 * 24. Replacing Unidirectional communication with Bidirectional
 * communication (Change Unidirectional Association to Bidirectional)
 *
 * BEFORE: One-way association
 */
class CustomerUni {
private:
    std::vector<std::string> orderIds; // Simplified - just storing IDs

public:
    void addOrder(const std::string& orderId) {
        orderIds.push_back(orderId);
        // Order doesn't know about customer
    }
};

class OrderUni {
private:
    std::vector<std::string> items;
};

/**
 * AFTER: Bidirectional association
 */
class CustomerBi;
class OrderBi;

class CustomerBi {
private:
    std::vector<OrderBi*> orders;

public:
    void addOrder(OrderBi* order);
};

class OrderBi {
private:
    CustomerBi* customer;
    std::vector<std::string> items;

public:
    void setCustomer(CustomerBi* customer) {
        this->customer = customer;
    }

    CustomerBi* getCustomer() const {
        return customer;
    }
};

void CustomerBi::addOrder(OrderBi* order) {
    orders.push_back(order);
    order->setCustomer(this);
}

/**
 * 25. Replacing Bidirectional communication with Unidirectional
 * communication (Change Bidirectional Association to Unidirectional)
 *
 * BEFORE: Unnecessary bidirectional association
 */
class CustomerBidirectional;
class OrderBidirectional;

class CustomerBidirectional {
private:
    std::vector<OrderBidirectional*> orders;

public:
    void addOrder(OrderBidirectional* order);
};

class OrderBidirectional {
private:
    CustomerBidirectional* customer;

public:
    void setCustomer(CustomerBidirectional* customer) {
        this->customer = customer;
    }

    CustomerBidirectional* getCustomer() const {
        return customer;
    }
};

void CustomerBidirectional::addOrder(OrderBidirectional* order) {
    orders.push_back(order);
    order->setCustomer(this);
}

/**
 * AFTER: Remove bidirectional link
 */
class CustomerUnidirectional {
private:
    std::vector<std::string> orderIds;

public:
    void addOrder(const std::string& orderId) {
        orderIds.push_back(orderId);
    }
};

class OrderUnidirectional {
private:
    std::string customerId;

public:
    OrderUnidirectional(const std::string& customerId) : customerId(customerId) {}

    std::string getCustomerId() const {
        return customerId;
    }
};

/**
 * 26. Replacing the magic number with a symbolic constant
 * (Replace Magic Number with Symbolic Constant)
 *
 * BEFORE: Magic numbers
 */
class GeometryBefore {
public:
    double calculateCircleArea(double radius) {
        return 3.14159 * radius * radius; // Magic number
    }

    double calculateCircleCircumference(double radius) {
        return 2 * 3.14159 * radius; // Same magic number
    }
};

/**
 * AFTER: Use symbolic constant
 */
class GeometryAfter {
public:
    static constexpr double PI = 3.14159;

    double calculateCircleArea(double radius) {
        return PI * radius * radius;
    }

    double calculateCircleCircumference(double radius) {
        return 2 * PI * radius;
    }
};

/**
 * 27. Encapsulate Field
 *
 * BEFORE: Public field
 */
class PersonPublic {
public:
    std::string name;
};

/**
 * AFTER: Encapsulated field
 */
class PersonEncapsulated {
private:
    std::string name;

public:
    std::string getName() const {
        return name;
    }

    void setName(const std::string& name) {
        this->name = name;
    }
};

/**
 * 28. Encapsulate Collection
 *
 * BEFORE: Direct access to collection
 */
class TeamBefore {
public:
    std::vector<std::string> players; // Direct access

    void addPlayer(const std::string& player) {
        players.push_back(player);
    }
};

/**
 * AFTER: Encapsulated collection
 */
class TeamAfter {
private:
    std::vector<std::string> players;

public:
    void addPlayer(const std::string& player) {
        players.push_back(player);
    }

    void removePlayer(const std::string& player) {
        auto it = std::find(players.begin(), players.end(), player);
        if (it != players.end()) {
            players.erase(it);
        }
    }

    std::vector<std::string> getPlayers() const {
        return players; // Return copy
    }

    size_t getPlayerCount() const {
        return players.size();
    }
};

/**
 * 29. Replacing a record with a Data Class
 *
 * BEFORE: Using array as data structure
 */
class EmployeeArray {
public:
    std::unordered_map<std::string, std::string> createEmployee(const std::unordered_map<std::string, std::string>& data) {
        return {
            {"name", data.at("name")},
            {"salary", data.at("salary")},
            {"department", data.at("department")}
        };
    }

    std::string getSalary(const std::unordered_map<std::string, std::string>& employee) {
        return employee.at("salary");
    }
};

/**
 * AFTER: Use data class
 */
class Employee {
private:
    std::string name;
    std::string salary;
    std::string department;

public:
    Employee(const std::string& name, const std::string& salary, const std::string& department)
        : name(name), salary(salary), department(department) {}

    std::string getName() const {
        return name;
    }

    std::string getSalary() const {
        return salary;
    }

    std::string getDepartment() const {
        return department;
    }
};

class EmployeeDataClass {
public:
    Employee createEmployee(const std::string& name, const std::string& salary, const std::string& department) {
        return Employee(name, salary, department);
    }

    std::string getSalary(const Employee& employee) {
        return employee.getSalary();
    }
};

/**
 * 30. Replacing Type Code with Class
 *
 * BEFORE: Type code as constants
 */
class EmployeeTypeCode {
public:
    static const int ENGINEER = 0;
    static const int SALESMAN = 1;
    static const int MANAGER = 2;

private:
    int type;

public:
    EmployeeTypeCode(int type) : type(type) {}

    int getTypeCode() const {
        return type;
    }

    double getMonthlySalary() const {
        switch (type) {
            case ENGINEER: return 5000;
            case SALESMAN: return 4000;
            case MANAGER: return 6000;
            default: return 0;
        }
    }
};

/**
 * AFTER: Replace type code with class
 */
class EmployeeType {
public:
    virtual ~EmployeeType() = default;
    virtual double getMonthlySalary() const = 0;

    static std::unique_ptr<EmployeeType> createEngineer();
    static std::unique_ptr<EmployeeType> createSalesman();
    static std::unique_ptr<EmployeeType> createManager();
};

class EngineerType : public EmployeeType {
public:
    double getMonthlySalary() const override {
        return 5000;
    }
};

class SalesmanType : public EmployeeType {
public:
    double getMonthlySalary() const override {
        return 4000;
    }
};

class ManagerType : public EmployeeType {
public:
    double getMonthlySalary() const override {
        return 6000;
    }
};

std::unique_ptr<EmployeeType> EmployeeType::createEngineer() {
    return std::make_unique<EngineerType>();
}

std::unique_ptr<EmployeeType> EmployeeType::createSalesman() {
    return std::make_unique<SalesmanType>();
}

std::unique_ptr<EmployeeType> EmployeeType::createManager() {
    return std::make_unique<ManagerType>();
}

class EmployeeTypeClass {
private:
    std::unique_ptr<EmployeeType> type;

public:
    EmployeeTypeClass(std::unique_ptr<EmployeeType> type) : type(std::move(type)) {}

    double getMonthlySalary() const {
        return type->getMonthlySalary();
    }
};

/**
 * 31. Replacing Type Code with Subclasses
 *
 * BEFORE: Type code in base class
 */
class EmployeeSubBefore {
public:
    static const int ENGINEER = 0;
    static const int SALESMAN = 1;
    static const int MANAGER = 2;

private:
    int type;
    double salary;

public:
    EmployeeSubBefore(int type, double salary) : type(type), salary(salary) {}

    double getSalary() const {
        return salary;
    }

    int getType() const {
        return type;
    }
};

/**
 * AFTER: Replace type code with subclasses
 */
class EmployeeSubAfter {
protected:
    double salary;

public:
    EmployeeSubAfter(double salary) : salary(salary) {}

    double getSalary() const {
        return salary;
    }

    virtual std::string getType() const = 0;
};

class Engineer : public EmployeeSubAfter {
public:
    Engineer(double salary) : EmployeeSubAfter(salary) {}

    std::string getType() const override {
        return "engineer";
    }
};

class Salesman : public EmployeeSubAfter {
public:
    Salesman(double salary) : EmployeeSubAfter(salary) {}

    std::string getType() const override {
        return "salesman";
    }
};

class Manager : public EmployeeSubAfter {
public:
    Manager(double salary) : EmployeeSubAfter(salary) {}

    std::string getType() const override {
        return "manager";
    }
};

/**
 * 32. Replacing Type Code with State/Strategy
 *
 * BEFORE: Type code with behavior
 */
class EmployeeStateBefore {
public:
    static const int JUNIOR = 0;
    static const int SENIOR = 1;
    static const int LEAD = 2;

private:
    int level;

public:
    EmployeeStateBefore(int level) : level(level) {}

    double getSalaryMultiplier() const {
        switch (level) {
            case JUNIOR: return 1.0;
            case SENIOR: return 1.5;
            case LEAD: return 2.0;
            default: return 1.0;
        }
    }
};

/**
 * AFTER: Use state/strategy pattern
 */
class EmployeeLevel {
public:
    virtual ~EmployeeLevel() = default;
    virtual double getSalaryMultiplier() const = 0;
};

class JuniorLevel : public EmployeeLevel {
public:
    double getSalaryMultiplier() const override {
        return 1.0;
    }
};

class SeniorLevel : public EmployeeLevel {
public:
    double getSalaryMultiplier() const override {
        return 1.5;
    }
};

class LeadLevel : public EmployeeLevel {
public:
    double getSalaryMultiplier() const override {
        return 2.0;
    }
};

class EmployeeStateAfter {
private:
    std::unique_ptr<EmployeeLevel> level;

public:
    EmployeeStateAfter(std::unique_ptr<EmployeeLevel> level) : level(std::move(level)) {}

    double getSalaryMultiplier() const {
        return level->getSalaryMultiplier();
    }
};

/**
 * 33. Replacing Subclass with Fields
 *
 * BEFORE: Unnecessary subclasses
 */
class PersonSub {
protected:
    std::string name;
    std::string gender;

public:
    PersonSub(const std::string& name, const std::string& gender)
        : name(name), gender(gender) {}

    std::string getName() const {
        return name;
    }

    virtual bool isMale() const = 0;
};

class Male : public PersonSub {
public:
    Male(const std::string& name) : PersonSub(name, "male") {}

    bool isMale() const override {
        return true;
    }
};

class Female : public PersonSub {
public:
    Female(const std::string& name) : PersonSub(name, "female") {}

    bool isMale() const override {
        return false;
    }
};

/**
 * AFTER: Replace subclass with field
 */
class PersonField {
private:
    std::string name;
    std::string gender; // 'male' or 'female'

public:
    PersonField(const std::string& name, const std::string& gender)
        : name(name), gender(gender) {}

    std::string getName() const {
        return name;
    }

    bool isMale() const {
        return gender == "male";
    }

    std::string getGender() const {
        return gender;
    }
};
