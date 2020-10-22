package jwtAdmin

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWT签名结构
type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error  = errors.New("token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "sxtsgz2020"
)

// 荷载
type Customclaims struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

//设置SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

//获取SignKey
func GetSignKey() string {
	return SignKey
}

// 生成SignKey
func NewJwt() *JWT {
	return &JWT{[]byte(GetSignKey())}
}

// 创建token
func (j *JWT) CreateToken(claims Customclaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析token
// 验证Token过程中，如果Token生成过程中，指定了iat与exp参数值，将会自动根据时间戳进行时间验证
func (j *JWT) ParseToken(tokenString string) (*Customclaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Customclaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*Customclaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 生成rmt自有token
func (j *JWT) GenerateToken(user string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":     time.Now().Unix(),                     // Token颁发时间
		"nbf":     time.Now().Unix(),                     // Token生效时间
		"exp":     time.Now().Add(time.Hour * 12).Unix(), // Token过期时间
		"iss":     "stgz",                                // 颁发者
		"userkey": user,
	})
	return token.SignedString(j.SigningKey)
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Customclaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Customclaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
