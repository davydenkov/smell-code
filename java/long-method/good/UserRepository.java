import java.sql.*;
import java.util.Map;

public class UserRepository {
    private Connection db;

    public UserRepository(Connection db) {
        this.db = db;
    }

    public boolean userExists(String email) throws SQLException {
        PreparedStatement stmt = db.prepareStatement("SELECT id FROM users WHERE email = ?");
        stmt.setString(1, email);
        ResultSet rs = stmt.executeQuery();
        return rs.next();
    }

    public int createUser(Map<String, String> userData) throws SQLException {
        PreparedStatement stmt = db.prepareStatement(
            "INSERT INTO users (email, password, first_name, last_name, verification_token, created_at) VALUES (?, ?, ?, ?, ?, NOW())",
            Statement.RETURN_GENERATED_KEYS
        );
        stmt.setString(1, userData.get("email"));
        stmt.setString(2, userData.get("password"));
        stmt.setString(3, userData.get("firstName"));
        stmt.setString(4, userData.get("lastName"));
        stmt.setString(5, userData.get("verificationToken"));
        stmt.executeUpdate();

        ResultSet rs = stmt.getGeneratedKeys();
        rs.next();
        return rs.getInt(1);
    }

    public void createUserProfile(int userId, Map<String, String> userData) throws SQLException {
        PreparedStatement stmt = db.prepareStatement(
            "INSERT INTO user_profiles (user_id, bio, avatar_url) VALUES (?, '', '')"
        );
        stmt.setInt(1, userId);
        stmt.executeUpdate();
    }

    public void createUserSettings(int userId) throws SQLException {
        PreparedStatement stmt = db.prepareStatement(
            "INSERT INTO user_settings (user_id, theme, notifications_enabled) VALUES (?, 'light', true)"
        );
        stmt.setInt(1, userId);
        stmt.executeUpdate();
    }
}
