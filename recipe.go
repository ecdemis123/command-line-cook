package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"

	"github.com/pkg/errors"
)

// Response is the main data from edamamAPI
type Response struct {
	Q    string `json:q`
	Hits []Hit  `json:hits`
}

// Hit is the sub response from edamamAPI
type Hit struct {
	Recipe Recipe `json:recipe`
}

// Recipe is the indiviual recipe
type Recipe struct {
	Yield       float64      `json:yield`
	Label       string       `json:label`
	Calories    float64      `json:calories`
	URL         string       `json:url`
	Ingredients []Ingredient `json:ingredients`
}

// Ingredient is the individual ingredient
type Ingredient struct {
	Text   string  `json:text`
	Weight float64 `json:weight`
}

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

func printRecipe(recipe Recipe, instructions string) {
	calories := int(recipe.Calories + math.Copysign(0.5, recipe.Calories))
	fmt.Println("Name:", recipe.Label)
	fmt.Println("Yield:", recipe.Yield, "Calories:", calories)
	fmt.Println("Ingredients:")
	for _, ingredient := range recipe.Ingredients {
		fmt.Println("*", ingredient.Text)
	}
	fmt.Println("Instructions:")
	fmt.Println(instructions)
	fmt.Println(recipe.URL)
	fmt.Println("--------")
}
