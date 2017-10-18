package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var res = &Response{
	Q: "chicken",
	Hits: []Hit{
		Hit{
			Recipe{
				Yield:    6,
				Label:    "healthy",
				Calories: 1300,
				Url:      "chicken_recipe.com",
				Ingredients: []Ingredient{
					Ingredient{
						Text:   "4 c chicken",
						Weight: 40,
					},
				},
			},
		},
	},
}

func TestGetRecipe(t *testing.T) {
	res2B, _ := json.Marshal(res)
	var edamamStub = httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(res2B))
			}))

	_, err := getRecipe(edamamStub.URL)
	if err != nil {
		t.Errorf("Error getting and parsing API data")
	}
}
