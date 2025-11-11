using System;
using System.Collections.Generic;

/// <summary>
/// 9. Substitution Algorithm
///
/// BEFORE: Complex algorithm that can be simplified
/// </summary>
public class PricingServiceBefore
{
    public double CalculatePrice(List<Dictionary<string, object>> items)
    {
        double total = 0;
        foreach (var item in items)
        {
            if ((string)item["type"] == "book")
            {
                total += (double)item["price"] * 0.9; // 10% discount for books
            }
            else if ((string)item["type"] == "electronics")
            {
                total += (double)item["price"] * 1.1; // 10% markup for electronics
            }
            else
            {
                total += (double)item["price"];
            }
        }
        return total;
    }
}

/// <summary>
/// AFTER: Substitute with a simpler algorithm
/// </summary>
public class PricingServiceAfter
{
    private readonly Dictionary<string, double> discounts = new Dictionary<string, double>
    {
        ["book"] = 0.9,
        ["electronics"] = 1.1,
        ["default"] = 1.0
    };

    public double CalculatePrice(List<Dictionary<string, object>> items)
    {
        double total = 0;
        foreach (var item in items)
        {
            double multiplier = discounts.GetValueOrDefault((string)item["type"], discounts["default"]);
            total += (double)item["price"] * multiplier;
        }
        return total;
    }
}

/// <summary>
/// 10. Moving functions between objects (Move Method)
///
/// BEFORE: Method in wrong class
/// </summary>
public class AccountBefore
{
    private double balance;

    public AccountBefore(double balance)
    {
        this.balance = balance;
    }

    public double GetBalance()
    {
        return balance;
    }

    // This method belongs in Bank class, not Account
    public bool TransferTo(AccountBefore target, double amount)
    {
        if (balance >= amount)
        {
            balance -= amount;
            target.balance += amount;
            return true;
        }
        return false;
    }
}

/// <summary>
/// AFTER: Move method to appropriate class
/// </summary>
public class AccountAfter
{
    private double balance;

    public AccountAfter(double balance)
    {
        this.balance = balance;
    }

    public double GetBalance()
    {
        return balance;
    }

    public void DecreaseBalance(double amount)
    {
        balance -= amount;
    }

    public void IncreaseBalance(double amount)
    {
        balance += amount;
    }
}

public class Bank
{
    public bool Transfer(AccountAfter from, AccountAfter to, double amount)
    {
        if (from.GetBalance() >= amount)
        {
            from.DecreaseBalance(amount);
            to.IncreaseBalance(amount);
            return true;
        }
        return false;
    }
}

/// <summary>
/// 11. Moving the field (Move Field)
///
/// BEFORE: Field in wrong class
/// </summary>
public class CustomerBefore
{
    private string name;
    private Dictionary<string, string> address; // This should be in Address class

    public CustomerBefore(string name, string street, string city, string zipCode)
    {
        this.name = name;
        this.address = new Dictionary<string, string>
        {
            ["street"] = street,
            ["city"] = city,
            ["zipCode"] = zipCode
        };
    }

    public string GetAddress()
    {
        return $"{address["street"]}, {address["city"]} {address["zipCode"]}";
    }
}

/// <summary>
/// AFTER: Move field to dedicated class
/// </summary>
public class Address
{
    private string street;
    private string city;
    private string zipCode;

    public Address(string street, string city, string zipCode)
    {
        this.street = street;
        this.city = city;
        this.zipCode = zipCode;
    }

    public string GetFullAddress()
    {
        return $"{street}, {city} {zipCode}";
    }
}

public class CustomerAfter
{
    private string name;
    private Address address;

    public CustomerAfter(string name, Address address)
    {
        this.name = name;
        this.address = address;
    }

    public string GetAddress()
    {
        return address.GetFullAddress();
    }
}

/// <summary>
/// 12. Class Allocation (Extract Class)
///
/// BEFORE: Class has too many responsibilities
/// </summary>
public class PersonBefore
{
    private string name;
    private string phoneNumber;
    private string officeAreaCode;
    private string officeNumber;

    public string GetTelephoneNumber()
    {
        return $"({officeAreaCode}) {officeNumber}";
    }
}

/// <summary>
/// AFTER: Extract telephone number to separate class
/// </summary>
public class TelephoneNumber
{
    private string areaCode;
    private string number;

    public TelephoneNumber(string areaCode, string number)
    {
        this.areaCode = areaCode;
        this.number = number;
    }

    public string GetTelephoneNumber()
    {
        return $"({areaCode}) {number}";
    }
}

public class PersonAfter
{
    private string name;
    private string phoneNumber;
    private TelephoneNumber officeTelephone;

    public PersonAfter(string name)
    {
        this.name = name;
    }

    public string GetOfficeTelephone()
    {
        return officeTelephone?.GetTelephoneNumber();
    }

    public void SetOfficeTelephone(TelephoneNumber telephone)
    {
        officeTelephone = telephone;
    }
}

/// <summary>
/// 13. Embedding a class (Inline Class)
///
/// BEFORE: Unnecessary class with single responsibility
/// </summary>
public class OrderProcessorBefore
{
    private OrderValidator validator;

    public OrderProcessorBefore()
    {
        validator = new OrderValidator();
    }

    public void Process(Dictionary<string, object> order)
    {
        if (validator.IsValid(order))
        {
            // Process order
        }
    }
}

public class OrderValidator
{
    public bool IsValid(Dictionary<string, object> order)
    {
        return (double)order["total"] > 0;
    }
}

/// <summary>
/// AFTER: Inline the class
/// </summary>
public class OrderProcessorAfter
{
    public void Process(Dictionary<string, object> order)
    {
        if (IsValidOrder(order))
        {
            // Process order
        }
    }

    private bool IsValidOrder(Dictionary<string, object> order)
    {
        return (double)order["total"] > 0;
    }
}

/// <summary>
/// 14. Hiding delegation (Hide Delegate)
///
/// BEFORE: Client has to know about delegation
/// </summary>
public class DepartmentBefore
{
    private Person personManager;

    public DepartmentBefore(Person personManager)
    {
        this.personManager = personManager;
    }

    public Person GetManager()
    {
        return personManager;
    }
}

public class Person
{
    private DepartmentBefore department;

    public DepartmentBefore GetDepartment()
    {
        return department;
    }
}

// Client code
// var manager = person.GetDepartment().GetManager();

/// <summary>
/// AFTER: Hide the delegation
/// </summary>
public class DepartmentAfter
{
    private PersonAfter personManager;

    public DepartmentAfter(PersonAfter personManager)
    {
        this.personManager = personManager;
    }

    public PersonAfter GetManager()
    {
        return personManager;
    }
}

public class PersonAfter
{
    private DepartmentAfter department;

    public DepartmentAfter GetDepartment()
    {
        return department;
    }

    public PersonAfter GetManager()
    {
        return department.GetManager();
    }
}

// Client code - much cleaner
// var manager = person.GetManager();

/// <summary>
/// 15. Removing the intermediary (Remove Middle Man)
///
/// BEFORE: Too much delegation
/// </summary>
public class PersonWithMiddleMan
{
    private DepartmentAfter department;

    public DepartmentAfter GetDepartment()
    {
        return department;
    }

    public PersonAfter GetManager()
    {
        return department.GetManager();
    }

    public string GetDepartmentName()
    {
        return department.GetName();
    }
}

/// <summary>
/// AFTER: Remove middle man if delegation is too heavy
/// </summary>
public class PersonDirect
{
    private DepartmentAfter department;
    private PersonAfter manager; // Direct reference

    public PersonAfter GetManager()
    {
        return manager;
    }

    public DepartmentAfter GetDepartment()
    {
        return department;
    }
}

/// <summary>
/// 16. Introduction of an external method (Introduce Foreign Method)
///
/// BEFORE: Using external class method in wrong place
/// </summary>
public class ReportGeneratorBefore
{
    public void GenerateReport()
    {
        var date = DateTime.Now;
        var nextMonth = date.AddMonths(1); // Foreign method usage

        // Generate report for next month
    }
}

/// <summary>
/// AFTER: Introduce foreign method
/// </summary>
public class ReportGeneratorAfter
{
    public void GenerateReport()
    {
        var date = DateTime.Now;
        var nextMonth = NextMonth(date);

        // Generate report for next month
    }

    private DateTime NextMonth(DateTime date)
    {
        return date.AddMonths(1);
    }
}

/// <summary>
/// 17. The introduction of local extension (Introduce Local Extension)
///
/// BEFORE: Adding methods to external class (not possible)
/// </summary>
public static class DateUtil
{
    public static DateTime NextMonth(DateTime date)
    {
        return date.AddMonths(1);
    }

    public static DateTime PreviousMonth(DateTime date)
    {
        return date.AddMonths(-1);
    }
}

/// <summary>
/// AFTER: Create local extension class
/// </summary>
public class DateTimeExtension : DateTime
{
    public DateTimeExtension() : base() { }

    public DateTime NextMonth()
    {
        return AddMonths(1);
    }

    public DateTime PreviousMonth()
    {
        return AddMonths(-1);
    }
}
