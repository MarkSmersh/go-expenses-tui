package components

import (
	"crypto/sha256"
	"encoding/base64"
)

func HashPassword(raw string) string {
	h := sha256.New()
	h.Write([]byte(raw))
	hashedPassword := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	return string(hashedPassword)
}
