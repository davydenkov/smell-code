import java.time.LocalDate;
import java.util.List;
import java.util.Map;
import java.util.HashMap;

/**
 * 9. Substitution Algorithm
 *
 * BEFORE: Complex algorithm that can be simplified
 */
class PricingServiceBefore {
    public double calculatePrice(List<Map<String, Object>> items) {
        double total = 0;
        for (Map<String, Object> item : items) {
            if ("book".equals(item.get("type"))) {
                total += (Double) item.get("price") * 0.9; // 10% discount for books
            } else if ("electronics".equals(item.get("type"))) {
                total += (Double) item.get("price") * 1.1; // 10% markup for electronics
            } else {
                total += (Double) item.get("price");
            }
        }
        return total;
    }
}

/**
 * AFTER: Substitute with a simpler algorithm
 */
class PricingServiceAfter {
    private Map<String, Double> discounts = Map.of(
        "book", 0.9,
        "electronics", 1.1,
        "default", 1.0
    );

    public double calculatePrice(List<Map<String, Object>> items) {
        double total = 0;
        for (Map<String, Object> item : items) {
            String type = (String) item.get("type");
            double multiplier = discounts.getOrDefault(type, discounts.get("default"));
            total += (Double) item.get("price") * multiplier;
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
    private double balance;

    public AccountBefore(double balance) {
        this.balance = balance;
    }

    public double getBalance() {
        return balance;
    }

    // This method belongs in Bank class, not Account
    public boolean transferTo(AccountBefore target, double amount) {
        if (balance >= amount) {
            balance -= amount;
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
    private double balance;

    public AccountAfter(double balance) {
        this.balance = balance;
    }

    public double getBalance() {
        return balance;
    }

    public void decreaseBalance(double amount) {
        balance -= amount;
    }

    public void increaseBalance(double amount) {
        balance += amount;
    }
}

class Bank {
    public boolean transfer(AccountAfter from, AccountAfter to, double amount) {
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
    private String name;
    private Map<String, String> address; // This should be in Address class

    public CustomerBefore(String name, String street, String city, String zipCode) {
        this.name = name;
        this.address = new HashMap<>();
        this.address.put("street", street);
        this.address.put("city", city);
        this.address.put("zipCode", zipCode);
    }

    public String getAddress() {
        return address.get("street") + ", " + address.get("city") + " " + address.get("zipCode");
    }
}

/**
 * AFTER: Move field to dedicated class
 */
class Address {
    private String street;
    private String city;
    private String zipCode;

    public Address(String street, String city, String zipCode) {
        this.street = street;
        this.city = city;
        this.zipCode = zipCode;
    }

    public String getFullAddress() {
        return street + ", " + city + " " + zipCode;
    }
}

class CustomerAfter {
    private String name;
    private Address address;

    public CustomerAfter(String name, Address address) {
        this.name = name;
        this.address = address;
    }

    public String getAddress() {
        return address.getFullAddress();
    }
}

/**
 * 12. Class Allocation (Extract Class)
 *
 * BEFORE: Class has too many responsibilities
 */
class PersonBefore {
    private String name;
    private String phoneNumber;
    private String officeAreaCode;
    private String officeNumber;

    public String getTelephoneNumber() {
        return "(" + officeAreaCode + ") " + officeNumber;
    }
}

/**
 * AFTER: Extract telephone number to separate class
 */
class TelephoneNumber {
    private String areaCode;
    private String number;

    public TelephoneNumber(String areaCode, String number) {
        this.areaCode = areaCode;
        this.number = number;
    }

    public String getTelephoneNumber() {
        return "(" + areaCode + ") " + number;
    }
}

class PersonAfter {
    private String name;
    private String phoneNumber;
    private TelephoneNumber officeTelephone;

    public PersonAfter(String name) {
        this.name = name;
    }

    public String getOfficeTelephone() {
        return officeTelephone.getTelephoneNumber();
    }

    public void setOfficeTelephone(TelephoneNumber telephone) {
        this.officeTelephone = telephone;
    }
}

/**
 * 13. Embedding a class (Inline Class)
 *
 * BEFORE: Unnecessary class with single responsibility
 */
class OrderProcessorBefore {
    private OrderValidator validator;

    public OrderProcessorBefore() {
        this.validator = new OrderValidator();
    }

    public void process(Map<String, Object> order) {
        if (validator.isValid(order)) {
            // Process order
        }
    }
}

class OrderValidator {
    public boolean isValid(Map<String, Object> order) {
        return (Double) order.get("total") > 0;
    }
}

/**
 * AFTER: Inline the class
 */
class OrderProcessorAfter {
    public void process(Map<String, Object> order) {
        if (isValidOrder(order)) {
            // Process order
        }
    }

    private boolean isValidOrder(Map<String, Object> order) {
        return (Double) order.get("total") > 0;
    }
}

/**
 * 14. Hiding delegation (Hide Delegate)
 *
 * BEFORE: Client has to know about delegation
 */
class DepartmentBefore {
    private Person personManager;

    public DepartmentBefore(Person personManager) {
        this.personManager = personManager;
    }

    public Person getManager() {
        return personManager;
    }
}

class Person {
    private DepartmentBefore department;

    public DepartmentBefore getDepartment() {
        return department;
    }
}

// Client code would be: Person manager = person.getDepartment().getManager();

/**
 * AFTER: Hide the delegation
 */
class DepartmentAfter {
    private PersonAfter personManager;

    public DepartmentAfter(PersonAfter personManager) {
        this.personManager = personManager;
    }

    public PersonAfter getManager() {
        return personManager;
    }
}

class PersonAfter {
    private DepartmentAfter department;

    public DepartmentAfter getDepartment() {
        return department;
    }

    public PersonAfter getManager() {
        return department.getManager();
    }
}

// Client code - much cleaner: PersonAfter manager = person.getManager();

/**
 * 15. Removing the intermediary (Remove Middle Man)
 *
 * BEFORE: Too much delegation
 */
class PersonWithMiddleMan {
    private Department department;

    public Department getDepartment() {
        return department;
    }

    public Person getManager() {
        return department.getManager();
    }

    public String getDepartmentName() {
        return department.getName();
    }
}

/**
 * AFTER: Remove middle man if delegation is too heavy
 */
class PersonDirect {
    private Department department;
    private Person manager; // Direct reference

    public Person getManager() {
        return manager;
    }

    public Department getDepartment() {
        return department;
    }
}

/**
 * 16. Introduction of an external method (Introduce Foreign Method)
 *
 * BEFORE: Using external class method in wrong place
 */
class ReportGeneratorBefore {
    public void generateReport() {
        LocalDate date = LocalDate.now();
        LocalDate nextMonth = date.plusMonths(1); // Foreign method usage

        // Generate report for next month
    }
}

/**
 * AFTER: Introduce foreign method
 */
class ReportGeneratorAfter {
    public void generateReport() {
        LocalDate date = LocalDate.now();
        LocalDate nextMonth = nextMonth(date);

        // Generate report for next month
    }

    private LocalDate nextMonth(LocalDate date) {
        return date.plusMonths(1);
    }
}

/**
 * 17. The introduction of local extension (Introduce Local Extension)
 *
 * BEFORE: Adding methods to external class (not possible)
 */
class DateUtil {
    public static LocalDate nextMonth(LocalDate date) {
        return date.plusMonths(1);
    }

    public static LocalDate previousMonth(LocalDate date) {
        return date.minusMonths(1);
    }
}

/**
 * AFTER: Create local extension class
 */
class LocalDateExtension {
    private LocalDate date;

    public LocalDateExtension(LocalDate date) {
        this.date = date;
    }

    public LocalDate nextMonth() {
        return date.plusMonths(1);
    }

    public LocalDate previousMonth() {
        return date.minusMonths(1);
    }

    public LocalDate getDate() {
        return date;
    }
}
