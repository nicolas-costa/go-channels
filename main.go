package main

import (
	"github.com/levigross/grequests"
)

func main() {
	c := make(chan *grequests.Response)

	go performRequest(c)

	for response := range c {
		responseStruct := response
		print(responseStruct.StatusCode)
	}
}

func performRequest(c chan *grequests.Response) {
	httpResponse, err := grequests.Get("https://httpbin.org/get", nil)

	if err != nil {
		print("error: ", err.Error())
	}

	c <- httpResponse
	close(c)
}
