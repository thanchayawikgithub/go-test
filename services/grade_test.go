package services_test

import (
	"fmt"
	"go-test/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGrade(t *testing.T) {

	tests := []struct {
		name     string
		score    int
		expected string
	}{
		{name: "A", score: 90, expected: "A"},
		{name: "B", score: 70, expected: "B"},
		{name: "C", score: 60, expected: "C"},
		{name: "D", score: 50, expected: "D"},
		{name: "F", score: 40, expected: "F"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			grade := services.CheckGrade(test.score)

			assert.Equal(t, test.expected, grade)
		})
	}
}

func BenchmarkCheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}
}

func ExampleCheckGrade() {
	grade := services.CheckGrade(90)
	fmt.Println(grade)
	// Output: A
}
