package request

type Address struct {
	Number       int    `json:"number"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Cep          string `json:"cep"`
}
