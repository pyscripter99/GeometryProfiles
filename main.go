/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"geometry_profiles/cmd"
	"os"

	"github.com/spf13/viper"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	viper.SetDefault("ProfileDir", home+"/.GeometryProfiles/")
	viper.SetDefault("GeometryData", "\"\" # Please enter geometry dash AppData folder. e.g. AppData local")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(home + "/.config/GeometryProfiles/")

	if _, err := os.Stat("~/.GeometryProfiles/"); os.IsNotExist(err) {
		os.MkdirAll(home+"/.GeometryProfiles/", 0744)
	}

	if _, err := os.Stat(home + "/.config/GeometryProfiles/"); os.IsNotExist(err) {
		os.MkdirAll(home+"/.config/GeometryProfiles/", 0744)
		os.WriteFile(home+"/.config/GeometryProfiles/config.yaml", []byte{}, 0744)
		viper.WriteConfigAs(home + "/.config/GeometryProfiles/config.yaml")
	}

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	cmd.Execute()
}
