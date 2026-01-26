package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {

	suserid := userID.String()
	regtdclaims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
		Subject:   suserid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, regtdclaims)
	str, err := token.SignedString(tokenSecret)
	return str, err //yahan chudke waps aaunga
}
