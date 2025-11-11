/**
 * 57. Lifting the field (Pull Up Field)
 *
 * BEFORE: Duplicate fields in subclasses
 */
class EmployeePullBefore {
    constructor() {
        this.name = null;
    }
}

class ManagerPullBefore extends EmployeePullBefore {
    constructor() {
        super();
        this.name = null; // Duplicate
        this.budget = null;
    }
}

class EngineerPullBefore extends EmployeePullBefore {
    constructor() {
        super();
        this.name = null; // Duplicate
        this.skills = null;
    }
}

/**
 * AFTER: Pull up field to superclass
 */
class EmployeePullAfter {
    constructor() {
        this.name = null;
    }
}

class ManagerPullAfter extends EmployeePullAfter {
    constructor() {
        super();
        this.budget = null;
    }
}

class EngineerPullAfter extends EmployeePullAfter {
    constructor() {
        super();
        this.skills = null;
    }
}

/**
 * 58. Lifting the method (Pull Up Method)
 *
 * BEFORE: Duplicate methods in subclasses
 */
class ShapeMethodBefore {
    area() {
        throw new Error('Abstract method must be implemented');
    }
}

class CircleMethodBefore extends ShapeMethodBefore {
    constructor(radius) {
        super();
        this.radius = radius;
    }

    area() {
        return Math.PI * this.radius * this.radius;
    }

    circumference() {
        return 2 * Math.PI * this.radius;
    }
}

class SquareMethodBefore extends ShapeMethodBefore {
    constructor(side) {
        super();
        this.side = side;
    }

    area() {
        return this.side * this.side;
    }

    circumference() { // Duplicate logic
        return 4 * this.side;
    }
}

/**
 * AFTER: Pull up method to superclass
 */
class ShapeMethodAfter {
    area() {
        throw new Error('Abstract method must be implemented');
    }

    circumference() {
        throw new Error('Abstract method must be implemented');
    }
}

class CircleMethodAfter extends ShapeMethodAfter {
    constructor(radius) {
        super();
        this.radius = radius;
    }

    area() {
        return Math.PI * this.radius * this.radius;
    }

    circumference() {
        return 2 * Math.PI * this.radius;
    }
}

class SquareMethodAfter extends ShapeMethodAfter {
    constructor(side) {
        super();
        this.side = side;
    }

    area() {
        return this.side * this.side;
    }

    circumference() {
        return 4 * this.side;
    }
}

/**
 * 59. Lifting the constructor Body (Pull Up Constructor Body)
 *
 * BEFORE: Duplicate constructor code
 */
class VehicleConstructorBefore {
    constructor(make, model, year) {
        this.make = make;
        this.model = model;
        this.year = year;
    }
}

class CarConstructorBefore extends VehicleConstructorBefore {
    constructor(make, model, year, doors) {
        super(make, model, year); // Duplicate calls
        this.doors = doors;
    }
}

class TruckConstructorBefore extends VehicleConstructorBefore {
    constructor(make, model, year, payload) {
        super(make, model, year); // Duplicate calls
        this.payload = payload;
    }
}

/**
 * AFTER: Pull up constructor body (already properly structured in JS)
 */
class VehicleConstructorAfter {
    constructor(make, model, year) {
        this.make = make;
        this.model = model;
        this.year = year;
    }
}

class CarConstructorAfter extends VehicleConstructorAfter {
    constructor(make, model, year, doors) {
        super(make, model, year);
        this.doors = doors;
    }
}

class TruckConstructorAfter extends VehicleConstructorAfter {
    constructor(make, model, year, payload) {
        super(make, model, year);
        this.payload = payload;
    }
}

/**
 * 60. Method Descent (Push Down Method)
 *
 * BEFORE: Method in wrong class hierarchy level
 */
class AnimalPushBefore {
    speak() {
        // Generic implementation
        return 'Generic animal sound';
    }
}

class DogPushBefore extends AnimalPushBefore {
    speak() {
        return 'Woof';
    }
}

class CatPushBefore extends AnimalPushBefore {
    speak() {
        return 'Meow';
    }
}

class FishPushBefore extends AnimalPushBefore {
    // Fish don't speak, but inherits speak method
}

/**
 * AFTER: Push down method to appropriate subclasses
 */
class AnimalPushAfter {
    // No speak method here
}

class DogPushAfter extends AnimalPushAfter {
    speak() {
        return 'Woof';
    }
}

class CatPushAfter extends AnimalPushAfter {
    speak() {
        return 'Meow';
    }
}

class FishPushAfter extends AnimalPushAfter {
    // No speak method - appropriate for Fish
}

/**
 * 61. Field Descent (Push Down Field)
 *
 * BEFORE: Field in wrong hierarchy level
 */
class EmployeeFieldBefore {
    constructor() {
        this.salary = null; // Not all employees have salary
    }
}

class SalariedEmployeeFieldBefore extends EmployeeFieldBefore {
    // Uses salary
}

class ContractorFieldBefore extends EmployeeFieldBefore {
    // Doesn't use salary, but inherits it
}

/**
 * AFTER: Push down field
 */
class EmployeeFieldAfter {
    // No salary field
}

class SalariedEmployeeFieldAfter extends EmployeeFieldAfter {
    constructor() {
        super();
        this.salary = null;
    }
}

class ContractorFieldAfter extends EmployeeFieldAfter {
    constructor() {
        super();
        this.hourlyRate = null;
    }
}

/**
 * 62. Subclass extraction (Extract Subclass)
 *
 * BEFORE: Class with conditional behavior
 */
class JobExtractBefore {
    constructor(type, rate, commission = null) {
        this.type = type;
        this.rate = rate;
        this.commission = commission;
    }

    getPay() {
        if (this.type === 'salaried') {
            return this.rate;
        } else {
            return this.rate + this.commission;
        }
    }
}

/**
 * AFTER: Extract subclass
 */
class JobExtractAfter {
    constructor(rate) {
        this.rate = rate;
    }

    getPay() {
        throw new Error('Abstract method must be implemented');
    }
}

class SalariedJob extends JobExtractAfter {
    getPay() {
        return this.rate;
    }
}

class CommissionedJob extends JobExtractAfter {
    constructor(rate, commission) {
        super(rate);
        this.commission = commission;
    }

    getPay() {
        return this.rate + this.commission;
    }
}

/**
 * 63. Allocation of the parent class (Extract Superclass)
 *
 * BEFORE: Duplicate code in classes
 */
class DepartmentSuperBefore {
    constructor(name, head) {
        this.name = name;
        this.head = head;
    }

    getName() {
        return this.name;
    }

    getHead() {
        return this.head;
    }
}

class CompanySuperBefore {
    constructor(name, head) {
        this.name = name;
        this.head = head;
    }

    getName() {
        return this.name;
    }

    getHead() {
        return this.head;
    }
}

/**
 * AFTER: Extract superclass
 */
class Party {
    constructor(name, head) {
        this.name = name;
        this.head = head;
    }

    getName() {
        return this.name;
    }

    getHead() {
        return this.head;
    }
}

class DepartmentSuperAfter extends Party {}

class CompanySuperAfter extends Party {}

/**
 * 64. Interface extraction (Extract Interface)
 *
 * BEFORE: Clients depend on concrete class
 */
class PrinterInterfaceBefore {
    print(document) {
        // Print logic
    }

    getStatus() {
        // Status logic
    }

    cancelJob(jobId) {
        // Cancel logic
    }
}

/**
 * AFTER: Extract interface (using class as interface simulation)
 */
class Printer {
    print(document) {
        throw new Error('Abstract method must be implemented');
    }

    getStatus() {
        throw new Error('Abstract method must be implemented');
    }
}

class LaserPrinter extends Printer {
    print(document) {
        // Print logic
    }

    getStatus() {
        // Status logic
    }

    cancelJob(jobId) {
        // Cancel logic - not part of interface
    }
}

class InkjetPrinter extends Printer {
    print(document) {
        // Print logic
    }

    getStatus() {
        // Status logic
    }
}

/**
 * 65. Collapse Hierarchy
 *
 * BEFORE: Unnecessary class hierarchy
 */
class EmployeeCollapseBefore {
    constructor() {
        this.name = null;
    }
}

class ManagerCollapseBefore extends EmployeeCollapseBefore {
    constructor() {
        super();
        this.department = null;
    }
}

/**
 * AFTER: Collapse hierarchy if only one subclass
 */
class EmployeeCollapseAfter {
    constructor() {
        this.name = null;
        this.department = null; // Moved up
    }
}

/**
 * 66. Formation of the method template (Form Template Method)
 *
 * BEFORE: Duplicate algorithm structure
 */
class ReportGeneratorTemplateBefore {
    generateHTMLReport() {
        const data = this.getData();
        const header = this.formatHeader();
        const body = this.formatBody(data);
        const footer = this.formatFooter();
        return header + body + footer;
    }

    generatePDFReport() {
        const data = this.getData(); // Duplicate
        const header = this.formatPDFHeader(); // Different
        const body = this.formatPDFBody(data); // Different
        const footer = this.formatPDFFooter(); // Different
        return header + body + footer;
    }

    getData() {
        return ['item1', 'item2'];
    }

    formatHeader() {
        return '<h1>Report</h1>';
    }

    formatBody(data) {
        return '<body>' + data.join('') + '</body>';
    }

    formatFooter() {
        return '<footer>End</footer>';
    }

    formatPDFHeader() {
        return 'PDF Report Header';
    }

    formatPDFBody(data) {
        return 'PDF Body: ' + data.join('');
    }

    formatPDFFooter() {
        return 'PDF Footer';
    }
}

/**
 * AFTER: Form template method
 */
class ReportGeneratorTemplateAfter {
    generateReport() {
        const data = this.getData();
        const header = this.formatHeader();
        const body = this.formatBody(data);
        const footer = this.formatFooter();
        return this.assembleReport(header, body, footer);
    }

    getData() {
        return ['item1', 'item2'];
    }

    formatHeader() {
        throw new Error('Abstract method must be implemented');
    }

    formatBody(data) {
        throw new Error('Abstract method must be implemented');
    }

    formatFooter() {
        throw new Error('Abstract method must be implemented');
    }

    assembleReport(header, body, footer) {
        throw new Error('Abstract method must be implemented');
    }
}

class HTMLReportGenerator extends ReportGeneratorTemplateAfter {
    formatHeader() {
        return '<h1>Report</h1>';
    }

    formatBody(data) {
        return '<body>' + data.join('') + '</body>';
    }

    formatFooter() {
        return '<footer>End</footer>';
    }

    assembleReport(header, body, footer) {
        return header + body + footer;
    }
}

class PDFReportGenerator extends ReportGeneratorTemplateAfter {
    formatHeader() {
        return 'PDF Report Header';
    }

    formatBody(data) {
        return 'PDF Body: ' + data.join('');
    }

    formatFooter() {
        return 'PDF Footer';
    }

    assembleReport(header, body, footer) {
        return header + body + footer;
    }
}

/**
 * 67. Replacement of inheritance by delegation (Replace Inheritance with Delegation)
 *
 * BEFORE: Inheritance where delegation would be better
 */
class StackInheritanceBefore extends Array {
    push(item) {
        super.push(item);
    }

    pop() {
        if (this.length === 0) {
            throw new Error('Stack is empty');
        }
        return super.pop();
    }
}

/**
 * AFTER: Replace inheritance with delegation
 */
class StackDelegationAfter {
    constructor() {
        this.items = [];
    }

    push(item) {
        this.items.push(item);
    }

    pop() {
        if (this.items.length === 0) {
            throw new Error('Stack is empty');
        }
        return this.items.pop();
    }

    get length() {
        return this.items.length;
    }
}

/**
 * 68. Replacement of delegation by inheritance (Replace Delegation with Inheritance)
 *
 * BEFORE: Delegation where inheritance would be simpler
 */
class MyStringDelegateBefore {
    constructor(string) {
        this.string = string;
    }

    length() {
        return this.string.length;
    }

    substr(start, length = undefined) {
        return this.string.substr(start, length);
    }

    indexOf(needle) {
        return this.string.indexOf(needle);
    }
}

/**
 * AFTER: Replace delegation with inheritance
 */
class MyStringInheritAfter extends String {
    constructor(string) {
        super(string);
    }

    toMyString() {
        return this.toString();
    }
}

module.exports = {
    EmployeePullBefore,
    ManagerPullBefore,
    EngineerPullBefore,
    EmployeePullAfter,
    ManagerPullAfter,
    EngineerPullAfter,
    ShapeMethodBefore,
    CircleMethodBefore,
    SquareMethodBefore,
    ShapeMethodAfter,
    CircleMethodAfter,
    SquareMethodAfter,
    VehicleConstructorBefore,
    CarConstructorBefore,
    TruckConstructorBefore,
    VehicleConstructorAfter,
    CarConstructorAfter,
    TruckConstructorAfter,
    AnimalPushBefore,
    DogPushBefore,
    CatPushBefore,
    FishPushBefore,
    AnimalPushAfter,
    DogPushAfter,
    CatPushAfter,
    FishPushAfter,
    EmployeeFieldBefore,
    SalariedEmployeeFieldBefore,
    ContractorFieldBefore,
    EmployeeFieldAfter,
    SalariedEmployeeFieldAfter,
    ContractorFieldAfter,
    JobExtractBefore,
    JobExtractAfter,
    SalariedJob,
    CommissionedJob,
    DepartmentSuperBefore,
    CompanySuperBefore,
    Party,
    DepartmentSuperAfter,
    CompanySuperAfter,
    PrinterInterfaceBefore,
    Printer,
    LaserPrinter,
    InkjetPrinter,
    EmployeeCollapseBefore,
    ManagerCollapseBefore,
    EmployeeCollapseAfter,
    ReportGeneratorTemplateBefore,
    ReportGeneratorTemplateAfter,
    HTMLReportGenerator,
    PDFReportGenerator,
    StackInheritanceBefore,
    StackDelegationAfter,
    MyStringDelegateBefore,
    MyStringInheritAfter
};
