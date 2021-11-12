package customer

import "time"

type Customer struct {
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	BirthDate time.Time `json:"birthDate"`
	Address   Address   `json:"address"`
}

type Address struct {
	Number       int    `json:"number"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Cep          string `json:"cep"`
}

func new() *Customer {
	return &Customer{
		Name:      "teste",
		Cpf:       "123",
		BirthDate: parseDate("2000-01-31"),
		Address: Address{
			Number:       1,
			Street:       "rua a",
			Neighborhood: "bairro a",
			City:         "cidade a",
			State:        "estado a",
			Country:      "país a",
			Cep:          "cep 123",
		},
	}
}

func parseDate(date string) time.Time {
	parsedDate, _ := time.Parse("2006-01-02", date)
	return parsedDate
}
