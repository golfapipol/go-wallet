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
		"reciever": "1",
		"giver":    "",
		"amount":   1000,
	}
	requestBodyString, _ := json.Marshal(requestBody)
	request := httptest.NewRequest("POST", url, bytes.NewBuffer(requestBodyString))
	responseRecorder := httptest.NewRecorder()
	expected := `{"status":""}`

	TransferHandler(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if expected != string(actual) {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}
