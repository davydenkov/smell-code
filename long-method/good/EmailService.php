<?php

class EmailService {
    public function sendVerificationEmail(string $email, string $firstName, string $verificationToken): void {
        $subject = 'Please verify your email address';
        $message = "Hello {$firstName},\n\n";
        $message .= "Thank you for registering. Please click the link below to verify your email:\n\n";
        $message .= "http://example.com/verify?token={$verificationToken}\n\n";
        $message .= "Best regards,\nThe Team";

        $headers = 'From: noreply@example.com' . "\r\n" .
                   'Reply-To: noreply@example.com' . "\r\n" .
                   'X-Mailer: PHP/' . phpversion();

        mail($email, $subject, $message, $headers);
    }
}
