package transfer

import (
	"errors"

	"github.com/bank-account/internal/account"
	"gorm.io/gorm"
)

func create(DB *gorm.DB, senderN, receiverN int, amount float64) (senderBalance float64, err error) {
	s, err := account.UseCaseReadByNumber(DB, senderN)
	if err != nil {
		return 0.0, err
	}

	r, err := account.UseCaseReadByNumber(DB, receiverN)
	if err != nil {
		return 0.0, err
	}

	senderHasAmount := s.Balance > amount
	if senderHasAmount {
		sNewBalance := s.Balance - amount
		account.UpdateBalance(DB, int(s.ID), sNewBalance)

		rNewBalance := r.Balance + amount
		account.UpdateBalance(DB, int(r.ID), rNewBalance)

		return sNewBalance, nil
	}
	return 0.0, errors.New("Insufficient balance")
}
