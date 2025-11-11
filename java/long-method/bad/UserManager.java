import java.sql.*;
import java.util.Map;

public class UserManager {
    private Connection db;

    public UserManager(Connection db) {
        this.db = db;
    }

    public void registerUser(Map<String, String> userData) throws Exception {
        // Validate input data
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

        // Check if user already exists
        PreparedStatement stmt = db.prepareStatement("SELECT id FROM users WHERE email = ?");
        stmt.setString(1, userData.get("email"));
        ResultSet rs = stmt.executeQuery();
        if (rs.next()) {
            throw new Exception("User already exists");
        }

        // Hash password
        String hashedPassword = hashPassword(userData.get("password"));

        // Generate verification token
        String verificationToken = generateVerificationToken();

        // Insert user into database
        stmt = db.prepareStatement(
            "INSERT INTO users (email, password, first_name, last_name, verification_token, created_at) VALUES (?, ?, ?, ?, ?, NOW())"
        );
        stmt.setString(1, userData.get("email"));
        stmt.setString(2, hashedPassword);
        stmt.setString(3, userData.get("firstName"));
        stmt.setString(4, userData.get("lastName"));
        stmt.setString(5, verificationToken);
        stmt.executeUpdate();

        // Get user ID
        Statement idStmt = db.createStatement();
        ResultSet idRs = idStmt.executeQuery("SELECT LAST_INSERT_ID()");
        idRs.next();
        int userId = idRs.getInt(1);

        // Create user profile
        stmt = db.prepareStatement(
            "INSERT INTO user_profiles (user_id, bio, avatar_url) VALUES (?, '', '')"
        );
        stmt.setInt(1, userId);
        stmt.executeUpdate();

        // Create user settings
        stmt = db.prepareStatement(
            "INSERT INTO user_settings (user_id, theme, notifications_enabled) VALUES (?, 'light', true)"
        );
        stmt.setInt(1, userId);
        stmt.executeUpdate();

        // Send verification email
        sendVerificationEmail(userData.get("email"), userData.get("firstName"), verificationToken);

        // Send welcome notification
        sendWelcomeNotification(userId);

        // Log registration
        logRegistration(userData.get("email"));
    }

    private boolean isValidEmail(String email) {
        return email != null && email.contains("@");
    }

    private String hashPassword(String password) {
        // Simplified password hashing
        return "hashed_" + password;
    }

    private String generateVerificationToken() {
        return "token_" + System.currentTimeMillis();
    }

    private void sendVerificationEmail(String email, String firstName, String token) {
        // Email sending logic
        System.out.println("Sending verification email to " + email);
    }

    private void sendWelcomeNotification(int userId) {
        // Notification logic
        System.out.println("Sending welcome notification to user " + userId);
    }

    private void logRegistration(String email) {
        // Logging logic
        System.out.println("User registered: " + email);
    }
}
