package customer

import "time"

type customer struct {
	name      string
	cpf       string
	birthDate time.Time
	address   address
}

type address struct {
	number       int
	street       string
	neighborhood string
	city         string
	state        string
	country      string
	cep          string
}

func new() *customer {
	return &customer{
		name:      "teste",
		cpf:       "123",
		birthDate: parseDate("2000-01-31"),
		address: address{
			number:       1,
			street:       "rua a",
			neighborhood: "bairro a",
			city:         "cidade a",
			state:        "estado a",
			country:      "país a",
			cep:          "cep 123",
		},
	}
}

func parseDate(date string) time.Time {
	parsedDate, _ := time.Parse("2006-01-02", date)
	return parsedDate
}
