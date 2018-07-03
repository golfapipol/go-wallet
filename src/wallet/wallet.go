package wallet

const MAX_PER_TRANACTION = float64(20000)
const MAX_PER_DAY = float64(100000)

type TransferResult struct {
	Status   string  `json:"status"`
	Receiver Account `json:"receiver"`
	Giver    Account `json:"giver"`
	Fee      float64 `json:"fee"`
}

type Account struct {
	Name          string  `json:"name"`
	Id            string  `json:"id"`
	Amount        float64 `json:"amount"`
	transferLimit float64
}

func TransferProcess(giverId, receiverId string, transferAmount float64) TransferResult {
	giver := getAccount(giverId)
	receiver := getAccount(receiverId)
	return Transfer(giver, receiver, transferAmount)
}

func Transfer(giver, receiver Account, transferAmount float64) TransferResult {
	if checkEnoughMoney(giver, transferAmount) {
		return TransferResult{Status: "Not Enough", Giver: giver, Receiver: receiver}
	}

	if checkLimitPerTransaction(transferAmount) {
		return TransferResult{Status: "Maximum Transfer per time is 20,000", Giver: giver, Receiver: receiver}
	}

	if checkLimitPerDay(giver, transferAmount) {
		return TransferResult{Status: "Maximum Transfer per Day is 100,000", Giver: giver, Receiver: receiver}
	}

	// TODO: ย้ายเป็น function
	giver.Amount -= transferAmount
	fee := float64(0)
	// TODO: ย้ายเป็น function
	receiver.Amount += transferAmount
	// TODO: ย้ายเป็น function
	giver.transferLimit += transferAmount

	return TransferResult{Status: "Success", Giver: giver, Receiver: receiver, Fee: fee}
}

func getAccount(id string) Account {
	users := map[string]Account{
		"123456789": Account{
			Name:   "Oat",
			Id:     "123456789",
			Amount: float64(12000),
		},
		"987654321": Account{
			Name:   "Nut",
			Id:     "987654321",
			Amount: float64(5000),
		},
	}
	return users[id]
}

func checkEnoughMoney(acc Account, transferAmount float64) bool {
	return acc.Amount < transferAmount
}

func checkLimitPerTransaction(transferAmount float64) bool {
	return transferAmount < MAX_PER_TRANACTION
}

func checkLimitPerDay(acc Account, transferAmount float64) bool {
	return (acc.Amount + transferAmount) > MAX_PER_DAY
}
