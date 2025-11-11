class NotificationService {
    constructor(db) {
        this.db = db;
    }

    async sendWelcomeNotification(userId) {
        await this.db.query(
            'INSERT INTO notifications (user_id, type, message, created_at) VALUES ($1, $2, $3, NOW())',
            [userId, 'welcome', 'Welcome to our platform!']
        );
    }
}

class Logger {
    static logRegistration(email) {
        const logMessage = `User registered: ${email} at ${new Date().toISOString()}`;
        console.log(logMessage);
    }
}

module.exports = {
    NotificationService,
    Logger
};
