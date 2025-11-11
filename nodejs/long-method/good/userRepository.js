class UserRepository {
    constructor(db) {
        this.db = db;
    }

    async userExists(email) {
        const result = await this.db.query('SELECT id FROM users WHERE email = $1', [email]);
        return result.rows.length > 0;
    }

    async createUser(userData) {
        const result = await this.db.query(
            'INSERT INTO users (email, password, first_name, last_name, verification_token, created_at) VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING id',
            [userData.email, userData.password, userData.firstName, userData.lastName, userData.verificationToken]
        );
        return result.rows[0].id;
    }

    async createUserProfile(userId, profileData) {
        await this.db.query(
            'INSERT INTO user_profiles (user_id, phone, address, city, state, zip_code) VALUES ($1, $2, $3, $4, $5, $6)',
            [userId, profileData.phone || null, profileData.address || null, profileData.city || null, profileData.state || null, profileData.zipCode || null]
        );
    }

    async createUserSettings(userId) {
        await this.db.query(
            'INSERT INTO user_settings (user_id, theme, notifications_enabled) VALUES ($1, $2, $3)',
            [userId, 'light', true]
        );
    }
}

module.exports = UserRepository;
