package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() (string, error) {
	var v []byte
	rand.Read(v)
	str := hex.EncodeToString(v)

	return str, nil

}
