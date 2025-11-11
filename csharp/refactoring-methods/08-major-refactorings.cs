using System;
using System.Collections.Generic;

/// <summary>
/// 69. Separation of inheritance (Tease Apart Inheritance)
///
/// BEFORE: Class hierarchy mixing two different responsibilities
/// </summary>
public class EmployeeTeaseBefore
{
    protected string name;
    protected double rate;

    public EmployeeTeaseBefore(string name, double rate)
    {
        this.name = name;
        this.rate = rate;
    }

    public string GetName()
    {
        return name;
    }
}

public class SalariedEmployeeTeaseBefore : EmployeeTeaseBefore
{
    public SalariedEmployeeTeaseBefore(string name, double rate) : base(name, rate) { }

    public double GetPay()
    {
        return rate;
    }
}

public class CommissionedEmployeeTeaseBefore : EmployeeTeaseBefore
{
    private double commission;

    public CommissionedEmployeeTeaseBefore(string name, double rate, double commission)
        : base(name, rate)
    {
        this.commission = commission;
    }

    public double GetPay()
    {
        return rate + commission;
    }
}

/// <summary>
/// AFTER: Tease apart inheritance into two separate hierarchies
/// </summary>
public interface IPayable
{
    double GetPay();
}

public class EmployeeTeaseAfter
{
    protected string name;

    public EmployeeTeaseAfter(string name)
    {
        this.name = name;
    }

    public string GetName()
    {
        return name;
    }
}

public class SalariedEmployeeTeaseAfter : EmployeeTeaseAfter, IPayable
{
    private double salary;

    public SalariedEmployeeTeaseAfter(string name, double salary) : base(name)
    {
        this.salary = salary;
    }

    public double GetPay()
    {
        return salary;
    }
}

public class CommissionedEmployeeTeaseAfter : EmployeeTeaseAfter, IPayable
{
    private double baseSalary;
    private double commission;

    public CommissionedEmployeeTeaseAfter(string name, double baseSalary, double commission)
        : base(name)
    {
        this.baseSalary = baseSalary;
        this.commission = commission;
    }

    public double GetPay()
    {
        return baseSalary + commission;
    }
}

/// <summary>
/// 70. Converting a procedural project into objects (Convert Procedural Design to Objects)
///
/// BEFORE: Procedural code with global functions and data
/// </summary>
public class ProceduralDesignBefore
{
    private static Dictionary<string, double> accounts = new Dictionary<string, double>();

    public static void CreateAccount(string id, double balance)
    {
        accounts[id] = balance;
    }

    public static double GetBalance(string id)
    {
        return accounts.ContainsKey(id) ? accounts[id] : 0;
    }

    public static void Deposit(string id, double amount)
    {
        if (accounts.ContainsKey(id))
        {
            accounts[id] += amount;
        }
    }

    public static bool Withdraw(string id, double amount)
    {
        if (accounts.ContainsKey(id) && accounts[id] >= amount)
        {
            accounts[id] -= amount;
            return true;
        }
        return false;
    }
}

/// <summary>
/// AFTER: Convert to object-oriented design
/// </summary>
public class Account
{
    private string id;
    private double balance;

    public Account(string id, double balance = 0)
    {
        this.id = id;
        this.balance = balance;
    }

    public string GetId()
    {
        return id;
    }

    public double GetBalance()
    {
        return balance;
    }

    public void Deposit(double amount)
    {
        balance += amount;
    }

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

public class Bank
{
    private Dictionary<string, Account> accounts = new Dictionary<string, Account>();

    public Account CreateAccount(string id, double balance = 0)
    {
        var account = new Account(id, balance);
        accounts[id] = account;
        return account;
    }

    public Account GetAccount(string id)
    {
        return accounts.ContainsKey(id) ? accounts[id] : null;
    }

    public double GetBalance(string id)
    {
        var account = GetAccount(id);
        return account?.GetBalance() ?? 0;
    }

    public void Deposit(string id, double amount)
    {
        var account = GetAccount(id);
        account?.Deposit(amount);
    }

    public bool Withdraw(string id, double amount)
    {
        var account = GetAccount(id);
        return account?.Withdraw(amount) ?? false;
    }
}

/// <summary>
/// 71. Separating the domain from the representation (Separate Domain from Presentation)
///
/// BEFORE: Domain logic mixed with presentation
/// </summary>
public class OrderPresentationBefore
{
    private List<Dictionary<string, object>> items = new List<Dictionary<string, object>>();
    private double total = 0;

    public void AddItem(string name, double price, int quantity)
    {
        items.Add(new Dictionary<string, object>
        {
            ["name"] = name,
            ["price"] = price,
            ["quantity"] = quantity
        });
        total += price * quantity;

        // Presentation logic mixed in
        Console.WriteLine($"Added {quantity} x {name} to order");
        Console.WriteLine($"Current total: ${total:F2}");
    }

    public double GetTotal()
    {
        return total;
    }

    public void DisplayOrder()
    {
        Console.WriteLine("Order Summary:");
        foreach (var item in items)
        {
            Console.WriteLine($"- {item["quantity"]} x {item["name"]} @ ${Convert.ToDouble(item["price"]):F2}");
        }
        Console.WriteLine($"Total: ${total:F2}");
    }
}

/// <summary>
/// AFTER: Separate domain from presentation
/// </summary>
public class OrderItem
{
    private string name;
    private double price;
    private int quantity;

    public OrderItem(string name, double price, int quantity)
    {
        this.name = name;
        this.price = price;
        this.quantity = quantity;
    }

    public string GetName() => name;
    public double GetPrice() => price;
    public int GetQuantity() => quantity;
    public double GetTotal() => price * quantity;
}

public class OrderDomainAfter
{
    private List<OrderItem> items = new List<OrderItem>();

    public void AddItem(string name, double price, int quantity)
    {
        var item = new OrderItem(name, price, quantity);
        items.Add(item);
    }

    public List<OrderItem> GetItems() => items;

    public double GetTotal()
    {
        double total = 0;
        foreach (var item in items)
        {
            total += item.GetTotal();
        }
        return total;
    }
}

public class OrderPresenter
{
    public void DisplayItemAdded(OrderItem item)
    {
        Console.WriteLine($"Added {item.GetQuantity()} x {item.GetName()} to order");
    }

    public void DisplayOrderSummary(OrderDomainAfter order)
    {
        Console.WriteLine("Order Summary:");
        foreach (var item in order.GetItems())
        {
            Console.WriteLine($"- {item.GetQuantity()} x {item.GetName()} @ ${item.GetPrice():F2}");
        }
        Console.WriteLine($"Total: ${order.GetTotal():F2}");
    }
}

public class OrderService
{
    private OrderDomainAfter order;
    private OrderPresenter presenter;

    public OrderService()
    {
        order = new OrderDomainAfter();
        presenter = new OrderPresenter();
    }

    public void AddItem(string name, double price, int quantity)
    {
        var item = new OrderItem(name, price, quantity);
        order.AddItem(name, price, quantity);
        presenter.DisplayItemAdded(item);
        presenter.DisplayOrderSummary(order);
    }

    public OrderDomainAfter GetOrder() => order;
}

/// <summary>
/// 72. Hierarchy Extraction (Extract Hierarchy)
///
/// BEFORE: Single class handling multiple responsibilities
/// </summary>
public class ComputerExtractBefore
{
    private string type;
    private double cpu;
    private double ram;
    private double storage;

    public ComputerExtractBefore(string type, double cpu, double ram, double storage)
    {
        this.type = type;
        this.cpu = cpu;
        this.ram = ram;
        this.storage = storage;
    }

    public string GetSpecs()
    {
        var specs = $"CPU: {cpu}\n";
        specs += $"RAM: {ram}GB\n";
        specs += $"Storage: {storage}GB\n";

        if (type == "desktop")
        {
            specs += "Form Factor: Desktop\n";
        }
        else if (type == "laptop")
        {
            specs += "Form Factor: Laptop\n";
            specs += "Battery Life: 8 hours\n";
        }
        else if (type == "server")
        {
            specs += "Form Factor: Server Rack\n";
            specs += "Redundancy: RAID 10\n";
        }

        return specs;
    }
}

/// <summary>
/// AFTER: Extract hierarchy
/// </summary>
public abstract class ComputerExtractAfter
{
    protected double cpu;
    protected double ram;
    protected double storage;

    public ComputerExtractAfter(double cpu, double ram, double storage)
    {
        this.cpu = cpu;
        this.ram = ram;
        this.storage = storage;
    }

    public abstract string GetFormFactor();
    public abstract string GetSpecialFeatures();

    public string GetBasicSpecs()
    {
        return $"CPU: {cpu}\n" +
               $"RAM: {ram}GB\n" +
               $"Storage: {storage}GB\n";
    }

    public string GetSpecs()
    {
        return GetBasicSpecs() +
               $"Form Factor: {GetFormFactor()}\n" +
               GetSpecialFeatures();
    }
}

public class DesktopComputer : ComputerExtractAfter
{
    public DesktopComputer(double cpu, double ram, double storage) : base(cpu, ram, storage) { }

    public override string GetFormFactor() => "Desktop";

    public override string GetSpecialFeatures() => "Expansion Slots: Multiple PCI\n";
}

public class LaptopComputer : ComputerExtractAfter
{
    public LaptopComputer(double cpu, double ram, double storage) : base(cpu, ram, storage) { }

    public override string GetFormFactor() => "Laptop";

    public override string GetSpecialFeatures() => "Battery Life: 8 hours\nWeight: 2.5 lbs\n";
}

public class ServerComputer : ComputerExtractAfter
{
    public ServerComputer(double cpu, double ram, double storage) : base(cpu, ram, storage) { }

    public override string GetFormFactor() => "Server Rack";

    public override string GetSpecialFeatures() => "Redundancy: RAID 10\nHot Swap Drives: Yes\n";
}
