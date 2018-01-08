package main

import "testing"

func TestScrapeInstructions(t *testing.T) {
	URL := "http://www.saveur.com/classic-chicken-marsala-recipe"
	_, err := scrapeInstructions(URL)
	if err != nil {
		t.Errorf("error retrieving instructions")
	}
}
