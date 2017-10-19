package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"

	multierror "github.com/hashicorp/go-multierror"
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

	res, err := http.Get(queryString)

	if err != nil {
		multierror.Append(errorResult, err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		multierror.Append(errorResult, err)
	}

	var r Response

	json.Unmarshal([]byte(body), &r)

	return r, errorResult
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
