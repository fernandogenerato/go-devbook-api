package models

import (
	"errors"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (u *User) Validate() error {

	if u.Name == "" {
		return errors.New("Name is required")
	}

	if u.Email == "" {
		return errors.New("Email is required")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return err
	}

	if u.Password == "" {
		return errors.New("Invalid password")
	}

	return nil
}
