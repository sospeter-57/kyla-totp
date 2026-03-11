package utils

import (
	"crypto/rand"
	"log"
)

func GenerateRandomBytes(length int) []byte {
	if length%8 != 0 {
		log.Fatal("incorrent length of a secret key, try 16 or 32 or 64 or 128 or 256 or 512")
	}
	var buffer []byte = make([]byte, length)

	_, err := rand.Read(buffer)
	if err != nil {
		log.Fatal("Error: unable to generate random bytes.")
	}

	return buffer
}
