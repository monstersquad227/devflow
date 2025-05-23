package utils

import (
	"devflow/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey = []byte(config.GlobalConfig.Jwt.Secret)

// Claims 结构体用于存储 JWT 的声明信息
type Claims struct {
	Account  string    `json:"account"`
	CreateAt time.Time `json:"create_at"`
	jwt.StandardClaims
}

// GenerateToken 生成 JWT 令牌
func GenerateToken(account string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(config.GlobalConfig.Jwt.Expire) * time.Hour)
	claims := &Claims{
		Account:  account,
		CreateAt: time.Now(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ParseToken 解析 JWT 令牌
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的token")
}
