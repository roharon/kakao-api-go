# Kakao REST API wrapper for Golang
![Build status](https://github.com/roharon/kakao-api-go/workflows/build-test/badge.svg)


This is a wrapper library for [Kakao API](https://developers.kakao.com/docs) written in Golang.
Based on [meinside/kakao-api-go](https://github.com/meinside/kakao-api-go)

## How to get

```bash
$ go get -u github.com/roharon/kakao-api-go
```

## How to use

```go
// sample
package main

import (
	kakaoapi "github.com/roharon/kakao-api-go"
)

const (
	apiKey = "0123456789abcdefghijklmnopqrstuvwxyz" // XXX - change this to yours
)

func main() {
	client := kakaoapi.NewClient(apiKey)
	//client.Verbose = true

	// TODO - do something with `client`
	// ...
}
```

See the [samples here](https://github.com/roharon/kakao-api-go/tree/master/samples).

## API coverages

- [ ] [KakaoLogin](https://developers.kakao.com/docs/latest/ko/kakaologin/rest-api)
- [X] [KakaotalkSocial](https://developers.kakao.com/docs/latest/ko/kakaotalk-social/rest-api), last update: 2021.02.
- [ ] [Message](https://developers.kakao.com/docs/latest/ko/message/rest-api)
- [ ] [KakaoStory](https://developers.kakao.com/docs/latest/ko/kakaostory/rest-api)
- [ ] [KakaoTalkChannel](https://developers.kakao.com/docs/latest/ko/kakaotalk-channel/rest-api)
- [ ] [KakaoSync](https://developers.kakao.com/docs/latest/ko/kakaosync/common)
- [ ] [KakaoPay](https://developers.kakao.com/docs/latest/ko/kakaopay/common)
- [ ] [KakaoMoment](https://developers.kakao.com/docs/latest/ko/kakaomoment/common)
- [ ] [PushNotification](https://developers.kakao.com/docs/latest/ko/push/rest-api)
- [ ] [DaumSearch](https://developers.kakao.com/docs/latest/ko/daum-search/dev-guide)
- [ ] [Map](https://developers.kakao.com/docs/latest/ko/kakaomap/common)
- [ ] [Local](https://developers.kakao.com/docs/latest/ko/local/dev-guide)
- [ ] [KakaoNavi](https://developers.kakao.com/docs/latest/ko/kakaonavi/common)
- [X] [Vision](https://developers.kakao.com/docs/latest/ko/vision/dev-guide), last update: 2020.08.
- [X] [Pose](https://developers.kakao.com/docs/latest/ko/pose/dev-guide), last update: 2020.06.
- [X] [Translation](https://developers.kakao.com/docs/latest/ko/translate/dev-guide), last update: 2020.08.
- [X] [Speech](https://developers.kakao.com/docs/latest/ko/voice/rest-api), last update: 2020.06.


## License

MIT

