const bcrypt = require('bcrypt');

class UserRepository {
    constructor(db) {
        this.db = db;
    }

    async createUser(userData) {
        // Validate user data
        this.validateUserData(userData);

        // Hash password
        userData.password = await bcrypt.hash(userData.password, 10);

        // Insert user
        const result = await this.db.query(
            'INSERT INTO users (email, password, name) VALUES ($1, $2, $3) RETURNING id',
            [userData.email, userData.password, userData.name]
        );

        return result.rows[0].id;
    }

    validateUserData(userData) {
        if (!userData.email || !userData.password || !userData.name) {
            throw new Error('Missing required fields');
        }
    }

    async getUser(userId) {
        const result = await this.db.query('SELECT * FROM users WHERE id = $1', [userId]);
        return result.rows[0];
    }
}

module.exports = UserRepository;
