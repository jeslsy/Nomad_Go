package main

import (
	"errors"
	"fmt"
	"net/http"
)

// 지금부터 URLchecker를 채널로 보낸대
// 진짜 최고 빠름.

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.co.kr/",
		"https://www.google.co.kr/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.instagram.com/",
	}

	//각 URL을 hitURL()한걸 채널에 넣어줌
	for _, url := range urls {
		go hitURL(url, c)
	}

	//채널에서 result라는 곳에 값(statusCode) 받아옴
	for i := 0; i < len(urls); i++ {
		result := <-c
		// results 맵에 값 저장
		results[result.url] = result.status
	}

	// URL이랑 statusCode 값 출력
	for url, status := range results {
		fmt.Println(url, status)
	}
}

// chan <- result 하면 sending only라는걸 보여주는 것!
func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
