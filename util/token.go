package util

import (
	"Gin-Basic-Service/global"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

const (
	secretKey = "custom-secret-key" // 密钥，用于签名和验证 token
)

// Claims 结构体用于定义 JWT 的声明
type Claims struct {
	Operator *global.Operator `json:"operator"` // 操作员ID
	jwt.StandardClaims
}

// GenerateToken 生成 token
func GenerateToken(operator *global.Operator) (string, error) {
	expiration, err := strconv.ParseInt(global.Conf.Server.TokenExpiration, 10, 64)
	// 创建声明
	claims := &Claims{
		Operator: operator,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expiration) * time.Hour).Unix(), // 设置过期时间
		},
	}

	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名 token
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析 token
func ParseToken(tokenString string) (*Claims, error) {
	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// 验证 token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
