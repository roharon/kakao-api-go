package api_message

type TemplateObject struct {
	ObjectType  string   `json:"object_type"`
	ButtonTitle string   `json:"button_title"`
	Buttons     []Button `json:"buttons"`
}

type RequestApiMessageMe struct {
	TemplateObject TemplateObject `json:"template_object"`
}

type RequestApiMessageMeWithUrl struct {
	RequestUrl   string                 `json:"request_url"`
	TemplateId   string                 `json:"template_id"`
	TemplateArgs map[string]interface{} `json:"template_args"`
}

type ResponseApiMessageMe struct {
	ResultCode int `json:"result_code"`
}

type Content struct {
	Title       string `json:"title"`
	ImageUrl    string `json:"image_url"`
	ImageWidth  int    `json:"image_width"`
	ImageHeight int    `json:"image_height"`
	Description string `json:"description"`
	Link        Link   `json:"link"`
}

type Button struct {
	Title string `json:"title"`
	Link  Link   `json:"link"`
}

// Must set 1 attribute at least
type Link struct {
	WebUrl                 string `json:"web_url"`
	MobileWebUrl           string `json:"mobile_web_url"`
	AndroidExecutionParams string `json:"android_execution_params"`
	IosExecutionParams     string `json:"ios_execution_params"`
}

// Object for represent social information likes like, comment, etc...
// represent same or under thatn 3 attributes
// Priority: Like > Comment > Shared > View > Subscriber
type Social struct {
	LikeCount       int `json:"like_count"`
	CommentCount    int `json:"comment_count"`
	SharedCount     int `json:"shared_count"`
	ViewCount       int `json:"view_count"`
	SubscriberCount int `json:"subscriber_count"`
}

type Commerce struct {
	ProductName        string `json:"product_name"`
	RegularPrice       int    `json:"regular_price"`
	DiscountPrice      int    `json:"discount_price"`
	DiscountRate       int    `json:"discount_rate"`
	FixedDiscountPrice int    `json:"fixed_discount_price"`
}

type FeedObject struct {
	TemplateObject
	Content Content `json:"content"`
	Social  Social  `json:"social"`
}

type ListObject struct {
	TemplateObject
	HeaderTitle string    `json:"header_title"`
	HeaderLink  Link      `json:"header_link"`
	Contents    []Content `json:"contents"`
}

type LocationObject struct {
	TemplateObject
	Address      string  `json:"address"`
	AddressTitle string  `json:"address_title"`
	Content      Content `json:"content"`
	Social       Social  `json:"social"`
}

type CommerceObject struct {
	TemplateObject
	Content  Content  `json:"content"`
	Commerce Commerce `json:"commerce"`
}

type TextObject struct {
	TemplateObject
	Text string `json:"text"`
	Link Link   `json:"link"`
}
