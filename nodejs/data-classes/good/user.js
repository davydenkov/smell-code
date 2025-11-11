class EmailValidator {
    isValid(email) {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
    }
}

class User {
    constructor(id, name, email, age, emailValidator = null) {
        this.id = id;
        this.name = name;
        this.age = age;
        this.emailValidator = emailValidator || new EmailValidator();

        this.setEmail(email); // Use setter for validation
    }

    getId() {
        return this.id;
    }

    getName() {
        return this.name;
    }

    getEmail() {
        return this.email;
    }

    setEmail(email) {
        if (!this.emailValidator.isValid(email)) {
            throw new Error('Invalid email address');
        }
        this.email = email;
    }

    getAge() {
        return this.age;
    }

    setAge(age) {
        if (age < 0 || age > 150) {
            throw new Error('Age must be between 0 and 150');
        }
        this.age = age;
    }

    getDisplayName() {
        return `${this.name} (${this.age} years old)`;
    }

    canVote() {
        return this.age >= 18;
    }

    isAdult() {
        return this.age >= 18;
    }

    getAgeCategory() {
        if (this.age < 13) return 'child';
        if (this.age < 20) return 'teenager';
        if (this.age < 65) return 'adult';
        return 'senior';
    }
}

module.exports = {
    User,
    EmailValidator
};
