const nodemailer = require('nodemailer');

class EmailService {
    constructor(emailConfig) {
        this.transporter = nodemailer.createTransporter({
            host: emailConfig.host,
            port: emailConfig.port,
            auth: {
                user: emailConfig.username,
                pass: emailConfig.password
            }
        });
    }

    async sendWelcomeEmail(userId) {
        // In a real implementation, you'd fetch user data first
        // For this example, we'll assume userId is used to get user data
        const user = await this.getUserById(userId); // You'd implement this

        await this.transporter.sendMail({
            from: process.env.EMAIL_FROM || 'noreply@example.com',
            to: user.email,
            subject: 'Welcome!',
            text: `Welcome ${user.name}!`
        });
    }

    async getUserById(userId) {
        // This would typically use a repository or database call
        // Simplified for the example
        return { id: userId, email: 'user@example.com', name: 'User' };
    }
}

module.exports = EmailService;
