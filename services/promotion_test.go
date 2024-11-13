package services_test

import (
	"go-test/services"
	"testing"
)

func TestPromotionService_CalculateDiscount(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		expected int
		err      error
	}{
		{name: "zero amount", amount: 0, expected: 0, err: services.ErrZeroAmount},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			promotionService := services.NewPromotionService(nil)

			discount, err := promotionService.CalculateDiscount(test.amount)
			if err != test.err {
				t.Errorf("Expected error %v, got %v", test.err, err)
			}

			if discount != test.expected {
				t.Errorf("Expected discount %d, got %d", test.expected, discount)
			}
		})
	}
}
