package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ovh/tat"
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
	flag.StringVar(&taturl, "url", "http://127.0.0.1:8080", "URL of Tat Engine")
	flag.StringVar(&username, "username", "tat", "tat username")
	flag.StringVar(&password, "password", "b43916450754f993ae2a180cf7748ccc509dc87e7d2008d6d7ff143404c274e013b92188a8da29626bd4a705fe7fe3a9a99eeaa3a5795c702c1ac549ef7c8132", "tat password")
	flag.Parse()

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
