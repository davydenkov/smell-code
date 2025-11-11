import java.util.Map;

/**
 * 1. Method Extraction (Extract Method)
 *
 * BEFORE: A method contains too much logic, making it hard to understand
 */
class OrderProcessorBefore {
    public double processOrder(Map<String, Object> order) throws Exception {
        // Validate order
        if ((Double) order.get("total") <= 0) {
            throw new Exception("Invalid order total");
        }

        // Calculate tax
        double tax = (Double) order.get("subtotal") * 0.08;

        // Calculate shipping
        double shipping = (Double) order.get("weight") > 10 ? 15.00 : 5.00;

        // Calculate total
        double total = (Double) order.get("subtotal") + tax + shipping;

        // Save to database
        saveOrder(order, total);

        return total;
    }

    private void saveOrder(Map<String, Object> order, double total) {
        // Database save logic
    }
}

/**
 * AFTER: Extract methods to separate concerns
 */
class OrderProcessorAfter {
    public double processOrder(Map<String, Object> order) throws Exception {
        validateOrder(order);

        double tax = calculateTax(order);
        double shipping = calculateShipping(order);
        double total = calculateTotal(order, tax, shipping);

        saveOrder(order, total);

        return total;
    }

    private void validateOrder(Map<String, Object> order) throws Exception {
        if ((Double) order.get("total") <= 0) {
            throw new Exception("Invalid order total");
        }
    }

    private double calculateTax(Map<String, Object> order) {
        return (Double) order.get("subtotal") * 0.08;
    }

    private double calculateShipping(Map<String, Object> order) {
        return (Double) order.get("weight") > 10 ? 15.00 : 5.00;
    }

    private double calculateTotal(Map<String, Object> order, double tax, double shipping) {
        return (Double) order.get("subtotal") + tax + shipping;
    }

    private void saveOrder(Map<String, Object> order, double total) {
        // Database save logic
    }
}

/**
 * 2. Embedding a method (Inline Method)
 *
 * BEFORE: A method is too simple and adds no value
 */
class UserBefore {
    private String firstName;
    private String lastName;

    public String getFullName() {
        return getFirstName() + " " + getLastName();
    }

    public String getFirstName() {
        return firstName;
    }

    public String getLastName() {
        return lastName;
    }
}

/**
 * AFTER: Inline the simple method
 */
class UserAfter {
    private String firstName;
    private String lastName;

    public String getFullName() {
        return firstName + " " + lastName;
    }

    // getFirstName() and getLastName() methods removed
}
