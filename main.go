package main

import (
	"fmt"
	// "golang.org/x/net/html"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://en.wikipedia.org/wiki/List_of_S%26P_500_companies")
	bytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("HTML:\n\n", string(bytes))

	resp.Body.Close()
}
