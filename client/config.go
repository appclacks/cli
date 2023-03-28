package client

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type ConfigFileProfile struct {
	OrganizationID string `yaml:"organization-id"`
	APIToken       string `yaml:"api-token"`
}

type ConfigFileProfiles struct {
	DefaultProfile string                       `yaml:"default-profile"`
	Profiles       map[string]ConfigFileProfile `yaml:"profiles"`
}

func GetConfigDirPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return path.Join(configDir, "appclacks"), nil

}

func GetConfigFilePath() (string, error) {
	configDir, err := GetConfigDirPath()
	if err != nil {
		return "", err
	}
	return path.Join(configDir, "appclacks.yaml"), err
}

func ReadConfig(path string) (ConfigFileProfiles, error) {
	configFile, err := os.ReadFile(path)
	if err != nil {
		return ConfigFileProfiles{}, err
	}
	var profiles ConfigFileProfiles
	err = yaml.Unmarshal(configFile, &profiles)
	if err != nil {
		return ConfigFileProfiles{}, fmt.Errorf("Invalid YAML file at %s: %w", path, err)
	}
	if profiles.Profiles == nil {
		profiles.Profiles = make(map[string]ConfigFileProfile)
	}
	return profiles, nil
}
