package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Cant Request")

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.naver.com/",
		"https://www.google.co.kr/",
		"https://soundcloud.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://www.youtube.com/",
		"https://www.amazon.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	// waiting message(channel)
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

// data send only
func hitURL(url string, c chan<- requestResult) error {
	res, err := http.Get(url)
	status := "OK"

	if err != nil || res.StatusCode >= 400 {
		c <- requestResult{url: url, status: "FAILED"}
	}
	c <- requestResult{url: url, status: status}

	return nil
}
