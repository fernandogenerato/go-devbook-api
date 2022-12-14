package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"go-devbook-api/src/models"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *Users {
	return &Users{
		db: db,
	}
}

func (u Users) Create(user models.User) (uint64, error) {
	st, err := u.db.Prepare("insert into users (name, email, password, createdAt) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer st.Close()

	result, err := st.Exec(user.Name, user.Email, user.Password, time.Now())
	if err != nil {
		return 0, fmt.Errorf("error to exec insert : %w", err)
	}

	id, _ := result.LastInsertId()
	return uint64(id), nil
}
