using System;
using System.Collections.Generic;

/// <summary>
/// 1. Method Extraction (Extract Method)
///
/// BEFORE: A method contains too much logic, making it hard to understand
/// </summary>
public class OrderProcessorBefore
{
    public double ProcessOrder(Dictionary<string, object> order)
    {
        // Validate order
        if ((double)order["total"] <= 0)
        {
            throw new ArgumentException("Invalid order total");
        }

        // Calculate tax
        double tax = (double)order["subtotal"] * 0.08;

        // Calculate shipping
        double shipping = (double)order["weight"] > 10 ? 15.00 : 5.00;

        // Calculate total
        double total = (double)order["subtotal"] + tax + shipping;

        // Save to database
        SaveOrder(order, total);

        return total;
    }

    private void SaveOrder(Dictionary<string, object> order, double total)
    {
        // Database save logic
    }
}

/// <summary>
/// AFTER: Extract methods to separate concerns
/// </summary>
public class OrderProcessorAfter
{
    public double ProcessOrder(Dictionary<string, object> order)
    {
        ValidateOrder(order);

        double tax = CalculateTax(order);
        double shipping = CalculateShipping(order);
        double total = CalculateTotal(order, tax, shipping);

        SaveOrder(order, total);

        return total;
    }

    private void ValidateOrder(Dictionary<string, object> order)
    {
        if ((double)order["total"] <= 0)
        {
            throw new ArgumentException("Invalid order total");
        }
    }

    private double CalculateTax(Dictionary<string, object> order)
    {
        return (double)order["subtotal"] * 0.08;
    }

    private double CalculateShipping(Dictionary<string, object> order)
    {
        return (double)order["weight"] > 10 ? 15.00 : 5.00;
    }

    private double CalculateTotal(Dictionary<string, object> order, double tax, double shipping)
    {
        return (double)order["subtotal"] + tax + shipping;
    }

    private void SaveOrder(Dictionary<string, object> order, double total)
    {
        // Database save logic
    }
}

/// <summary>
/// 2. Embedding a method (Inline Method)
///
/// BEFORE: A method is too simple and adds no value
/// </summary>
public class UserBefore
{
    private string firstName;
    private string lastName;

    public string GetFullName()
    {
        return GetFirstName() + " " + GetLastName();
    }

    public string GetFirstName()
    {
        return firstName;
    }

    public string GetLastName()
    {
        return lastName;
    }
}

/// <summary>
/// AFTER: Inline the simple method
/// </summary>
public class UserAfter
{
    private string firstName;
    private string lastName;

    public string GetFullName()
    {
        return firstName + " " + lastName;
    }

    // GetFirstName() and GetLastName() methods removed
}
