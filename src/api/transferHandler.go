package api

import (
	"encoding/json"
	"net/http"
)

type TransferRequest struct {
	Receiver string  `json:"receiver"`
	Giver    string  `json:"giver"`
	Amount   float64 `json:"amount"`
}

type transferResponse struct {
	Status string `json:"status"`
}

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	transferRequest := TransferRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&transferRequest)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	transferResponse := transferResponse{}
	response, _ := json.Marshal(transferResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
