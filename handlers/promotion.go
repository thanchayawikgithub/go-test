package handlers

import (
	"go-test/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type (
	PromotionHandler interface {
		CalculateDiscount(c *fiber.Ctx) error
	}

	promotionHandler struct {
		promotionService services.PromotionService
	}
)

func NewPromotionHandler(promotionService services.PromotionService) PromotionHandler {
	return &promotionHandler{promotionService: promotionService}
}

func (h promotionHandler) CalculateDiscount(c *fiber.Ctx) error {
	amountStr := c.Query("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid amount"})
	}

	discount, err := h.promotionService.CalculateDiscount(amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendString(strconv.Itoa(discount))
}
