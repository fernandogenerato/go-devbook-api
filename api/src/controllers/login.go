package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"go-devbook-api/src/db"
	"go-devbook-api/src/models"
	"go-devbook-api/src/repositories"
	"go-devbook-api/src/response"
	"go-devbook-api/src/secure"
)

func Login(w http.ResponseWriter, r *http.Request) {
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
	dbHelper := repositories.NewUsersRepository(db)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	data, err := dbHelper.FindUserByEmail(u.Email)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	err = secure.CompareSecret(u.Password, data.Password)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}
	response.JSON(w, http.StatusOK, nil)
}
