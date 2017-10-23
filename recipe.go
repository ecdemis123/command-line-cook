package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"

	"github.com/pkg/errors"
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

var errorResult error

func getRecipe(queryString string) (response Response, err error) {

	var r Response

	res, err := http.Get(queryString)
	if err != nil {
		return r, errors.Wrap(err, "error performing HTTP request")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, errors.Wrap(err, "error reading response body")
	}

	err = json.Unmarshal([]byte(body), &r)
	if err != nil {
		return r, errors.Wrap(err, "error unmarshaling response")
	}

	return r, err
}

func printRecipe(recipe Recipe) {
	calories := int(recipe.Calories + math.Copysign(0.5, recipe.Calories))
	fmt.Println("Name:", recipe.Label)
	fmt.Println("Yield:", recipe.Yield, "Calories:", calories)
	fmt.Println("Ingredients:")
	for _, ingredient := range recipe.Ingredients {
		fmt.Println("*", ingredient.Text)
	}
	fmt.Println("Instructions:")
	fmt.Println(recipe.Url)
	fmt.Println("--------")
}
