package util

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var privKey *rsa.PrivateKey
var pubKey rsa.PublicKey

func init() {
	var err error
	privKey, err = rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic("failed to generate RSA key: " + err.Error())
	}
	pubKey = privKey.PublicKey
}

func GenerateToken(email, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.RegisteredClaims{
		Subject: userId,
		Issuer:  "http://localhost:8080",
		Audience: jwt.ClaimStrings{
			email,
		},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 30)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ID:        uuid.NewString(),
	})

	return token.SignedString(privKey)
}

func ValidateToken(token string) (*jwt.Token, *jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		return pubKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodRS512.Alg()}))

	if err != nil {
		return nil, nil, errors.New("invalid token")
	}

	return jwtToken, claims, nil
}
