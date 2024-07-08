package server

import (
	"pertamina_abi/database"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App
	DB         database.Service
	serverName string
}

func New(serverName string, dbConfig database.Config) *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: serverName,
			AppName:      serverName,
		}),
		DB:         database.New(dbConfig),
		serverName: serverName,
	}

	return server
}
