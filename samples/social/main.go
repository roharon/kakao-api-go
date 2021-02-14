package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	kakaoapi "github.com/roharon/kakao-api-go"
	"github.com/roharon/kakao-api-go/schema/api_social"
)

func get_friends(client *kakaoapi.Client) {
	data := []byte(`{
		"elements": [
			{
				"profile_nickname": "김땡땡",
				"profile_thumbnail_image": "",
				"id": 1379141686,
				"uuid": "VmBXbaclUWSQfa91E2AcCUF2T2NSAlFoXTx",
				"favorite": false
			},
			{
				"profile_nickname": "박ㅇㅇ",
				"profile_thumbnail_image": "",
				"id": 1364106081,
				"uuid": "VmNTZ1JrXmdLf093QCUF2T1hUGBTal8B",
				"favorite": false
			},
			{
				"profile_nickname": "신ㅇㅇ",
				"profile_thumbnail_image": "",
				"id": 1554498953,
				"uuid": "VmVVYldhV29ZdUBCUF2Tfk06VmdXZF1oDw",
				"favorite": false
			},
			{
				"profile_nickname": "한ㅇㅇ",
				"profile_thumbnail_image": "",
				"id": 1454431081,
				"uuid": "VmNWZlZmVWZKe05CUF2Td096VmdXZF1oAQ",
				"favorite": false
			}
		],
		"total_count": 4,
		"after_url": null,
		"result_id": "zbnYtN-BuIu3ibuNtYXa69vo0eS7yazLotGlwILXx-yF512yg1md-KeWyajEqPec81az6Qw",
		"favorite_count": 0
	}`)

	response := api_social.ResponseApiFriends{}
	if err := json.Unmarshal(data, &response); err != nil {
		log.Printf("* Failed to decode bytes: %s", err)
	} else {
		log.Printf("Response: %s", prettify(response))
		log.Print(response.Elements[0].Id)
	}
}

func main() {
	apiKey := os.Getenv("apiKey")
	client := kakaoapi.NewClient(apiKey)

	if response, err := client.GetProfile(); err != nil {
		log.Printf("Error: %s", err)
	} else {
		log.Printf("Response: %s", prettify(response))
	}

	if response, err := client.GetFriends(0, 5, "asc", "favorite"); err != nil {
		log.Printf("Error: %s", err)
	} else {
		log.Printf("Response: %s", prettify(response))
	}

	get_friends(client)
}

func prettify(obj interface{}) string {
	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		return string(bytes)
	}

	return fmt.Sprintf("%v", obj)
}
