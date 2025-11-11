/**
 * 42. Renaming a method (Rename Method)
 *
 * BEFORE: Poorly named method
 */
class CalculatorBefore {
    public double calc(double a, double b) { // Unclear name
        return a + b;
    }
}

/**
 * AFTER: Rename method to be more descriptive
 */
class CalculatorAfter {
    public double add(double a, double b) {
        return a + b;
    }
}

/**
 * 43. Adding a parameter (Add Parameter)
 *
 * BEFORE: Method missing required parameter
 */
class EmailSenderBefore {
    public void sendEmail(String to, String subject, String body) {
        // Send email with default priority
        String priority = "normal";
        // Send logic
    }
}

/**
 * AFTER: Add parameter
 */
class EmailSenderAfter {
    public void sendEmail(String to, String subject, String body, String priority) {
        // Send logic with priority
    }

    // Overloaded method for backward compatibility
    public void sendEmail(String to, String subject, String body) {
        sendEmail(to, subject, body, "normal");
    }
}

/**
 * 44. Deleting a parameter (Remove Parameter)
 *
 * BEFORE: Unnecessary parameter
 */
class ReportGeneratorBefore {
    public void generateReport(Object data, String format, boolean includeHeader) {
        if ("html".equals(format)) {
            // Always include header for HTML
            includeHeader = true;
        }
        // Generate report
    }
}

/**
 * AFTER: Remove unnecessary parameter
 */
class ReportGeneratorAfter {
    public void generateReport(Object data, String format) {
        boolean includeHeader = "html".equals(format);
        // Generate report
    }
}

/**
 * 45. Separation of Query and Modifier (Separate Query from Modifier)
 *
 * BEFORE: Method that both queries and modifies
 */
class BankAccountBefore {
    private double balance = 0;

    public boolean withdraw(double amount) {
        if (balance >= amount) {
            balance -= amount;
            return true;
        }
        return false;
    }
}

/**
 * AFTER: Separate query from modifier
 */
class BankAccountAfter {
    private double balance = 0;

    public boolean canWithdraw(double amount) {
        return balance >= amount;
    }

    public boolean withdraw(double amount) {
        if (canWithdraw(amount)) {
            balance -= amount;
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
    public Object generateWeeklyReport() {
        return generateReport(7);
    }

    public Object generateMonthlyReport() {
        return generateReport(30);
    }

    public Object generateQuarterlyReport() {
        return generateReport(90);
    }

    private Object generateReport(int days) {
        // Generate report for specified days
        return null;
    }
}

/**
 * AFTER: Parameterize method
 */
class ReportGeneratorParamAfter {
    public Object generateReport(int days) {
        // Generate report for specified days
        return null;
    }

    public Object generateWeeklyReport() {
        return generateReport(7);
    }

    public Object generateMonthlyReport() {
        return generateReport(30);
    }

    public Object generateQuarterlyReport() {
        return generateReport(90);
    }
}

/**
 * 47. Replacing a parameter with explicit methods (Replace Parameter with Explicit Methods)
 *
 * BEFORE: Parameter determines behavior
 */
class EmployeeExplicitBefore {
    public static final int ENGINEER = 0;
    public static final int SALESMAN = 1;
    public static final int MANAGER = 2;

    private int type;

    public EmployeeExplicitBefore(int type) {
        this.type = type;
    }

    public double getSalary(double baseSalary) {
        switch (type) {
            case ENGINEER:
                return baseSalary * 1.0;
            case SALESMAN:
                return baseSalary * 1.1;
            case MANAGER:
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
    public double getEngineerSalary(double baseSalary) {
        return baseSalary * 1.0;
    }

    public double getSalesmanSalary(double baseSalary) {
        return baseSalary * 1.1;
    }

    public double getManagerSalary(double baseSalary) {
        return baseSalary * 1.2;
    }
}

/**
 * 48. Save the Whole Object
 *
 * BEFORE: Passing individual fields
 */
class OrderWholeBefore {
    private Map<String, String> customer;

    public OrderWholeBefore(String customerName, String customerAddress) {
        this.customer = new java.util.HashMap<>();
        this.customer.put("name", customerName);
        this.customer.put("address", customerAddress);
    }

    public double calculateShipping() {
        return getShippingCost(customer.get("name"), customer.get("address"));
    }

    private double getShippingCost(String name, String address) {
        // Calculate based on name and address
        return 10.0;
    }
}

/**
 * AFTER: Pass whole object
 */
class CustomerWhole {
    private String name;
    private String address;

    public CustomerWhole(String name, String address) {
        this.name = name;
        this.address = address;
    }

    public String getName() {
        return name;
    }

    public String getAddress() {
        return address;
    }
}

class OrderWholeAfter {
    private CustomerWhole customer;

    public OrderWholeAfter(CustomerWhole customer) {
        this.customer = customer;
    }

    public double calculateShipping() {
        return getShippingCost(customer);
    }

    private double getShippingCost(CustomerWhole customer) {
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
    public double calculateDiscount(double price, String customerType) {
        // customerType passed in
        return price * getDiscountRate(customerType);
    }

    private double getDiscountRate(String customerType) {
        switch (customerType) {
            case "premium":
                return 0.1;
            case "regular":
                return 0.05;
            default:
                return 0.0;
        }
    }
}

class CustomerParam {
    private String type;

    public String getType() {
        return type;
    }
}

class OrderParamBefore {
    private CustomerParam customer;

    public double getDiscountedPrice(double price) {
        DiscountCalculatorParamBefore calculator = new DiscountCalculatorParamBefore();
        return calculator.calculateDiscount(price, customer.getType());
    }
}

/**
 * AFTER: Replace parameter with method call
 */
class DiscountCalculatorParamAfter {
    public double calculateDiscount(double price, CustomerParam customer) {
        return price * getDiscountRate(customer.getType());
    }

    private double getDiscountRate(String customerType) {
        switch (customerType) {
            case "premium":
                return 0.1;
            case "regular":
                return 0.05;
            default:
                return 0.0;
        }
    }
}

class OrderParamAfter {
    private CustomerParam customer;

    public double getDiscountedPrice(double price) {
        DiscountCalculatorParamAfter calculator = new DiscountCalculatorParamAfter();
        return calculator.calculateDiscount(price, customer);
    }
}
