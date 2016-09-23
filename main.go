package main

import (
	"fmt"
	"os"

	"github.com/ovh/tat"
	"github.com/spf13/viper"
)

// taturl, username / password of tat engine
var (
	taturl   string
	username string
	password string
)

/*
Usage:
 go get -u github.com/ovh/tat
 build && ./mycli-minimal -url=http://url-tat-engine -username=<tatUsername> -password=<tatPassword> /Internal/your/topic your message
*/

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	taturl = viper.GetString("taturl")
	username = viper.GetString("username")
	password = viper.GetString("password")

	client, err := tat.NewClient(tat.Options{
		URL:      taturl,
		Username: username,
		Password: password,
		Referer:  "main.v0",
	})

	if err != nil {
		fmt.Printf("Error while create new Tat Client: %s\n", err)
		os.Exit(1)
	}

	stats, err := client.StatsInstance()

	if err != nil {
		fmt.Printf("Error:%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("instances: %s\n", stats)

}
