class UserManager {
    constructor(db) {
        this.db = db;
    }

    async registerUser(userData) {
        // Validate input data
        if (!userData.email) {
            throw new Error('Email is required');
        }
        if (!userData.password) {
            throw new Error('Password is required');
        }
        if (userData.password.length < 8) {
            throw new Error('Password must be at least 8 characters');
        }
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(userData.email)) {
            throw new Error('Invalid email format');
        }

        // Check if user already exists
        const existingUser = await this.db.query('SELECT id FROM users WHERE email = $1', [userData.email]);
        if (existingUser.rows.length > 0) {
            throw new Error('User already exists');
        }

        // Hash password
        const bcrypt = require('bcrypt');
        const hashedPassword = await bcrypt.hash(userData.password, 10);

        // Generate verification token
        const crypto = require('crypto');
        const verificationToken = crypto.randomBytes(32).toString('hex');

        // Insert user into database
        const userResult = await this.db.query(
            'INSERT INTO users (email, password, first_name, last_name, verification_token, created_at) VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING id',
            [userData.email, hashedPassword, userData.firstName, userData.lastName, verificationToken]
        );

        const userId = userResult.rows[0].id;

        // Create user profile
        await this.db.query(
            'INSERT INTO user_profiles (user_id, phone, address, city, state, zip_code) VALUES ($1, $2, $3, $4, $5, $6)',
            [userId, userData.phone || null, userData.address || null, userData.city || null, userData.state || null, userData.zipCode || null]
        );

        // Send verification email
        const nodemailer = require('nodemailer');
        const transporter = nodemailer.createTransporter({
            service: 'gmail',
            auth: { user: 'noreply@example.com', pass: 'password' }
        });

        const subject = 'Please verify your email address';
        const message = `Hello ${userData.firstName},\n\nThank you for registering. Please click the link below to verify your email:\n\nhttp://example.com/verify?token=${verificationToken}\n\nBest regards,\nThe Team`;

        await transporter.sendMail({
            from: 'noreply@example.com',
            to: userData.email,
            subject: subject,
            text: message
        });

        // Log registration
        console.log(`User registered: ${userData.email} at ${new Date().toISOString()}`);

        // Create default settings
        await this.db.query(
            'INSERT INTO user_settings (user_id, theme, notifications_enabled) VALUES ($1, $2, $3)',
            [userId, 'light', true]
        );

        // Send welcome notification
        await this.db.query(
            'INSERT INTO notifications (user_id, type, message, created_at) VALUES ($1, $2, $3, NOW())',
            [userId, 'welcome', 'Welcome to our platform!']
        );

        return userId;
    }
}

module.exports = UserManager;
