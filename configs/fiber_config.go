package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		EnableTrustedProxyCheck: true,
		ReadTimeout:             time.Second * time.Duration(readTimeoutSecondsCount),
		JSONEncoder:             json.Marshal,
		JSONDecoder:             json.Unmarshal,
	}
}
