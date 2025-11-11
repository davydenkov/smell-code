import java.util.HashMap;
import java.util.Map;

/**
 * 3. Embedding a temporary variable (Inline Temp)
 *
 * BEFORE: Unnecessary temporary variable
 */
class PriceCalculatorBefore {
    private int quantity;
    private double itemPrice;

    public double getPrice() {
        double basePrice = quantity * itemPrice;
        if (basePrice > 1000) {
            return basePrice * 0.95;
        } else {
            return basePrice * 0.98;
        }
    }
}

/**
 * AFTER: Inline the temporary variable
 */
class PriceCalculatorAfter {
    private int quantity;
    private double itemPrice;

    public double getPrice() {
        if (quantity * itemPrice > 1000) {
            return quantity * itemPrice * 0.95;
        } else {
            return quantity * itemPrice * 0.98;
        }
    }
}

/**
 * 4. Replacing a temporary variable with a method call (Replace Temp with Query)
 *
 * BEFORE: Temporary variable used multiple times
 */
class OrderBefore {
    private int quantity;
    private double itemPrice;

    public double getPrice() {
        double basePrice = quantity * itemPrice;
        return basePrice - getDiscount(basePrice);
    }

    private double getDiscount(double basePrice) {
        return Math.max(0, basePrice - 500) * 0.05;
    }
}

/**
 * AFTER: Replace temp with query
 */
class OrderAfter {
    private int quantity;
    private double itemPrice;

    public double getPrice() {
        return getBasePrice() - getDiscount();
    }

    private double getBasePrice() {
        return quantity * itemPrice;
    }

    private double getDiscount() {
        return Math.max(0, getBasePrice() - 500) * 0.05;
    }
}

/**
 * 5. Introduction of an explanatory variable (Introduce Explaining Variable)
 *
 * BEFORE: Complex expression hard to understand
 */
class PerformanceCalculatorBefore {
    private int goals;
    private int assists;
    private int minutesPlayed;

    public double getPerformance() {
        return (goals * 2) + (assists * 1.5) + (minutesPlayed / 60) * 0.1;
    }
}

/**
 * AFTER: Introduce explaining variables for clarity
 */
class PerformanceCalculatorAfter {
    private int goals;
    private int assists;
    private int minutesPlayed;

    public double getPerformance() {
        double goalPoints = goals * 2;
        double assistPoints = assists * 1.5;
        double playingTimeBonus = (minutesPlayed / 60.0) * 0.1;

        return goalPoints + assistPoints + playingTimeBonus;
    }
}

/**
 * 6. Splitting a Temporary Variable
 *
 * BEFORE: Same variable used for different purposes
 */
class TemperatureMonitorBefore {
    public Map<String, Double> getReading() {
        double temp = getCurrentTemperature();

        // First use: get initial reading
        double initialTemp = temp;

        // Later: temp is reused for different calculation
        temp = temp + getAdjustment();
        double adjustedTemp = temp;

        Map<String, Double> result = new HashMap<>();
        result.put("initial", initialTemp);
        result.put("adjusted", adjustedTemp);
        return result;
    }

    private double getCurrentTemperature() {
        return 25.0; // dummy implementation
    }

    private double getAdjustment() {
        return 2.0; // dummy implementation
    }
}

/**
 * AFTER: Split the temporary variable
 */
class TemperatureMonitorAfter {
    public Map<String, Double> getReading() {
        double temp = getCurrentTemperature();
        double initialTemp = temp;

        double adjustedTemp = temp + getAdjustment();

        Map<String, Double> result = new HashMap<>();
        result.put("initial", initialTemp);
        result.put("adjusted", adjustedTemp);
        return result;
    }

    private double getCurrentTemperature() {
        return 25.0; // dummy implementation
    }

    private double getAdjustment() {
        return 2.0; // dummy implementation
    }
}

/**
 * 7. Removing parameter Assignments (Remove Assignments to Parameters)
 *
 * BEFORE: Parameter is modified inside method
 */
class DiscountCalculatorBefore {
    public double applyDiscount(double price) {
        if (price > 100) {
            price = price * 0.9; // Modifying parameter
        }
        return price;
    }
}

/**
 * AFTER: Use a local variable instead
 */
class DiscountCalculatorAfter {
    public double applyDiscount(double price) {
        double result = price;
        if (price > 100) {
            result = price * 0.9;
        }
        return result;
    }
}

/**
 * 8. Replacing a method with a method Object (Replace Method with Method Object)
 *
 * BEFORE: Method with many parameters and local variables
 */
class AccountBefore {
    public double calculateInterest(double principal, double rate, int time, int compoundingFrequency) {
        double amount = principal * Math.pow(1 + (rate / compoundingFrequency), compoundingFrequency * time);
        double interest = amount - principal;
        return interest;
    }
}

/**
 * AFTER: Extract to a method object
 */
class InterestCalculation {
    private double principal;
    private double rate;
    private int time;
    private int compoundingFrequency;

    public InterestCalculation(double principal, double rate, int time, int compoundingFrequency) {
        this.principal = principal;
        this.rate = rate;
        this.time = time;
        this.compoundingFrequency = compoundingFrequency;
    }

    public double calculate() {
        double amount = principal * Math.pow(1 + (rate / compoundingFrequency),
                                           compoundingFrequency * time);
        return amount - principal;
    }
}

class AccountAfter {
    public double calculateInterest(double principal, double rate, int time, int compoundingFrequency) {
        InterestCalculation calculation = new InterestCalculation(principal, rate, time, compoundingFrequency);
        return calculation.calculate();
    }
}
