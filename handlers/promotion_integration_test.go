//go:build integration

package handlers_test

import (
	"fmt"
	"go-test/handlers"
	"go-test/repositories"
	"go-test/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscountIntegrationService(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		amount := 100
		expected := 80

		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promotionService := services.NewPromotionService(promoRepo)
		promotionHandler := handlers.NewPromotionHandler(promotionService)

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
}
