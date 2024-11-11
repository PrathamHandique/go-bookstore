package utils

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"log"
)

func ParseBody(r *http.Request, x interface{}) {
	// Read the body of the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		return
	}

	// Unmarshal the JSON data into the provided struct
	err = json.Unmarshal(body, x)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return
	}
}
