/**
 * 42. Renaming a method (Rename Method)
 *
 * BEFORE: Poorly named method
 */
class CalculatorBefore {
    calc(a, b) { // Unclear name
        return a + b;
    }
}

/**
 * AFTER: Rename method to be more descriptive
 */
class CalculatorAfter {
    add(a, b) {
        return a + b;
    }
}

/**
 * 43. Adding a parameter (Add Parameter)
 *
 * BEFORE: Method missing required parameter
 */
class EmailSenderBefore {
    sendEmail(to, subject, body) {
        // Send email with default priority
        const priority = 'normal';
        // Send logic
    }
}

/**
 * AFTER: Add parameter
 */
class EmailSenderAfter {
    sendEmail(to, subject, body, priority = 'normal') {
        // Send logic with priority
    }
}

/**
 * 44. Deleting a parameter (Remove Parameter)
 *
 * BEFORE: Unnecessary parameter
 */
class ReportGeneratorBefore {
    generateReport(data, format, includeHeader = true) {
        if (format === 'html') {
            // Always include header for HTML
            // includeHeader = true; // This would be ignored
        }
        // Generate report
    }
}

/**
 * AFTER: Remove unnecessary parameter
 */
class ReportGeneratorAfter {
    generateReport(data, format) {
        const includeHeader = (format === 'html');
        // Generate report
    }
}

/**
 * 45. Separation of Query and Modifier (Separate Query from Modifier)
 *
 * BEFORE: Method that both queries and modifies
 */
class BankAccountBefore {
    constructor() {
        this.balance = 0;
    }

    withdraw(amount) {
        if (this.balance >= amount) {
            this.balance -= amount;
            return true;
        }
        return false;
    }
}

/**
 * AFTER: Separate query from modifier
 */
class BankAccountAfter {
    constructor() {
        this.balance = 0;
    }

    canWithdraw(amount) {
        return this.balance >= amount;
    }

    withdraw(amount) {
        if (this.canWithdraw(amount)) {
            this.balance -= amount;
            return true;
        }
        return false;
    }
}

/**
 * 46. Parameterization of the method (Parameterize Method)
 *
 * BEFORE: Similar methods with different values
 */
class ReportGeneratorParamBefore {
    generateWeeklyReport() {
        return this.generateReport(7);
    }

    generateMonthlyReport() {
        return this.generateReport(30);
    }

    generateQuarterlyReport() {
        return this.generateReport(90);
    }

    generateReport(days) {
        // Generate report for specified days
    }
}

/**
 * AFTER: Parameterize method
 */
class ReportGeneratorParamAfter {
    generateReport(days) {
        // Generate report for specified days
    }

    generateWeeklyReport() {
        return this.generateReport(7);
    }

    generateMonthlyReport() {
        return this.generateReport(30);
    }

    generateQuarterlyReport() {
        return this.generateReport(90);
    }
}

/**
 * 47. Replacing a parameter with explicit methods (Replace Parameter with Explicit Methods)
 *
 * BEFORE: Parameter determines behavior
 */
class EmployeeExplicitBefore {
    static ENGINEER = 0;
    static SALESMAN = 1;
    static MANAGER = 2;

    getSalary(baseSalary, type) {
        switch (type) {
            case EmployeeExplicitBefore.ENGINEER:
                return baseSalary * 1.0;
            case EmployeeExplicitBefore.SALESMAN:
                return baseSalary * 1.1;
            case EmployeeExplicitBefore.MANAGER:
                return baseSalary * 1.2;
            default:
                return baseSalary;
        }
    }
}

/**
 * AFTER: Replace parameter with explicit methods
 */
class EmployeeExplicitAfter {
    getEngineerSalary(baseSalary) {
        return baseSalary * 1.0;
    }

    getSalesmanSalary(baseSalary) {
        return baseSalary * 1.1;
    }

    getManagerSalary(baseSalary) {
        return baseSalary * 1.2;
    }
}

/**
 * 48. Save the Whole Object
 *
 * BEFORE: Passing individual fields
 */
class OrderWholeBefore {
    constructor(customerName, customerAddress) {
        this.customer = { name: customerName, address: customerAddress };
    }

    calculateShipping() {
        return this.getShippingCost(this.customer.name, this.customer.address);
    }

    getShippingCost(name, address) {
        // Calculate based on name and address
        return 10.0;
    }
}

/**
 * AFTER: Pass whole object
 */
class Customer {
    constructor(name, address) {
        this.name = name;
        this.address = address;
    }

    getName() {
        return this.name;
    }

    getAddress() {
        return this.address;
    }
}

class OrderWholeAfter {
    constructor(customer) {
        this.customer = customer;
    }

    calculateShipping() {
        return this.getShippingCost(this.customer);
    }

    getShippingCost(customer) {
        // Calculate based on customer object
        return 10.0;
    }
}

/**
 * 49. Replacing a parameter with a method call (Replace Parameter with Method)
 *
 * BEFORE: Parameter calculated outside method
 */
class DiscountCalculatorParamBefore {
    calculateDiscount(price, customerType) {
        // customerType passed in
        return price * this.getDiscountRate(customerType);
    }

    getDiscountRate(customerType) {
        switch (customerType) {
            case 'premium':
                return 0.1;
            case 'regular':
                return 0.05;
            default:
                return 0.0;
        }
    }
}

class OrderParamBefore {
    constructor() {
        this.customer = { type: 'regular' };
    }

    getDiscountedPrice(price) {
        const calculator = new DiscountCalculatorParamBefore();
        return calculator.calculateDiscount(price, this.customer.type);
    }
}

/**
 * AFTER: Replace parameter with method call
 */
class DiscountCalculatorParamAfter {
    calculateDiscount(price, customer) {
        return price * this.getDiscountRate(customer.getType());
    }

    getDiscountRate(customerType) {
        switch (customerType) {
            case 'premium':
                return 0.1;
            case 'regular':
                return 0.05;
            default:
                return 0.0;
        }
    }
}

class OrderParamAfter {
    constructor() {
        this.customer = {
            getType: () => 'regular'
        };
    }

    getDiscountedPrice(price) {
        const calculator = new DiscountCalculatorParamAfter();
        return calculator.calculateDiscount(price, this.customer);
    }
}

/**
 * 50. Introduction of the boundary object (Introduce Parameter Object)
 *
 * BEFORE: Multiple parameters
 */
class TemperatureRangeBefore {
    withinRange(minTemp, maxTemp, currentTemp) {
        return currentTemp >= minTemp && currentTemp <= maxTemp;
    }

    getAverageTemp(minTemp, maxTemp) {
        return (minTemp + maxTemp) / 2;
    }
}

/**
 * AFTER: Introduce parameter object
 */
class TemperatureRange {
    constructor(minTemp, maxTemp) {
        this.minTemp = minTemp;
        this.maxTemp = maxTemp;
    }

    getMinTemp() {
        return this.minTemp;
    }

    getMaxTemp() {
        return this.maxTemp;
    }

    withinRange(currentTemp) {
        return currentTemp >= this.minTemp && currentTemp <= this.maxTemp;
    }

    getAverageTemp() {
        return (this.minTemp + this.maxTemp) / 2;
    }
}

/**
 * 51. Removing the Value Setting Method
 *
 * BEFORE: Setter that's not needed
 */
class SensorBefore {
    constructor(temperature) {
        this.temperature = temperature;
    }

    getTemperature() {
        return this.temperature;
    }

    setTemperature(temperature) { // Not needed if immutable
        this.temperature = temperature;
    }
}

/**
 * AFTER: Remove setter for immutable object
 */
class SensorAfter {
    constructor(temperature) {
        this.temperature = temperature;
    }

    getTemperature() {
        return this.temperature;
    }

    // setTemperature removed
}

/**
 * 52. Hiding a method (Hide Method)
 *
 * BEFORE: Public method that should be private
 */
class DataProcessorHideBefore {
    isValidData(data) { // Should be private
        return data !== null && typeof data === 'object';
    }

    processData(data) {
        if (this.isValidData(data)) {
            // Process data
        }
    }
}

/**
 * AFTER: Hide method
 */
class DataProcessorHideAfter {
    #isValidData(data) { // Private method
        return data !== null && typeof data === 'object';
    }

    processData(data) {
        if (this.#isValidData(data)) {
            // Process data
        }
    }
}

/**
 * 53. Replacing the constructor with the factory method (Replace Constructor with Factory Method)
 *
 * BEFORE: Complex constructor
 */
class ComplexObjectBefore {
    constructor(type, config = {}) {
        this.type = type;
        this.config = config;

        if (type === 'database') {
            this.config = { ...{ host: 'localhost', port: 3306 }, ...config };
        } else if (type === 'file') {
            this.config = { ...{ path: '/tmp', format: 'json' }, ...config };
        }
    }
}

/**
 * AFTER: Replace constructor with factory method
 */
class ComplexObjectAfter {
    constructor(type, config) {
        this.type = type;
        this.config = config;
    }

    static createDatabaseConnection(config = {}) {
        const finalConfig = { ...{ host: 'localhost', port: 3306 }, ...config };
        return new ComplexObjectAfter('database', finalConfig);
    }

    static createFileHandler(config = {}) {
        const finalConfig = { ...{ path: '/tmp', format: 'json' }, ...config };
        return new ComplexObjectAfter('file', finalConfig);
    }
}

/**
 * 54. Encapsulation of top-down type conversion (Encapsulate Downcast)
 *
 * BEFORE: Downcast in client code
 */
class ShapeCollectionBefore {
    constructor() {
        this.shapes = [];
    }

    addShape(shape) {
        this.shapes.push(shape);
    }

    getShapes() {
        return this.shapes;
    }
}

// Client code
// const collection = new ShapeCollectionBefore();
// // ... add shapes
// const circles = collection.getShapes().filter(shape => shape instanceof Circle);

/**
 * AFTER: Encapsulate downcast
 */
class ShapeCollectionAfter {
    constructor() {
        this.shapes = [];
    }

    addShape(shape) {
        this.shapes.push(shape);
    }

    getCircles() {
        return this.shapes.filter(shape => shape.constructor.name === 'Circle');
    }

    getSquares() {
        return this.shapes.filter(shape => shape.constructor.name === 'Square');
    }
}

/**
 * 55. Replacing the error code with an exceptional situation (Replace Error Code with Exception)
 *
 * BEFORE: Error codes
 */
class FileReaderErrorBefore {
    static FILE_NOT_FOUND = 1;
    static PERMISSION_DENIED = 2;

    readFile(filename) {
        if (!require('fs').existsSync(filename)) {
            return FileReaderErrorBefore.FILE_NOT_FOUND;
        }

        try {
            require('fs').accessSync(filename, require('fs').constants.R_OK);
        } catch (e) {
            return FileReaderErrorBefore.PERMISSION_DENIED;
        }

        return require('fs').readFileSync(filename, 'utf8');
    }
}

// Client code
// const reader = new FileReaderErrorBefore();
// const result = reader.readFile('test.txt');
// if (result === FileReaderErrorBefore.FILE_NOT_FOUND) {
//     // Handle error
// } else if (result === FileReaderErrorBefore.PERMISSION_DENIED) {
//     // Handle error
// } else {
//     // Use content
// }

/**
 * AFTER: Replace error codes with exceptions
 */
class FileNotFoundException extends Error {
    constructor(message) {
        super(message);
        this.name = 'FileNotFoundException';
    }
}

class PermissionDeniedException extends Error {
    constructor(message) {
        super(message);
        this.name = 'PermissionDeniedException';
    }
}

class FileReaderExceptionAfter {
    readFile(filename) {
        const fs = require('fs');

        if (!fs.existsSync(filename)) {
            throw new FileNotFoundException(`File not found: ${filename}`);
        }

        try {
            fs.accessSync(filename, fs.constants.R_OK);
        } catch (e) {
            throw new PermissionDeniedException(`Permission denied: ${filename}`);
        }

        return fs.readFileSync(filename, 'utf8');
    }
}

// Client code
// try {
//     const reader = new FileReaderExceptionAfter();
//     const content = reader.readFile('test.txt');
//     // Use content
// } catch (e) {
//     if (e instanceof FileNotFoundException) {
//         // Handle file not found
//     } else if (e instanceof PermissionDeniedException) {
//         // Handle permission denied
//     }
// }

/**
 * 56. Replacing an exceptional situation with a check (Replace Exception with Test)
 *
 * BEFORE: Using exception for control flow
 */
class StackExceptionBefore {
    constructor() {
        this.items = [];
    }

    pop() {
        if (this.items.length === 0) {
            throw new Error('Stack is empty');
        }
        return this.items.pop();
    }
}

/**
 * AFTER: Replace exception with test
 */
class StackTestAfter {
    constructor() {
        this.items = [];
    }

    isEmpty() {
        return this.items.length === 0;
    }

    pop() {
        return this.items.pop();
    }
}

// Client code
// const stack = new StackTestAfter();
// const item = stack.isEmpty() ? null : stack.pop();

module.exports = {
    CalculatorBefore,
    CalculatorAfter,
    EmailSenderBefore,
    EmailSenderAfter,
    ReportGeneratorBefore,
    ReportGeneratorAfter,
    BankAccountBefore,
    BankAccountAfter,
    ReportGeneratorParamBefore,
    ReportGeneratorParamAfter,
    EmployeeExplicitBefore,
    EmployeeExplicitAfter,
    OrderWholeBefore,
    Customer,
    OrderWholeAfter,
    DiscountCalculatorParamBefore,
    OrderParamBefore,
    DiscountCalculatorParamAfter,
    OrderParamAfter,
    TemperatureRangeBefore,
    TemperatureRange,
    SensorBefore,
    SensorAfter,
    DataProcessorHideBefore,
    DataProcessorHideAfter,
    ComplexObjectBefore,
    ComplexObjectAfter,
    ShapeCollectionBefore,
    ShapeCollectionAfter,
    FileReaderErrorBefore,
    FileNotFoundException,
    PermissionDeniedException,
    FileReaderExceptionAfter,
    StackExceptionBefore,
    StackTestAfter
};
