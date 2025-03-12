package restapi

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var AdminPassword string

// JWT token titkosításához.
var JwtSecret []byte

// Ellenőrzi a kapott base64-ba kódolt jelszót.
// GET kérés Authorization fejléc alapján.
//
// Visszaad egy JWT tokent, ha megegyezik a megadott jelszóval.
//
// https://pkg.go.dev/github.com/golang-jwt/jwt/v5#section-readme
//
// https://ututuv.medium.com/building-user-authentication-and-authorisation-api-in-go-using-gin-and-gorm-93dfe38e0612
func CheckAdminPassword(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Authorization header"})
		return
	}

	password, err := base64.StdEncoding.DecodeString(authParts[1])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if string(password) != AdminPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}
	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := generatedToken.SignedString(JwtSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate token"})
	}

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}

// Ellenőrzi a kapott base64-ba kódolt JWT tokent
func CheckAdminToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	// c.AbortWithStatus needed for checker functions
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Authorization header"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authToken, err := base64.StdEncoding.DecodeString(authParts[1])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-ErrorChecking
	accessToken, err := jwt.Parse(string(authToken), func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JwtSecret), nil
	})
	if err != nil || !accessToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	claims, ok := accessToken.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Succesful check, go to api function
	c.Next()
}
