package users

import (
	"users/internal/helpers"
	"users/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	users := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.ID, &data.Username, &data.Email, &data.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return users, err
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE id=?", id.String())
	helpers.CloseDB(db)

	var user models.User
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Age)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func AddUser(user *models.User) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	id, err := uuid.NewV4() // Generate a new UUID

	// Convert UUID to string directly before inserting into the database
	_, err = db.Exec("INSERT INTO users (id, username, email, age) VALUES (?, ?, ?, ?)", id.String(), user.Username, user.Email, user.Age)
	if err != nil {
		return err
	}

	return nil
}

func EditUser(user *models.User) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE users SET username = ?, email = ?, age = ? WHERE id = ?", user.Username, user.Email, user.Age, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(userID uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		return err
	}

	return nil
}
