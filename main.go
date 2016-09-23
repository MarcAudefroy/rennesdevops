package main

import (
	"fmt"
	"os"

	"github.com/ovh/tat"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	var mainCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	mainCmd.AddCommand(statsCmd)

	mainCmd.Execute()
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Stats Instances",
	Run: func(cmd *cobra.Command, args []string) {

		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.SetConfigType("json")

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		taturl := viper.GetString("taturl")
		username := viper.GetString("username")
		password := viper.GetString("password")

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
	},
}
