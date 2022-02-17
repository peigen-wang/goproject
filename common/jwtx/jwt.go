package jwtx

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

//获取jwt token
func GetToken(secreKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	accesstoken, err := token.SignedString([]byte(secreKey))
	logx.Info(accesstoken, err)
	return accesstoken, err
}
