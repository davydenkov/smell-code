class UserValidator {
    validateRegistrationData(userData) {
        this.validateRequiredFields(userData);
        this.validateEmailFormat(userData.email);
        this.validatePasswordStrength(userData.password);
    }

    validateRequiredFields(userData) {
        if (!userData.email) {
            throw new Error('Email is required');
        }
        if (!userData.password) {
            throw new Error('Password is required');
        }
    }

    validateEmailFormat(email) {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(email)) {
            throw new Error('Invalid email format');
        }
    }

    validatePasswordStrength(password) {
        if (password.length < 8) {
            throw new Error('Password must be at least 8 characters');
        }
    }
}

module.exports = UserValidator;
