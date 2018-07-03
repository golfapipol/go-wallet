package api

import (
	"encoding/json"
	"net/http"
)

type transferRequest struct {
	Receiver string
	Giver    string
	Amount   float64
}

type transferResponse struct {
	Status string `json:"status"`
}

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	var transferRequest transferRequest
	json.NewDecoder(r.Body).Decode(&transferRequest)

	transferResponse := transferResponse{}
	response, _ := json.Marshal(transferResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
