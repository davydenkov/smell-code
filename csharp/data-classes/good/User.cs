using System;

public class User
{
    private int id;
    private string name;
    private string email;
    private int age;
    private EmailValidator emailValidator;

    public User(int id, string name, string email, int age, EmailValidator emailValidator = null)
    {
        this.id = id;
        this.name = name;
        this.age = age;
        this.emailValidator = emailValidator ?? new EmailValidator();

        SetEmail(email); // Use setter for validation
    }

    public int GetId()
    {
        return id;
    }

    public string GetName()
    {
        return name;
    }

    public string GetEmail()
    {
        return email;
    }

    public void SetEmail(string email)
    {
        if (!emailValidator.IsValid(email))
        {
            throw new ArgumentException("Invalid email address");
        }
        this.email = email;
    }

    public int GetAge()
    {
        return age;
    }

    public void SetAge(int age)
    {
        if (age < 0 || age > 150)
        {
            throw new ArgumentException("Age must be between 0 and 150");
        }
        this.age = age;
    }

    public string GetDisplayName()
    {
        return $"{name} ({age} years old)";
    }

    public bool CanVote()
    {
        return age >= 18;
    }

    public bool IsAdult()
    {
        return age >= 18;
    }

    public string GetAgeCategory()
    {
        if (age < 13) return "child";
        if (age < 20) return "teenager";
        if (age < 65) return "adult";
        return "senior";
    }
}

public class EmailValidator
{
    public bool IsValid(string email)
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
