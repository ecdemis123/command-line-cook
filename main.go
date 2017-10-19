package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/url"
	"time"

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

	// api does not return a random result
	seed := rand.NewSource(time.Now().UnixNano())
	rn := rand.New(seed)
	randomIndex := rn.Intn(100)

	recipe := r.Hits[randomIndex].Recipe
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
