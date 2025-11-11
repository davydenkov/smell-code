public class EmailService {
    public void sendVerificationEmail(String email, String firstName, String token) {
        // Email sending logic
        System.out.println("Sending verification email to " + email + " for " + firstName + " with token " + token);
    }
}
