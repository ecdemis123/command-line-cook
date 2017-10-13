package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
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

var (
	queryParam string
)

func init() {
	flag.StringVar(&queryParam, "queryParam", "chicken", "High-level query param you are looking for.")
}

func main() {
	viper.AutomaticEnv()
	app_id := viper.Get("edamam_app_id").(string)
	app_key := viper.Get("edamam_app_key").(string)

	flag.Parse()

	values := url.Values{
		"app_id":  {app_id},
		"app_key": {app_key},
		"q":       {queryParam},
		"from":    {"0"},
		"to":      {"1"},
	}

	u, _ := url.Parse("https://api.edamam.com/search")
	u.RawQuery = values.Encode()
	res, _ := http.Get(u.String())
	body, _ := ioutil.ReadAll(res.Body)
	var r = new(Response)
	json.Unmarshal([]byte(body), &r)
	fmt.Println(r)
}
