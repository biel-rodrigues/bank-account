package request

type Transfer struct {
	SN int     `json:"sender_number"`
	RN int     `json:"receiver_number"`
	A  float64 `json:"amount"`
}
