package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func scrapeInstructions(url string) (instructions string, err error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return instructions, err
	}
	doc.Find(".instruction").Each(func(_ int, s *goquery.Selection) {
		text := s.Text()
		fmt.Println(text)
	})
	return instructions, err
}
