package util

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

var privKey *rsa.PrivateKey
var pubKey *rsa.PublicKey

func init() {
	var err error
	privKey, err = rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic("failed to generate RSA key: " + err.Error())
	}
	pubKey = &privKey.PublicKey
}

func GenerateToken(email, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.RegisteredClaims{
		Subject: userId,
		Issuer:  "http://localhost:8080",
		Audience: jwt.ClaimStrings{
			email,
		},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
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
		return nil, nil, fmt.Errorf("token error: %w", err)
	}

	return jwtToken, claims, nil
}

func ExtractTokenFromHeader(ctx *context.Context) (*string, error) {
	md, ok := metadata.FromIncomingContext(*ctx)
	if !ok {
		return nil, errors.New("invalid request metadata")
	}

	authValues := md.Get("authorization")
	if len(authValues) < 1 {
		return nil, errors.New("empty token")
	}

	token, found := strings.CutPrefix(authValues[0], "Bearer ")

	if !found {
		return nil, errors.New("invalid authorization prefix")
	}

	return &token, nil
}
