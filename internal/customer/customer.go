package customer

import (
	"time"

	"github.com/bank-account/internal/dtos/request"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	BirthDate time.Time `json:"birthDate" gorm:"type:datetime"`
	AddressID uint      `gorm:"column:address_id"`
	Address   *Address  `json:"address"`
}

type Address struct {
	gorm.Model
	Number       int    `json:"number"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Cep          string `json:"cep"`
}

func convert(DTO request.Customer) *Customer {
	a := Address{
		Number:       DTO.Address.Number,
		Street:       DTO.Address.Street,
		Neighborhood: DTO.Address.Neighborhood,
		City:         DTO.Address.City,
		State:        DTO.Address.State,
		Country:      DTO.Address.Country,
		Cep:          DTO.Address.Country,
	}

	c := Customer{
		Name:    DTO.Name,
		Cpf:     DTO.Cpf,
		Address: &a,
	}

	return &c
}
