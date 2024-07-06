package middlewares

import (
	"context"
	"konsulin-service/internal/pkg/constvars"
	"konsulin-service/internal/pkg/exceptions"
	"konsulin-service/internal/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func AuthMiddleware(redisClient *redis.Client, jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return exceptions.WrapWithoutError(constvars.StatusUnauthorized, constvars.ErrClientNotAuthorized, constvars.ErrDevAuthTokenMissing)
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		sessionID, err := utils.ParseJWT(token, jwtSecret)
		if err != nil {
			return err
		}

		sessionData, err := redisClient.Get(context.Background(), sessionID).Result()
		if err != nil {
			return exceptions.WrapWithoutError(constvars.StatusUnauthorized, constvars.ErrClientNotAuthorized, constvars.ErrDevAuthInvalidSession)
		}

		c.Locals("sessionData", sessionData)
		return c.Next()
	}
}
