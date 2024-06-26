package handlers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Validatable interface {
	Validate() error
}

type Handler struct {
	Logger    *zap.Logger
	Validator *validator.Validate
}

func (h *Handler) BodyParserValidator(c *fiber.Ctx, out any) error {
	if err := c.BodyParser(out); err != nil {
		return fmt.Errorf("can't parse body: %w", err)
	}

	return h.validateStruct(out)
}

func (h *Handler) QueryParserValidator(c *fiber.Ctx, out any) error {
	if err := c.QueryParser(out); err != nil {
		return fmt.Errorf("can't parse query: %w", err)
	}

	return h.validateStruct(out)
}

func (h *Handler) ParamsParserValidator(c *fiber.Ctx, out any) error {
	if err := c.ParamsParser(out); err != nil {
		return fmt.Errorf("can't parse params: %w", err)
	}

	return h.validateStruct(out)
}

func (h *Handler) validateStruct(out any) error {
	if h.Validator != nil {
		if err := h.Validator.Struct(out); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	}

	if req, ok := out.(Validatable); ok {
		if err := req.Validate(); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	}

	return nil
}
