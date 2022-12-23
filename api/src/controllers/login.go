package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"go-devbook-api/src/auth"
	"go-devbook-api/src/db"
	"go-devbook-api/src/models"
	"go-devbook-api/src/repositories"
	"go-devbook-api/src/response"
	"go-devbook-api/src/secure"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var u models.User
	err = json.Unmarshal(body, &u)
	if err != nil {
		log.Printf("login.json.Unmarshal: %v", err)
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connection()
	if err != nil {
		log.Printf("login.db.Connection: %v", err)
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

	token, err := auth.CreateToken(data.ID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	res := map[string]any{
		"token": token,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}

	d, _ := json.Marshal(res)

	w.Write(d)
}
