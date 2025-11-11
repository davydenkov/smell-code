<?php

class EmailService {
    private string $smtpHost;
    private int $smtpPort;
    private string $smtpUsername;
    private string $smtpPassword;

    public function __construct(array $emailConfig) {
        $this->smtpHost = $emailConfig['host'];
        $this->smtpPort = $emailConfig['port'];
        $this->smtpUsername = $emailConfig['username'];
        $this->smtpPassword = $emailConfig['password'];
    }

    public function sendWelcomeEmail(string $email, string $name): bool {
        $subject = 'Welcome to our platform!';
        $message = "Hello {$name},\n\nWelcome to our platform!";

        $headers = 'From: ' . $this->smtpUsername . "\r\n";
        $headers .= 'Reply-To: ' . $this->smtpUsername . "\r\n";

        return mail($email, $subject, $message, $headers);
    }

    public function sendPasswordResetEmail(string $email, string $resetToken): bool {
        $subject = 'Password Reset';
        $message = "Click here to reset your password: http://example.com/reset?token={$resetToken}";

        $headers = 'From: ' . $this->smtpUsername . "\r\n";

        return mail($email, $subject, $message, $headers);
    }

    public function sendNotificationEmail(string $email, string $subject, string $message): bool {
        $headers = 'From: ' . $this->smtpUsername . "\r\n";

        return mail($email, $subject, $message, $headers);
    }
}
