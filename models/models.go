package models

type ChangeRequest struct {
	Amount    int   `json:"amount" `
	Banknotes []int `json:"banknotes" `
}

type Exchanges struct {
	Exchanges [][]int `json:"exchanges"`
}
