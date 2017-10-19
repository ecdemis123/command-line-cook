package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"net/url"

	"github.com/spf13/viper"
)

var (
	search string
)

func init() {
	flag.StringVar(&search, "search", "chicken", "High-level query param you are looking for.")
}

func main() {
	viper.AutomaticEnv()
	app_id := viper.Get("edamam_app_id").(string)
	app_key := viper.Get("edamam_app_key").(string)

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.Parse()

	values := url.Values{
		"app_id":  {app_id},
		"app_key": {app_key},
		"q":       {search},
		"from":    {"0"},
		"to":      {"100"},
	}

	u, err := url.Parse("https://api.edamam.com/search")

	if err != nil {
		log.Fatalf("Error parsing url: %s\n", err)
	}

	u.RawQuery = values.Encode()
	queryString := u.String()

	r, err := getRecipe(queryString)
	if err != nil {
		log.Fatalf("Error getting recipe data: %s\n", err)
	}

	for i, _ := range r.Hits {
		recipe := r.Hits[i].Recipe
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
}
