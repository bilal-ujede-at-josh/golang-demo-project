package auth_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendOtp(t *testing.T) {
	// Define the JSON payload to be sent in the request body
	payload := map[string]string{
		"mobile": "7276490862",
	}

	// Serialize the JSON payload into a bytes.Buffer
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8000/send-otp", bytes.NewReader(payloadBytes))

	if err != nil {
		log.Fatal(err)
	}

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Call your HTTP handler with the request and response recorder
	handler := responseHandler
	handler.ServeHTTP(rr, req)
}

func responseHandler(req *http.Request, r *http.Response) {

}
