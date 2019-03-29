package helper

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Claims struct
type Claims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(name, password string) (string, error) {
	var jwtSecret = []byte("yq-starter")
	nowTime := time.Now()
	expireTime := nowTime.Add(6 * time.Hour)

	claims := Claims{
		name,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "yq-starter-issuer",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	var jwtSecret = []byte("yq-starter")
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// JWTMiddleware : add user on header.
func JWTMiddleware() gin.HandlerFunc {
	// Set out header value for each response
	return func(c *gin.Context) {
		token := c.Query(Token)
		claims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				Code:  http.StatusUnauthorized,
				Error: err.Error(),
			})
			c.Abort()
			return
		}
		c.Header("X-App-User", claims.Name)
		c.Next()
	}
}
