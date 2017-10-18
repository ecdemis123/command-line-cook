package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/spf13/viper"
)

var (
	queryParam  string
	queryString string
)

func init() {
	flag.StringVar(&queryParam, "queryParam", "chicken", "High-level query param you are looking for.")
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
		"q":       {queryParam},
		"from":    {"0"},
		"to":      {"1"},
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
	fmt.Println(r)
}
