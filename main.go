package main

import (
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
)

func getFirstElementByClassName(className string, n *html.Node) (element *html.Node, ok bool) {
	for _, a := range n.Attr {
		if a.Key == "class" && strings.Contains(a.Val, className) {
			return n, true
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if element, ok = getFirstElementByClassName(className, c); ok {
			return
		}
	}
	return
}

func main() {
	response, err := http.Get("https://en.wikipedia.org/wiki/List_of_S%26P_500_companies")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	root, err := html.Parse(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	el, ok := getFirstElementByClassName("wikitable", root)
	if !ok {
		log.Fatal(err)
	}

	log.Printf("%+v", el)
}
