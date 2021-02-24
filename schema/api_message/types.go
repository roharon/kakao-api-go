package api_message

const (
	FeedType ObjectType = iota
	ListType
	LocationType
	ComemrceType
	TextType
)

type ObjectType int

type RequestApiMessageMe struct {
	TemplateObject interface{} `json:"template_object"`
}

type RequestApiMessageMeWithUrl struct {
	RequestUrl   string            `json:"request_url"`
	TemplateId   string            `json:"template_id"`
	TemplateArgs TemplateArguments `json:"template_args"`
}

type RequestApiMessageWithTemplateId struct {
	TemplateId   string            `json:"template_id"`
	TemplateArgs TemplateArguments `json:"template_args"`
}

type ResponseApiMessageMe struct {
	ResultCode int `json:"result_code"`
}

type TemplateArguments struct {
	ScrapImage         string `json:"${SCRAP_IMAGE}"`
	ScrapImageWidth    string `json:"${SCRAP_IMAGE_WIDTH}"`
	ScrapImageHeight   string `json:"${SCRAP_IMAGE_HEIGHT}"`
	ScrapImageDuration string `json:"${SCRAP_IMAGE_DURATION}"`
	ScrapTitle         string `json:"${SCRAP_TITLE}"`
	ScrapDescription   string `json:"${SCRAP_DESCRIPTION}"`
	ScrapHost          string `json:"${SCRAP_HOST}"`
	ScrapRequestedUrl  string `json:"${SCRAP_REQUESTED_URL}"`
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
	ObjectType  string   `json:"object_type" default:"feed"`
	ButtonTitle string   `json:"button_title"`
	Buttons     []Button `json:"buttons"`
	Content     Content  `json:"content"`
	Social      Social   `json:"social"`
}

type ListObject struct {
	ObjectType  string    `json:"object_type" default:"list"`
	ButtonTitle string    `json:"button_title"`
	Buttons     []Button  `json:"buttons"`
	HeaderTitle string    `json:"header_title"`
	HeaderLink  Link      `json:"header_link"`
	Contents    []Content `json:"contents"`
}

type LocationObject struct {
	ObjectType   string   `json:"object_type" default:"location"`
	ButtonTitle  string   `json:"button_title"`
	Buttons      []Button `json:"buttons"`
	Address      string   `json:"address"`
	AddressTitle string   `json:"address_title"`
	Content      Content  `json:"content"`
	Social       Social   `json:"social"`
}

type CommerceObject struct {
	ObjectType  string   `json:"object_type" default:"commerce"`
	ButtonTitle string   `json:"button_title"`
	Buttons     []Button `json:"buttons"`
	Content     Content  `json:"content"`
	Commerce    Commerce `json:"commerce"`
}

type TextObject struct {
	ObjectType  string   `json:"object_type" default:"text"`
	ButtonTitle string   `json:"button_title"`
	Buttons     []Button `json:"buttons"`
	Text        string   `json:"text"`
	Link        Link     `json:"link"`
}

func NewContent(title string, imageUrl string, imageWidth int,
	imageHeight int, description string, link Link) Content {
	return Content{
		Title: title, ImageUrl: imageUrl,
		ImageWidth: imageWidth, ImageHeight: imageHeight,
		Description: description, Link: link,
	}
}

func NewButton(title string, link Link) Button {
	return Button{Title: title, Link: link}
}

// Generate Link Object
//
// 0: WebUrl, 1: MobileWebUrl
// 2: AndroidExecutionsParams, 3: IosExecutionParams
func NewLink(webUrl, mobileWebUrl,
	androidExecutionParams, iosExecutionParams string) Link {
	return Link{
		WebUrl:                 webUrl,
		MobileWebUrl:           mobileWebUrl,
		AndroidExecutionParams: androidExecutionParams,
		IosExecutionParams:     iosExecutionParams,
	}
}

func NewSocial(likeCount int, commentCount int, sharedCount int,
	viewCount int, subscriberCount int) Social {
	return Social{
		LikeCount:       likeCount,
		CommentCount:    commentCount,
		SharedCount:     sharedCount,
		ViewCount:       viewCount,
		SubscriberCount: subscriberCount,
	}
}

// Generate Commerce Object
//
// args[0]: DiscountPrice, args[1]: DiscountRate
// args[2]: FixedDiscountPrice
func NewCommerce(productName string, regularPrice int, args ...int) Commerce {
	return Commerce{
		ProductName:   productName,
		RegularPrice:  regularPrice,
		DiscountPrice: args[0], DiscountRate: args[1], FixedDiscountPrice: args[2],
	}
}

func NewTemplateObject(ObjectType ObjectType) interface{} {

	switch ObjectType {
	case FeedType:
		return FeedObject{ObjectType: "feed"}
	case ListType:
		return ListObject{ObjectType: "list"}
	case LocationType:
		return LocationObject{ObjectType: "location"}
	case ComemrceType:
		return CommerceObject{ObjectType: "commerce"}
	case TextType:
		return TextObject{ObjectType: "text"}
	}

	return nil
}

func (o *FeedObject) SetButtons(buttons ...Button) {
	o.Buttons = buttons
}

func (o *FeedObject) SetButtonTitle(title string) {
	o.ButtonTitle = title
}

func (o *FeedObject) SetContent(content Content) {
	o.Content = content
}

func (o *FeedObject) SetSocial(social Social) {
	o.Social = social
}

func (o *ListObject) SetButtons(buttons ...Button) {
	o.Buttons = buttons
}

func (o *ListObject) SetButtonTitle(title string) {
	o.ButtonTitle = title
}

func (o *ListObject) SetHeaderTitle(title string) {
	o.HeaderTitle = title
}

func (o *ListObject) SetHeaderLink(link Link) {
	o.HeaderLink = link
}

func (o *ListObject) SetContents(content []Content) {
	o.Contents = content
}

func (o *LocationObject) SetButtons(buttons ...Button) {
	o.Buttons = buttons
}

func (o *LocationObject) SetButtonTitle(title string) {
	o.ButtonTitle = title
}

func (o *LocationObject) SetAddress(address string) {
	o.Address = address
}

func (o *LocationObject) SetAddressTitle(title string) {
	o.AddressTitle = title
}

func (o *LocationObject) SetContent(content Content) {
	o.Content = content
}

func (o *LocationObject) SetSocial(social Social) {
	o.Social = social
}

func (o *CommerceObject) SetButtons(buttons ...Button) {
	o.Buttons = buttons
}

func (o *CommerceObject) SetButtonTitle(title string) {
	o.ButtonTitle = title
}

func (o *CommerceObject) SetContent(content Content) {
	o.Content = content
}

func (o *CommerceObject) SetCommerce(commerce Commerce) {
	o.Commerce = commerce
}

func (o *TextObject) SetButtons(buttons ...Button) {
	o.Buttons = buttons
}

func (o *TextObject) SetButtonTitle(title string) {
	o.ButtonTitle = title
}

func (o *TextObject) SetText(text string) {
	o.Text = text
}

func (o *TextObject) SetLink(link Link) {
	o.Link = link
}
