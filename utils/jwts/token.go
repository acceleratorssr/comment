package jwts

import (
	"comment/global"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtPayload struct {
	Username    string `json:"username"`
	UserID      uint   `json:"user_id"`
	NickName    string `json:"nick_name"`
	Permissions int    `json:"permissions"`
}

type CustomClaims struct {
	JwtPayload
	jwt.RegisteredClaims
}

var MySecret []byte

// GenToken 创建Token
func GenToken(user JwtPayload) (string, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

func ParseToken(tokenStr string) (*CustomClaims, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("token parse err: %s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
