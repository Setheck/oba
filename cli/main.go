package main

import (
	"fmt"

	"github.com/Setheck/oba/cli/cmd"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.obacli")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
	}
}

func main() {
	cmd.Execute()
}
