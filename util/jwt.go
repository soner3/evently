package util

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(privKey)
}
