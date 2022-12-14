package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"go-devbook-api/src/db"
	"go-devbook-api/src/models"
	"go-devbook-api/src/repositories"
	"go-devbook-api/src/response"
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

	db, err := db.Connection()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	rep := repositories.NewUsersRepository(db)
	id, err := rep.Create(u)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	u.ID = id
	response.JSON(w, http.StatusCreated, u)
}
func FindUsers(w http.ResponseWriter, r *http.Request)  {}
func FindUser(w http.ResponseWriter, r *http.Request)   {}
func UpdateUser(w http.ResponseWriter, r *http.Request) {}
func DeleteUser(w http.ResponseWriter, r *http.Request) {}
