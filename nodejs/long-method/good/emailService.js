const nodemailer = require('nodemailer');

class EmailService {
    constructor() {
        // In a real application, you would configure this properly
        this.transporter = nodemailer.createTransporter({
            service: 'gmail', // or your email service
            auth: {
                user: process.env.EMAIL_USER || 'noreply@example.com',
                pass: process.env.EMAIL_PASS || 'password'
            }
        });
    }

    async sendVerificationEmail(email, firstName, verificationToken) {
        const subject = 'Please verify your email address';
        const message = `Hello ${firstName},\n\nThank you for registering. Please click the link below to verify your email:\n\nhttp://example.com/verify?token=${verificationToken}\n\nBest regards,\nThe Team`;

        await this.transporter.sendMail({
            from: 'noreply@example.com',
            to: email,
            subject: subject,
            text: message
        });
    }
}

module.exports = EmailService;
