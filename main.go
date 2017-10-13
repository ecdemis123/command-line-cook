package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	Yield       int          `json:yield`
	Label       string       `json:label`
	Calories    float64      `json:calories`
	Ingredients []Ingredient `json:ingredients`
}

type Ingredient struct {
	Text   string `json:text`
	Weight string `json:weight`
}

func main() {
	viper.AutomaticEnv()
	app_id := viper.Get("edamam_app_id")
	app_key := viper.Get("edamam_app_key")
	url := "https://api.edamam.com/search"
	queryString := fmt.Sprintf("%v?q=chicken&app_id=%v&app_key=%v&from=0&to=1", url, app_id, app_key)
	res, _ := http.Get(queryString)
	body, _ := ioutil.ReadAll(res.Body)
	var r = new(Response)
	json.Unmarshal([]byte(body), &r)
	fmt.Println(r)
}
