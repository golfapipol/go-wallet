package wallet

type Account struct {
	Name          string  `json: "name"`
	Id            string  `json: "id"`
	Amount        float64 `json: "amount"`
	TransferLimit float64
}

const MAX_PER_TRANACTION = float64(20000)
const MAX_PER_DAY = float64(100000)

type TransferResult struct {
	Status   string  `json: "status"`
	Receiver Account `json: "receiver"`
	Giver    Account `json: "giver"`
	Fee      float64 `json:"fee"`
}

func Transfer(giver, receiver Account, transferAmount float64) TransferResult {
	if giver.Amount < transferAmount {
		return TransferResult{Status: "Not Enough", Giver: giver, Receiver: receiver}
	}

	if transferAmount > MAX_PER_TRANACTION {
		return TransferResult{Status: "Maximum Transfer per time is 20,000", Giver: giver, Receiver: receiver}
	}

	if (transferAmount + transferAmount) > MAX_PER_DAY {
		return TransferResult{Status: "Maximum Transfer per Day is 100,000", Giver: giver, Receiver: receiver}
	}

	giver.Amount -= transferAmount
	receiver.Amount += transferAmount
	fee := float64(0)
	giver.TransferLimit += transferAmount

	return TransferResult{Status: "Success", Giver: giver, Receiver: receiver, Fee: fee}
}
