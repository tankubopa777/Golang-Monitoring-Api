package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

func GenerateJWT(email string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateToken(signedToken string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(
        signedToken,
        &Claims{},
        func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        },
    )
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, err
    }

    return claims, nil
}
