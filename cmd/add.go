/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"geometry_profiles/utils"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a profile",
	Long:  `Creates a geometry dash profile`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Only one profile can be created")
			return
		} else if len(args) < 1 {
			fmt.Println("Please specify a name for the profile")
			return
		}

		os.MkdirAll(viper.GetString("ProfileDir")+"/"+utils.HashName(args[0]), 0744)

		if _, err := os.Stat(viper.GetString("ProfileDir") + "/" + "profiles.json"); os.IsNotExist(err) {
			bytearr, err := json.Marshal(&utils.DotProfile{
				Current: utils.HashName(args[0]),
				Profiles: []utils.Profile{
					{Name: args[0], Hash: utils.HashName(args[0])},
				},
			})
			if err != nil {
				fmt.Println(err)
				return
			}
			os.WriteFile(viper.GetString("ProfileDir")+"/"+"profiles.json", bytearr, 0744)
		} else {
			bytearr, err := os.ReadFile(viper.GetString("ProfileDir") + "/" + "profiles.json")
			if err != nil {
				fmt.Println(err)
				return
			}

			var profiles utils.DotProfile
			if err := json.Unmarshal(bytearr, &profiles); err != nil {
				fmt.Println(err)
				return
			}

			profiles.Profiles = append(profiles.Profiles, utils.Profile{
				Name: args[0],
				Hash: utils.HashName(args[0]),
			})

			bytearr, err = json.Marshal(profiles)
			if err != nil {
				fmt.Println(err)
				return
			}

			if err = os.WriteFile(viper.GetString("ProfileDir")+"/"+"profiles.json", bytearr, 0744); err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
