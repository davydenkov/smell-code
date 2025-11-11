/**
 * 18. Self-Encapsulate Field
 *
 * BEFORE: Direct field access
 */
class PersonBefore {
    constructor(name) {
        this.name = name; // Direct access
    }

    getName() {
        return this.name;
    }

    setName(name) {
        this.name = name;
    }
}

/**
 * AFTER: Self-encapsulate field
 */
class PersonAfter {
    constructor(name) {
        this._name = name;
    }

    getName() {
        return this._name;
    }

    setName(name) {
        this._name = name;
    }
}

/**
 * 19. Replacing the data value with an object (Replace Data Value with Object)
 *
 * BEFORE: Primitive data type that should be an object
 */
class OrderBefore {
    constructor(customerName) {
        this.customer = customerName; // Just a string
    }

    getCustomerName() {
        return this.customer;
    }

    setCustomer(customer) {
        this.customer = customer;
    }
}

/**
 * AFTER: Replace with object
 */
class Customer {
    constructor(name) {
        this.name = name;
    }

    getName() {
        return this.name;
    }
}

class OrderAfter {
    constructor(customer) {
        this.customer = customer;
    }

    getCustomer() {
        return this.customer;
    }

    setCustomer(customer) {
        this.customer = customer;
    }

    getCustomerName() {
        return this.customer.getName();
    }
}

/**
 * 20. Replacing the value with a reference (Change Value to Reference)
 *
 * BEFORE: Multiple instances of same object
 */
class CustomerValue {
    constructor(name) {
        this.name = name;
    }

    getName() {
        return this.name;
    }
}

class OrderValue {
    constructor(customerName) {
        this.customer = new CustomerValue(customerName); // New instance for each order
    }
}

/**
 * AFTER: Use reference to single instance
 */
class CustomerReference {
    constructor(name) {
        this.name = name;
    }

    getName() {
        return this.name;
    }

    static create(name) {
        if (!CustomerReference.instances) {
            CustomerReference.instances = new Map();
        }
        if (!CustomerReference.instances.has(name)) {
            CustomerReference.instances.set(name, new CustomerReference(name));
        }
        return CustomerReference.instances.get(name);
    }
}

class OrderReference {
    constructor(customerName) {
        this.customer = CustomerReference.create(customerName);
    }
}

/**
 * 21. Replacing a reference with a value (Change Reference to Value)
 *
 * BEFORE: Unnecessary reference when value would suffice
 */
class CurrencyReference {
    constructor(code) {
        this.code = code;
    }

    getCode() {
        return this.code;
    }
}

class ProductReference {
    constructor(price, currency) {
        this.price = price;
        this.currency = currency; // Reference object
    }
}

/**
 * AFTER: Use value object instead
 */
class CurrencyValue {
    constructor(code) {
        this.code = code;
    }

    getCode() {
        return this.code;
    }
}

class ProductValue {
    constructor(price, currencyCode) {
        this.price = price;
        this.currencyCode = currencyCode; // Just the value
    }

    getCurrencyCode() {
        return this.currencyCode;
    }
}

/**
 * 22. Replacing an array with an object (Replace Array with Object)
 *
 * BEFORE: Using array for structured data
 */
class PerformanceArray {
    getPerformanceData() {
        return {
            goals: 10,
            assists: 5,
            minutes: 120
        };
    }

    calculateScore(data) {
        return (data.goals * 2) + (data.assists * 1.5) + (data.minutes / 60);
    }
}

/**
 * AFTER: Replace array with object
 */
class PerformanceData {
    constructor(goals, assists, minutes) {
        this.goals = goals;
        this.assists = assists;
        this.minutes = minutes;
    }

    getGoals() {
        return this.goals;
    }

    getAssists() {
        return this.assists;
    }

    getMinutes() {
        return this.minutes;
    }

    calculateScore() {
        return (this.goals * 2) + (this.assists * 1.5) + (this.minutes / 60);
    }
}

class PerformanceObject {
    getPerformanceData() {
        return new PerformanceData(10, 5, 120);
    }

    calculateScore(data) {
        return data.calculateScore();
    }
}

/**
 * 23. Duplication of visible data (Duplicate Observed Data)
 *
 * BEFORE: Domain data mixed with presentation
 */
class OrderDomain {
    constructor() {
        this.total = 0;
    }

    addItem(price) {
        this.total += price;
        // Have to update UI here too
        this.updateDisplay();
    }

    updateDisplay() {
        // Update UI elements
    }
}

/**
 * AFTER: Separate domain and presentation data
 */
class OrderDomainSeparated {
    constructor() {
        this.total = 0;
        this.observers = [];
    }

    addItem(price) {
        this.total += price;
        this.notifyObservers();
    }

    getTotal() {
        return this.total;
    }

    addObserver(observer) {
        this.observers.push(observer);
    }

    notifyObservers() {
        this.observers.forEach(observer => observer.update(this.total));
    }
}

class OrderDisplay {
    constructor(order) {
        this.order = order;
        this.order.addObserver(this);
    }

    update(total) {
        // Update display with new total
    }
}

/**
 * 24. Replacing Unidirectional communication with Bidirectional
 * communication (Change Unidirectional Association to Bidirectional)
 *
 * BEFORE: One-way association
 */
class CustomerUni {
    constructor() {
        this.orders = [];
    }

    addOrder(order) {
        this.orders.push(order);
        // Order doesn't know about customer
    }
}

class OrderUni {
    constructor() {
        this.items = [];
    }
}

/**
 * AFTER: Bidirectional association
 */
class CustomerBi {
    constructor() {
        this.orders = [];
    }

    addOrder(order) {
        this.orders.push(order);
        order.setCustomer(this);
    }
}

class OrderBi {
    constructor() {
        this.customer = null;
        this.items = [];
    }

    setCustomer(customer) {
        this.customer = customer;
    }

    getCustomer() {
        return this.customer;
    }
}

/**
 * 25. Replacing Bidirectional communication with Unidirectional
 * communication (Change Bidirectional Association to Unidirectional)
 *
 * BEFORE: Unnecessary bidirectional association
 */
class CustomerBidirectional {
    constructor() {
        this.orders = [];
    }

    addOrder(order) {
        this.orders.push(order);
        order.setCustomer(this);
    }
}

class OrderBidirectional {
    constructor() {
        this.customer = null;
    }

    setCustomer(customer) {
        this.customer = customer;
    }

    getCustomer() {
        return this.customer;
    }
}

/**
 * AFTER: Remove bidirectional link
 */
class CustomerUnidirectional {
    constructor() {
        this.orders = [];
    }

    addOrder(order) {
        this.orders.push(order);
    }
}

class OrderUnidirectional {
    constructor(customerId) {
        this.customerId = customerId;
    }

    getCustomerId() {
        return this.customerId;
    }
}

/**
 * 26. Replacing the magic number with a symbolic constant
 * (Replace Magic Number with Symbolic Constant)
 *
 * BEFORE: Magic numbers
 */
class GeometryBefore {
    calculateCircleArea(radius) {
        return 3.14159 * radius * radius; // Magic number
    }

    calculateCircleCircumference(radius) {
        return 2 * 3.14159 * radius; // Same magic number
    }
}

/**
 * AFTER: Use symbolic constant
 */
class GeometryAfter {
    static PI = 3.14159;

    calculateCircleArea(radius) {
        return GeometryAfter.PI * radius * radius;
    }

    calculateCircleCircumference(radius) {
        return 2 * GeometryAfter.PI * radius;
    }
}

/**
 * 27. Encapsulate Field
 *
 * BEFORE: Public field
 */
class PersonPublic {
    constructor() {
        this.name = null;
    }
}

/**
 * AFTER: Encapsulated field
 */
class PersonEncapsulated {
    constructor() {
        this._name = null;
    }

    getName() {
        return this._name;
    }

    setName(name) {
        this._name = name;
    }
}

/**
 * 28. Encapsulate Collection
 *
 * BEFORE: Direct access to collection
 */
class TeamBefore {
    constructor() {
        this.players = []; // Direct access
    }

    addPlayer(player) {
        this.players.push(player);
    }
}

/**
 * AFTER: Encapsulated collection
 */
class TeamAfter {
    constructor() {
        this._players = [];
    }

    addPlayer(player) {
        this._players.push(player);
    }

    removePlayer(player) {
        const index = this._players.indexOf(player);
        if (index !== -1) {
            this._players.splice(index, 1);
        }
    }

    getPlayers() {
        return [...this._players]; // Return copy
    }

    getPlayerCount() {
        return this._players.length;
    }
}

/**
 * 29. Replacing a record with a Data Class
 *
 * BEFORE: Using array as data structure
 */
class EmployeeArray {
    createEmployee(data) {
        return {
            name: data.name,
            salary: data.salary,
            department: data.department
        };
    }

    getSalary(employee) {
        return employee.salary;
    }
}

/**
 * AFTER: Use data class
 */
class Employee {
    constructor(name, salary, department) {
        this.name = name;
        this.salary = salary;
        this.department = department;
    }

    getName() {
        return this.name;
    }

    getSalary() {
        return this.salary;
    }

    getDepartment() {
        return this.department;
    }
}

class EmployeeDataClass {
    createEmployee(name, salary, department) {
        return new Employee(name, salary, department);
    }

    getSalary(employee) {
        return employee.getSalary();
    }
}

/**
 * 30. Replacing Type Code with Class
 *
 * BEFORE: Type code as constants
 */
class EmployeeTypeCode {
    static ENGINEER = 0;
    static SALESMAN = 1;
    static MANAGER = 2;

    constructor(type) {
        this.type = type;
    }

    getTypeCode() {
        return this.type;
    }

    getMonthlySalary() {
        switch (this.type) {
            case EmployeeTypeCode.ENGINEER:
                return 5000;
            case EmployeeTypeCode.SALESMAN:
                return 4000;
            case EmployeeTypeCode.MANAGER:
                return 6000;
            default:
                return 0;
        }
    }
}

/**
 * AFTER: Replace type code with class
 */
class EmployeeType {
    getMonthlySalary() {
        throw new Error('Abstract method must be implemented');
    }

    static createEngineer() {
        return new EngineerType();
    }

    static createSalesman() {
        return new SalesmanType();
    }

    static createManager() {
        return new ManagerType();
    }
}

class EngineerType extends EmployeeType {
    getMonthlySalary() {
        return 5000;
    }
}

class SalesmanType extends EmployeeType {
    getMonthlySalary() {
        return 4000;
    }
}

class ManagerType extends EmployeeType {
    getMonthlySalary() {
        return 6000;
    }
}

class EmployeeTypeClass {
    constructor(type) {
        this.type = type;
    }

    getMonthlySalary() {
        return this.type.getMonthlySalary();
    }
}

/**
 * 31. Replacing Type Code with Subclasses
 *
 * BEFORE: Type code in base class
 */
class EmployeeSubBefore {
    static ENGINEER = 0;
    static SALESMAN = 1;
    static MANAGER = 2;

    constructor(type, salary) {
        this.type = type;
        this.salary = salary;
    }

    getSalary() {
        return this.salary;
    }

    getType() {
        return this.type;
    }
}

/**
 * AFTER: Replace type code with subclasses
 */
class EmployeeSubAfter {
    constructor(salary) {
        this.salary = salary;
    }

    getSalary() {
        return this.salary;
    }

    getType() {
        throw new Error('Abstract method must be implemented');
    }
}

class Engineer extends EmployeeSubAfter {
    getType() {
        return 'engineer';
    }
}

class Salesman extends EmployeeSubAfter {
    getType() {
        return 'salesman';
    }
}

class Manager extends EmployeeSubAfter {
    getType() {
        return 'manager';
    }
}

/**
 * 32. Replacing Type Code with State/Strategy
 *
 * BEFORE: Type code with behavior
 */
class EmployeeStateBefore {
    static JUNIOR = 0;
    static SENIOR = 1;
    static LEAD = 2;

    constructor(level) {
        this.level = level;
    }

    getSalaryMultiplier() {
        switch (this.level) {
            case EmployeeStateBefore.JUNIOR:
                return 1.0;
            case EmployeeStateBefore.SENIOR:
                return 1.5;
            case EmployeeStateBefore.LEAD:
                return 2.0;
            default:
                return 1.0;
        }
    }
}

/**
 * AFTER: Use state/strategy pattern
 */
class EmployeeLevel {
    getSalaryMultiplier() {
        throw new Error('Abstract method must be implemented');
    }
}

class JuniorLevel extends EmployeeLevel {
    getSalaryMultiplier() {
        return 1.0;
    }
}

class SeniorLevel extends EmployeeLevel {
    getSalaryMultiplier() {
        return 1.5;
    }
}

class LeadLevel extends EmployeeLevel {
    getSalaryMultiplier() {
        return 2.0;
    }
}

class EmployeeStateAfter {
    constructor(level) {
        this.level = level;
    }

    getSalaryMultiplier() {
        return this.level.getSalaryMultiplier();
    }
}

/**
 * 33. Replacing Subclass with Fields
 *
 * BEFORE: Unnecessary subclasses
 */
class PersonSub {
    constructor(name, gender) {
        this.name = name;
        this.gender = gender;
    }

    getName() {
        return this.name;
    }

    isMale() {
        throw new Error('Abstract method must be implemented');
    }
}

class Male extends PersonSub {
    isMale() {
        return true;
    }
}

class Female extends PersonSub {
    isMale() {
        return false;
    }
}

/**
 * AFTER: Replace subclass with field
 */
class PersonField {
    constructor(name, gender) {
        this.name = name;
        this.gender = gender; // 'male' or 'female'
    }

    getName() {
        return this.name;
    }

    isMale() {
        return this.gender === 'male';
    }

    getGender() {
        return this.gender;
    }
}

module.exports = {
    PersonBefore,
    PersonAfter,
    OrderBefore,
    Customer,
    OrderAfter,
    CustomerValue,
    OrderValue,
    CustomerReference,
    OrderReference,
    CurrencyReference,
    ProductReference,
    CurrencyValue,
    ProductValue,
    PerformanceArray,
    PerformanceData,
    PerformanceObject,
    OrderDomain,
    OrderDomainSeparated,
    OrderDisplay,
    CustomerUni,
    OrderUni,
    CustomerBi,
    OrderBi,
    CustomerBidirectional,
    OrderBidirectional,
    CustomerUnidirectional,
    OrderUnidirectional,
    GeometryBefore,
    GeometryAfter,
    PersonPublic,
    PersonEncapsulated,
    TeamBefore,
    TeamAfter,
    EmployeeArray,
    Employee,
    EmployeeDataClass,
    EmployeeTypeCode,
    EmployeeType,
    EngineerType,
    SalesmanType,
    ManagerType,
    EmployeeTypeClass,
    EmployeeSubBefore,
    EmployeeSubAfter,
    Engineer,
    Salesman,
    Manager,
    EmployeeStateBefore,
    EmployeeLevel,
    JuniorLevel,
    SeniorLevel,
    LeadLevel,
    EmployeeStateAfter,
    PersonSub,
    Male,
    Female,
    PersonField
};
