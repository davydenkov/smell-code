import java.util.Map;

public class UserManager {
    private UserValidator validator;
    private UserRepository repository;
    private EmailService emailService;
    private NotificationService notificationService;

    public UserManager(java.sql.Connection db) {
        this.validator = new UserValidator();
        this.repository = new UserRepository(db);
        this.emailService = new EmailService();
        this.notificationService = new NotificationService(db);
    }

    public int registerUser(Map<String, String> userData) throws Exception {
        validator.validateRegistrationData(userData);

        if (repository.userExists(userData.get("email"))) {
            throw new Exception("User already exists");
        }

        Map<String, String> preparedData = prepareUserData(userData);

        int userId = repository.createUser(preparedData);
        repository.createUserProfile(userId, preparedData);
        repository.createUserSettings(userId);

        emailService.sendVerificationEmail(
            preparedData.get("email"),
            preparedData.get("firstName"),
            preparedData.get("verificationToken")
        );

        notificationService.sendWelcomeNotification(userId);
        Logger.logRegistration(preparedData.get("email"));

        return userId;
    }

    private Map<String, String> prepareUserData(Map<String, String> userData) {
        Map<String, String> prepared = new java.util.HashMap<>(userData);
        prepared.put("password", hashPassword(userData.get("password")));
        prepared.put("verificationToken", generateVerificationToken());
        return prepared;
    }

    private String hashPassword(String password) {
        // Simplified password hashing
        return "hashed_" + password;
    }

    private String generateVerificationToken() {
        return "token_" + System.currentTimeMillis();
    }
}

class Logger {
    public static void logRegistration(String email) {
        System.out.println("User registered: " + email);
    }
}
