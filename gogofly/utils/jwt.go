package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type JWTCustomClaims struct {
	ID   uint
	Name string
	jwt.RegisteredClaims
}

var signKey string
var signMethod = jwt.SigningMethodHS256

func GenerateToken(id uint, name string) (string, error) {

	claim := JWTCustomClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                                                   // 签发者
			Subject:   "Token",                                                                         // 签发对象
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.Expire") * time.Hour)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                  //签发时间
		},
	}
	if signKey == "" {
		signKey = viper.GetString("jwt.SignKey")
	}
	token, err := jwt.NewWithClaims(signMethod, claim).SignedString([]byte(signKey))
	return token, err
}

func ParseToken(tokenStr string) (*JWTCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*JWTCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}

func IsValidToken(token string) bool {
	_, err := ParseToken(token)
	return err == nil
}
