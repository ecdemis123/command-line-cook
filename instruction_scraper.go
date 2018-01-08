package main

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

func scrapeInstructions(url string) (instructions string, err error) {
	var buffer bytes.Buffer

	doc, err := goquery.NewDocument(url)

	if err != nil {
		return instructions, err
	}
	doc.Find(".instruction .instructions").Each(func(_ int, s *goquery.Selection) {
		text := s.Text()
		buffer.WriteString(text)
	})
	return buffer.String(), err
}
