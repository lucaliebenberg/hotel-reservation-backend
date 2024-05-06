package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lucaliebenberg/hotel-reservation/db"
)

func JWTAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {

		headers := c.GetReqHeaders()         // returns a map of strings
		tokens, ok := headers["X-Api-Token"] // returns a slice of strings ([]strings)
		if !ok || len(tokens) == 0 {
			return ErrUnauthorized()
		}

		token := tokens[0]
		// fmt.Println("---token", token)
		claims, err := validateToken(token)
		if err != nil {
			return err
		}
		// fmt.Println("--- claims", claims)

		expiresStr := claims["expires"].(string)
		expiresTime, err := time.Parse(time.RFC3339, expiresStr)
		if err != nil {
			return fmt.Errorf("failed to parse expiration time: %v", err)
		}

		// Check token expiration
		if time.Now().After(expiresTime) {
			return NewError(http.StatusUnauthorized, "token expired")
		}
		userID := claims["id"].(string)
		user, err := userStore.GetUserByID(c.Context(), userID)
		if err != nil {
			return ErrUnauthorized()
		}
		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Invalid signing method:", token.Header["alg"])
			return nil, ErrUnauthorized()
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse token:", err)
		return nil, ErrUnauthorized()
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, ErrUnauthorized()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrUnauthorized()
	}

	return claims, nil
}
