package account

import (
	"errors"

	"github.com/bank-account/internal/dtos/request"
	"gorm.io/gorm"
)

func create(DB *gorm.DB, DTO request.Account) (response Account, err error) {
	a := convert(DTO)

	err = DB.Create(&a).Error
	if err != nil {
		return Account{}, err
	}
	return a, nil
}

func read(DB *gorm.DB, accountID int) (response Account, err error) {
	account := Account{}
	DB.Find(&account, accountID)
	if account.ID == 0 {
		return Account{}, errors.New("account not found")
	}
	return account, nil
}

func UseCaseReadByNumber(DB *gorm.DB, n int) (response Account, err error) {
	account := Account{}
	DB.Where("number = ?", n).Find(&account)
	if account.ID == 0 {
		return Account{}, errors.New("account not found")
	}
	return account, nil
}

func readAllByCustomerID(DB *gorm.DB, customerID int) (accounts []Account, err error) {
	DB.Where("customer_id = ?", customerID).Find(&accounts)
	if len(accounts) == 0 {
		return accounts, errors.New("no account found")
	}
	return accounts, nil
}

func UpdateBalance(DB *gorm.DB, id int, newBalance float64) {
	var a Account
	DB.Where("id=?", id).Find(&a)

	a.Balance = newBalance
	DB.Save(&a)
}
