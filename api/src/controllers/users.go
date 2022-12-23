package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go-devbook-api/src/auth"
	"go-devbook-api/src/db"
	"go-devbook-api/src/models"
	"go-devbook-api/src/repositories"
	"go-devbook-api/src/response"
	"go-devbook-api/src/secure"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var u models.User
	err = json.Unmarshal(body, &u)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := u.Validate(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	var pass []byte
	if pass, err = secure.Hash(u.Password); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	u.Password = string(pass)

	db, err := db.Connection()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	dbHelper := repositories.NewUsersRepository(db)
	id, err := dbHelper.Create(u)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	u, _ = dbHelper.FindUserByID(id)
	response.JSON(w, http.StatusCreated, u)
}
func FindUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connection()
	if err != nil {
		response.Error(w, 500, err)
		return
	}
	defer db.Close()

	us := repositories.NewUsersRepository(db).FindUsers()
	response.JSON(w, 200, us)
}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 64)

	db, err := db.Connection()
	if err != nil {
		response.Error(w, 500, err)
		return
	}
	defer db.Close()

	user, ok := repositories.NewUsersRepository(db).FindUserByID(id)
	if !ok {
		response.Error(w, 404, errors.New("user not found"))
	}
	response.JSON(w, 200, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	ut, err := auth.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != ut {
		response.Error(w, http.StatusForbidden, errors.New("user has no permission to update"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var u models.User
	err = json.Unmarshal(body, &u)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connection()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	u.ID = userID
	err = repositories.NewUsersRepository(db).Update(u)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	ut, err := auth.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	if id != ut {
		response.Error(w, http.StatusForbidden, errors.New("user has no permission to delete"))
		return
	}

	db, err := db.Connection()
	if err != nil {
		response.Error(w, 500, err)
		return
	}
	defer db.Close()

	err = repositories.NewUsersRepository(db).DeleteUser(id)
	if err != nil {
		response.Error(w, 404, errors.New("user not found"))
	}
	response.JSON(w, http.StatusNoContent, nil)

}
