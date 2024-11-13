package handlers_test

import (
	"errors"
	"fmt"
	"go-test/handlers"
	"go-test/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionHandler_CalculateDiscount(t *testing.T) {
	//Arrange
	amount := 100
	expected := 80

	t.Run("success", func(t *testing.T) {
		//Arrange
		promotionService := services.NewPromotionServiceMock()
		promotionService.On("CalculateDiscount", amount).Return(expected, nil)

		promotionHandler := handlers.NewPromotionHandler(promotionService)

		//Act
		app := fiber.New()
		app.Get("/calculate-discount", promotionHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate-discount?amount=%d", amount), nil)
		res, _ := app.Test(req)

		//Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}
	})

	t.Run("invalid amount", func(t *testing.T) {
		//Arrange
		promotionService := services.NewPromotionServiceMock()
		promotionHandler := handlers.NewPromotionHandler(promotionService)

		//Act
		app := fiber.New()
		app.Get("/calculate-discount", promotionHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", "/calculate-discount?amount=ssd", nil)
		res, _ := app.Test(req)

		//Assert
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	t.Run("error from service", func(t *testing.T) {
		//Arrange
		promotionService := services.NewPromotionServiceMock()
		promotionService.On("CalculateDiscount", amount).Return(0, errors.New("service error"))
		promotionHandler := handlers.NewPromotionHandler(promotionService)

		//Act
		app := fiber.New()
		app.Get("/calculate-discount", promotionHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", "/calculate-discount?amount=100", nil)
		res, _ := app.Test(req)

		//Assert
		assert.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
	})
}
