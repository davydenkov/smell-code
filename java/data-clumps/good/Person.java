public class Person {
    private String firstName;
    private String lastName;
    private String email;
    private String phone;
    private String dateOfBirth;

    public Person(String firstName, String lastName, String email) {
        this(firstName, lastName, email, null, null);
    }

    public Person(String firstName, String lastName, String email, String phone, String dateOfBirth) {
        this.firstName = firstName;
        this.lastName = lastName;
        this.email = email;
        this.phone = phone;
        this.dateOfBirth = dateOfBirth;
    }

    public String getFirstName() {
        return firstName;
    }

    public String getLastName() {
        return lastName;
    }

    public String getFullName() {
        return firstName + " " + lastName;
    }

    public String getEmail() {
        return email;
    }

    public String getPhone() {
        return phone;
    }

    public String getDateOfBirth() {
        return dateOfBirth;
    }

    public boolean isValid() {
        if (firstName == null || firstName.isEmpty() ||
            lastName == null || lastName.isEmpty()) {
            return false;
        }
        if (!isValidEmail(email)) {
            return false;
        }
        return true;
    }

    private boolean isValidEmail(String email) {
        // Simple email validation
        return email != null && email.contains("@") && email.contains(".");
    }
}
