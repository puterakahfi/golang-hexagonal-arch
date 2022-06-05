package api

import (
	"golang-hexagonal-arch/module/book/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type fiberHandler struct {
	engine  *fiber.App
	service usecase.BookUseCaseInterface
}
