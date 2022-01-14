package account

import (
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
