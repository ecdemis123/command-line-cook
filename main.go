package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

type Recipe struct {
	Q string `json:q`
}

func main() {
	viper.AutomaticEnv()
	app_id := viper.Get("edamam_app_id")
	app_key := viper.Get("edamam_app_key")
	url := "https://api.edamam.com/search"

	queryString := fmt.Sprintf("%v?q=chicken&app_id=%v&app_key=%v", url, app_id, app_key)
	res, _ := http.Get(queryString)
	body, _ := ioutil.ReadAll(res.Body)
	var r = new(Recipe)
	json.Unmarshal([]byte(body), &r)
	fmt.Println(r)
}
