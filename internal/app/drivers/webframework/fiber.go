package webframework

import (
	"fmt"
	"konsulin-service/internal/app/config"
	"konsulin-service/internal/app/delivery/http/middlewares"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewFiber(driverConfig *config.DriverConfig) *fiber.App {
	fiberConfig := fiber.Config{
		// Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		AppName:       fmt.Sprintf("Konsulin Service %s", driverConfig.App.Version),
		BodyLimit:     driverConfig.App.RequestBodyLimitInMegabyte * 1024 * 1024,
		ErrorHandler:  middlewares.ErrorHandler,
	}
	app := fiber.New(fiberConfig)

	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        10,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"error":   "Too many requests on single time-frame",
			})
		},
	}))
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${ip} | ${status} ==> ${latency} | ${method} | ${path}\n",
		TimeFormat: time.RFC850,
		TimeZone:   driverConfig.App.Timezone,
	}))
	return app
}
