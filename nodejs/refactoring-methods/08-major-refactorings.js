/**
 * 69. Separation of inheritance (Tease Apart Inheritance)
 *
 * BEFORE: Class hierarchy mixing two different responsibilities
 */
class EmployeeTeaseBefore {
    constructor(name, rate) {
        this.name = name;
        this.rate = rate;
    }

    getName() {
        return this.name;
    }
}

class SalariedEmployeeTeaseBefore extends EmployeeTeaseBefore {
    getPay() {
        return this.rate;
    }
}

class CommissionedEmployeeTeaseBefore extends EmployeeTeaseBefore {
    constructor(name, rate, commission) {
        super(name, rate);
        this.commission = commission;
    }

    getPay() {
        return this.rate + this.commission;
    }
}

/**
 * AFTER: Tease apart inheritance into two separate hierarchies
 */
interface Payable {
    getPay();
}

class EmployeeTeaseAfter {
    constructor(name) {
        this.name = name;
    }

    getName() {
        return this.name;
    }
}

class SalariedEmployeeTeaseAfter extends EmployeeTeaseAfter {
    constructor(name, salary) {
        super(name);
        this.salary = salary;
    }

    getPay() {
        return this.salary;
    }
}

class CommissionedEmployeeTeaseAfter extends EmployeeTeaseAfter {
    constructor(name, baseSalary, commission) {
        super(name);
        this.baseSalary = baseSalary;
        this.commission = commission;
    }

    getPay() {
        return this.baseSalary + this.commission;
    }
}

/**
 * 70. Converting a procedural project into objects (Convert Procedural Design to Objects)
 *
 * BEFORE: Procedural code with global functions and data
 */
class ProceduralDesignBefore {
    static accounts = new Map();

    static createAccount(id, balance) {
        ProceduralDesignBefore.accounts.set(id, balance);
    }

    static getBalance(id) {
        return ProceduralDesignBefore.accounts.get(id) || 0;
    }

    static deposit(id, amount) {
        if (ProceduralDesignBefore.accounts.has(id)) {
            const current = ProceduralDesignBefore.accounts.get(id);
            ProceduralDesignBefore.accounts.set(id, current + amount);
        }
    }

    static withdraw(id, amount) {
        if (ProceduralDesignBefore.accounts.has(id) && ProceduralDesignBefore.accounts.get(id) >= amount) {
            const current = ProceduralDesignBefore.accounts.get(id);
            ProceduralDesignBefore.accounts.set(id, current - amount);
            return true;
        }
        return false;
    }
}

/**
 * AFTER: Convert to object-oriented design
 */
class Account {
    constructor(id, balance = 0) {
        this.id = id;
        this.balance = balance;
    }

    getId() {
        return this.id;
    }

    getBalance() {
        return this.balance;
    }

    deposit(amount) {
        this.balance += amount;
    }

    withdraw(amount) {
        if (this.balance >= amount) {
            this.balance -= amount;
            return true;
        }
        return false;
    }
}

class Bank {
    constructor() {
        this.accounts = new Map();
    }

    createAccount(id, balance = 0) {
        const account = new Account(id, balance);
        this.accounts.set(id, account);
        return account;
    }

    getAccount(id) {
        return this.accounts.get(id) || null;
    }

    getBalance(id) {
        const account = this.getAccount(id);
        return account ? account.getBalance() : 0;
    }

    deposit(id, amount) {
        const account = this.getAccount(id);
        if (account) {
            account.deposit(amount);
        }
    }

    withdraw(id, amount) {
        const account = this.getAccount(id);
        return account ? account.withdraw(amount) : false;
    }
}

/**
 * 71. Separating the domain from the representation (Separate Domain from Presentation)
 *
 * BEFORE: Domain logic mixed with presentation
 */
class OrderPresentationBefore {
    constructor() {
        this.items = [];
        this.total = 0;
    }

    addItem(name, price, quantity) {
        this.items.push({ name, price, quantity });
        this.total += price * quantity;

        // Presentation logic mixed in
        console.log(`Added ${quantity} x ${name} to order`);
        console.log(`Current total: $${this.total.toFixed(2)}`);
    }

    getTotal() {
        return this.total;
    }

    displayOrder() {
        console.log('Order Summary:');
        this.items.forEach(item => {
            console.log(`- ${item.quantity} x ${item.name} @ $${item.price.toFixed(2)}`);
        });
        console.log(`Total: $${this.total.toFixed(2)}`);
    }
}

/**
 * AFTER: Separate domain from presentation
 */
class OrderItem {
    constructor(name, price, quantity) {
        this.name = name;
        this.price = price;
        this.quantity = quantity;
    }

    getName() {
        return this.name;
    }

    getPrice() {
        return this.price;
    }

    getQuantity() {
        return this.quantity;
    }

    getTotal() {
        return this.price * this.quantity;
    }
}

class OrderDomainAfter {
    constructor() {
        this.items = [];
    }

    addItem(name, price, quantity) {
        const item = new OrderItem(name, price, quantity);
        this.items.push(item);
    }

    getItems() {
        return this.items;
    }

    getTotal() {
        return this.items.reduce((total, item) => total + item.getTotal(), 0);
    }
}

class OrderPresenter {
    displayItemAdded(item) {
        console.log(`Added ${item.getQuantity()} x ${item.getName()} to order`);
    }

    displayOrderSummary(order) {
        console.log(`Current total: $${order.getTotal().toFixed(2)}`);
        console.log('Order Summary:');
        order.getItems().forEach(item => {
            console.log(`- ${item.getQuantity()} x ${item.getName()} @ $${item.getPrice().toFixed(2)}`);
        });
        console.log(`Total: $${order.getTotal().toFixed(2)}`);
    }
}

class OrderService {
    constructor() {
        this.order = new OrderDomainAfter();
        this.presenter = new OrderPresenter();
    }

    addItem(name, price, quantity) {
        const item = new OrderItem(name, price, quantity);
        this.order.addItem(name, price, quantity);
        this.presenter.displayItemAdded(item);
        this.presenter.displayOrderSummary(this.order);
    }

    getOrder() {
        return this.order;
    }
}

/**
 * 72. Hierarchy Extraction (Extract Hierarchy)
 *
 * BEFORE: Single class handling multiple responsibilities
 */
class ComputerExtractBefore {
    constructor(type, cpu, ram, storage) {
        this.type = type;
        this.cpu = cpu;
        this.ram = ram;
        this.storage = storage;
    }

    getSpecs() {
        let specs = `CPU: ${this.cpu}\n`;
        specs += `RAM: ${this.ram}GB\n`;
        specs += `Storage: ${this.storage}GB\n`;

        if (this.type === 'desktop') {
            specs += 'Form Factor: Desktop\n';
        } else if (this.type === 'laptop') {
            specs += 'Form Factor: Laptop\n';
            specs += 'Battery Life: 8 hours\n';
        } else if (this.type === 'server') {
            specs += 'Form Factor: Server Rack\n';
            specs += 'Redundancy: RAID 10\n';
        }

        return specs;
    }
}

/**
 * AFTER: Extract hierarchy
 */
class ComputerExtractAfter {
    constructor(cpu, ram, storage) {
        this.cpu = cpu;
        this.ram = ram;
        this.storage = storage;
    }

    getFormFactor() {
        throw new Error('Abstract method must be implemented');
    }

    getSpecialFeatures() {
        throw new Error('Abstract method must be implemented');
    }

    getBasicSpecs() {
        return `CPU: ${this.cpu}\n` +
               `RAM: ${this.ram}GB\n` +
               `Storage: ${this.storage}GB\n`;
    }

    getSpecs() {
        return this.getBasicSpecs() +
               `Form Factor: ${this.getFormFactor()}\n` +
               this.getSpecialFeatures();
    }
}

class DesktopComputer extends ComputerExtractAfter {
    getFormFactor() {
        return 'Desktop';
    }

    getSpecialFeatures() {
        return 'Expansion Slots: Multiple PCI\n';
    }
}

class LaptopComputer extends ComputerExtractAfter {
    getFormFactor() {
        return 'Laptop';
    }

    getSpecialFeatures() {
        return 'Battery Life: 8 hours\nWeight: 2.5 lbs\n';
    }
}

class ServerComputer extends ComputerExtractAfter {
    getFormFactor() {
        return 'Server Rack';
    }

    getSpecialFeatures() {
        return 'Redundancy: RAID 10\nHot Swap Drives: Yes\n';
    }
}

module.exports = {
    EmployeeTeaseBefore,
    SalariedEmployeeTeaseBefore,
    CommissionedEmployeeTeaseBefore,
    EmployeeTeaseAfter,
    SalariedEmployeeTeaseAfter,
    CommissionedEmployeeTeaseAfter,
    ProceduralDesignBefore,
    Account,
    Bank,
    OrderPresentationBefore,
    OrderItem,
    OrderDomainAfter,
    OrderPresenter,
    OrderService,
    ComputerExtractBefore,
    ComputerExtractAfter,
    DesktopComputer,
    LaptopComputer,
    ServerComputer
};
