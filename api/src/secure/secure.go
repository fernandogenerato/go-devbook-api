package secure

import "golang.org/x/crypto/bcrypt"

func Hash(secret string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
}

func CompareSecret(secret, secrethash string) error {
	return bcrypt.CompareHashAndPassword([]byte(secrethash), []byte(secret))
}
