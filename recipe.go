package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Q    string `json:q`
	Hits []Hit  `json:hits`
}

type Hit struct {
	Recipe Recipe `json:recipe`
}

type Recipe struct {
	Yield       float64      `json:yield`
	Label       string       `json:label`
	Calories    float64      `json:calories`
	Url         string       `json:url`
	Ingredients []Ingredient `json:ingredients`
}

type Ingredient struct {
	Text   string  `json:text`
	Weight float64 `json:weight`
}

func getRecipe(queryString string) Response {

	res, err := http.Get(queryString)

	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	var r Response
	json.Unmarshal([]byte(body), &r)

	return r
}
