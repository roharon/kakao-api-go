package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	kakaoapi "github.com/roharon/kakao-api-go"
	"github.com/roharon/kakao-api-go/schema/api_social"
)

const (
	nickName     string = "개발자"
	thumbnailUrl string = "https://xxx.kakao.co.kr/.../aaa.jpg"
	user1Uuid    string = "abcdefg0002"
	user2Image   string = "https://xxx.kakao.co.kr/.../bbb.jpg"
	totalCount   int    = 11
)

func TestApiSocialProfile(t *testing.T) {
	client := kakaoapi.NewClient("abcdefg")

	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", kakaoapi.APIKakaoURL+kakaoapi.ApiGetProfile,
		func(req *http.Request) (*http.Response, error) {
			json_str := fmt.Sprintf(`
			{
				"nickName":"%s",
				"profileImageURL":"https://xxx.kakao.co.kr/.../aaa.jpg",
				"thumbnailURL":"%s",
				"countryISO":"KR"
			}`, nickName, thumbnailUrl)

			data := make(map[string]interface{})
			err := json.Unmarshal([]byte(json_str), &data)

			if err != nil {
				t.Log(err)
			}
			resp, err := httpmock.NewJsonResponse(200, data)

			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	res, err := client.GetProfile()
	if err != nil {
		t.Error(err)
	} else {
		assertEqual(t, res.NickName, nickName, "")
		assertEqual(t, res.ThumbNailUrl, thumbnailUrl, "")
	}
}

func TestApiSocialFriends(t *testing.T) {
	client := kakaoapi.NewClient("abcdefg")

	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", kakaoapi.APIKakaoURL+kakaoapi.ApiGetFriends,
		func(req *http.Request) (*http.Response, error) {
			json_str := fmt.Sprintf(`
			{
				"after_url": "https://kapi.kakao.com/v1/api/talk/friends?offset=3&limit=3&order=asc",
				"elements": [
						{
								"id": 1,
								"uuid": "%s",
								"favorite": true,
								"profile_nickname": "이수민",
								"profile_thumbnail_image": "https://xxx.kakao.co.kr/.../aaa.jpg"
						},
						{
								"id": 2,
								"uuid": "abcdefg0002",
								"favorite": false,
								"profile_nickname": "홍길동",
								"profile_thumbnail_image": "%s"
						},
						 {
								"id": 3,
								"uuid": "abcdefg0003",
								"favorite": false,
								"profile_nickname": "김철수",
								"profile_thumbnail_image": "https://xxx.kakao.co.kr/.../ccc.jpg"
						}
				],
				"total_count": %d,
				"favorite_count": 1
			}`, user1Uuid, user2Image, totalCount)

			data := make(map[string]interface{})
			err := json.Unmarshal([]byte(json_str), &data)

			if err != nil {
				t.Log(err)
			}
			resp, err := httpmock.NewJsonResponse(200, data)

			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	res, err := client.GetFriends(0, 3, api_social.Asc, api_social.Favorite)
	if err != nil {
		t.Error(err)
	} else {
		assertEqual(t, res.TotalCount, totalCount, "")
		assertEqual(t, res.Elements[0].Uuid, user1Uuid, "")
		assertEqual(t, res.Elements[1].ProfileThumbnailImage, user2Image, "")
	}
}
