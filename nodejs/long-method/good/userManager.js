const UserValidator = require('./userValidator');
const UserRepository = require('./userRepository');
const EmailService = require('./emailService');
const { NotificationService, Logger } = require('./notificationService');
const crypto = require('crypto');
const bcrypt = require('bcrypt');

class UserManager {
    constructor(db) {
        this.validator = new UserValidator();
        this.repository = new UserRepository(db);
        this.emailService = new EmailService();
        this.notificationService = new NotificationService(db);
    }

    async registerUser(userData) {
        this.validator.validateRegistrationData(userData);

        if (await this.repository.userExists(userData.email)) {
            throw new Error('User already exists');
        }

        const preparedUserData = this.prepareUserData(userData);

        const userId = await this.repository.createUser(preparedUserData);
        await this.repository.createUserProfile(userId, preparedUserData);
        await this.repository.createUserSettings(userId);

        await this.emailService.sendVerificationEmail(
            preparedUserData.email,
            preparedUserData.firstName,
            preparedUserData.verificationToken
        );

        await this.notificationService.sendWelcomeNotification(userId);
        Logger.logRegistration(preparedUserData.email);

        return userId;
    }

    prepareUserData(userData) {
        const preparedData = { ...userData };
        preparedData.password = bcrypt.hashSync(preparedData.password, 10);
        preparedData.verificationToken = crypto.randomBytes(32).toString('hex');
        return preparedData;
    }
}

module.exports = UserManager;
