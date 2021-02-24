package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	kakaoapi "github.com/roharon/kakao-api-go"
	msg "github.com/roharon/kakao-api-go/schema/api_message"
)

const resultCode int = 0

func TestMessageSendMe(t *testing.T) {
	client := kakaoapi.NewClient("abcedfg")

	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", kakaoapi.APIKakaoURL+kakaoapi.ApiSendMe,
		func(req *http.Request) (*http.Response, error) {
			jsonStr := fmt.Sprintf(`
			{
				"result_code": %d
			}
			`, resultCode)

			data := make(map[string]interface{})
			if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
				t.Log(err)
			}

			resp, err := httpmock.NewJsonResponse(200, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		})

	WEB_URL := "https://google.com"

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

	assertEqual(t, data.ButtonTitle, "sample Title", "")
	assertEqual(t, data.ObjectType, "feed", "")
	assertEqual(t, data.Buttons[0].Link.WebUrl, WEB_URL, "")
	res, err := client.SendMe(data)

	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, res.ResultCode, resultCode, "")
}
