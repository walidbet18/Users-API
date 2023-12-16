package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	ID       *uuid.UUID `json:"id"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Age      string     `json:"age"`
}
