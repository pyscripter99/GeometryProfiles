package utils

import (
	"os"

	"github.com/spf13/viper"
)

func LoadProfiles() ([]string, error) {
	files, err := os.ReadDir(viper.GetString("ProfileDir"))
	if err != nil {
		return []string{}, err
	}

	profiles := []string{}

	for _, f := range files {
		if f.Name() == "profiles.json" {
			continue
		}

		profiles = append(profiles, f.Name())
	}

	return profiles, nil
}
