package main

import (
	"encoding/json"
	"fmt"
	"log"

	kakaoapi "github.com/roharon/kakao-api-go"
)

const (
	apiKey = "test_key" // XXX - change this to User's oauth token
)

func main() {
	client := kakaoapi.NewClient(apiKey)

	if response, err := client.GetProfile(); err != nil {
		log.Printf("Error: %s", err)
	} else {
		log.Printf("Response: %s", prettify(response))
	}

}

func prettify(obj interface{}) string {
	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		return string(bytes)
	}

	return fmt.Sprintf("%v", obj)
}
