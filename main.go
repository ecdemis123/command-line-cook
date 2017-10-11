package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	key := viper.Get("api_key")
	fmt.Println(key)
}
