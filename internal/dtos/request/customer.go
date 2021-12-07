package request

import "time"

type Customer struct {
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	BirthDate time.Time `json:"birthDate"`
	Address   Address   `json:"address"`
}
