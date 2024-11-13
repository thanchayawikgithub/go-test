package services_test

import (
	"go-test/repositories"
	"go-test/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name            string
	purchaseMin     int
	discountPercent int
	amount          int
	expected        int
	err             error
}

func TestPromotionService_CalculateDiscount(t *testing.T) {
	//Arrange
	promotionRepo := repositories.NewPromotionRepositoryMock()
	promotionRepo.On("GetPromotion").Return(repositories.Promotion{ID: 1, PurchaseMin: 100, DiscountPercent: 20}, nil)

	promotionService := services.NewPromotionService(promotionRepo)

	//Act
	discount, _ := promotionService.CalculateDiscount(100)
	expected := 80

	//Assert
	assert.Equal(t, expected, discount)

	tests := []testCase{
		{name: "amount less than purchase min", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80, err: nil},
		{name: "amount more than purchase min", purchaseMin: 100, discountPercent: 20, amount: 90, expected: 90, err: nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			promotionRepo := repositories.NewPromotionRepositoryMock()
			promotionRepo.On("GetPromotion").Return(repositories.Promotion{ID: 1, PurchaseMin: test.purchaseMin, DiscountPercent: test.discountPercent}, nil)

			promotionService := services.NewPromotionService(promotionRepo)

			//Act
			discount, _ := promotionService.CalculateDiscount(test.amount)

			//Assert
			assert.Equal(t, test.expected, discount)
		})
	}

	t.Run("zero amount", func(t *testing.T) {
		//Arrange
		promotionRepo := repositories.NewPromotionRepositoryMock()
		promotionRepo.On("GetPromotion").Return(repositories.Promotion{ID: 1, PurchaseMin: 100, DiscountPercent: 20}, services.ErrZeroAmount)

		promotionService := services.NewPromotionService(promotionRepo)

		//Act
		_, err := promotionService.CalculateDiscount(0)

		//Assert
		assert.ErrorIs(t, err, services.ErrZeroAmount)
		promotionRepo.AssertNotCalled(t, "GetPromotion")
	})

	t.Run("error repo", func(t *testing.T) {
		//Arrange
		promotionRepo := repositories.NewPromotionRepositoryMock()
		promotionRepo.On("GetPromotion").Return(repositories.Promotion{}, services.ErrRepository)

		promotionService := services.NewPromotionService(promotionRepo)

		//Act
		_, err := promotionService.CalculateDiscount(100)

		//Assert
		assert.ErrorIs(t, err, services.ErrRepository)
	})
}
