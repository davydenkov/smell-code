import java.sql.Connection;

public class NotificationService {
    private Connection db;

    public NotificationService(Connection db) {
        this.db = db;
    }

    public void sendWelcomeNotification(int userId) {
        // Notification logic
        System.out.println("Sending welcome notification to user " + userId);
    }
}
