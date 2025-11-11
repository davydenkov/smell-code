using System;
using System.Collections.Generic;
using System.Linq;

/// <summary>
/// 34. Decomposition of a conditional operator (Decompose Conditional)
///
/// BEFORE: Complex conditional logic
/// </summary>
public class PaymentProcessorBefore
{
    public double CalculateFee(double amount, bool isInternational, bool isPremium)
    {
        double fee = 0;
        if (amount > 100 && isInternational && isPremium)
        {
            fee = amount * 0.05 + 10;
        }
        else if (amount > 100 && isInternational && !isPremium)
        {
            fee = amount * 0.05 + 15;
        }
        else if (amount <= 100 && isInternational)
        {
            fee = amount * 0.03 + 5;
        }
        else
        {
            fee = amount * 0.02;
        }
        return fee;
    }
}

/// <summary>
/// AFTER: Decompose conditional
/// </summary>
public class PaymentProcessorAfter
{
    public double CalculateFee(double amount, bool isInternational, bool isPremium)
    {
        if (IsHighValueInternationalPremium(amount, isInternational, isPremium))
        {
            return CalculateHighValueInternationalPremiumFee(amount);
        }
        else if (IsHighValueInternationalStandard(amount, isInternational, isPremium))
        {
            return CalculateHighValueInternationalStandardFee(amount);
        }
        else if (IsLowValueInternational(amount, isInternational))
        {
            return CalculateLowValueInternationalFee(amount);
        }
        else
        {
            return CalculateDomesticFee(amount);
        }
    }

    private bool IsHighValueInternationalPremium(double amount, bool isInternational, bool isPremium)
    {
        return amount > 100 && isInternational && isPremium;
    }

    private bool IsHighValueInternationalStandard(double amount, bool isInternational, bool isPremium)
    {
        return amount > 100 && isInternational && !isPremium;
    }

    private bool IsLowValueInternational(double amount, bool isInternational)
    {
        return amount <= 100 && isInternational;
    }

    private double CalculateHighValueInternationalPremiumFee(double amount)
    {
        return amount * 0.05 + 10;
    }

    private double CalculateHighValueInternationalStandardFee(double amount)
    {
        return amount * 0.05 + 15;
    }

    private double CalculateLowValueInternationalFee(double amount)
    {
        return amount * 0.03 + 5;
    }

    private double CalculateDomesticFee(double amount)
    {
        return amount * 0.02;
    }
}

/// <summary>
/// 35. Consolidation of a conditional expression (Consolidate Conditional Expression)
///
/// BEFORE: Multiple conditionals with same result
/// </summary>
public class InsuranceCalculatorBefore
{
    public bool IsEligibleForDiscount(int age, bool isStudent, bool hasGoodRecord)
    {
        if (age < 25) return false;
        if (isStudent) return true;
        if (hasGoodRecord) return true;
        return false;
    }
}

/// <summary>
/// AFTER: Consolidate conditionals
/// </summary>
public class InsuranceCalculatorAfter
{
    public bool IsEligibleForDiscount(int age, bool isStudent, bool hasGoodRecord)
    {
        return age >= 25 && (isStudent || hasGoodRecord);
    }
}

/// <summary>
/// 36. Consolidation of duplicate conditional fragments
/// (Consolidate Duplicate Conditional Fragments)
///
/// BEFORE: Duplicate code in conditional branches
/// </summary>
public class FileProcessorBefore
{
    public void ProcessFile(object file)
    {
        if (IsValidFile(file))
        {
            LogProcessing(file);
            ValidateContent(file);
            SaveToDatabase(file);
            SendNotification(file);
        }
        else
        {
            LogError(file);
            SendNotification(file); // Duplicate
        }
    }

    private void SendNotification(object file)
    {
        // Send notification logic
    }

    private bool IsValidFile(object file) => true;
    private void LogProcessing(object file) { }
    private void ValidateContent(object file) { }
    private void SaveToDatabase(object file) { }
    private void LogError(object file) { }
}

/// <summary>
/// AFTER: Consolidate duplicate fragments
/// </summary>
public class FileProcessorAfter
{
    public void ProcessFile(object file)
    {
        SendNotification(file); // Moved outside conditional

        if (IsValidFile(file))
        {
            LogProcessing(file);
            ValidateContent(file);
            SaveToDatabase(file);
        }
        else
        {
            LogError(file);
        }
    }

    private void SendNotification(object file)
    {
        // Send notification logic
    }

    private bool IsValidFile(object file) => true;
    private void LogProcessing(object file) { }
    private void ValidateContent(object file) { }
    private void SaveToDatabase(object file) { }
    private void LogError(object file) { }
}

/// <summary>
/// 37. Remove Control Flag
///
/// BEFORE: Control flag to break out of loop
/// </summary>
public class DataProcessorBefore
{
    public object FindPerson(List<Dictionary<string, object>> people, string name)
    {
        object found = null;
        foreach (var person in people)
        {
            if ((string)person["name"] == name)
            {
                found = person;
                break; // Control flag usage
            }
        }
        return found;
    }
}

/// <summary>
/// AFTER: Remove control flag
/// </summary>
public class DataProcessorAfter
{
    public Dictionary<string, object> FindPerson(List<Dictionary<string, object>> people, string name)
    {
        foreach (var person in people)
        {
            if ((string)person["name"] == name)
            {
                return person; // Direct return
            }
        }
        return null;
    }
}

/// <summary>
/// 38. Replacing Nested Conditional statements with a boundary operator
/// (Replace Nested Conditional with Guard Clauses)
///
/// BEFORE: Nested conditionals
/// </summary>
public class PaymentValidatorBefore
{
    public bool IsValidPayment(Dictionary<string, object> payment)
    {
        if ((double)payment["amount"] > 0)
        {
            if (payment["cardNumber"] != null)
            {
                if (((string)payment["cardNumber"]).Length == 16)
                {
                    if (IsValidExpiry((string)payment["expiry"]))
                    {
                        return true;
                    }
                }
            }
        }
        return false;
    }

    private bool IsValidExpiry(string expiry)
    {
        return DateTime.Parse(expiry) > DateTime.Now;
    }
}

/// <summary>
/// AFTER: Replace with guard clauses
/// </summary>
public class PaymentValidatorAfter
{
    public bool IsValidPayment(Dictionary<string, object> payment)
    {
        if ((double)payment["amount"] <= 0)
        {
            return false;
        }

        if (payment["cardNumber"] == null)
        {
            return false;
        }

        if (((string)payment["cardNumber"]).Length != 16)
        {
            return false;
        }

        if (!IsValidExpiry((string)payment["expiry"]))
        {
            return false;
        }

        return true;
    }

    private bool IsValidExpiry(string expiry)
    {
        return DateTime.Parse(expiry) > DateTime.Now;
    }
}

/// <summary>
/// 39. Replacing a conditional operator with polymorphism (Replace Conditional with Polymorphism)
///
/// BEFORE: Type checking with conditionals
/// </summary>
public class BirdBefore
{
    public const string EUROPEAN = "european";
    public const string AFRICAN = "african";
    public const string NORWEGIAN_BLUE = "norwegian_blue";

    private string type;
    private int voltage;
    private bool isNailed;

    public BirdBefore(string type)
    {
        this.type = type;
    }

    public double GetSpeed()
    {
        switch (type)
        {
            case EUROPEAN:
                return GetBaseSpeed();
            case AFRICAN:
                return GetBaseSpeed() - voltage * 2;
            case NORWEGIAN_BLUE:
                return isNailed ? 0 : GetBaseSpeed();
            default:
                return GetBaseSpeed();
        }
    }

    private double GetBaseSpeed()
    {
        return 10;
    }
}

/// <summary>
/// AFTER: Replace conditional with polymorphism
/// </summary>
public abstract class BirdAfter
{
    public abstract double GetSpeed();

    protected double GetBaseSpeed()
    {
        return 10;
    }
}

public class EuropeanSwallow : BirdAfter
{
    public override double GetSpeed()
    {
        return GetBaseSpeed();
    }
}

public class AfricanSwallow : BirdAfter
{
    private int voltage;

    public AfricanSwallow(int voltage)
    {
        this.voltage = voltage;
    }

    public override double GetSpeed()
    {
        return GetBaseSpeed() - voltage * 2;
    }
}

public class NorwegianBlueParrot : BirdAfter
{
    private bool isNailed;

    public NorwegianBlueParrot(bool isNailed)
    {
        this.isNailed = isNailed;
    }

    public override double GetSpeed()
    {
        return isNailed ? 0 : GetBaseSpeed();
    }
}

/// <summary>
/// 40. Introduction of the object (Introduce Object)
///
/// BEFORE: Primitive obsession with conditionals
/// </summary>
public class UserValidatorBefore
{
    public object ValidateUser(Dictionary<string, object> user)
    {
        if (string.IsNullOrEmpty((string)user["name"]))
        {
            return "Name is required";
        }

        if (((string)user["name"]).Length < 2)
        {
            return "Name must be at least 2 characters";
        }

        if (!IsValidEmail((string)user["email"]))
        {
            return "Invalid email format";
        }

        return true;
    }

    private bool IsValidEmail(string email)
    {
        try
        {
            var addr = new System.Net.Mail.MailAddress(email);
            return addr.Address == email;
        }
        catch
        {
            return false;
        }
    }
}

/// <summary>
/// AFTER: Introduce validation result object
/// </summary>
public class ValidationResult
{
    private bool isValid;
    private List<string> errors;

    public ValidationResult(bool isValid = true, List<string> errors = null)
    {
        this.isValid = isValid;
        this.errors = errors ?? new List<string>();
    }

    public bool IsValid()
    {
        return isValid;
    }

    public List<string> GetErrors()
    {
        return new List<string>(errors);
    }

    public ValidationResult AddError(string error)
    {
        isValid = false;
        errors.Add(error);
        return this;
    }
}

public class UserValidatorAfter
{
    public ValidationResult ValidateUser(Dictionary<string, object> user)
    {
        var result = new ValidationResult();

        if (string.IsNullOrEmpty((string)user["name"]))
        {
            result.AddError("Name is required");
        }

        if (((string)user["name"]).Length < 2)
        {
            result.AddError("Name must be at least 2 characters");
        }

        if (!IsValidEmail((string)user["email"]))
        {
            result.AddError("Invalid email format");
        }

        return result;
    }

    private bool IsValidEmail(string email)
    {
        try
        {
            var addr = new System.Net.Mail.MailAddress(email);
            return addr.Address == email;
        }
        catch
        {
            return false;
        }
    }
}

/// <summary>
/// 41. Introduction of the statement (Introduction Statement)
///
/// BEFORE: Magic assertion
/// </summary>
public class AccountAssertion
{
    private double balance;

    public void Withdraw(double amount)
    {
        System.Diagnostics.Debug.Assert(amount > 0 && amount <= balance);
        balance -= amount;
    }
}

/// <summary>
/// AFTER: Introduce assertion method
/// </summary>
public class AccountAssertionAfter
{
    private double balance;

    public void Withdraw(double amount)
    {
        AssertValidWithdrawal(amount);
        balance -= amount;
    }

    private void AssertValidWithdrawal(double amount)
    {
        System.Diagnostics.Debug.Assert(amount > 0 && amount <= balance, "Invalid withdrawal amount");
    }
}
