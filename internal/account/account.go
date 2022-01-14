package account

import (
	"github.com/bank-account/internal/customer"
	"github.com/bank-account/internal/dtos/request"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Number     int               `json:"number" gorm:"unique"`
	Balance    float64           `json:"balance"`
	CustomerID int               `json:"customer_id"`
	Customer   customer.Customer `json:"customer"`
}

func convert(DTO request.Account) Account {
	return Account{
		Number:     DTO.Number,
		CustomerID: DTO.CustomerID,
		Balance:    DTO.Balance,
	}
}
