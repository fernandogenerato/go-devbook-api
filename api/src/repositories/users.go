package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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

func (u Users) FindUser(name string) (models.User, bool) {
	rows, err := u.db.Query("select * from users where name = ?", name)
	if err != nil {
		return models.User{}, false
	}

	defer rows.Close()
	us := models.User{}
	for rows.Next() {
		if err := rows.Scan(&us.ID, &us.Name, &us.Email, &us.Password, &us.CreatedAt); err != nil {
			return models.User{}, false
		}
		return us, true
	}
	return us, false
}

func (u Users) FindUsers() []models.User {
	rows, err := u.db.Query("select id, name, email, createdAt from users")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	data := []models.User{}

	for rows.Next() {
		us := models.User{}
		if err := rows.Scan(&us.ID, &us.Name, &us.Email, &us.CreatedAt); err != nil {
			log.Println(err)
			return []models.User{}
		}
		data = append(data, us)
	}
	return data
}

func (u Users) FindUserByID(id uint64) (models.User, bool) {
	rows, err := u.db.Query("select * from users where id = ? ", id)
	if err != nil {
		return models.User{}, false
	}

	defer rows.Close()

	for rows.Next() {
		us := models.User{}
		if err := rows.Scan(&us.ID, &us.Name, &us.Email, &us.Password, &us.CreatedAt); err != nil {
			return models.User{}, false
		}
		return us, true
	}
	return models.User{}, false
}

func (u Users) Update(user models.User) error {
	st, err := u.db.Prepare("update users set name = ? , email = ? where id = ?")
	if err != nil {
		return err
	}
	_, err = st.Exec(user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u Users) DeleteUser(id uint64) error {
	st, err := u.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	res, err := st.Exec(id)
	if err != nil {
		return err
	}

	r, err := res.RowsAffected()
	if r == 0 || err != nil {
		return errors.New("user not found")
	}

	return nil
}

func (u Users) FindUserByEmail(email string) (models.User, error) {
	rows, err := u.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()
	var user models.User
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}
