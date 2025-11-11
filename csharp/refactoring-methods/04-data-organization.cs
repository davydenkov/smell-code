using System;
using System.Collections.Generic;

/// <summary>
/// 18. Self-Encapsulate Field
///
/// BEFORE: Direct field access
/// </summary>
public class PersonBefore
{
    public string Name; // Direct access

    public string GetName()
    {
        return Name;
    }

    public void SetName(string name)
    {
        Name = name;
    }
}

/// <summary>
/// AFTER: Self-encapsulate field
/// </summary>
public class PersonAfter
{
    private string name;

    public string GetName()
    {
        return name;
    }

    public void SetName(string name)
    {
        this.name = name;
    }
}

/// <summary>
/// 19. Replacing the data value with an object (Replace Data Value with Object)
///
/// BEFORE: Primitive data type that should be an object
/// </summary>
public class OrderBefore
{
    private string customer; // Just a string

    public string GetCustomerName()
    {
        return customer;
    }

    public void SetCustomer(string customer)
    {
        this.customer = customer;
    }
}

/// <summary>
/// AFTER: Replace with object
/// </summary>
public class Customer
{
    private string name;

    public Customer(string name)
    {
        this.name = name;
    }

    public string GetName()
    {
        return name;
    }
}

public class OrderAfter
{
    private Customer customer;

    public Customer GetCustomer()
    {
        return customer;
    }

    public void SetCustomer(Customer customer)
    {
        this.customer = customer;
    }

    public string GetCustomerName()
    {
        return customer?.GetName();
    }
}

/// <summary>
/// 20. Replacing the value with a reference (Change Value to Reference)
///
/// BEFORE: Multiple instances of same object
/// </summary>
public class CustomerValue
{
    private string name;

    public CustomerValue(string name)
    {
        this.name = name;
    }

    public string GetName()
    {
        return name;
    }
}

public class OrderValue
{
    private CustomerValue customer; // New instance for each order

    public OrderValue(string customerName)
    {
        customer = new CustomerValue(customerName);
    }
}

/// <summary>
/// AFTER: Use reference to single instance
/// </summary>
public class CustomerReference
{
    private string name;
    private static Dictionary<string, CustomerReference> instances = new Dictionary<string, CustomerReference>();

    private CustomerReference(string name)
    {
        this.name = name;
    }

    public static CustomerReference Create(string name)
    {
        if (!instances.ContainsKey(name))
        {
            instances[name] = new CustomerReference(name);
        }
        return instances[name];
    }

    public string GetName()
    {
        return name;
    }
}

public class OrderReference
{
    private CustomerReference customer;

    public OrderReference(string customerName)
    {
        customer = CustomerReference.Create(customerName);
    }
}

/// <summary>
/// 21. Replacing a reference with a value (Change Reference to Value)
///
/// BEFORE: Unnecessary reference when value would suffice
/// </summary>
public class CurrencyReference
{
    private string code;

    public CurrencyReference(string code)
    {
        this.code = code;
    }

    public string GetCode()
    {
        return code;
    }
}

public class ProductReference
{
    private double price;
    private CurrencyReference currency; // Reference object

    public ProductReference(double price, CurrencyReference currency)
    {
        this.price = price;
        this.currency = currency;
    }
}

/// <summary>
/// AFTER: Use value object instead
/// </summary>
public class CurrencyValue
{
    private string code;

    public CurrencyValue(string code)
    {
        this.code = code;
    }

    public string GetCode()
    {
        return code;
    }
}

public class ProductValue
{
    private double price;
    private string currencyCode; // Just the value

    public ProductValue(double price, string currencyCode)
    {
        this.price = price;
        this.currencyCode = currencyCode;
    }

    public string GetCurrencyCode()
    {
        return currencyCode;
    }
}

/// <summary>
/// 22. Replacing an array with an object (Replace Array with Object)
///
/// BEFORE: Using array for structured data
/// </summary>
public class PerformanceArray
{
    public Dictionary<string, object> GetPerformanceData()
    {
        return new Dictionary<string, object>
        {
            ["goals"] = 10,
            ["assists"] = 5,
            ["minutes"] = 120
        };
    }

    public double CalculateScore(Dictionary<string, object> data)
    {
        return ((int)data["goals"] * 2) + ((int)data["assists"] * 1.5) + ((int)data["minutes"] / 60.0);
    }
}

/// <summary>
/// AFTER: Replace array with object
/// </summary>
public class PerformanceData
{
    private int goals;
    private int assists;
    private int minutes;

    public PerformanceData(int goals, int assists, int minutes)
    {
        this.goals = goals;
        this.assists = assists;
        this.minutes = minutes;
    }

    public int GetGoals()
    {
        return goals;
    }

    public int GetAssists()
    {
        return assists;
    }

    public int GetMinutes()
    {
        return minutes;
    }

    public double CalculateScore()
    {
        return (goals * 2) + (assists * 1.5) + (minutes / 60.0);
    }
}

public class PerformanceObject
{
    public PerformanceData GetPerformanceData()
    {
        return new PerformanceData(10, 5, 120);
    }

    public double CalculateScore(PerformanceData data)
    {
        return data.CalculateScore();
    }
}

/// <summary>
/// 23. Duplication of visible data (Duplicate Observed Data)
///
/// BEFORE: Domain data mixed with presentation
/// </summary>
public class OrderDomain
{
    private double total = 0;

    public void AddItem(double price)
    {
        total += price;
        // Have to update UI here too
        UpdateDisplay();
    }

    private void UpdateDisplay()
    {
        // Update UI elements
    }
}

/// <summary>
/// AFTER: Separate domain and presentation data
/// </summary>
public interface IOrderObserver
{
    void Update(double total);
}

public class OrderDomainSeparated
{
    private double total = 0;
    private List<IOrderObserver> observers = new List<IOrderObserver>();

    public void AddItem(double price)
    {
        total += price;
        NotifyObservers();
    }

    public double GetTotal()
    {
        return total;
    }

    public void AddObserver(IOrderObserver observer)
    {
        observers.Add(observer);
    }

    private void NotifyObservers()
    {
        foreach (var observer in observers)
        {
            observer.Update(total);
        }
    }
}

public class OrderDisplay : IOrderObserver
{
    private OrderDomainSeparated order;

    public OrderDisplay(OrderDomainSeparated order)
    {
        this.order = order;
        this.order.AddObserver(this);
    }

    public void Update(double total)
    {
        // Update display with new total
    }
}

/// <summary>
/// 24. Replacing Unidirectional communication with Bidirectional
/// communication (Change Unidirectional Association to Bidirectional)
///
/// BEFORE: One-way association
/// </summary>
public class CustomerUni
{
    private List<object> orders = new List<object>();

    public void AddOrder(object order)
    {
        orders.Add(order);
        // Order doesn't know about customer
    }
}

public class OrderUni
{
    private List<object> items = new List<object>();
}

/// <summary>
/// AFTER: Bidirectional association
/// </summary>
public class CustomerBi
{
    private List<OrderBi> orders = new List<OrderBi>();

    public void AddOrder(OrderBi order)
    {
        orders.Add(order);
        order.SetCustomer(this);
    }
}

public class OrderBi
{
    private CustomerBi customer;
    private List<object> items = new List<object>();

    public void SetCustomer(CustomerBi customer)
    {
        this.customer = customer;
    }

    public CustomerBi GetCustomer()
    {
        return customer;
    }
}

/// <summary>
/// 25. Replacing Bidirectional communication with Unidirectional
/// communication (Change Bidirectional Association to Unidirectional)
///
/// BEFORE: Unnecessary bidirectional association
/// </summary>
public class CustomerBidirectional
{
    private List<OrderBidirectional> orders = new List<OrderBidirectional>();

    public void AddOrder(OrderBidirectional order)
    {
        orders.Add(order);
        order.SetCustomer(this);
    }
}

public class OrderBidirectional
{
    private CustomerBidirectional customer;

    public void SetCustomer(CustomerBidirectional customer)
    {
        this.customer = customer;
    }

    public CustomerBidirectional GetCustomer()
    {
        return customer;
    }
}

/// <summary>
/// AFTER: Remove bidirectional link
/// </summary>
public class CustomerUnidirectional
{
    private List<OrderUnidirectional> orders = new List<OrderUnidirectional>();

    public void AddOrder(OrderUnidirectional order)
    {
        orders.Add(order);
    }
}

public class OrderUnidirectional
{
    private string customerId;

    public OrderUnidirectional(string customerId)
    {
        this.customerId = customerId;
    }

    public string GetCustomerId()
    {
        return customerId;
    }
}

/// <summary>
/// 26. Replacing the magic number with a symbolic constant
/// (Replace Magic Number with Symbolic Constant)
///
/// BEFORE: Magic numbers
/// </summary>
public class GeometryBefore
{
    public double CalculateCircleArea(double radius)
    {
        return 3.14159 * radius * radius; // Magic number
    }

    public double CalculateCircleCircumference(double radius)
    {
        return 2 * 3.14159 * radius; // Same magic number
    }
}

/// <summary>
/// AFTER: Use symbolic constant
/// </summary>
public class GeometryAfter
{
    private const double PI = 3.14159;

    public double CalculateCircleArea(double radius)
    {
        return PI * radius * radius;
    }

    public double CalculateCircleCircumference(double radius)
    {
        return 2 * PI * radius;
    }
}

/// <summary>
/// 27. Encapsulate Field
///
/// BEFORE: Public field
/// </summary>
public class PersonPublic
{
    public string Name;
}

/// <summary>
/// AFTER: Encapsulated field
/// </summary>
public class PersonEncapsulated
{
    private string name;

    public string GetName()
    {
        return name;
    }

    public void SetName(string name)
    {
        this.name = name;
    }
}

/// <summary>
/// 28. Encapsulate Collection
///
/// BEFORE: Direct access to collection
/// </summary>
public class TeamBefore
{
    public List<string> Players = new List<string>(); // Direct access

    public void AddPlayer(string player)
    {
        Players.Add(player);
    }
}

/// <summary>
/// AFTER: Encapsulated collection
/// </summary>
public class TeamAfter
{
    private List<string> players = new List<string>();

    public void AddPlayer(string player)
    {
        players.Add(player);
    }

    public void RemovePlayer(string player)
    {
        players.Remove(player);
    }

    public List<string> GetPlayers()
    {
        return new List<string>(players); // Return copy
    }

    public int GetPlayerCount()
    {
        return players.Count;
    }
}

/// <summary>
/// 29. Replacing a record with a Data Class
///
/// BEFORE: Using array as data structure
/// </summary>
public class EmployeeArray
{
    public Dictionary<string, object> CreateEmployee(Dictionary<string, object> data)
    {
        return new Dictionary<string, object>
        {
            ["name"] = data["name"],
            ["salary"] = data["salary"],
            ["department"] = data["department"]
        };
    }

    public object GetSalary(Dictionary<string, object> employee)
    {
        return employee["salary"];
    }
}

/// <summary>
/// AFTER: Use data class
/// </summary>
public class Employee
{
    private string name;
    private double salary;
    private string department;

    public Employee(string name, double salary, string department)
    {
        this.name = name;
        this.salary = salary;
        this.department = department;
    }

    public string GetName()
    {
        return name;
    }

    public double GetSalary()
    {
        return salary;
    }

    public string GetDepartment()
    {
        return department;
    }
}

public class EmployeeDataClass
{
    public Employee CreateEmployee(string name, double salary, string department)
    {
        return new Employee(name, salary, department);
    }

    public double GetSalary(Employee employee)
    {
        return employee.GetSalary();
    }
}

/// <summary>
/// 30. Replacing Type Code with Class
///
/// BEFORE: Type code as constants
/// </summary>
public class EmployeeTypeCode
{
    public const int ENGINEER = 0;
    public const int SALESMAN = 1;
    public const int MANAGER = 2;

    private int type;

    public EmployeeTypeCode(int type)
    {
        this.type = type;
    }

    public int GetTypeCode()
    {
        return type;
    }

    public double GetMonthlySalary()
    {
        switch (type)
        {
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

/// <summary>
/// AFTER: Replace type code with class
/// </summary>
public abstract class EmployeeType
{
    public abstract double GetMonthlySalary();

    public static EngineerType CreateEngineer()
    {
        return new EngineerType();
    }

    public static SalesmanType CreateSalesman()
    {
        return new SalesmanType();
    }

    public static ManagerType CreateManager()
    {
        return new ManagerType();
    }
}

public class EngineerType : EmployeeType
{
    public override double GetMonthlySalary()
    {
        return 5000;
    }
}

public class SalesmanType : EmployeeType
{
    public override double GetMonthlySalary()
    {
        return 4000;
    }
}

public class ManagerType : EmployeeType
{
    public override double GetMonthlySalary()
    {
        return 6000;
    }
}

public class EmployeeTypeClass
{
    private EmployeeType type;

    public EmployeeTypeClass(EmployeeType type)
    {
        this.type = type;
    }

    public double GetMonthlySalary()
    {
        return type.GetMonthlySalary();
    }
}

/// <summary>
/// 31. Replacing Type Code with Subclasses
///
/// BEFORE: Type code in base class
/// </summary>
public class EmployeeSubBefore
{
    public const int ENGINEER = 0;
    public const int SALESMAN = 1;
    public const int MANAGER = 2;

    private int type;
    private double salary;

    public EmployeeSubBefore(int type, double salary)
    {
        this.type = type;
        this.salary = salary;
    }

    public double GetSalary()
    {
        return salary;
    }

    public int GetType()
    {
        return type;
    }
}

/// <summary>
/// AFTER: Replace type code with subclasses
/// </summary>
public abstract class EmployeeSubAfter
{
    protected double salary;

    public EmployeeSubAfter(double salary)
    {
        this.salary = salary;
    }

    public double GetSalary()
    {
        return salary;
    }

    public abstract string GetType();
}

public class Engineer : EmployeeSubAfter
{
    public Engineer(double salary) : base(salary) { }

    public override string GetType()
    {
        return "engineer";
    }
}

public class Salesman : EmployeeSubAfter
{
    public Salesman(double salary) : base(salary) { }

    public override string GetType()
    {
        return "salesman";
    }
}

public class Manager : EmployeeSubAfter
{
    public Manager(double salary) : base(salary) { }

    public override string GetType()
    {
        return "manager";
    }
}

/// <summary>
/// 32. Replacing Type Code with State/Strategy
///
/// BEFORE: Type code with behavior
/// </summary>
public class EmployeeStateBefore
{
    public const int JUNIOR = 0;
    public const int SENIOR = 1;
    public const int LEAD = 2;

    private int level;

    public EmployeeStateBefore(int level)
    {
        this.level = level;
    }

    public double GetSalaryMultiplier()
    {
        switch (level)
        {
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

/// <summary>
/// AFTER: Use state/strategy pattern
/// </summary>
public interface IEmployeeLevel
{
    double GetSalaryMultiplier();
}

public class JuniorLevel : IEmployeeLevel
{
    public double GetSalaryMultiplier()
    {
        return 1.0;
    }
}

public class SeniorLevel : IEmployeeLevel
{
    public double GetSalaryMultiplier()
    {
        return 1.5;
    }
}

public class LeadLevel : IEmployeeLevel
{
    public double GetSalaryMultiplier()
    {
        return 2.0;
    }
}

public class EmployeeStateAfter
{
    private IEmployeeLevel level;

    public EmployeeStateAfter(IEmployeeLevel level)
    {
        this.level = level;
    }

    public double GetSalaryMultiplier()
    {
        return level.GetSalaryMultiplier();
    }
}

/// <summary>
/// 33. Replacing Subclass with Fields
///
/// BEFORE: Unnecessary subclasses
/// </summary>
public abstract class PersonSub
{
    protected string name;
    protected string gender;

    public PersonSub(string name, string gender)
    {
        this.name = name;
        this.gender = gender;
    }

    public string GetName()
    {
        return name;
    }

    public abstract bool IsMale();
}

public class Male : PersonSub
{
    public Male(string name) : base(name, "male") { }

    public override bool IsMale()
    {
        return true;
    }
}

public class Female : PersonSub
{
    public Female(string name) : base(name, "female") { }

    public override bool IsMale()
    {
        return false;
    }
}

/// <summary>
/// AFTER: Replace subclass with field
/// </summary>
public class PersonField
{
    private string name;
    private string gender; // 'male' or 'female'

    public PersonField(string name, string gender)
    {
        this.name = name;
        this.gender = gender;
    }

    public string GetName()
    {
        return name;
    }

    public bool IsMale()
    {
        return gender == "male";
    }

    public string GetGender()
    {
        return gender;
    }
}
