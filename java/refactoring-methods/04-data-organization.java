import java.util.*;

/**
 * 18. Self-Encapsulate Field
 *
 * BEFORE: Direct field access
 */
class PersonBefore {
    public String name; // Direct access

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}

/**
 * AFTER: Self-encapsulate field
 */
class PersonAfter {
    private String name;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}

/**
 * 19. Replacing the data value with an object (Replace Data Value with Object)
 *
 * BEFORE: Primitive data type that should be an object
 */
class OrderBefore {
    private String customer; // Just a string

    public String getCustomerName() {
        return customer;
    }

    public void setCustomer(String customer) {
        this.customer = customer;
    }
}

/**
 * AFTER: Replace with object
 */
class Customer {
    private String name;

    public Customer(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }
}

class OrderAfter {
    private Customer customer;

    public Customer getCustomer() {
        return customer;
    }

    public void setCustomer(Customer customer) {
        this.customer = customer;
    }

    public String getCustomerName() {
        return customer.getName();
    }
}

/**
 * 20. Replacing the value with a reference (Change Value to Reference)
 *
 * BEFORE: Multiple instances of same object
 */
class CustomerValue {
    private String name;

    public CustomerValue(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }
}

class OrderValue {
    private CustomerValue customer; // New instance for each order

    public OrderValue(String customerName) {
        this.customer = new CustomerValue(customerName);
    }
}

/**
 * AFTER: Use reference to single instance
 */
class CustomerReference {
    private String name;
    private static Map<String, CustomerReference> instances = new HashMap<>();

    private CustomerReference(String name) {
        this.name = name;
    }

    public static CustomerReference create(String name) {
        if (!instances.containsKey(name)) {
            instances.put(name, new CustomerReference(name));
        }
        return instances.get(name);
    }

    public String getName() {
        return name;
    }
}

class OrderReference {
    private CustomerReference customer;

    public OrderReference(String customerName) {
        this.customer = CustomerReference.create(customerName);
    }
}

/**
 * 21. Replacing a reference with a value (Change Reference to Value)
 *
 * BEFORE: Unnecessary reference when value would suffice
 */
class CurrencyReference {
    private String code;

    public CurrencyReference(String code) {
        this.code = code;
    }

    public String getCode() {
        return code;
    }
}

class ProductReference {
    private double price;
    private CurrencyReference currency; // Reference object

    public ProductReference(double price, CurrencyReference currency) {
        this.price = price;
        this.currency = currency;
    }
}

/**
 * AFTER: Use value object instead
 */
class CurrencyValue {
    private String code;

    public CurrencyValue(String code) {
        this.code = code;
    }

    public String getCode() {
        return code;
    }
}

class ProductValue {
    private double price;
    private String currencyCode; // Just the value

    public ProductValue(double price, String currencyCode) {
        this.price = price;
        this.currencyCode = currencyCode;
    }

    public String getCurrencyCode() {
        return currencyCode;
    }
}

/**
 * 22. Replacing an array with an object (Replace Array with Object)
 *
 * BEFORE: Using array for structured data
 */
class PerformanceArray {
    public Map<String, Object> getPerformanceData() {
        Map<String, Object> data = new HashMap<>();
        data.put("goals", 10);
        data.put("assists", 5);
        data.put("minutes", 120);
        return data;
    }

    public double calculateScore(Map<String, Object> data) {
        int goals = (Integer) data.get("goals");
        int assists = (Integer) data.get("assists");
        int minutes = (Integer) data.get("minutes");
        return (goals * 2) + (assists * 1.5) + (minutes / 60.0);
    }
}

/**
 * AFTER: Replace array with object
 */
class PerformanceData {
    private int goals;
    private int assists;
    private int minutes;

    public PerformanceData(int goals, int assists, int minutes) {
        this.goals = goals;
        this.assists = assists;
        this.minutes = minutes;
    }

    public int getGoals() {
        return goals;
    }

    public int getAssists() {
        return assists;
    }

    public int getMinutes() {
        return minutes;
    }

    public double calculateScore() {
        return (goals * 2) + (assists * 1.5) + (minutes / 60.0);
    }
}

class PerformanceObject {
    public PerformanceData getPerformanceData() {
        return new PerformanceData(10, 5, 120);
    }

    public double calculateScore(PerformanceData data) {
        return data.calculateScore();
    }
}

/**
 * 23. Duplication of visible data (Duplicate Observed Data)
 *
 * BEFORE: Domain data mixed with presentation
 */
class OrderDomain {
    private double total = 0;

    public void addItem(double price) {
        total += price;
        // Have to update UI here too
        updateDisplay();
    }

    private void updateDisplay() {
        // Update UI elements
    }
}

/**
 * AFTER: Separate domain and presentation data
 */
interface OrderObserver {
    void update(double total);
}

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
            observer.update(total);
        }
    }
}

class OrderDisplay implements OrderObserver {
    private OrderDomainSeparated order;

    public OrderDisplay(OrderDomainSeparated order) {
        this.order = order;
        order.addObserver(this);
    }

    public void update(double total) {
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
    private List<OrderUni> orders = new ArrayList<>();

    public void addOrder(OrderUni order) {
        orders.add(order);
        // Order doesn't know about customer
    }
}

class OrderUni {
    private List<String> items = new ArrayList<>();
}

/**
 * AFTER: Bidirectional association
 */
class CustomerBi {
    private List<OrderBi> orders = new ArrayList<>();

    public void addOrder(OrderBi order) {
        orders.add(order);
        order.setCustomer(this);
    }
}

class OrderBi {
    private CustomerBi customer;
    private List<String> items = new ArrayList<>();

    public void setCustomer(CustomerBi customer) {
        this.customer = customer;
    }

    public CustomerBi getCustomer() {
        return customer;
    }
}

/**
 * 25. Replacing Bidirectional communication with Unidirectional
 * communication (Change Bidirectional Association to Unidirectional)
 *
 * BEFORE: Unnecessary bidirectional association
 */
class CustomerBidirectional {
    private List<OrderBidirectional> orders = new ArrayList<>();

    public void addOrder(OrderBidirectional order) {
        orders.add(order);
        order.setCustomer(this);
    }
}

class OrderBidirectional {
    private CustomerBidirectional customer;

    public void setCustomer(CustomerBidirectional customer) {
        this.customer = customer;
    }

    public CustomerBidirectional getCustomer() {
        return customer;
    }
}

/**
 * AFTER: Remove bidirectional link
 */
class CustomerUnidirectional {
    private List<OrderUnidirectional> orders = new ArrayList<>();

    public void addOrder(OrderUnidirectional order) {
        orders.add(order);
    }
}

class OrderUnidirectional {
    private String customerId;

    public OrderUnidirectional(String customerId) {
        this.customerId = customerId;
    }

    public String getCustomerId() {
        return customerId;
    }
}

/**
 * 26. Replacing the magic number with a symbolic constant
 * (Replace Magic Number with Symbolic Constant)
 *
 * BEFORE: Magic numbers
 */
class GeometryBefore {
    public double calculateCircleArea(double radius) {
        return 3.14159 * radius * radius; // Magic number
    }

    public double calculateCircleCircumference(double radius) {
        return 2 * 3.14159 * radius; // Same magic number
    }
}

/**
 * AFTER: Use symbolic constant
 */
class GeometryAfter {
    public static final double PI = 3.14159;

    public double calculateCircleArea(double radius) {
        return PI * radius * radius;
    }

    public double calculateCircleCircumference(double radius) {
        return 2 * PI * radius;
    }
}

/**
 * 27. Encapsulate Field
 *
 * BEFORE: Public field
 */
class PersonPublic {
    public String name;
}

/**
 * AFTER: Encapsulated field
 */
class PersonEncapsulated {
    private String name;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}

/**
 * 28. Encapsulate Collection
 *
 * BEFORE: Direct access to collection
 */
class TeamBefore {
    public List<String> players = new ArrayList<>(); // Direct access

    public void addPlayer(String player) {
        players.add(player);
    }
}

/**
 * AFTER: Encapsulated collection
 */
class TeamAfter {
    private List<String> players = new ArrayList<>();

    public void addPlayer(String player) {
        players.add(player);
    }

    public void removePlayer(String player) {
        players.remove(player);
    }

    public List<String> getPlayers() {
        return new ArrayList<>(players); // Return copy
    }

    public int getPlayerCount() {
        return players.size();
    }
}

/**
 * 29. Replacing a record with a Data Class
 *
 * BEFORE: Using array as data structure
 */
class EmployeeArray {
    public Map<String, Object> createEmployee(Map<String, Object> data) {
        Map<String, Object> employee = new HashMap<>();
        employee.put("name", data.get("name"));
        employee.put("salary", data.get("salary"));
        employee.put("department", data.get("department"));
        return employee;
    }

    public double getSalary(Map<String, Object> employee) {
        return (Double) employee.get("salary");
    }
}

/**
 * AFTER: Use data class
 */
class Employee {
    private String name;
    private double salary;
    private String department;

    public Employee(String name, double salary, String department) {
        this.name = name;
        this.salary = salary;
        this.department = department;
    }

    public String getName() {
        return name;
    }

    public double getSalary() {
        return salary;
    }

    public String getDepartment() {
        return department;
    }
}

class EmployeeDataClass {
    public Employee createEmployee(String name, double salary, String department) {
        return new Employee(name, salary, department);
    }

    public double getSalary(Employee employee) {
        return employee.getSalary();
    }
}

/**
 * 30. Replacing Type Code with Class
 *
 * BEFORE: Type code as constants
 */
class EmployeeTypeCode {
    public static final int ENGINEER = 0;
    public static final int SALESMAN = 1;
    public static final int MANAGER = 2;

    private int type;

    public EmployeeTypeCode(int type) {
        this.type = type;
    }

    public int getTypeCode() {
        return type;
    }

    public double getMonthlySalary() {
        switch (type) {
            case ENGINEER:
                return 5000;
            case SALESMAN:
                return 4000;
            case MANAGER:
                return 6000;
            default:
                return 0;
        }
    }
}

/**
 * AFTER: Replace type code with class
 */
abstract class EmployeeType {
    public abstract double getMonthlySalary();

    public static EngineerType createEngineer() {
        return new EngineerType();
    }

    public static SalesmanType createSalesman() {
        return new SalesmanType();
    }

    public static ManagerType createManager() {
        return new ManagerType();
    }
}

class EngineerType extends EmployeeType {
    public double getMonthlySalary() {
        return 5000;
    }
}

class SalesmanType extends EmployeeType {
    public double getMonthlySalary() {
        return 4000;
    }
}

class ManagerType extends EmployeeType {
    public double getMonthlySalary() {
        return 6000;
    }
}

class EmployeeTypeClass {
    private EmployeeType type;

    public EmployeeTypeClass(EmployeeType type) {
        this.type = type;
    }

    public double getMonthlySalary() {
        return type.getMonthlySalary();
    }
}

/**
 * 31. Replacing Type Code with Subclasses
 *
 * BEFORE: Type code in base class
 */
class EmployeeSubBefore {
    public static final int ENGINEER = 0;
    public static final int SALESMAN = 1;
    public static final int MANAGER = 2;

    private int type;
    private double salary;

    public EmployeeSubBefore(int type, double salary) {
        this.type = type;
        this.salary = salary;
    }

    public double getSalary() {
        return salary;
    }

    public int getType() {
        return type;
    }
}

/**
 * AFTER: Replace type code with subclasses
 */
abstract class EmployeeSubAfter {
    protected double salary;

    public EmployeeSubAfter(double salary) {
        this.salary = salary;
    }

    public double getSalary() {
        return salary;
    }

    public abstract String getType();
}

class Engineer extends EmployeeSubAfter {
    public Engineer(double salary) {
        super(salary);
    }

    public String getType() {
        return "engineer";
    }
}

class Salesman extends EmployeeSubAfter {
    public Salesman(double salary) {
        super(salary);
    }

    public String getType() {
        return "salesman";
    }
}

class Manager extends EmployeeSubAfter {
    public Manager(double salary) {
        super(salary);
    }

    public String getType() {
        return "manager";
    }
}

/**
 * 32. Replacing Type Code with State/Strategy
 *
 * BEFORE: Type code with behavior
 */
class EmployeeStateBefore {
    public static final int JUNIOR = 0;
    public static final int SENIOR = 1;
    public static final int LEAD = 2;

    private int level;

    public EmployeeStateBefore(int level) {
        this.level = level;
    }

    public double getSalaryMultiplier() {
        switch (level) {
            case JUNIOR:
                return 1.0;
            case SENIOR:
                return 1.5;
            case LEAD:
                return 2.0;
            default:
                return 1.0;
        }
    }
}

/**
 * AFTER: Use state/strategy pattern
 */
interface EmployeeLevel {
    double getSalaryMultiplier();
}

class JuniorLevel implements EmployeeLevel {
    public double getSalaryMultiplier() {
        return 1.0;
    }
}

class SeniorLevel implements EmployeeLevel {
    public double getSalaryMultiplier() {
        return 1.5;
    }
}

class LeadLevel implements EmployeeLevel {
    public double getSalaryMultiplier() {
        return 2.0;
    }
}

class EmployeeStateAfter {
    private EmployeeLevel level;

    public EmployeeStateAfter(EmployeeLevel level) {
        this.level = level;
    }

    public double getSalaryMultiplier() {
        return level.getSalaryMultiplier();
    }
}

/**
 * 33. Replacing Subclass with Fields
 *
 * BEFORE: Unnecessary subclasses
 */
abstract class PersonSub {
    protected String name;
    protected String gender;

    public PersonSub(String name, String gender) {
        this.name = name;
        this.gender = gender;
    }

    public String getName() {
        return name;
    }

    public abstract boolean isMale();
}

class Male extends PersonSub {
    public Male(String name) {
        super(name, "male");
    }

    public boolean isMale() {
        return true;
    }
}

class Female extends PersonSub {
    public Female(String name) {
        super(name, "female");
    }

    public boolean isMale() {
        return false;
    }
}

/**
 * AFTER: Replace subclass with field
 */
class PersonField {
    private String name;
    private String gender; // 'male' or 'female'

    public PersonField(String name, String gender) {
        this.name = name;
        this.gender = gender;
    }

    public String getName() {
        return name;
    }

    public boolean isMale() {
        return "male".equals(gender);
    }

    public String getGender() {
        return gender;
    }
}
