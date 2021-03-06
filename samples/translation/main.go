package main

import (
	"encoding/json"
	"fmt"
	"log"

	kakaoapi "github.com/roharon/kakao-api-go"
)

const (
	apiKey = "0123456789abcdefghijklmnopqrstuvwxyz" // XXX - change this to your REST API key

	koreanText = "테스트 좀 합시다."
)

func main() {
	client := kakaoapi.NewClient(apiKey)
	//client.Verbose = true

	// translate text
	if translated, err := client.TranslateText(koreanText, kakaoapi.LanguageKorean, kakaoapi.LanguageEnglish); err == nil {
		log.Printf("Translated: %s", prettify(translated))
	} else {
		log.Printf("Failed to translate text: %s", err)
	}
	if translated, err := client.TranslateText(koreanText, kakaoapi.LanguageKorean, kakaoapi.LanguageChinese); err == nil {
		log.Printf("Translated: %s", prettify(translated))
	} else {
		log.Printf("Failed to translate text: %s", err)
	}
	if translated, err := client.TranslateText(koreanText, kakaoapi.LanguageKorean, kakaoapi.LanguageJapanese); err == nil {
		log.Printf("Translated: %s", prettify(translated))
	} else {
		log.Printf("Failed to translate text: %s", err)
	}

	// detect language
	if detected, err := client.DetectLanguage(koreanText); err == nil {
		log.Printf("Detected languages: %s", prettify(detected))
	} else {
		log.Printf("Failed to detect language from text: %s", err)
	}
}

func prettify(obj interface{}) string {
	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		return string(bytes)
	}

	return fmt.Sprintf("%v", obj)
}
