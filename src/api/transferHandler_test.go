package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_TransferHandler_Input_giver_oat_receiver_nut_should_be_success(t *testing.T) {
	url := "/transfer"
	requestBody := map[string]interface{}{
		"giver":    "123456789",
		"receiver": "987654321",
		"amount":   1000,
	}
	requestBodyString, _ := json.Marshal(requestBody)
	request := httptest.NewRequest("POST", url, bytes.NewBuffer(requestBodyString))
	responseRecorder := httptest.NewRecorder()
	expected := `{"status":"Success","receiver":{"name":"Nut","id":"987654321","amount":6000},"giver":{"name":"Oat","id":"123456789","amount":11000},"fee":0}`

	TransferHandler(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if expected != string(actual) {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}
