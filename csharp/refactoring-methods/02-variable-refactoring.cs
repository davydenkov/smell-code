using System;
using System.Collections.Generic;

/// <summary>
/// 3. Embedding a temporary variable (Inline Temp)
///
/// BEFORE: Unnecessary temporary variable
/// </summary>
public class PriceCalculatorBefore
{
    private int quantity;
    private double itemPrice;

    public double GetPrice()
    {
        double basePrice = quantity * itemPrice;
        if (basePrice > 1000)
        {
            return basePrice * 0.95;
        }
        else
        {
            return basePrice * 0.98;
        }
    }
}

/// <summary>
/// AFTER: Inline the temporary variable
/// </summary>
public class PriceCalculatorAfter
{
    private int quantity;
    private double itemPrice;

    public double GetPrice()
    {
        if (quantity * itemPrice > 1000)
        {
            return quantity * itemPrice * 0.95;
        }
        else
        {
            return quantity * itemPrice * 0.98;
        }
    }
}

/// <summary>
/// 4. Replacing a temporary variable with a method call (Replace Temp with Query)
///
/// BEFORE: Temporary variable used multiple times
/// </summary>
public class OrderBefore
{
    private int quantity;
    private double itemPrice;

    public double GetPrice()
    {
        double basePrice = quantity * itemPrice;
        return basePrice - GetDiscount(basePrice);
    }

    private double GetDiscount(double basePrice)
    {
        return Math.Max(0, basePrice - 500) * 0.05;
    }
}

/// <summary>
/// AFTER: Replace temp with query
/// </summary>
public class OrderAfter
{
    private int quantity;
    private double itemPrice;

    public double GetPrice()
    {
        return GetBasePrice() - GetDiscount();
    }

    private double GetBasePrice()
    {
        return quantity * itemPrice;
    }

    private double GetDiscount()
    {
        return Math.Max(0, GetBasePrice() - 500) * 0.05;
    }
}

/// <summary>
/// 5. Introduction of an explanatory variable (Introduce Explaining Variable)
///
/// BEFORE: Complex expression hard to understand
/// </summary>
public class PerformanceCalculatorBefore
{
    private int goals;
    private int assists;
    private int minutesPlayed;

    public double GetPerformance()
    {
        return (goals * 2) + (assists * 1.5) + (minutesPlayed / 60) * 0.1;
    }
}

/// <summary>
/// AFTER: Introduce explaining variables for clarity
/// </summary>
public class PerformanceCalculatorAfter
{
    private int goals;
    private int assists;
    private int minutesPlayed;

    public double GetPerformance()
    {
        double goalPoints = goals * 2;
        double assistPoints = assists * 1.5;
        double playingTimeBonus = (minutesPlayed / 60) * 0.1;

        return goalPoints + assistPoints + playingTimeBonus;
    }
}

/// <summary>
/// 6. Splitting a Temporary Variable
///
/// BEFORE: Same variable used for different purposes
/// </summary>
public class TemperatureMonitorBefore
{
    public Dictionary<string, double> GetReading()
    {
        double temp = GetCurrentTemperature();

        // First use: get initial reading
        double initialTemp = temp;

        // Later: temp is reused for different calculation
        temp = temp + GetAdjustment();
        double adjustedTemp = temp;

        return new Dictionary<string, double>
        {
            ["initial"] = initialTemp,
            ["adjusted"] = adjustedTemp
        };
    }

    private double GetCurrentTemperature() => 25.0;
    private double GetAdjustment() => 2.0;
}

/// <summary>
/// AFTER: Split the temporary variable
/// </summary>
public class TemperatureMonitorAfter
{
    public Dictionary<string, double> GetReading()
    {
        double temp = GetCurrentTemperature();
        double initialTemp = temp;

        double adjustedTemp = temp + GetAdjustment();

        return new Dictionary<string, double>
        {
            ["initial"] = initialTemp,
            ["adjusted"] = adjustedTemp
        };
    }

    private double GetCurrentTemperature() => 25.0;
    private double GetAdjustment() => 2.0;
}

/// <summary>
/// 7. Removing parameter Assignments (Remove Assignments to Parameters)
///
/// BEFORE: Parameter is modified inside method
/// </summary>
public class DiscountCalculatorBefore
{
    public double ApplyDiscount(double price)
    {
        if (price > 100)
        {
            price = price * 0.9; // Modifying parameter
        }
        return price;
    }
}

/// <summary>
/// AFTER: Use a local variable instead
/// </summary>
public class DiscountCalculatorAfter
{
    public double ApplyDiscount(double price)
    {
        double result = price;
        if (price > 100)
        {
            result = price * 0.9;
        }
        return result;
    }
}

/// <summary>
/// 8. Replacing a method with a method Object (Replace Method with Method Object)
///
/// BEFORE: Method with many parameters and local variables
/// </summary>
public class AccountBefore
{
    public double CalculateInterest(double principal, double rate, double time, int compoundingFrequency)
    {
        double amount = principal * Math.Pow(1 + (rate / compoundingFrequency), compoundingFrequency * time);
        double interest = amount - principal;
        return interest;
    }
}

/// <summary>
/// AFTER: Extract to a method object
/// </summary>
public class InterestCalculation
{
    private double principal;
    private double rate;
    private double time;
    private int compoundingFrequency;

    public InterestCalculation(double principal, double rate, double time, int compoundingFrequency)
    {
        this.principal = principal;
        this.rate = rate;
        this.time = time;
        this.compoundingFrequency = compoundingFrequency;
    }

    public double Calculate()
    {
        double amount = principal * Math.Pow(1 + (rate / compoundingFrequency),
                                           compoundingFrequency * time);
        return amount - principal;
    }
}

public class AccountAfter
{
    public double CalculateInterest(double principal, double rate, double time, int compoundingFrequency)
    {
        var calculation = new InterestCalculation(principal, rate, time, compoundingFrequency);
        return calculation.Calculate();
    }
}
