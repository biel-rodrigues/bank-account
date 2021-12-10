package account

import (
	"github.com/bank-account/internal/dtos/request"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Number     int     `json:"number"`
	CustomerID int     `json:"customer_id"`
	Balance    float64 `json:"balance"`
}

func convert(DTO request.Account) Account {
	return Account{
		Number:     DTO.Number,
		CustomerID: DTO.CustomerID,
		Balance:    DTO.Balance,
	}
}
