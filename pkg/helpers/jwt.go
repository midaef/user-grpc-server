package helpers

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
)

// JWTManager ...
type JWTManager interface {
	NewJWT(userID string, ttl time.Duration) (string, error)
	NewJWTParser(accessToken string) (string, error)
	NewRefreshToken() string
}

// PrivateKey ...
type PrivateKey struct {
	privateKey string
}

// NewPrivateKey ...
func NewPrivateKey(privateKey string) (*PrivateKey, error) {
	if privateKey == "" {
		return nil, errors.New("jwt: private key is empty")
	}
	return &PrivateKey{
		privateKey: privateKey,
	}, nil
}

// NewJWT ...
func (pk *PrivateKey) NewJWT(userID string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   userID,
	})
	return token.SignedString([]byte(pk.privateKey))
}

// NewJWTParser ...
func (pk *PrivateKey) NewJWTParser(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("jwt: unexpected signing method")
		}
		return []byte(pk.privateKey), nil
	})
	if err != nil {
		return "", errors.New("jwt: parser error")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("jwt: claims error")
	}
	return claims["sub"].(string), nil
}

// NewRefreshToken ...
func (pk *PrivateKey) NewRefreshToken() string {
	token := securecookie.GenerateRandomKey(64)
	strToken := base64.StdEncoding.EncodeToString(token)
	return strToken
}
