package utils

import (
	"encoding/base32"
)

func EncodeBase64(bytes []byte) string {
	encoded := base32.StdEncoding.EncodeToString(bytes)
	return encoded
}
