using System;
using System.Collections.Generic;

/// <summary>
/// 42. Renaming a method (Rename Method)
///
/// BEFORE: Poorly named method
/// </summary>
public class CalculatorBefore
{
    public double Calc(double a, double b) // Unclear name
    {
        return a + b;
    }
}

/// <summary>
/// AFTER: Rename method to be more descriptive
/// </summary>
public class CalculatorAfter
{
    public double Add(double a, double b)
    {
        return a + b;
    }
}

/// <summary>
/// 43. Adding a parameter (Add Parameter)
///
/// BEFORE: Method missing required parameter
/// </summary>
public class EmailSenderBefore
{
    public void SendEmail(string to, string subject, string body)
    {
        // Send email with default priority
        string priority = "normal";
        // Send logic
    }
}

/// <summary>
/// AFTER: Add parameter
/// </summary>
public class EmailSenderAfter
{
    public void SendEmail(string to, string subject, string body, string priority = "normal")
    {
        // Send logic with priority
    }
}

/// <summary>
/// 44. Deleting a parameter (Remove Parameter)
///
/// BEFORE: Unnecessary parameter
/// </summary>
public class ReportGeneratorBefore
{
    public void GenerateReport(object data, string format, bool includeHeader = true)
    {
        if (format == "html")
        {
            // Always include header for HTML
            includeHeader = true;
        }
        // Generate report
    }
}

/// <summary>
/// AFTER: Remove unnecessary parameter
/// </summary>
public class ReportGeneratorAfter
{
    public void GenerateReport(object data, string format)
    {
        bool includeHeader = (format == "html");
        // Generate report
    }
}

/// <summary>
/// 45. Separation of Query and Modifier (Separate Query from Modifier)
///
/// BEFORE: Method that both queries and modifies
/// </summary>
public class BankAccountBefore
{
    private double balance = 0;

    public bool Withdraw(double amount)
    {
        if (balance >= amount)
        {
            balance -= amount;
            return true;
        }
        return false;
    }
}

/// <summary>
/// AFTER: Separate query from modifier
/// </summary>
public class BankAccountAfter
{
    private double balance = 0;

    public bool CanWithdraw(double amount)
    {
        return balance >= amount;
    }

    public bool Withdraw(double amount)
    {
        if (CanWithdraw(amount))
        {
            balance -= amount;
            return true;
        }
        return false;
    }
}

/// <summary>
/// 46. Parameterization of the method (Parameterize Method)
///
/// BEFORE: Similar methods with different values
/// </summary>
public class ReportGeneratorParamBefore
{
    public object GenerateWeeklyReport()
    {
        return GenerateReport(7);
    }

    public object GenerateMonthlyReport()
    {
        return GenerateReport(30);
    }

    public object GenerateQuarterlyReport()
    {
        return GenerateReport(90);
    }

    private object GenerateReport(int days)
    {
        // Generate report for specified days
        return null;
    }
}

/// <summary>
/// AFTER: Parameterize method
/// </summary>
public class ReportGeneratorParamAfter
{
    public object GenerateReport(int days)
    {
        // Generate report for specified days
        return null;
    }

    public object GenerateWeeklyReport()
    {
        return GenerateReport(7);
    }

    public object GenerateMonthlyReport()
    {
        return GenerateReport(30);
    }

    public object GenerateQuarterlyReport()
    {
        return GenerateReport(90);
    }
}

/// <summary>
/// 47. Replacing a parameter with explicit methods (Replace Parameter with Explicit Methods)
///
/// BEFORE: Parameter determines behavior
/// </summary>
public class EmployeeExplicitBefore
{
    public const int ENGINEER = 0;
    public const int SALESMAN = 1;
    public const int MANAGER = 2;

    private int type;

    public EmployeeExplicitBefore(int type)
    {
        this.type = type;
    }

    public double GetSalary(double baseSalary)
    {
        switch (type)
        {
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

/// <summary>
/// AFTER: Replace parameter with explicit methods
/// </summary>
public class EmployeeExplicitAfter
{
    public double GetEngineerSalary(double baseSalary)
    {
        return baseSalary * 1.0;
    }

    public double GetSalesmanSalary(double baseSalary)
    {
        return baseSalary * 1.1;
    }

    public double GetManagerSalary(double baseSalary)
    {
        return baseSalary * 1.2;
    }
}

/// <summary>
/// 48. Save the Whole Object
///
/// BEFORE: Passing individual fields
/// </summary>
public class OrderWholeBefore
{
    private Dictionary<string, string> customer;

    public OrderWholeBefore(string customerName, string customerAddress)
    {
        customer = new Dictionary<string, string>
        {
            ["name"] = customerName,
            ["address"] = customerAddress
        };
    }

    public double CalculateShipping()
    {
        return GetShippingCost(customer["name"], customer["address"]);
    }

    private double GetShippingCost(string name, string address)
    {
        // Calculate based on name and address
        return 10.0;
    }
}

/// <summary>
/// AFTER: Pass whole object
/// </summary>
public class CustomerWhole
{
    private string name;
    private string address;

    public CustomerWhole(string name, string address)
    {
        this.name = name;
        this.address = address;
    }

    public string GetName()
    {
        return name;
    }

    public string GetAddress()
    {
        return address;
    }
}

public class OrderWholeAfter
{
    private CustomerWhole customer;

    public OrderWholeAfter(CustomerWhole customer)
    {
        this.customer = customer;
    }

    public double CalculateShipping()
    {
        return GetShippingCost(customer);
    }

    private double GetShippingCost(CustomerWhole customer)
    {
        // Calculate based on customer object
        return 10.0;
    }
}

/// <summary>
/// 49. Replacing a parameter with a method call (Replace Parameter with Method)
///
/// BEFORE: Parameter calculated outside method
/// </summary>
public class DiscountCalculatorParamBefore
{
    public double CalculateDiscount(double price, string customerType)
    {
        // customerType passed in
        return price * GetDiscountRate(customerType);
    }

    private double GetDiscountRate(string customerType)
    {
        switch (customerType)
        {
            case "premium":
                return 0.1;
            case "regular":
                return 0.05;
            default:
                return 0.0;
        }
    }
}

public class OrderParamBefore
{
    private CustomerParam customer;

    public double GetDiscountedPrice(double price)
    {
        var calculator = new DiscountCalculatorParamBefore();
        return calculator.CalculateDiscount(price, customer.GetType());
    }
}

public class CustomerParam
{
    public string GetType() => "regular";
}

/// <summary>
/// AFTER: Replace parameter with method call
/// </summary>
public class DiscountCalculatorParamAfter
{
    public double CalculateDiscount(double price, CustomerParam customer)
    {
        return price * GetDiscountRate(customer.GetType());
    }

    private double GetDiscountRate(string customerType)
    {
        switch (customerType)
        {
            case "premium":
                return 0.1;
            case "regular":
                return 0.05;
            default:
                return 0.0;
        }
    }
}

public class OrderParamAfter
{
    private CustomerParam customer;

    public double GetDiscountedPrice(double price)
    {
        var calculator = new DiscountCalculatorParamAfter();
        return calculator.CalculateDiscount(price, customer);
    }
}

/// <summary>
/// 50. Introduction of the boundary object (Introduce Parameter Object)
///
/// BEFORE: Multiple parameters
/// </summary>
public class TemperatureRangeBefore
{
    public bool WithinRange(double minTemp, double maxTemp, double currentTemp)
    {
        return currentTemp >= minTemp && currentTemp <= maxTemp;
    }

    public double GetAverageTemp(double minTemp, double maxTemp)
    {
        return (minTemp + maxTemp) / 2;
    }
}

/// <summary>
/// AFTER: Introduce parameter object
/// </summary>
public class TemperatureRange
{
    private double minTemp;
    private double maxTemp;

    public TemperatureRange(double minTemp, double maxTemp)
    {
        this.minTemp = minTemp;
        this.maxTemp = maxTemp;
    }

    public double GetMinTemp()
    {
        return minTemp;
    }

    public double GetMaxTemp()
    {
        return maxTemp;
    }

    public bool WithinRange(double currentTemp)
    {
        return currentTemp >= minTemp && currentTemp <= maxTemp;
    }

    public double GetAverageTemp()
    {
        return (minTemp + maxTemp) / 2;
    }
}

/// <summary>
/// 51. Removing the Value Setting Method
///
/// BEFORE: Setter that's not needed
/// </summary>
public class SensorBefore
{
    private double temperature;

    public SensorBefore(double temperature)
    {
        this.temperature = temperature;
    }

    public double GetTemperature()
    {
        return temperature;
    }

    public void SetTemperature(double temperature) // Not needed if immutable
    {
        this.temperature = temperature;
    }
}

/// <summary>
/// AFTER: Remove setter for immutable object
/// </summary>
public class SensorAfter
{
    private double temperature;

    public SensorAfter(double temperature)
    {
        this.temperature = temperature;
    }

    public double GetTemperature()
    {
        return temperature;
    }

    // SetTemperature removed
}

/// <summary>
/// 52. Hiding a method (Hide Method)
///
/// BEFORE: Public method that should be private
/// </summary>
public class DataProcessorHideBefore
{
    public bool ValidateData(object data) // Should be private
    {
        return data != null && data is Array;
    }

    public void ProcessData(object data)
    {
        if (ValidateData(data))
        {
            // Process data
        }
    }
}

/// <summary>
/// AFTER: Hide method
/// </summary>
public class DataProcessorHideAfter
{
    private bool ValidateData(object data)
    {
        return data != null && data is Array;
    }

    public void ProcessData(object data)
    {
        if (ValidateData(data))
        {
            // Process data
        }
    }
}

/// <summary>
/// 53. Replacing the constructor with the factory method (Replace Constructor with Factory Method)
///
/// BEFORE: Complex constructor
/// </summary>
public class ComplexObjectBefore
{
    private string type;
    private Dictionary<string, object> config;

    public ComplexObjectBefore(string type, Dictionary<string, object> config = null)
    {
        this.type = type;
        this.config = config ?? new Dictionary<string, object>();

        if (type == "database")
        {
            // Merge default database config
            var defaults = new Dictionary<string, object>
            {
                ["host"] = "localhost",
                ["port"] = 3306
            };
            foreach (var kvp in defaults)
            {
                if (!this.config.ContainsKey(kvp.Key))
                {
                    this.config[kvp.Key] = kvp.Value;
                }
            }
        }
        else if (type == "file")
        {
            // Merge default file config
            var defaults = new Dictionary<string, object>
            {
                ["path"] = "/tmp",
                ["format"] = "json"
            };
            foreach (var kvp in defaults)
            {
                if (!this.config.ContainsKey(kvp.Key))
                {
                    this.config[kvp.Key] = kvp.Value;
                }
            }
        }
    }
}

/// <summary>
/// AFTER: Replace constructor with factory method
/// </summary>
public class ComplexObjectAfter
{
    private string type;
    private Dictionary<string, object> config;

    private ComplexObjectAfter(string type, Dictionary<string, object> config)
    {
        this.type = type;
        this.config = config;
    }

    public static ComplexObjectAfter CreateDatabaseConnection(Dictionary<string, object> config = null)
    {
        var finalConfig = config ?? new Dictionary<string, object>();
        var defaults = new Dictionary<string, object>
        {
            ["host"] = "localhost",
            ["port"] = 3306
        };

        foreach (var kvp in defaults)
        {
            if (!finalConfig.ContainsKey(kvp.Key))
            {
                finalConfig[kvp.Key] = kvp.Value;
            }
        }

        return new ComplexObjectAfter("database", finalConfig);
    }

    public static ComplexObjectAfter CreateFileHandler(Dictionary<string, object> config = null)
    {
        var finalConfig = config ?? new Dictionary<string, object>();
        var defaults = new Dictionary<string, object>
        {
            ["path"] = "/tmp",
            ["format"] = "json"
        };

        foreach (var kvp in defaults)
        {
            if (!finalConfig.ContainsKey(kvp.Key))
            {
                finalConfig[kvp.Key] = kvp.Value;
            }
        }

        return new ComplexObjectAfter("file", finalConfig);
    }
}

/// <summary>
/// 54. Encapsulation of top-down type conversion (Encapsulate Downcast)
///
/// BEFORE: Downcast in client code
/// </summary>
public class ShapeCollectionBefore
{
    private List<object> shapes = new List<object>();

    public void AddShape(object shape)
    {
        shapes.Add(shape);
    }

    public List<object> GetShapes()
    {
        return shapes;
    }
}

public class Circle { }
public class Square { }

// Client code
// var collection = new ShapeCollectionBefore();
// ... add shapes
// var circles = collection.GetShapes().Where(shape => shape is Circle).ToList();

/// <summary>
/// AFTER: Encapsulate downcast
/// </summary>
public class ShapeCollectionAfter
{
    private List<object> shapes = new List<object>();

    public void AddShape(object shape)
    {
        shapes.Add(shape);
    }

    public List<Circle> GetCircles()
    {
        return shapes.Where(shape => shape is Circle).Cast<Circle>().ToList();
    }

    public List<Square> GetSquares()
    {
        return shapes.Where(shape => shape is Square).Cast<Square>().ToList();
    }
}

/// <summary>
/// 55. Replacing the error code with an exceptional situation (Replace Error Code with Exception)
///
/// BEFORE: Error codes
/// </summary>
public class FileReaderErrorBefore
{
    public const int FILE_NOT_FOUND = 1;
    public const int PERMISSION_DENIED = 2;

    public object ReadFile(string filename)
    {
        if (!System.IO.File.Exists(filename))
        {
            return FILE_NOT_FOUND;
        }

        // For simplicity, assuming readable if exists
        if (!System.IO.File.Exists(filename)) // Simplified check
        {
            return PERMISSION_DENIED;
        }

        return System.IO.File.ReadAllText(filename);
    }
}

// Client code
// var reader = new FileReaderErrorBefore();
// var result = reader.ReadFile("test.txt");
// if ((int)result == FileReaderErrorBefore.FILE_NOT_FOUND) {
//     // Handle error
// } else if ((int)result == FileReaderErrorBefore.PERMISSION_DENIED) {
//     // Handle error
// } else {
//     // Use content
// }

/// <summary>
/// AFTER: Replace error codes with exceptions
/// </summary>
public class FileNotFoundException : Exception
{
    public FileNotFoundException(string message) : base(message) { }
}

public class PermissionDeniedException : Exception
{
    public PermissionDeniedException(string message) : base(message) { }
}

public class FileReaderExceptionAfter
{
    public string ReadFile(string filename)
    {
        if (!System.IO.File.Exists(filename))
        {
            throw new FileNotFoundException($"File not found: {filename}");
        }

        // Simplified permission check
        if (!System.IO.File.Exists(filename)) // Would normally check permissions
        {
            throw new PermissionDeniedException($"Permission denied: {filename}");
        }

        return System.IO.File.ReadAllText(filename);
    }
}

// Client code
// try {
//     var reader = new FileReaderExceptionAfter();
//     var content = reader.ReadFile("test.txt");
//     // Use content
// } catch (FileNotFoundException e) {
//     // Handle file not found
// } catch (PermissionDeniedException e) {
//     // Handle permission denied
// }

/// <summary>
/// 56. Replacing an exceptional situation with a check (Replace Exception with Test)
///
/// BEFORE: Using exception for control flow
/// </summary>
public class EmptyStackException : Exception { }

public class StackExceptionBefore
{
    private List<object> items = new List<object>();

    public object Pop()
    {
        if (items.Count == 0)
        {
            throw new EmptyStackException();
        }
        var item = items[items.Count - 1];
        items.RemoveAt(items.Count - 1);
        return item;
    }
}

// Client code
// var stack = new StackExceptionBefore();
// try {
//     var item = stack.Pop();
// } catch (EmptyStackException e) {
//     item = null; // Default value
// }

/// <summary>
/// AFTER: Replace exception with test
/// </summary>
public class StackTestAfter
{
    private List<object> items = new List<object>();

    public bool IsEmpty()
    {
        return items.Count == 0;
    }

    public object Pop()
    {
        if (items.Count == 0) return null;
        var item = items[items.Count - 1];
        items.RemoveAt(items.Count - 1);
        return item;
    }
}

// Client code
// var stack = new StackTestAfter();
// var item = stack.IsEmpty() ? null : stack.Pop();
