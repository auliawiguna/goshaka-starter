package helpers

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServer(a *fiber.App) {
	fiberConnectionUrl, _ := ConnectionURLBuilder("fiber")

	if err := a.Listen(fiberConnectionUrl); err != nil {
		log.Printf("Cannot run server %v", err)
	}
}

func StartServerWithGracefulShutdown(a *fiber.App) {
	idleConnectionsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Cannot shutdown server %v", err)
		}

		close(idleConnectionsClosed)
	}()

	fiberConnectionUrl, _ := ConnectionURLBuilder("fiber")

	if err := a.Listen(fiberConnectionUrl); err != nil {
		log.Printf("Cannot run server %v", err)
	}

	<-idleConnectionsClosed
}
