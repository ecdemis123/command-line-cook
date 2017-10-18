package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var res = &Response{
	Q: "chicken",
}

func TestGetRecipe(t *testing.T) {
	res2B, _ := json.Marshal(res)
	var edamamStub = httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(res2B))
			}))

	res := getRecipe(edamamStub.URL)
	fmt.Println(res)
}
