package main

import "testing"

func TestScrapeInstructions(t *testing.T) {
	URL := "http://www.saveur.com/classic-chicken-marsala-recipe"
	scrapeInstructions(URL)
}
