package users

import (
	"database/sql"
	"net/http"
	"users/internal/models"
	repository "users/internal/repositories/users"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllUsers() ([]models.User, error) {
	var err error
	// calling repository
	users, err := repository.GetAllUsers()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return users, nil
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	user, err := repository.GetUserById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return user, err
}

func AddUser(user *models.User) (*models.User, error) {
	err := repository.AddUser(user)
	if err != nil {
		logrus.Errorf("error adding user: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Failed to add user",
			Code:    http.StatusInternalServerError,
		}
	}

	return user, nil
}
func UpdateUser(user *models.User) (*models.User, error) {
	err := repository.EditUser(user)
	if err != nil {
		logrus.Errorf("error updating user: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Failed to update user",
			Code:    http.StatusInternalServerError,
		}
	}

	return user, nil
}

func DeleteUser(id uuid.UUID) error {
	err := repository.DeleteUser(id)
	if err != nil {
		logrus.Errorf("error deleting user: %s", err.Error())
		return &models.CustomError{
			Message: "Failed to delete user",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}
