package wallet

type Account struct {
	Name          string  `json: "name"`
	Id            string  `json: "id"`
	Amount        float64 `json: "amount"`
	TransferLimit float64
}

type TransferResult struct {
	Status   string  `json: "status"`
	Receiver Account `json: "receiver"`
	Giver    Account `json: "giver"`
	Fee      float64 `json:"fee"`
}

func Transfer(giver, receiver Account, amount float64) TransferResult {
	return TransferResult{}
}
