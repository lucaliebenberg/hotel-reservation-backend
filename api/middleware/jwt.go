package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(c *fiber.Ctx) error {
	fmt.Println("---- JWT Auth ---")

	headers := c.GetReqHeaders()         // returns a map of strings
	tokens, ok := headers["X-Api-Token"] // returns a slice of strings ([]strings)
	if !ok || len(tokens) == 0 {
		return fmt.Errorf("unauthorised")
	}

	token := tokens[0]
	if err := parseToken(token); err != nil {
		return err
	}

	return nil
}

func parseToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Invalid signing method:", token.Header["alg"])
			return nil, fmt.Errorf("Unauthorised")
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse token:", err)
		return fmt.Errorf("unauthorised")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	}
	return fmt.Errorf("unauthorized")
}
