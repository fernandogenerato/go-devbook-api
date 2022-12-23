package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go-devbook-api/src/config"
)

func ExtractUserID(r *http.Request) (uint64, error) {
	t := extractToken(r)
	token, err := jwt.Parse(t, retrieveKey)
	if err != nil {
		return 0, err
	}

	if rules, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		u, err := strconv.ParseUint(fmt.Sprintf("%.0f", rules["userID"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return u, nil
	}
	return 0, errors.New("invalid_token")
}

func CreateToken(userID uint64) (string, error) {
	rules := jwt.MapClaims{}
	rules["authorized"] = true
	rules["userID"] = userID
	rules["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, rules)
	return token.SignedString(config.SecretKey)
}

func ValidateToken(r *http.Request) error {
	t := extractToken(r)
	token, err := jwt.Parse(t, retrieveKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Print(t)
		return nil
	}
	return errors.New("invalid_token")

}

func retrieveKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("sign method unexpected! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}
