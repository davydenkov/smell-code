import java.util.regex.Pattern;

public class User {
    private int id;
    private String name;
    private String email;
    private int age;
    private EmailValidator emailValidator;

    public User(int id, String name, String email, int age, EmailValidator emailValidator) {
        this.id = id;
        this.name = name;
        this.age = age;
        this.emailValidator = emailValidator != null ? emailValidator : new EmailValidator();

        setEmail(email); // Use setter for validation
    }

    public User(int id, String name, String email, int age) {
        this(id, name, email, age, null);
    }

    public int getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        if (!emailValidator.isValid(email)) {
            throw new IllegalArgumentException("Invalid email address");
        }
        this.email = email;
    }

    public int getAge() {
        return age;
    }

    public void setAge(int age) {
        if (age < 0 || age > 150) {
            throw new IllegalArgumentException("Age must be between 0 and 150");
        }
        this.age = age;
    }

    public String getDisplayName() {
        return name + " (" + age + " years old)";
    }

    public boolean canVote() {
        return age >= 18;
    }

    public boolean isAdult() {
        return age >= 18;
    }

    public String getAgeCategory() {
        if (age < 13) return "child";
        if (age < 20) return "teenager";
        if (age < 65) return "adult";
        return "senior";
    }
}

class EmailValidator {
    private static final Pattern EMAIL_PATTERN =
        Pattern.compile("^[A-Za-z0-9+_.-]+@[A-Za-z0-9.-]+$");

    public boolean isValid(String email) {
        if (email == null) return false;
        return EMAIL_PATTERN.matcher(email).matches();
    }
}
