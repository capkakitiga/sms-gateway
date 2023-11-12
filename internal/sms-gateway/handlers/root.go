package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type rootHandler struct {
}

func (h *rootHandler) Register(app *fiber.App) {
	app.Static("/", "static")
}

func newRootHandler() *rootHandler {
	return &rootHandler{}
}
