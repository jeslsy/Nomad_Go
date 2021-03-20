package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	//map 타입으로 results 정의
	var results = make(map[string]string)

	urls := []string{
		"https://www.airbnb.co.kr/",
		"https://www.google.co.kr/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.instagram.com/",
	}

	results["gello"] = "Hello"

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil { // err있으면
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

// hit 못하면 에러를 반환해야 겠죠 ?
func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)

	// 상태 코드 400부터는 페이지 오류니까..!
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
