package main

type userRepository struct {
	db Database
}

func NewUserRepository(db Database) UserRepository {
	return &userRepository{db: db}
}

func (ur userRepository) UpdateUserProfile(userId int, name, email string) error {
	_, err := ur.db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", name, email, userId)
	return err
}

func (ur userRepository) GetUserEmail(userId int) (string, error) {
	rows, err := ur.db.Query("SELECT email FROM users WHERE id = ?", userId)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			return "", err
		}
		return email, nil
	}

	return "", sql.ErrNoRows
}
