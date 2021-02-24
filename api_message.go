package kakaoapi

import (
	"encoding/json"
	"log"

	message "github.com/roharon/kakao-api-go/schema/api_message"
)

const (
	ApiSendMe               string = "/v2/api/talk/memo/default/send"
	ApiSendMeWithUrl        string = "/v2/api/talk/memo/scrap/send"
	ApiSendMeWithTemplateId string = "/v2/api/talk/memo/send"
)

// Send Message to me.
//
// https://developers.kakao.com/docs/latest/ko/message/rest-api#send-me
func (c *Client) SendMe(request interface{}) (message.ResponseApiMessageMe, error) {
	var requestInterface map[string]interface{}

	request = message.RequestApiMessageMe{
		TemplateObject: request,
	}

	params, err := json.Marshal(request)
	if err != nil {
		log.Printf("* Failed to encode struct: %s", err)
		return message.ResponseApiMessageMe{}, err
	}

	if err := json.Unmarshal(params, &requestInterface); err != nil {
		log.Printf("* Failed to decode struct to interface: %s", err)
		return message.ResponseApiMessageMe{}, err
	}

	bytes, err := c.post(APIKakaoURL+ApiSendMe, authTypeBearer, nil, requestInterface)
	if err != nil {
		return message.ResponseApiMessageMe{}, err
	}

	response := message.ResponseApiMessageMe{}
	err = json.Unmarshal(bytes, &response)

	if err != nil {
		log.Printf("* Failed to decode bytes: %s", string(bytes))
		return message.ResponseApiMessageMe{}, err
	}

	return response, nil
}

// Send Scrap Message to me
//
// https://developers.kakao.com/docs/latest/ko/message/rest-api#send-me-with-url
func (c *Client) SendMeWithUrl(requestUrl string, templateId string, templateArgs interface{}) (message.ResponseApiMessageMe, error) {
	var requestInerface map[string]interface{}

	if templateArgs == "" {
		templateArgs = message.TemplateArguments{}
	}

	request := message.RequestApiMessageMeWithUrl{
		RequestUrl:   requestUrl,
		TemplateId:   templateId,
		TemplateArgs: templateArgs.(message.TemplateArguments),
	}

	params, err := json.Marshal(request)
	if err != nil {
		log.Printf("* Failed to encode struct: %s", err)
		return message.ResponseApiMessageMe{}, err
	}

	if err := json.Unmarshal(params, &requestInerface); err != nil {
		log.Printf("* Failed to decode struct to interface: %s", err)
		return message.ResponseApiMessageMe{}, err
	}

	bytes, err := c.post(APIKakaoURL+ApiSendMeWithUrl, authTypeBearer, nil, requestInerface)
	if err != nil {
		return message.ResponseApiMessageMe{}, err
	}

	response := message.ResponseApiMessageMe{}
	if err := json.Unmarshal(bytes, &response); err != nil {
		log.Printf("* Failed to decode bytes: %s", string(bytes))
		return message.ResponseApiMessageMe{}, err
	}

	return response, nil
}

// Send Template Message to me
//
// /developers.kakao.com/docs/latest/ko/message/rest-api#send-me-with-template-id
func (c *Client) SendMeWithTemplateId(templateId string, templateArgs interface{}) (message.ResponseApiMessageMe, error) {
	var requestInterface map[string]interface{}

	if templateArgs == "" {
		templateArgs = message.TemplateArguments{}
	}

	request := message.RequestApiMessageWithTemplateId{
		TemplateId:   templateId,
		TemplateArgs: templateArgs.(message.TemplateArguments),
	}

	params, err := json.Marshal(request)
	if err != nil {
		log.Printf("* Failed to encode struct: %s", err)
		return message.ResponseApiMessageMe{}, err
	}

	if err := json.Unmarshal(params, &requestInterface); err != nil {
		log.Printf("* Failed to decode struct to interface: %s", err)
		return message.ResponseApiMessageMe{}, err
	}

	bytes, err := c.post(APIKakaoURL+ApiSendMeWithTemplateId, authTypeBearer, nil, requestInterface)
	if err != nil {
		return message.ResponseApiMessageMe{}, err
	}

	response := message.ResponseApiMessageMe{}
	if err := json.Unmarshal(bytes, &response); err != nil {
		log.Printf("* Failed to decode bytes: %s", string(bytes))
		return message.ResponseApiMessageMe{}, err
	}

	return response, nil
}
