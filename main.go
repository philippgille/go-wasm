package main

import (
	"io"
	"net/http"
)

func main() {
	println("Hello World")
}

func MakeHTTPRequest() {
	res, err := http.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
