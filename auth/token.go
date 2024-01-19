package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"kingkong-be/common"
	"kingkong-be/domain/user"
	"log"
	"net/http"
	"time"
)

func CreateToken(user *user.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	// Create the JWT claims with userID as a custom claim
	claims := jwt.MapClaims{
		"user_id":  user.UserID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      expirationTime.Unix(),
	}

	// Create the token with claims and sign it using a secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(common.JWTSigningKey) // Replace with your secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateTokenMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
		c.Abort()
		return
	}

	tokenString = tokenString[len("Bearer "):]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.JWTSigningKey), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
		c.Abort()
		return
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user_id in claims"})
		c.Abort()
		return
	}

	userID := int(userIDFloat) // Convert float64 to int

	username, ok := claims["username"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username in claims"})
		c.Abort()
		return
	}

	roleFloat, ok := claims["role"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role in claims"})
		c.Abort()
		return
	}

	role := int(roleFloat) // Convert float64 to int

	log.Println(userID, username, role)
	c.Set("user_id", userID)
	c.Set("username", username)
	c.Set("role", role)

	// Token is valid, proceed to the next handler
	c.Next()
}
