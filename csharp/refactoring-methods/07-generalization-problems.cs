using System;
using System.Collections.Generic;
using System.Linq;

/// <summary>
/// 57. Lifting the field (Pull Up Field)
///
/// BEFORE: Duplicate fields in subclasses
/// </summary>
public class EmployeePullBefore
{
    protected string name;
}

public class ManagerPullBefore : EmployeePullBefore
{
    protected string name; // Duplicate
    private double budget;
}

public class EngineerPullBefore : EmployeePullBefore
{
    protected string name; // Duplicate
    private List<string> skills;
}

/// <summary>
/// AFTER: Pull up field to superclass
/// </summary>
public class EmployeePullAfter
{
    protected string name;
}

public class ManagerPullAfter : EmployeePullAfter
{
    private double budget;
}

public class EngineerPullAfter : EmployeePullAfter
{
    private List<string> skills;
}

/// <summary>
/// 58. Lifting the method (Pull Up Method)
///
/// BEFORE: Duplicate methods in subclasses
/// </summary>
public abstract class ShapeMethodBefore
{
    public abstract double Area();
}

public class CircleMethodBefore : ShapeMethodBefore
{
    private double radius;

    public override double Area()
    {
        return Math.PI * radius * radius;
    }

    public double Circumference()
    {
        return 2 * Math.PI * radius;
    }
}

public class SquareMethodBefore : ShapeMethodBefore
{
    private double side;

    public override double Area()
    {
        return side * side;
    }

    public double Circumference() // Duplicate logic
    {
        return 4 * side;
    }
}

/// <summary>
/// AFTER: Pull up method to superclass
/// </summary>
public abstract class ShapeMethodAfter
{
    public abstract double Area();
    public abstract double Circumference();
}

public class CircleMethodAfter : ShapeMethodAfter
{
    private double radius;

    public override double Area()
    {
        return Math.PI * radius * radius;
    }

    public override double Circumference()
    {
        return 2 * Math.PI * radius;
    }
}

public class SquareMethodAfter : ShapeMethodAfter
{
    private double side;

    public override double Area()
    {
        return side * side;
    }

    public override double Circumference()
    {
        return 4 * side;
    }
}

/// <summary>
/// 59. Lifting the constructor Body (Pull Up Constructor Body)
///
/// BEFORE: Duplicate constructor code
/// </summary>
public class VehicleConstructorBefore
{
    protected string make;
    protected string model;
    protected int year;
}

public class CarConstructorBefore : VehicleConstructorBefore
{
    private int doors;

    public CarConstructorBefore(string make, string model, int year, int doors)
    {
        this.make = make; // Duplicate
        this.model = model; // Duplicate
        this.year = year; // Duplicate
        this.doors = doors;
    }
}

public class TruckConstructorBefore : VehicleConstructorBefore
{
    private double payload;

    public TruckConstructorBefore(string make, string model, int year, double payload)
    {
        this.make = make; // Duplicate
        this.model = model; // Duplicate
        this.year = year; // Duplicate
        this.payload = payload;
    }
}

/// <summary>
/// AFTER: Pull up constructor body
/// </summary>
public class VehicleConstructorAfter
{
    protected string make;
    protected string model;
    protected int year;

    public VehicleConstructorAfter(string make, string model, int year)
    {
        this.make = make;
        this.model = model;
        this.year = year;
    }
}

public class CarConstructorAfter : VehicleConstructorAfter
{
    private int doors;

    public CarConstructorAfter(string make, string model, int year, int doors)
        : base(make, model, year)
    {
        this.doors = doors;
    }
}

public class TruckConstructorAfter : VehicleConstructorAfter
{
    private double payload;

    public TruckConstructorAfter(string make, string model, int year, double payload)
        : base(make, model, year)
    {
        this.payload = payload;
    }
}

/// <summary>
/// 60. Method Descent (Push Down Method)
///
/// BEFORE: Method in wrong class hierarchy level
/// </summary>
public class AnimalPushBefore
{
    public virtual string Speak()
    {
        // Generic implementation
        return "";
    }
}

public class DogPushBefore : AnimalPushBefore
{
    public override string Speak()
    {
        return "Woof";
    }
}

public class CatPushBefore : AnimalPushBefore
{
    public override string Speak()
    {
        return "Meow";
    }
}

public class FishPushBefore : AnimalPushBefore
{
    // Fish don't speak, but inherits speak method
}

/// <summary>
/// AFTER: Push down method to appropriate subclasses
/// </summary>
public class AnimalPushAfter
{
    // No speak method here
}

public class DogPushAfter : AnimalPushAfter
{
    public string Speak()
    {
        return "Woof";
    }
}

public class CatPushAfter : AnimalPushAfter
{
    public string Speak()
    {
        return "Meow";
    }
}

public class FishPushAfter : AnimalPushAfter
{
    // No speak method - appropriate for Fish
}

/// <summary>
/// 61. Field Descent (Push Down Field)
///
/// BEFORE: Field in wrong hierarchy level
/// </summary>
public class EmployeeFieldBefore
{
    protected double salary; // Not all employees have salary
}

public class SalariedEmployeeFieldBefore : EmployeeFieldBefore
{
    // Uses salary
}

public class ContractorFieldBefore : EmployeeFieldBefore
{
    // Doesn't use salary, but inherits it
}

/// <summary>
/// AFTER: Push down field
/// </summary>
public class EmployeeFieldAfter
{
    // No salary field
}

public class SalariedEmployeeFieldAfter : EmployeeFieldAfter
{
    protected double salary;
}

public class ContractorFieldAfter : EmployeeFieldAfter
{
    protected double hourlyRate;
}

/// <summary>
/// 62. Subclass extraction (Extract Subclass)
///
/// BEFORE: Class with conditional behavior
/// </summary>
public class JobExtractBefore
{
    private string type;
    private double rate;
    private double? commission;

    public JobExtractBefore(string type, double rate, double? commission = null)
    {
        this.type = type;
        this.rate = rate;
        this.commission = commission;
    }

    public double GetPay()
    {
        if (type == "salaried")
        {
            return rate;
        }
        else
        {
            return rate + (commission ?? 0);
        }
    }
}

/// <summary>
/// AFTER: Extract subclass
/// </summary>
public abstract class JobExtractAfter
{
    protected double rate;

    public JobExtractAfter(double rate)
    {
        this.rate = rate;
    }

    public abstract double GetPay();
}

public class SalariedJob : JobExtractAfter
{
    public SalariedJob(double rate) : base(rate) { }

    public override double GetPay()
    {
        return rate;
    }
}

public class CommissionedJob : JobExtractAfter
{
    private double commission;

    public CommissionedJob(double rate, double commission) : base(rate)
    {
        this.commission = commission;
    }

    public override double GetPay()
    {
        return rate + commission;
    }
}

/// <summary>
/// 63. Allocation of the parent class (Extract Superclass)
///
/// BEFORE: Duplicate code in classes
/// </summary>
public class DepartmentSuperBefore
{
    private string name;
    private string head;

    public DepartmentSuperBefore(string name, string head)
    {
        this.name = name;
        this.head = head;
    }

    public string GetName()
    {
        return name;
    }

    public string GetHead()
    {
        return head;
    }
}

public class CompanySuperBefore
{
    private string name;
    private string head;

    public CompanySuperBefore(string name, string head)
    {
        this.name = name;
        this.head = head;
    }

    public string GetName()
    {
        return name;
    }

    public string GetHead()
    {
        return head;
    }
}

/// <summary>
/// AFTER: Extract superclass
/// </summary>
public abstract class Party
{
    private string name;
    private string head;

    public Party(string name, string head)
    {
        this.name = name;
        this.head = head;
    }

    public string GetName()
    {
        return name;
    }

    public string GetHead()
    {
        return head;
    }
}

public class DepartmentSuperAfter : Party
{
    public DepartmentSuperAfter(string name, string head) : base(name, head) { }
}

public class CompanySuperAfter : Party
{
    public CompanySuperAfter(string name, string head) : base(name, head) { }
}

/// <summary>
/// 64. Interface extraction (Extract Interface)
///
/// BEFORE: Clients depend on concrete class
/// </summary>
public class PrinterInterfaceBefore
{
    public void Print(object document)
    {
        // Print logic
    }

    public string GetStatus()
    {
        // Status logic
        return "Ready";
    }

    public void CancelJob(int jobId)
    {
        // Cancel logic
    }
}

/// <summary>
/// AFTER: Extract interface
/// </summary>
public interface IPrinter
{
    void Print(object document);
    string GetStatus();
}

public class LaserPrinter : IPrinter
{
    public void Print(object document)
    {
        // Print logic
    }

    public string GetStatus()
    {
        // Status logic
        return "Ready";
    }

    public void CancelJob(int jobId)
    {
        // Cancel logic - not part of interface
    }
}

public class InkjetPrinter : IPrinter
{
    public void Print(object document)
    {
        // Print logic
    }

    public string GetStatus()
    {
        // Status logic
        return "Ready";
    }
}

/// <summary>
/// 65. Collapse Hierarchy
///
/// BEFORE: Unnecessary class hierarchy
/// </summary>
public class EmployeeCollapseBefore
{
}

public class ManagerCollapseBefore : EmployeeCollapseBefore
{
    private string department;
}

/// <summary>
/// AFTER: Collapse hierarchy if only one subclass
/// </summary>
public class EmployeeCollapseAfter
{
    private string department; // Moved up
}

/// <summary>
/// 66. Formation of the method template (Form Template Method)
///
/// BEFORE: Duplicate algorithm structure
/// </summary>
public class ReportGeneratorTemplateBefore
{
    public string GenerateHTMLReport()
    {
        var data = GetData();
        var header = FormatHeader();
        var body = FormatBody(data);
        var footer = FormatFooter();
        return header + body + footer;
    }

    public string GeneratePDFReport()
    {
        var data = GetData(); // Duplicate
        var header = FormatPDFHeader(); // Different
        var body = FormatPDFBody(data); // Different
        var footer = FormatPDFFooter(); // Different
        return header + body + footer;
    }

    protected virtual List<string> GetData()
    {
        return new List<string> { "item1", "item2" };
    }

    protected virtual string FormatHeader()
    {
        return "<h1>Report</h1>";
    }

    protected virtual string FormatBody(List<string> data)
    {
        return "<body>" + string.Join("", data) + "</body>";
    }

    protected virtual string FormatFooter()
    {
        return "<footer>End</footer>";
    }

    protected virtual string FormatPDFHeader()
    {
        return "PDF Report Header";
    }

    protected virtual string FormatPDFBody(List<string> data)
    {
        return "PDF Body: " + string.Join("", data);
    }

    protected virtual string FormatPDFFooter()
    {
        return "PDF Footer";
    }
}

/// <summary>
/// AFTER: Form template method
/// </summary>
public abstract class ReportGeneratorTemplateAfter
{
    public string GenerateReport()
    {
        var data = GetData();
        var header = FormatHeader();
        var body = FormatBody(data);
        var footer = FormatFooter();
        return AssembleReport(header, body, footer);
    }

    protected virtual List<string> GetData()
    {
        return new List<string> { "item1", "item2" };
    }

    protected abstract string FormatHeader();
    protected abstract string FormatBody(List<string> data);
    protected abstract string FormatFooter();
    protected abstract string AssembleReport(string header, string body, string footer);
}

public class HTMLReportGenerator : ReportGeneratorTemplateAfter
{
    protected override string FormatHeader()
    {
        return "<h1>Report</h1>";
    }

    protected override string FormatBody(List<string> data)
    {
        return "<body>" + string.Join("", data) + "</body>";
    }

    protected override string FormatFooter()
    {
        return "<footer>End</footer>";
    }

    protected override string AssembleReport(string header, string body, string footer)
    {
        return header + body + footer;
    }
}

public class PDFReportGenerator : ReportGeneratorTemplateAfter
{
    protected override string FormatHeader()
    {
        return "PDF Report Header";
    }

    protected override string FormatBody(List<string> data)
    {
        return "PDF Body: " + string.Join("", data);
    }

    protected override string FormatFooter()
    {
        return "PDF Footer";
    }

    protected override string AssembleReport(string header, string body, string footer)
    {
        return header + body + footer;
    }
}

/// <summary>
/// 67. Replacement of inheritance by delegation (Replace Inheritance with Delegation)
///
/// BEFORE: Inheritance where delegation would be better
/// </summary>
public class StackInheritanceBefore : List<object>
{
    public void Push(object item)
    {
        Add(item);
    }

    public object Pop()
    {
        if (Count == 0)
        {
            throw new InvalidOperationException("Stack is empty");
        }
        var lastIndex = Count - 1;
        var item = this[lastIndex];
        RemoveAt(lastIndex);
        return item;
    }
}

/// <summary>
/// AFTER: Replace inheritance with delegation
/// </summary>
public class StackDelegationAfter
{
    private List<object> items;

    public StackDelegationAfter()
    {
        items = new List<object>();
    }

    public void Push(object item)
    {
        items.Add(item);
    }

    public object Pop()
    {
        if (items.Count == 0)
        {
            throw new InvalidOperationException("Stack is empty");
        }
        var lastIndex = items.Count - 1;
        var item = items[lastIndex];
        items.RemoveAt(lastIndex);
        return item;
    }

    public int Count => items.Count;
}

/// <summary>
/// 68. Replacement of delegation by inheritance (Replace Delegation with Inheritance)
///
/// BEFORE: Delegation where inheritance would be simpler
/// </summary>
public class MyStringDelegateBefore
{
    private string str;

    public MyStringDelegateBefore(string str)
    {
        this.str = str;
    }

    public int Length()
    {
        return str.Length;
    }

    public string Substr(int start, int? length = null)
    {
        if (length.HasValue)
        {
            return str.Substring(start, length.Value);
        }
        return str.Substring(start);
    }

    public int IndexOf(string needle)
    {
        return str.IndexOf(needle);
    }
}

/// <summary>
/// AFTER: Replace delegation with inheritance
/// </summary>
public class MyStringInheritAfter : List<char>
{
    public MyStringInheritAfter(string str) : base(str.ToCharArray())
    {
    }

    public string ToString()
    {
        return new string(this.ToArray());
    }
}
