package kakaoapi

import (
	"encoding/json"
	"log"

	social "github.com/roharon/kakao-api-go/schema/api_social"
)

const (
	ApiGetProfile string = "/v1/api/talk/profile"
	ApiGetFriends string = "/v1/api/talk/friends"
)

// Get user profile
//
// https://developers.kakao.com/docs/latest/ko/kakaotalk-social/rest-api#get-profile
func (c *Client) GetProfile() (social.ResponseApiSocial, error) {
	var bytes []byte
	var err error

	bytes, err = c.get(APIKakaoURL+ApiGetProfile, authTypeBearer, nil, nil)

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

// Get Friends list
//
// https://developers.kakao.com/docs/latest/ko/kakaotalk-social/rest-api#get-friends
func (c *Client) GetFriends(offset int, limit int,
	order social.Order, friendOrder social.FriendOrder) (social.ResponseApiFriends, error) {
	var bytes []byte
	var requestInterface map[string]interface{}

	data := &social.RequestApiFriends{
		Offset:      offset,
		Limit:       limit,
		Order:       order,
		FriendOrder: friendOrder,
	}

	params, err := json.Marshal(data)
	if err != nil {
		log.Printf("* Failed to encode struct: %s", err)
	}

	if err := json.Unmarshal(params, &requestInterface); err != nil {
		log.Printf("* Failed to decode struct to interface: %s", err)
	}

	bytes, err = c.get(APIKakaoURL+ApiGetFriends, authTypeBearer, nil, requestInterface)
	if err != nil {
		return social.ResponseApiFriends{}, err
	}

	response := social.ResponseApiFriends{}
	if err := json.Unmarshal(bytes, &response); err != nil {
		log.Printf("* Failed to decode bytes: %s", err)
	}

	return response, nil
}
