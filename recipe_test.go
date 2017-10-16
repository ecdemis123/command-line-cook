package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRecipe(t *testing.T) {
	var edamamStub = httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

	res := getRecipe(edamamStub.URL)
	fmt.Println(res)
}
