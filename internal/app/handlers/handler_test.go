package handlers

import (
	"fmt"
	"slices"
	"testing"
)

func TestCoinCombinations(t *testing.T) {

	t.Run("simple", func(t *testing.T) {
		var amount int = 5
		var banknotes = []int{1, 2, 3}
		var expectedResult = [][]int{
			{2, 3},
			{1, 1, 3},
			{1, 2, 2},
			{1, 1, 1, 2},
			{1, 1, 1, 1, 1},
		}

		realResult := coinCombinations(amount, banknotes)

		for i := 0; i < len(realResult); i++ {
			fmt.Printf("got %d want %d \n", realResult[i], expectedResult[i])

			areEqual := slices.Compare(expectedResult[i], realResult[i])
			if areEqual != 0 {
				t.Errorf("test fail, got %d want %d ", realResult[i], expectedResult[i])
			}
		}
	})

	t.Run("medium", func(t *testing.T) {
		var amount int = 400
		var banknotes = []int{5000, 2000, 1000, 500, 200, 100, 50}
		var expectedResult = [][]int{
			{200, 200},
			{100, 100, 200},
			{100, 100, 100, 100},
			{50, 50, 100, 200},
			{50, 50, 100, 100, 100},
			{50, 50, 50, 50, 200},
			{50, 50, 50, 50, 100, 100},
			{50, 50, 50, 50, 50, 50, 100},
			{50, 50, 50, 50, 50, 50, 50, 50},
		}

		realResult := coinCombinations(amount, banknotes)

		for i := 0; i < len(realResult); i++ {
			fmt.Printf("got %d want %d \n", realResult[i], expectedResult[i])

			areEqual := slices.Compare(expectedResult[i], realResult[i])
			if areEqual != 0 {
				t.Errorf("test fail, got %d want %d ", realResult[i], expectedResult[i])
			}
		}
	})

}
