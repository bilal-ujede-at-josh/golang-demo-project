package apiresponse

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func PrintJsonResponse(w http.ResponseWriter, msg string, statuscode int, data interface{}) {
	// Defined empty response structure
	res := response{
		false,
		msg,
		data,
	}

	// Convert response to bytes
	bres, err := json.Marshal(res)
	if err != nil {
		log.Println("Error prining response")
	}

	// Print output to console
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	w.Write(bres)
}
