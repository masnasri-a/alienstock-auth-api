package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(param string) string {
	hasher := sha256.New()

	// Menambahkan kata sandi ke hasher
	hasher.Write([]byte(param))

	// Mengambil hasil hash dalam bentuk byte
	hashedPassword := hasher.Sum(nil)

	// Mengonversi byte ke string dalam format hexadecimal
	return hex.EncodeToString(hashedPassword)
}
