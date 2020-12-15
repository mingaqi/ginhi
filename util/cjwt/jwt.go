package cjwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	TokenExpired     error  = errors.New("token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "daily_paper_mqP9YCmhOuhpLbbJ"
)

// 		"iat":    time.Now().Unix(),                     // Token颁发时间
//		"nbf":    time.Now().Unix(),                     // Token生效时间
//		"exp":    time.Now().Add(time.Hour * 12).Unix(), // Token过期时间，目前是24小时
//		"iss":    "",                                    // 颁发者
//		"sub":    "AuthToken",                           // 主题
type CustomClaims struct {
	UserId int64
	DeptId int64
	jwt.StandardClaims
}

// 创建token
func CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SignKey))
}

// 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	fmt.Println("new token: ", tokenString)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	fmt.Println("err: ", err)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("That's not even a token")
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				fmt.Println("token is expired")
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				fmt.Println("Token not active yet")
				return nil, TokenNotValidYet
			} else {
				fmt.Println("Couldn't handle this token:")
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		fmt.Println(claims)
		return claims, nil
	}
	fmt.Println("Couldn't handle this token: 2")
	return nil, TokenInvalid
}

// 更新token
func RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return CreateToken(*claims)
	}
	return "", TokenInvalid
}
