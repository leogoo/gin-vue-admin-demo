package utils

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// jwt包自带的jwt.StandardClaims只包含了官方字段
// 额外记录username,
type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var jwtSecret = []byte("夏天夏天悄悄过去")

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	// 过期时间
	expireTime := nowTime.Add(time.Hour * 3)

	claims := MyClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-vue-admin-demo", //签发人
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

func ParseToken(token string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
