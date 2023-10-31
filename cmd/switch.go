/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"geometry_profiles/utils"
	"os"

	cp "github.com/otiai10/copy"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch profiles",
	Long:  `Switches profiles`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Please only specify one profile name")
		}
		if len(args) < 1 {
			fmt.Println("Please specify profile name")
		}

		if viper.GetString("GeometryData") == "" {
			fmt.Println("Please set geometry dash data folder")
			return
		}

		profiles, err := utils.LoadProfiles()
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(profiles) < 2 {
			fmt.Println("At least to profiles are needed before switching")
		}

		if _, err := os.Stat(viper.GetString("ProfileDir") + "/" + utils.HashName(args[0])); os.IsNotExist(err) {
			fmt.Println("Profile with that name not found...")
			return
		}

		bytearr, err := os.ReadFile(viper.GetString("ProfileDir") + "/" + "profiles.json")
		if err != nil {
			fmt.Println(err)
			return
		}

		var dotProfiles utils.DotProfile
		if err := json.Unmarshal(bytearr, &dotProfiles); err != nil {
			fmt.Println(err)
			return
		}

		if dotProfiles.Current == utils.HashName(args[0]) {
			fmt.Println("Cannot switch to currently active profile")
			return
		}

		cp.Copy(viper.GetString("GeometryData"), viper.GetString("ProfileDir")+"/"+dotProfiles.Current)
		os.RemoveAll(viper.GetString("GeometryData"))
		os.MkdirAll(viper.GetString("GeometryData"), 0774)
		cp.Copy(viper.GetString("ProfileDir")+"/"+utils.HashName(args[0]), viper.GetString("GeometryData"))
		dotProfiles.Current = utils.HashName(args[0])

		bytearr, err = json.Marshal(dotProfiles)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err = os.WriteFile(viper.GetString("ProfileDir")+"/"+"profiles.json", bytearr, 0744); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// switchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// switchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
