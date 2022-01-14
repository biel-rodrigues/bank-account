package request

type Account struct {
	Number     int     `json:"number"`
	CustomerID int     `json:"customer_id"`
	Balance    float64 `json:"balance"`
}
