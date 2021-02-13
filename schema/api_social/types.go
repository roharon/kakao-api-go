package api_social

type ResponseApiSocial struct {
	NickName        string `json:"nickName"`
	ProfileImageUrl string `json:"profileImageUrl"`
	ThumbNailUrl    string `json:"thumbnailUrl"`
	CountryISO      string `json:"countryISO"`
}
