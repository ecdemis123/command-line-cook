package main

import (
	"net/http"

	"golang.org/x/net/html"
)

func scrapeInstructions(url string) (instructions string, err error) {

	res, err := http.Get(url)
	if err != nil {
	}

	b := res.Body
	defer b.Close()

	tokenizer := html.NewTokenizer(b)

	for {
		tt := tokenizer.Next()
		switch {
		case tt == html.ErrorToken:
			return instructions, err
		case tt == html.StartTagToken:
			t := tokenizer.Token()
			if t.Data != "ul" {
				continue
			}
		}
	}
}
