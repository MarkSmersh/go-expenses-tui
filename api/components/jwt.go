package components

import (
	"errors"
	"log/slog"
	"os"
)

func GetJwtSecretKey() (string, error) {
	key := os.Getenv("JWT_SIGN_KEY")

	if len(key) <= 0 {
		slog.Error("Enviroment variable JWT_SIGN_KEY is missing. Due to their missing creating an access token is impossible.")
		return "", errors.New("")
	}

	if len(key) != 32 {
		slog.Error("Wrong length of JWT_SIGN_KEY. The signing method requires 256 bytes. Every character in a string has 8 bytes. So the string needs to have 32 symbols.")
		return "", errors.New("")
	}

	return key, nil
}
