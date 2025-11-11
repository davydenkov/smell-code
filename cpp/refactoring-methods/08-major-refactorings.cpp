#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <memory>
#include <iomanip>

/**
 * 69. Separation of inheritance (Tease Apart Inheritance)
 *
 * BEFORE: Class hierarchy mixing two different responsibilities
 */
class EmployeeTeaseBefore {
protected:
    std::string name;
    double rate;

public:
    EmployeeTeaseBefore(const std::string& name, double rate)
        : name(name), rate(rate) {}

    std::string getName() const {
        return name;
    }
};

class SalariedEmployeeTeaseBefore : public EmployeeTeaseBefore {
public:
    SalariedEmployeeTeaseBefore(const std::string& name, double rate)
        : EmployeeTeaseBefore(name, rate) {}

    double getPay() const {
        return rate;
    }
};

class CommissionedEmployeeTeaseBefore : public EmployeeTeaseBefore {
private:
    double commission;

public:
    CommissionedEmployeeTeaseBefore(const std::string& name, double rate, double commission)
        : EmployeeTeaseBefore(name, rate), commission(commission) {}

    double getPay() const {
        return rate + commission;
    }
};

/**
 * AFTER: Tease apart inheritance into two separate hierarchies
 */
class Payable {
public:
    virtual ~Payable() = default;
    virtual double getPay() const = 0;
};

class EmployeeTeaseAfter {
protected:
    std::string name;

public:
    EmployeeTeaseAfter(const std::string& name) : name(name) {}

    std::string getName() const {
        return name;
    }
};

class SalariedEmployeeTeaseAfter : public EmployeeTeaseAfter, public Payable {
private:
    double salary;

public:
    SalariedEmployeeTeaseAfter(const std::string& name, double salary)
        : EmployeeTeaseAfter(name), salary(salary) {}

    double getPay() const override {
        return salary;
    }
};

class CommissionedEmployeeTeaseAfter : public EmployeeTeaseAfter, public Payable {
private:
    double baseSalary;
    double commission;

public:
    CommissionedEmployeeTeaseAfter(const std::string& name, double baseSalary, double commission)
        : EmployeeTeaseAfter(name), baseSalary(baseSalary), commission(commission) {}

    double getPay() const override {
        return baseSalary + commission;
    }
};

/**
 * 70. Converting a procedural project into objects (Convert Procedural Design to Objects)
 *
 * BEFORE: Procedural code with global functions and data
 */
class ProceduralDesignBefore {
private:
    static std::unordered_map<std::string, double> accounts;

public:
    static void createAccount(const std::string& id, double balance) {
        accounts[id] = balance;
    }

    static double getBalance(const std::string& id) {
        auto it = accounts.find(id);
        return (it != accounts.end()) ? it->second : 0.0;
    }

    static void deposit(const std::string& id, double amount) {
        auto it = accounts.find(id);
        if (it != accounts.end()) {
            it->second += amount;
        }
    }

    static bool withdraw(const std::string& id, double amount) {
        auto it = accounts.find(id);
        if (it != accounts.end() && it->second >= amount) {
            it->second -= amount;
            return true;
        }
        return false;
    }
};

std::unordered_map<std::string, double> ProceduralDesignBefore::accounts;

/**
 * AFTER: Convert to object-oriented design
 */
class Account {
private:
    std::string id;
    double balance;

public:
    Account(const std::string& id, double balance = 0.0)
        : id(id), balance(balance) {}

    std::string getId() const {
        return id;
    }

    double getBalance() const {
        return balance;
    }

    void deposit(double amount) {
        balance += amount;
    }

    bool withdraw(double amount) {
        if (balance >= amount) {
            balance -= amount;
            return true;
        }
        return false;
    }
};

class Bank {
private:
    std::unordered_map<std::string, std::shared_ptr<Account>> accounts;

public:
    std::shared_ptr<Account> createAccount(const std::string& id, double balance = 0.0) {
        auto account = std::make_shared<Account>(id, balance);
        accounts[id] = account;
        return account;
    }

    std::shared_ptr<Account> getAccount(const std::string& id) {
        auto it = accounts.find(id);
        return (it != accounts.end()) ? it->second : nullptr;
    }

    double getBalance(const std::string& id) {
        auto account = getAccount(id);
        return account ? account->getBalance() : 0.0;
    }

    void deposit(const std::string& id, double amount) {
        auto account = getAccount(id);
        if (account) {
            account->deposit(amount);
        }
    }

    bool withdraw(const std::string& id, double amount) {
        auto account = getAccount(id);
        return account ? account->withdraw(amount) : false;
    }
};

/**
 * 71. Separating the domain from the representation (Separate Domain from Presentation)
 *
 * BEFORE: Domain logic mixed with presentation
 */
class OrderPresentationBefore {
private:
    std::vector<std::unordered_map<std::string, std::string>> items;
    double total = 0.0;

public:
    void addItem(const std::string& name, double price, int quantity) {
        std::unordered_map<std::string, std::string> item;
        item["name"] = name;
        item["price"] = std::to_string(price);
        item["quantity"] = std::to_string(quantity);
        items.push_back(item);
        total += price * quantity;

        // Presentation logic mixed in
        std::cout << "Added " << quantity << " x " << name << " to order" << std::endl;
        std::cout << "Current total: $" << std::fixed << std::setprecision(2) << total << std::endl;
    }

    double getTotal() const {
        return total;
    }

    void displayOrder() {
        std::cout << "Order Summary:" << std::endl;
        for (const auto& item : items) {
            double price = std::stod(item.at("price"));
            int quantity = std::stoi(item.at("quantity"));
            std::cout << "- " << quantity << " x " << item.at("name")
                     << " @ $" << std::fixed << std::setprecision(2) << price << std::endl;
        }
        std::cout << "Total: $" << std::fixed << std::setprecision(2) << total << std::endl;
    }
};

/**
 * AFTER: Separate domain from presentation
 */
class OrderItem {
private:
    std::string name;
    double price;
    int quantity;

public:
    OrderItem(const std::string& name, double price, int quantity)
        : name(name), price(price), quantity(quantity) {}

    std::string getName() const {
        return name;
    }

    double getPrice() const {
        return price;
    }

    int getQuantity() const {
        return quantity;
    }

    double getTotal() const {
        return price * quantity;
    }
};

class OrderDomainAfter {
private:
    std::vector<OrderItem> items;

public:
    void addItem(const std::string& name, double price, int quantity) {
        OrderItem item(name, price, quantity);
        items.push_back(item);
    }

    const std::vector<OrderItem>& getItems() const {
        return items;
    }

    double getTotal() const {
        double total = 0.0;
        for (const auto& item : items) {
            total += item.getTotal();
        }
        return total;
    }
};

class OrderPresenter {
public:
    void displayItemAdded(const OrderItem& item) {
        std::cout << "Added " << item.getQuantity() << " x " << item.getName() << " to order" << std::endl;
    }

    void displayOrderSummary(const OrderDomainAfter& order) {
        std::cout << "Order Summary:" << std::endl;
        for (const auto& item : order.getItems()) {
            std::cout << "- " << item.getQuantity() << " x " << item.getName()
                     << " @ $" << std::fixed << std::setprecision(2) << item.getPrice() << std::endl;
        }
        std::cout << "Total: $" << std::fixed << std::setprecision(2) << order.getTotal() << std::endl;
    }
};

class OrderService {
private:
    OrderDomainAfter order;
    OrderPresenter presenter;

public:
    void addItem(const std::string& name, double price, int quantity) {
        OrderItem item(name, price, quantity);
        order.addItem(name, price, quantity);
        presenter.displayItemAdded(item);
        presenter.displayOrderSummary(order);
    }

    const OrderDomainAfter& getOrder() const {
        return order;
    }
};

/**
 * 72. Hierarchy Extraction (Extract Hierarchy)
 *
 * BEFORE: Single class handling multiple responsibilities
 */
class ComputerExtractBefore {
private:
    std::string type;
    std::string cpu;
    int ram;
    int storage;

public:
    ComputerExtractBefore(const std::string& type, const std::string& cpu, int ram, int storage)
        : type(type), cpu(cpu), ram(ram), storage(storage) {}

    std::string getSpecs() const {
        std::string specs = "CPU: " + cpu + "\n";
        specs += "RAM: " + std::to_string(ram) + "GB\n";
        specs += "Storage: " + std::to_string(storage) + "GB\n";

        if (type == "desktop") {
            specs += "Form Factor: Desktop\n";
        } else if (type == "laptop") {
            specs += "Form Factor: Laptop\n";
            specs += "Battery Life: 8 hours\n";
        } else if (type == "server") {
            specs += "Form Factor: Server Rack\n";
            specs += "Redundancy: RAID 10\n";
        }

        return specs;
    }
};

/**
 * AFTER: Extract hierarchy
 */
class ComputerExtractAfter {
protected:
    std::string cpu;
    int ram;
    int storage;

public:
    ComputerExtractAfter(const std::string& cpu, int ram, int storage)
        : cpu(cpu), ram(ram), storage(storage) {}

    virtual ~ComputerExtractAfter() = default;

    virtual std::string getFormFactor() const = 0;
    virtual std::string getSpecialFeatures() const = 0;

    std::string getBasicSpecs() const {
        return "CPU: " + cpu + "\n" +
               "RAM: " + std::to_string(ram) + "GB\n" +
               "Storage: " + std::to_string(storage) + "GB\n";
    }

    std::string getSpecs() const {
        return getBasicSpecs() +
               "Form Factor: " + getFormFactor() + "\n" +
               getSpecialFeatures();
    }
};

class DesktopComputer : public ComputerExtractAfter {
public:
    DesktopComputer(const std::string& cpu, int ram, int storage)
        : ComputerExtractAfter(cpu, ram, storage) {}

    std::string getFormFactor() const override {
        return "Desktop";
    }

    std::string getSpecialFeatures() const override {
        return "Expansion Slots: Multiple PCI\n";
    }
};

class LaptopComputer : public ComputerExtractAfter {
public:
    LaptopComputer(const std::string& cpu, int ram, int storage)
        : ComputerExtractAfter(cpu, ram, storage) {}

    std::string getFormFactor() const override {
        return "Laptop";
    }

    std::string getSpecialFeatures() const override {
        return "Battery Life: 8 hours\nWeight: 2.5 lbs\n";
    }
};

class ServerComputer : public ComputerExtractAfter {
public:
    ServerComputer(const std::string& cpu, int ram, int storage)
        : ComputerExtractAfter(cpu, ram, storage) {}

    std::string getFormFactor() const override {
        return "Server Rack";
    }

    std::string getSpecialFeatures() const override {
        return "Redundancy: RAID 10\nHot Swap Drives: Yes\n";
    }
};
