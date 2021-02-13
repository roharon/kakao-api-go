package kakaoapi

import (
	"encoding/json"
	"log"

	social "github.com/roharon/kakao-api-go/schema/api_social"
)

const (
	ApiTalkProfile string = "/v1/api/talk/profile"
)

func (c *Client) GetProfile() (social.ResponseApiSocial, error) {
	var bytes []byte
	var err error

	bytes, err = c.get(APIKakaoURL+ApiTalkProfile, authTypeBearer, nil, nil)

	if err != nil {
		return social.ResponseApiSocial{}, err
	}

	response := social.ResponseApiSocial{}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		log.Printf("* Failed to decode bytes: %s", string(bytes))
		return social.ResponseApiSocial{}, err
	}

	return response, nil
}
