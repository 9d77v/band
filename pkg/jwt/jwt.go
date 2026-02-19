package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Conf Conf
}

func NewJWT(conf Conf) JWT {
	return JWT{
		Conf: conf,
	}
}

type CustomClaims struct {
	jwt.RegisteredClaims
}

func (j *JWT) CreateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        id,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.Conf.Issuer,
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1 * time.Minute)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.Conf.ExpireTime) * time.Second)),
		},
	})
	return token.SignedString([]byte(j.Conf.SecretKey))
}

func (j *JWT) ParserToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(j.Conf.SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
