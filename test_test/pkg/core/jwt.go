package core

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"os"
	"strings"
	"time"
)

var DefaultExpired = time.Hour * 24

type JWTPayload struct {
	Sub   string   `json:"sub"`
	Roles []string `json:"roles"`
	Exp   int64    `json:"exp"`
}

func GetString(c *fiber.Ctx, key string) string {
	return GetStringBySecret(c, key, os.Getenv("JWT_SECRET"))
}

func GetSub(token string) string {
	if payload, err := getPayloadByToken(token, os.Getenv("JWT_SECRET")); err == nil {
		return payload.Sub
	}
	return ""
}

func GetRoleId(c *fiber.Ctx) string {
	payload := GetJWTPayload(c)
	if len(payload.Roles) > 0 {
		return payload.Roles[0]
	}
	return ""
}

func GetStringBySecret(c *fiber.Ctx, key string, secretKey string) string {
	authorization := Authorization(c)
	if authorization != "" {
		tokenStr := strings.Replace(authorization, "Bearer ", "", -1)
		hmacSecret := []byte(secretKey)
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return hmacSecret, nil
		})

		if err != nil {
			return ""
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			uid := claims[key]
			return uid.(string)
		}
	}
	return ""
}

func GetJWTPayload(c *fiber.Ctx) JWTPayload {
	payload, _ := getJWTPayload(c, os.Getenv("JWT_SECRET"))
	return payload
}

func getJWTPayload(c *fiber.Ctx, secret string) (JWTPayload, error) {
	authorization := Authorization(c)
	if authorization != "" {
		tokenStr := strings.Replace(authorization, "Bearer ", "", -1)
		return getPayloadByToken(tokenStr, secret)
	}
	return JWTPayload{}, errors.New("Authorization is empty")
}

func getPayloadByToken(tokenStr string, secret string) (JWTPayload, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return JWTPayload{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if p, err := json.Marshal(claims); err == nil {
			payload := JWTPayload{}
			if err = json.Unmarshal(p, &payload); err == nil {
				return payload, nil
			}
		}
	}
	return JWTPayload{}, errors.New("Token invalid")
}

func NewClaims() jwt.MapClaims {
	// Create token
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := jwtToken.Claims.(jwt.MapClaims)
	return claims
}

func GenerateJWT(secret string, claims map[string]interface{}) string {
	// Set claims
	payload := jwt.MapClaims{}
	for k, v := range claims {
		payload[k] = v
	}

	// Generate encoded token and send it as response.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return ""
	}
	return fmt.Sprintf("Bearer %s", tokenStr)
}
