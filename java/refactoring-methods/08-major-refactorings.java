import java.util.*;

/**
 * 69. Separation of inheritance (Tease Apart Inheritance)
 *
 * BEFORE: Class hierarchy mixing two different responsibilities
 */
class EmployeeTeaseBefore {
    protected String name;
    protected double rate;

    public EmployeeTeaseBefore(String name, double rate) {
        this.name = name;
        this.rate = rate;
    }

    public String getName() {
        return name;
    }
}

class SalariedEmployeeTeaseBefore extends EmployeeTeaseBefore {
    public SalariedEmployeeTeaseBefore(String name, double rate) {
        super(name, rate);
    }

    public double getPay() {
        return rate;
    }
}

class CommissionedEmployeeTeaseBefore extends EmployeeTeaseBefore {
    private double commission;

    public CommissionedEmployeeTeaseBefore(String name, double rate, double commission) {
        super(name, rate);
        this.commission = commission;
    }

    public double getPay() {
        return rate + commission;
    }
}

/**
 * AFTER: Tease apart inheritance into two separate hierarchies
 */
interface Payable {
    double getPay();
}

class EmployeeTeaseAfter {
    protected String name;

    public EmployeeTeaseAfter(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }
}

class SalariedEmployeeTeaseAfter extends EmployeeTeaseAfter implements Payable {
    private double salary;

    public SalariedEmployeeTeaseAfter(String name, double salary) {
        super(name);
        this.salary = salary;
    }

    public double getPay() {
        return salary;
    }
}

class CommissionedEmployeeTeaseAfter extends EmployeeTeaseAfter implements Payable {
    private double baseSalary;
    private double commission;

    public CommissionedEmployeeTeaseAfter(String name, double baseSalary, double commission) {
        super(name);
        this.baseSalary = baseSalary;
        this.commission = commission;
    }

    public double getPay() {
        return baseSalary + commission;
    }
}

/**
 * 70. Converting a procedural project into objects (Convert Procedural Design to Objects)
 *
 * BEFORE: Procedural code with global functions and data
 */
class ProceduralDesignBefore {
    private static Map<String, Double> accounts = new HashMap<>();

    public static void createAccount(String id, double balance) {
        accounts.put(id, balance);
    }

    public static double getBalance(String id) {
        return accounts.getOrDefault(id, 0.0);
    }

    public static void deposit(String id, double amount) {
        accounts.put(id, accounts.getOrDefault(id, 0.0) + amount);
    }
}

/**
 * AFTER: Convert to object-oriented design
 */
class Account {
    private String id;
    private double balance;

    public Account(String id, double balance) {
        this.id = id;
        this.balance = balance;
    }

    public String getId() {
        return id;
    }

    public double getBalance() {
        return balance;
    }

    public void deposit(double amount) {
        balance += amount;
    }

    public boolean withdraw(double amount) {
        if (balance >= amount) {
            balance -= amount;
            return true;
        }
        return false;
    }
}

class AccountService {
    private Map<String, Account> accounts = new HashMap<>();

    public void createAccount(String id, double balance) {
        accounts.put(id, new Account(id, balance));
    }

    public double getBalance(String id) {
        Account account = accounts.get(id);
        return account != null ? account.getBalance() : 0.0;
    }

    public void deposit(String id, double amount) {
        Account account = accounts.get(id);
        if (account != null) {
            account.deposit(amount);
        }
    }
}

/**
 * 71. Separation of domain from presentation (Separate Domain from Presentation)
 *
 * BEFORE: Domain logic mixed with presentation
 */
class OrderPresentationBefore {
    private double total = 0;

    public void addItem(double price) {
        total += price;
        updateUI();
    }

    public double getTotal() {
        return total;
    }

    private void updateUI() {
        // UI update logic mixed with domain logic
        System.out.println("Total: " + total);
    }
}

/**
 * AFTER: Separate domain from presentation
 */
class OrderDomainSeparated {
    private double total = 0;
    private List<OrderObserver> observers = new ArrayList<>();

    public void addItem(double price) {
        total += price;
        notifyObservers();
    }

    public double getTotal() {
        return total;
    }

    public void addObserver(OrderObserver observer) {
        observers.add(observer);
    }

    private void notifyObservers() {
        for (OrderObserver observer : observers) {
            observer.onTotalChanged(total);
        }
    }
}

interface OrderObserver {
    void onTotalChanged(double total);
}

class OrderPresentationAfter implements OrderObserver {
    private OrderDomainSeparated order;

    public OrderPresentationAfter(OrderDomainSeparated order) {
        this.order = order;
        order.addObserver(this);
    }

    public void onTotalChanged(double total) {
        // Pure presentation logic
        System.out.println("Total updated: " + total);
        updateDisplay(total);
    }

    private void updateDisplay(double total) {
        // UI update logic
    }
}

/**
 * 72. Extraction of hierarchy (Extract Hierarchy)
 *
 * BEFORE: Complex conditional logic
 */
class EmployeeHierarchyBefore {
    private String type;
    private double salary;
    private double commission;
    private double bonus;

    public double calculatePay() {
        double pay = salary;
        if ("salesman".equals(type)) {
            pay += commission;
        } else if ("manager".equals(type)) {
            pay += bonus;
        }
        return pay;
    }
}

/**
 * AFTER: Extract hierarchy
 */
abstract class EmployeeHierarchyAfter {
    protected double salary;

    public EmployeeHierarchyAfter(double salary) {
        this.salary = salary;
    }

    public abstract double calculatePay();
}

class SalesmanHierarchyAfter extends EmployeeHierarchyAfter {
    private double commission;

    public SalesmanHierarchyAfter(double salary, double commission) {
        super(salary);
        this.commission = commission;
    }

    public double calculatePay() {
        return salary + commission;
    }
}

class ManagerHierarchyAfter extends EmployeeHierarchyAfter {
    private double bonus;

    public ManagerHierarchyAfter(double salary, double bonus) {
        super(salary);
        this.bonus = bonus;
    }

    public double calculatePay() {
        return salary + bonus;
    }
}

class EngineerHierarchyAfter extends EmployeeHierarchyAfter {
    public EngineerHierarchyAfter(double salary) {
        super(salary);
    }

    public double calculatePay() {
        return salary;
    }
}
