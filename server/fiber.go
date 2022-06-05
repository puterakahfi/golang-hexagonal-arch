package server

import (
	"github.com/gofiber/fiber/v2"
)

// fiber server implementation (adapter)
type fiberServer struct {
	fiber *fiber.App
}

func InitFiberServer() *fiberServer {
	app := fiber.New()
	return &fiberServer{fiber: app}
}

func (f *fiberServer) Run() {
	f.fiber.Listen(":8099")
}

func (f *fiberServer) GetServerInstance() interface{} {
	return f.fiber
}
