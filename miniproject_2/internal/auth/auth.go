package auth

import (
	"encoding/base64"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}

func ParseBasicAuthHeader(header string) (string, string, error) {
	decoded, err := base64.StdEncoding.DecodeString(header)
	if err != nil {
		return "", "", errors.New("invalid authorization header")
	}

	credentials := strings.SplitN(string(decoded), ":", 2)
	if len(credentials) != 2 {
		return "", "", errors.New("invalid authorization header")
	}

	return credentials[0], credentials[1], nil
}
