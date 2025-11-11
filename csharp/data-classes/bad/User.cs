public class User
{
    private int id;
    private string name;
    private string email;
    private int age;

    public User(int id, string name, string email, int age)
    {
        this.id = id;
        this.name = name;
        this.email = email;
        this.age = age;
    }

    public int GetId()
    {
        return id;
    }

    public void SetId(int id)
    {
        this.id = id;
    }

    public string GetName()
    {
        return name;
    }

    public void SetName(string name)
    {
        this.name = name;
    }

    public string GetEmail()
    {
        return email;
    }

    public void SetEmail(string email)
    {
        this.email = email;
    }

    public int GetAge()
    {
        return age;
    }

    public void SetAge(int age)
    {
        this.age = age;
    }

    // No behavior, just data and getters/setters - this is a data class smell!
}
