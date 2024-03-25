package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var secret = []byte("epoch-Chain")

const TokenExpireDuration = 24 * time.Hour

func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "" {
			claim, err := ParseToken(token)
			if err != nil {
				c.JSON(http.StatusForbidden, "Invalid token")
				c.Abort()
				return
			}
			c.Set("email", claim.Email)
			c.Next()
		}
	}
}

// generate JWT token
func GenerateToken(email string) (string, error) {
	// claim
	c := AuthClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "epochChain",
		},
	}
	// HS256 signature
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// return the token string
	return token.SignedString(secret)
}

// Parse JWT token
func ParseToken(tokenStr string) (*AuthClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaim{}, func(tk *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claim, ok := token.Claims.(*AuthClaim); ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("invalid token")
}
