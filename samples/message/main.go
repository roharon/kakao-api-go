package main

import (
	"log"
	"os"

	kakaoapi "github.com/roharon/kakao-api-go"
	msg "github.com/roharon/kakao-api-go/schema/api_message"
)

const (
	WEB_URL = "https://google.com"
)

func main() {
	apiKey := os.Getenv("apiKey")
	client := kakaoapi.NewClient(apiKey)

	// Feed Object
	data := msg.NewTemplateObject(msg.FeedType).(msg.FeedObject)
	data.SetButtonTitle("sample Title")
	data.SetButtons(
		msg.NewButton("sample title",
			msg.NewLink(WEB_URL, "", "", ""),
		),
	)
	data.SetContent(
		msg.NewContent("content title", "https://image.png", 1, 1,
			"description", msg.NewLink("https://github.com", "", "", ""),
		),
	)
	data.SetSocial(
		msg.NewSocial(1, 2, 3, 4, 5),
	)

	res, err := client.SendMe(data)

	if err != nil {
		log.Print(err)
	}
	log.Print(res)

	// SendMeWithUrl example
	res, err = client.SendMeWithUrl("https://github.com/roharon/kakao-api-go", "", "")
	if err != nil {
		log.Print(err)
	}
	log.Print(res)
}
