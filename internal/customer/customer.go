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

func dtoToCustomer(dto request.Customer) *Customer {
	a := Address{
		Number:       dto.Address.Number,
		Street:       dto.Address.Street,
		Neighborhood: dto.Address.Neighborhood,
		City:         dto.Address.City,
		State:        dto.Address.State,
		Country:      dto.Address.Country,
		Cep:          dto.Address.Country,
	}

	c := Customer{
		Name:    dto.Name,
		Cpf:     dto.Cpf,
		Address: &a,
	}

	return &c
}
