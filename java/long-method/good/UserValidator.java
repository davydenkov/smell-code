import java.util.Map;

public class UserValidator {
    public void validateRegistrationData(Map<String, String> userData) throws Exception {
        if (userData.get("email") == null || userData.get("email").isEmpty()) {
            throw new Exception("Email is required");
        }
        if (userData.get("password") == null || userData.get("password").isEmpty()) {
            throw new Exception("Password is required");
        }
        if (userData.get("password").length() < 8) {
            throw new Exception("Password must be at least 8 characters");
        }
        if (!isValidEmail(userData.get("email"))) {
            throw new Exception("Invalid email format");
        }
    }

    private boolean isValidEmail(String email) {
        return email != null && email.contains("@");
    }
}
