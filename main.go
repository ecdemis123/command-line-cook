package main

import (
	"flag"
	"log"
	"math/rand"
	"net/url"
	"strconv"
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
	recipeCount := 100

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.Parse()

	values := url.Values{
		"app_id":  {app_id},
		"app_key": {app_key},
		"q":       {search},
		"from":    {"0"},
		"to":      {strconv.Itoa(recipeCount)},
	}

	u, _ := url.Parse("https://api.edamam.com/search")

	u.RawQuery = values.Encode()
	queryString := u.String()

	r, err := getRecipe(queryString)
	if err != nil {
		log.Fatalf("Error getting recipe data: %s\n", err)
	}

	seed := rand.NewSource(time.Now().UnixNano())
	rn := rand.New(seed)
	randomIndex := rn.Intn(recipeCount)
	recipe := r.Hits[randomIndex].Recipe

	printRecipe(recipe)
}
