// GOOD: Focused service that only handles user management
class UserService {
    constructor(userRepository, emailService) {
        this.userRepository = userRepository;
        this.emailService = emailService;
    }

    async createUser(userData) {
        const userId = await this.userRepository.createUser(userData);
        await this.emailService.sendWelcomeEmail(userId);
        return userId;
    }

    async getUser(userId) {
        return this.userRepository.getUser(userId);
    }
}

module.exports = UserService;
