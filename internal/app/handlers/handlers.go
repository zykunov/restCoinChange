package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zykunov/RESTCurrencyChange/models"
)

// Алгоритм нахождения размена
func coinCombinations(goal int, coins []int) [][]int {
	wallets := make([][]int, len(coins))
	for i, coin := range coins {
		wallets[i] = []int{coin}
	}
	var newWallets [][]int
	var collected [][]int

	for len(wallets) > 0 {
		for _, wallet := range wallets {
			s := sum(wallet)
			for _, coin := range coins {
				if coin >= wallet[len(wallet)-1] {
					if s+coin < goal {
						newWallet := make([]int, len(wallet)+1)
						copy(newWallet, wallet)
						newWallet[len(wallet)] = coin
						newWallets = append(newWallets, newWallet)
					} else if s+coin == goal {
						newWallet := make([]int, len(wallet)+1)
						copy(newWallet, wallet)
						newWallet[len(wallet)] = coin
						collected = append(collected, newWallet)
					}
				}
			}
		}
		wallets = newWallets
		newWallets = nil
	}
	return collected
}

func sum(wallet []int) int {
	s := 0
	for _, coin := range wallet {
		s += coin
	}
	return s
}

// Хэндлер для "/change"
func Change(c echo.Context) error {

	request := new(models.ChangeRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	changes := coinCombinations(request.Amount, request.Banknotes)

	response := models.Exchanges{
		Exchanges: changes,
	}

	return c.JSON(http.StatusCreated, response)

}
