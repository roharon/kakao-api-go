package api_social

type Order string
type FriendOrder string

const (
	Asc      Order       = "asc"
	Desc     Order       = "desc"
	Favorite FriendOrder = "favorite"
	Nickname FriendOrder = "nickname"
)

type Friend struct {
	Id                    int64  `json:"id"`
	ProfileNickname       string `json:"profile_nickname"`
	ProfileThumbnailImage string `json:"profile_thumbnail_image"`
	Uuid                  string `json:"uuid"`
	Favorite              bool   `json:"favorite"`
}

type ResponseApiSocial struct {
	NickName        string `json:"nickName"`
	ProfileImageUrl string `json:"profileImageUrl"`
	ThumbNailUrl    string `json:"thumbnailUrl"`
	CountryISO      string `json:"countryISO"`
}

type RequestApiFriends struct {
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	Order       Order       `json:"order"`
	FriendOrder FriendOrder `json:"friend_order"`
}

type ResponseApiFriends struct {
	Elements      []Friend `json:"elements"`
	TotalCount    int      `json:"total_count"`
	FavoriteCount int      `json:"favorite_count"`
	BeforeUrl     string   `json:"before_url"`
	AfterUrl      string   `json:"after_url"`
}

type ResponseError struct {
	Message string `json:"msg"`
	Code    string `json:"code"`
}
