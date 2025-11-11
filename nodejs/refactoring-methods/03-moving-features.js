/**
 * 9. Substitution Algorithm
 *
 * BEFORE: Complex algorithm that can be simplified
 */
class PricingServiceBefore {
    calculatePrice(items) {
        let total = 0;
        for (const item of items) {
            if (item.type === 'book') {
                total += item.price * 0.9; // 10% discount for books
            } else if (item.type === 'electronics') {
                total += item.price * 1.1; // 10% markup for electronics
            } else {
                total += item.price;
            }
        }
        return total;
    }
}

/**
 * AFTER: Substitute with a simpler algorithm
 */
class PricingServiceAfter {
    constructor() {
        this.discounts = {
            'book': 0.9,
            'electronics': 1.1,
            'default': 1.0
        };
    }

    calculatePrice(items) {
        let total = 0;
        for (const item of items) {
            const multiplier = this.discounts[item.type] || this.discounts['default'];
            total += item.price * multiplier;
        }
        return total;
    }
}

/**
 * 10. Moving functions between objects (Move Method)
 *
 * BEFORE: Method in wrong class
 */
class AccountBefore {
    constructor(balance) {
        this.balance = balance;
    }

    getBalance() {
        return this.balance;
    }

    // This method belongs in Bank class, not Account
    transferTo(target, amount) {
        if (this.balance >= amount) {
            this.balance -= amount;
            target.balance += amount;
            return true;
        }
        return false;
    }
}

/**
 * AFTER: Move method to appropriate class
 */
class AccountAfter {
    constructor(balance) {
        this.balance = balance;
    }

    getBalance() {
        return this.balance;
    }

    decreaseBalance(amount) {
        this.balance -= amount;
    }

    increaseBalance(amount) {
        this.balance += amount;
    }
}

class Bank {
    transfer(from, to, amount) {
        if (from.getBalance() >= amount) {
            from.decreaseBalance(amount);
            to.increaseBalance(amount);
            return true;
        }
        return false;
    }
}

/**
 * 11. Moving the field (Move Field)
 *
 * BEFORE: Field in wrong class
 */
class CustomerBefore {
    constructor(name, street, city, zipCode) {
        this.name = name;
        this.address = { street, city, zipCode };
    }

    getAddress() {
        return `${this.address.street}, ${this.address.city} ${this.address.zipCode}`;
    }
}

/**
 * AFTER: Move field to dedicated class
 */
class Address {
    constructor(street, city, zipCode) {
        this.street = street;
        this.city = city;
        this.zipCode = zipCode;
    }

    getFullAddress() {
        return `${this.street}, ${this.city} ${this.zipCode}`;
    }
}

class CustomerAfter {
    constructor(name, address) {
        this.name = name;
        this.address = address;
    }

    getAddress() {
        return this.address.getFullAddress();
    }
}

/**
 * 12. Class Allocation (Extract Class)
 *
 * BEFORE: Class has too many responsibilities
 */
class PersonBefore {
    constructor(name) {
        this.name = name;
        this.phoneNumber = null;
        this.officeAreaCode = null;
        this.officeNumber = null;
    }

    getTelephoneNumber() {
        return `(${this.officeAreaCode}) ${this.officeNumber}`;
    }
}

/**
 * AFTER: Extract telephone number to separate class
 */
class TelephoneNumber {
    constructor(areaCode, number) {
        this.areaCode = areaCode;
        this.number = number;
    }

    getTelephoneNumber() {
        return `(${this.areaCode}) ${this.number}`;
    }
}

class PersonAfter {
    constructor(name) {
        this.name = name;
        this.phoneNumber = null;
        this.officeTelephone = null;
    }

    getOfficeTelephone() {
        return this.officeTelephone ? this.officeTelephone.getTelephoneNumber() : null;
    }

    setOfficeTelephone(telephone) {
        this.officeTelephone = telephone;
    }
}

/**
 * 13. Embedding a class (Inline Class)
 *
 * BEFORE: Unnecessary class with single responsibility
 */
class OrderProcessorBefore {
    constructor() {
        this.validator = new OrderValidator();
    }

    process(order) {
        if (this.validator.isValid(order)) {
            // Process order
        }
    }
}

class OrderValidator {
    isValid(order) {
        return order.total > 0;
    }
}

/**
 * AFTER: Inline the class
 */
class OrderProcessorAfter {
    process(order) {
        if (this.isValidOrder(order)) {
            // Process order
        }
    }

    isValidOrder(order) {
        return order.total > 0;
    }
}

/**
 * 14. Hiding delegation (Hide Delegate)
 *
 * BEFORE: Client has to know about delegation
 */
class DepartmentBefore {
    constructor(manager) {
        this.manager = manager;
    }

    getManager() {
        return this.manager;
    }
}

class Person {
    constructor(department) {
        this.department = department;
    }

    getDepartment() {
        return this.department;
    }
}

// Client code
// const manager = person.getDepartment().getManager();

/**
 * AFTER: Hide the delegation
 */
class DepartmentAfter {
    constructor(manager) {
        this.manager = manager;
    }

    getManager() {
        return this.manager;
    }
}

class PersonAfter {
    constructor(department) {
        this.department = department;
    }

    getDepartment() {
        return this.department;
    }

    getManager() {
        return this.department.getManager();
    }
}

// Client code - much cleaner
// const manager = person.getManager();

/**
 * 15. Removing the intermediary (Remove Middle Man)
 *
 * BEFORE: Too much delegation
 */
class PersonWithMiddleMan {
    constructor(department) {
        this.department = department;
    }

    getDepartment() {
        return this.department;
    }

    getManager() {
        return this.department.getManager();
    }

    getDepartmentName() {
        return this.department.getName();
    }
}

/**
 * AFTER: Remove middle man if delegation is too heavy
 */
class PersonDirect {
    constructor(department, manager) {
        this.department = department;
        this.manager = manager; // Direct reference
    }

    getManager() {
        return this.manager;
    }

    getDepartment() {
        return this.department;
    }
}

/**
 * 16. Introduction of an external method (Introduce Foreign Method)
 *
 * BEFORE: Using external class method in wrong place
 */
class ReportGeneratorBefore {
    generateReport() {
        const date = new Date();
        const nextMonth = this.addMonths(date, 1); // Foreign method usage

        // Generate report for next month
    }

    addMonths(date, months) {
        const result = new Date(date);
        result.setMonth(result.getMonth() + months);
        return result;
    }
}

/**
 * AFTER: Introduce foreign method
 */
class ReportGeneratorAfter {
    generateReport() {
        const date = new Date();
        const nextMonth = this.nextMonth(date);

        // Generate report for next month
    }

    nextMonth(date) {
        const result = new Date(date);
        result.setMonth(result.getMonth() + 1);
        return result;
    }
}

/**
 * 17. The introduction of local extension (Introduce Local Extension)
 *
 * BEFORE: Adding methods to external class (not possible)
 */

// DateUtil would be a utility class
class DateUtil {
    static nextMonth(date) {
        const result = new Date(date);
        result.setMonth(result.getMonth() + 1);
        return result;
    }

    static previousMonth(date) {
        const result = new Date(date);
        result.setMonth(result.getMonth() - 1);
        return result;
    }
}

/**
 * AFTER: Create local extension class
 */
class DateTimeExtension extends Date {
    nextMonth() {
        const result = new Date(this);
        result.setMonth(result.getMonth() + 1);
        return result;
    }

    previousMonth() {
        const result = new Date(this);
        result.setMonth(result.getMonth() - 1);
        return result;
    }
}

// Example usage
// const extendedDate = new DateTimeExtension();
// const nextMonth = extendedDate.nextMonth();

module.exports = {
    PricingServiceBefore,
    PricingServiceAfter,
    AccountBefore,
    AccountAfter,
    Bank,
    CustomerBefore,
    Address,
    CustomerAfter,
    PersonBefore,
    TelephoneNumber,
    PersonAfter,
    OrderProcessorBefore,
    OrderValidator,
    OrderProcessorAfter,
    DepartmentBefore,
    Person,
    DepartmentAfter,
    PersonAfter,
    PersonWithMiddleMan,
    PersonDirect,
    ReportGeneratorBefore,
    ReportGeneratorAfter,
    DateUtil,
    DateTimeExtension
};
