package wallet

import "testing"

func Test_Transfer_Should_Be_Transfer_Success(t *testing.T) {
	giver := Account{
		Name:   "Oat",
		Id:     "123456789",
		Amount: float64(12000),
	}
	receiver := Account{
		Name:   "Nut",
		Id:     "987654321",
		Amount: float64(5000),
	}
	transferMoney := float64(10000)
	expected := TransferResult{
		Status: "Success",
		Receiver: Account{
			Name:   "Nut",
			Id:     "987654321",
			Amount: float64(15000),
		},
		Giver: Account{
			Name:          "Oat",
			Id:            "123456789",
			Amount:        float64(2000),
			transferLimit: float64(10000),
		},
	}
	actual := Transfer(giver, receiver, transferMoney)

	if expected != actual {
		t.Errorf("expected %v but it got %v", expected, actual)
	}
}
