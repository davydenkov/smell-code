#include <iostream>
#include <string>
#include <vector>
#include <memory>
#include <cmath>

/**
 * 57. Lifting the field (Pull Up Field)
 *
 * BEFORE: Duplicate fields in subclasses
 */
class EmployeePullBefore {
protected:
    std::string name;
};

class ManagerPullBefore : public EmployeePullBefore {
protected:
    std::string name; // Duplicate
    double budget;
};

class EngineerPullBefore : public EmployeePullBefore {
protected:
    std::string name; // Duplicate
    std::vector<std::string> skills;
};

/**
 * AFTER: Pull up field to superclass
 */
class EmployeePullAfter {
protected:
    std::string name;
};

class ManagerPullAfter : public EmployeePullAfter {
private:
    double budget;
};

class EngineerPullAfter : public EmployeePullAfter {
private:
    std::vector<std::string> skills;
};

/**
 * 58. Lifting the method (Pull Up Method)
 *
 * BEFORE: Duplicate methods in subclasses
 */
class ShapeMethodBefore {
public:
    virtual double area() const = 0;
};

class CircleMethodBefore : public ShapeMethodBefore {
private:
    double radius;

public:
    double area() const override {
        return M_PI * radius * radius;
    }

    double circumference() const {
        return 2 * M_PI * radius;
    }
};

class SquareMethodBefore : public ShapeMethodBefore {
private:
    double side;

public:
    double area() const override {
        return side * side;
    }

    double circumference() const { // Duplicate logic
        return 4 * side;
    }
};

/**
 * AFTER: Pull up method to superclass
 */
class ShapeMethodAfter {
public:
    virtual double area() const = 0;
    virtual double circumference() const = 0;
};

class CircleMethodAfter : public ShapeMethodAfter {
private:
    double radius;

public:
    double area() const override {
        return M_PI * radius * radius;
    }

    double circumference() const override {
        return 2 * M_PI * radius;
    }
};

class SquareMethodAfter : public ShapeMethodAfter {
private:
    double side;

public:
    double area() const override {
        return side * side;
    }

    double circumference() const override {
        return 4 * side;
    }
};

/**
 * 59. Lifting the constructor Body (Pull Up Constructor Body)
 *
 * BEFORE: Duplicate constructor code
 */
class VehicleConstructorBefore {
protected:
    std::string make;
    std::string model;
    int year;
};

class CarConstructorBefore : public VehicleConstructorBefore {
private:
    int doors;

public:
    CarConstructorBefore(const std::string& make, const std::string& model, int year, int doors) {
        this->make = make; // Duplicate
        this->model = model; // Duplicate
        this->year = year; // Duplicate
        this->doors = doors;
    }
};

class TruckConstructorBefore : public VehicleConstructorBefore {
private:
    double payload;

public:
    TruckConstructorBefore(const std::string& make, const std::string& model, int year, double payload) {
        this->make = make; // Duplicate
        this->model = model; // Duplicate
        this->year = year; // Duplicate
        this->payload = payload;
    }
};

/**
 * AFTER: Pull up constructor body
 */
class VehicleConstructorAfter {
protected:
    std::string make;
    std::string model;
    int year;

public:
    VehicleConstructorAfter(const std::string& make, const std::string& model, int year)
        : make(make), model(model), year(year) {}
};

class CarConstructorAfter : public VehicleConstructorAfter {
private:
    int doors;

public:
    CarConstructorAfter(const std::string& make, const std::string& model, int year, int doors)
        : VehicleConstructorAfter(make, model, year), doors(doors) {}
};

class TruckConstructorAfter : public VehicleConstructorAfter {
private:
    double payload;

public:
    TruckConstructorAfter(const std::string& make, const std::string& model, int year, double payload)
        : VehicleConstructorAfter(make, model, year), payload(payload) {}
};

/**
 * 60. Method Descent (Push Down Method)
 *
 * BEFORE: Method in wrong class hierarchy level
 */
class AnimalPushBefore {
public:
    virtual std::string speak() const {
        return ""; // Generic implementation
    }
};

class DogPushBefore : public AnimalPushBefore {
public:
    std::string speak() const override {
        return "Woof";
    }
};

class CatPushBefore : public AnimalPushBefore {
public:
    std::string speak() const override {
        return "Meow";
    }
};

class FishPushBefore : public AnimalPushBefore {
    // Fish don't speak, but inherits speak method
};

/**
 * AFTER: Push down method to appropriate subclasses
 */
class AnimalPushAfter {
    // No speak method here
};

class DogPushAfter : public AnimalPushAfter {
public:
    std::string speak() const {
        return "Woof";
    }
};

class CatPushAfter : public AnimalPushAfter {
public:
    std::string speak() const {
        return "Meow";
    }
};

class FishPushAfter : public AnimalPushAfter {
    // No speak method - appropriate for Fish
};

/**
 * 61. Field Descent (Push Down Field)
 *
 * BEFORE: Field in wrong hierarchy level
 */
class EmployeeFieldBefore {
protected:
    double salary; // Not all employees have salary
};

class SalariedEmployeeFieldBefore : public EmployeeFieldBefore {
    // Uses salary
};

class ContractorFieldBefore : public EmployeeFieldBefore {
    // Doesn't use salary, but inherits it
};

/**
 * AFTER: Push down field
 */
class EmployeeFieldAfter {
    // No salary field
};

class SalariedEmployeeFieldAfter : public EmployeeFieldAfter {
protected:
    double salary;
};

class ContractorFieldAfter : public EmployeeFieldAfter {
private:
    double hourlyRate;
};

/**
 * 62. Subclass extraction (Extract Subclass)
 *
 * BEFORE: Class with conditional behavior
 */
class JobExtractBefore {
private:
    std::string type;
    double rate;
    double commission;

public:
    JobExtractBefore(const std::string& type, double rate, double commission = 0.0)
        : type(type), rate(rate), commission(commission) {}

    double getPay() const {
        if (type == "salaried") {
            return rate;
        } else {
            return rate + commission;
        }
    }
};

/**
 * AFTER: Extract subclass
 */
class JobExtractAfter {
protected:
    double rate;

public:
    JobExtractAfter(double rate) : rate(rate) {}

    virtual double getPay() const = 0;
};

class SalariedJob : public JobExtractAfter {
public:
    SalariedJob(double rate) : JobExtractAfter(rate) {}

    double getPay() const override {
        return rate;
    }
};

class CommissionedJob : public JobExtractAfter {
private:
    double commission;

public:
    CommissionedJob(double rate, double commission)
        : JobExtractAfter(rate), commission(commission) {}

    double getPay() const override {
        return rate + commission;
    }
};

/**
 * 63. Allocation of the parent class (Extract Superclass)
 *
 * BEFORE: Duplicate code in classes
 */
class DepartmentSuperBefore {
private:
    std::string name;
    std::string head;

public:
    DepartmentSuperBefore(const std::string& name, const std::string& head)
        : name(name), head(head) {}

    std::string getName() const {
        return name;
    }

    std::string getHead() const {
        return head;
    }
};

class CompanySuperBefore {
private:
    std::string name;
    std::string head;

public:
    CompanySuperBefore(const std::string& name, const std::string& head)
        : name(name), head(head) {}

    std::string getName() const {
        return name;
    }

    std::string getHead() const {
        return head;
    }
};

/**
 * AFTER: Extract superclass
 */
class Party {
private:
    std::string name;
    std::string head;

public:
    Party(const std::string& name, const std::string& head)
        : name(name), head(head) {}

    std::string getName() const {
        return name;
    }

    std::string getHead() const {
        return head;
    }
};

class DepartmentSuperAfter : public Party {
public:
    DepartmentSuperAfter(const std::string& name, const std::string& head)
        : Party(name, head) {}
};

class CompanySuperAfter : public Party {
public:
    CompanySuperAfter(const std::string& name, const std::string& head)
        : Party(name, head) {}
};

/**
 * 64. Interface extraction (Extract Interface)
 *
 * BEFORE: Clients depend on concrete class
 */
class PrinterInterfaceBefore {
public:
    void print(const std::string& document) {
        // Print logic
        std::cout << "Printing: " << document << std::endl;
    }

    std::string getStatus() const {
        return "Ready";
    }

    void cancelJob(int jobId) {
        // Cancel logic
        std::cout << "Cancelling job: " << jobId << std::endl;
    }
};

/**
 * AFTER: Extract interface
 */
class Printer {
public:
    virtual ~Printer() = default;
    virtual void print(const std::string& document) = 0;
    virtual std::string getStatus() const = 0;
};

class LaserPrinter : public Printer {
public:
    void print(const std::string& document) override {
        // Print logic
        std::cout << "Laser printing: " << document << std::endl;
    }

    std::string getStatus() const override {
        return "Ready";
    }

    void cancelJob(int jobId) {
        // Cancel logic - not part of interface
        std::cout << "Cancelling laser job: " << jobId << std::endl;
    }
};

class InkjetPrinter : public Printer {
public:
    void print(const std::string& document) override {
        // Print logic
        std::cout << "Inkjet printing: " << document << std::endl;
    }

    std::string getStatus() const override {
        return "Ready";
    }
};

/**
 * 65. Collapse Hierarchy
 *
 * BEFORE: Unnecessary class hierarchy
 */
class EmployeeCollapseBefore {
};

class ManagerCollapseBefore : public EmployeeCollapseBefore {
private:
    std::string department;
};

/**
 * AFTER: Collapse hierarchy if only one subclass
 */
class EmployeeCollapseAfter {
private:
    std::string department; // Moved up
};

/**
 * 66. Formation of the method template (Form Template Method)
 *
 * BEFORE: Duplicate algorithm structure
 */
class ReportGeneratorTemplateBefore {
public:
    std::string generateHTMLReport() {
        std::vector<std::string> data = getData();
        std::string header = formatHeader();
        std::string body = formatBody(data);
        std::string footer = formatFooter();
        return header + body + footer;
    }

    std::string generatePDFReport() {
        std::vector<std::string> data = getData(); // Duplicate
        std::string header = formatPDFHeader(); // Different
        std::string body = formatPDFBody(data); // Different
        std::string footer = formatPDFFooter(); // Different
        return header + body + footer;
    }

protected:
    virtual std::vector<std::string> getData() const {
        return {"item1", "item2"};
    }

    virtual std::string formatHeader() const {
        return "<h1>Report</h1>";
    }

    virtual std::string formatBody(const std::vector<std::string>& data) const {
        std::string body = "<body>";
        for (const auto& item : data) {
            body += item;
        }
        body += "</body>";
        return body;
    }

    virtual std::string formatFooter() const {
        return "<footer>End</footer>";
    }

    virtual std::string formatPDFHeader() const {
        return "PDF Report Header";
    }

    virtual std::string formatPDFBody(const std::vector<std::string>& data) const {
        std::string body = "PDF Body: ";
        for (const auto& item : data) {
            body += item;
        }
        return body;
    }

    virtual std::string formatPDFFooter() const {
        return "PDF Footer";
    }
};

/**
 * AFTER: Form template method
 */
class ReportGeneratorTemplateAfter {
public:
    std::string generateReport() {
        std::vector<std::string> data = getData();
        std::string header = formatHeader();
        std::string body = formatBody(data);
        std::string footer = formatFooter();
        return assembleReport(header, body, footer);
    }

protected:
    virtual std::vector<std::string> getData() const {
        return {"item1", "item2"};
    }

    virtual std::string formatHeader() const = 0;
    virtual std::string formatBody(const std::vector<std::string>& data) const = 0;
    virtual std::string formatFooter() const = 0;
    virtual std::string assembleReport(const std::string& header, const std::string& body, const std::string& footer) const = 0;
};

class HTMLReportGenerator : public ReportGeneratorTemplateAfter {
protected:
    std::string formatHeader() const override {
        return "<h1>Report</h1>";
    }

    std::string formatBody(const std::vector<std::string>& data) const override {
        std::string body = "<body>";
        for (const auto& item : data) {
            body += item;
        }
        body += "</body>";
        return body;
    }

    std::string formatFooter() const override {
        return "<footer>End</footer>";
    }

    std::string assembleReport(const std::string& header, const std::string& body, const std::string& footer) const override {
        return header + body + footer;
    }
};

class PDFReportGenerator : public ReportGeneratorTemplateAfter {
protected:
    std::string formatHeader() const override {
        return "PDF Report Header";
    }

    std::string formatBody(const std::vector<std::string>& data) const override {
        std::string body = "PDF Body: ";
        for (const auto& item : data) {
            body += item;
        }
        return body;
    }

    std::string formatFooter() const override {
        return "PDF Footer";
    }

    std::string assembleReport(const std::string& header, const std::string& body, const std::string& footer) const override {
        return header + body + footer;
    }
};

/**
 * 67. Replacement of inheritance by delegation (Replace Inheritance with Delegation)
 *
 * BEFORE: Inheritance where delegation would be better
 */
class StackInheritanceBefore : public std::vector<int> {
public:
    void push(int item) {
        push_back(item);
    }

    int pop() {
        if (empty()) {
            throw std::runtime_error("Stack is empty");
        }
        int lastIndex = size() - 1;
        int item = at(lastIndex);
        erase(begin() + lastIndex);
        return item;
    }
};

/**
 * AFTER: Replace inheritance with delegation
 */
class StackDelegationAfter {
private:
    std::vector<int> items;

public:
    void push(int item) {
        items.push_back(item);
    }

    int pop() {
        if (items.empty()) {
            throw std::runtime_error("Stack is empty");
        }
        int lastIndex = items.size() - 1;
        int item = items[lastIndex];
        items.erase(items.begin() + lastIndex);
        return item;
    }

    size_t size() const {
        return items.size();
    }
};

/**
 * 68. Replacement of delegation by inheritance (Replace Delegation with Inheritance)
 *
 * BEFORE: Delegation where inheritance would be simpler
 */
class MyStringDelegateBefore {
private:
    std::string string;

public:
    MyStringDelegateBefore(const std::string& str) : string(str) {}

    size_t length() const {
        return string.length();
    }

    std::string substr(size_t start, size_t length = std::string::npos) const {
        return string.substr(start, length);
    }

    size_t strpos(const std::string& needle) const {
        return string.find(needle);
    }
};

/**
 * AFTER: Replace delegation with inheritance
 */
class MyStringInheritAfter : public std::vector<char> {
public:
    MyStringInheritAfter(const std::string& str) {
        assign(str.begin(), str.end());
    }

    std::string toString() const {
        return std::string(begin(), end());
    }
};
