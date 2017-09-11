package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"golang.org/x/net/html"
)

func main() {
	resp, _ := http.Get("http://www.daum.net")
	bytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("HTML:\n\n", string(bytes))

	resp.Body.Close()
}