package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	app_id := viper.Get("edamam_app_id")
	app_key := viper.Get("edamam_app_key")
	fmt.Println(app_id)
	fmt.Println(app_key)
}
