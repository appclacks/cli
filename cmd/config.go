package cmd

import (
	"fmt"
	"os"
	"path"
	"gopkg.in/yaml.v3"
)



type ConnexionInfos struct {
	TypeAuth string // "token" or "password" or "noAuth"
	identifier string 
	secret string
	baseURL string
	profile string
}

func (ci *ConnexionInfos) setPasswordAuth(email string, password string) {
	ci.TypeAuth = "password"
	ci.identifier = email
	ci.secret = password
}

func (ci *ConnexionInfos) setTokenAuth(orgID string, token string) {
	ci.TypeAuth = "token"
	ci.identifier = orgID
	ci.secret = token
}

func (ci *ConnexionInfos) setBaseURL(baseURL string){
	ci.baseURL = baseURL 
}

func CreateConnInfos() (*ConnexionInfos){
	return &ConnexionInfos{
		TypeAuth: "noAuth",
		identifier: "",
		secret: "",
		baseURL: "",
		profile: "",
	}
}

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

func loadConnInfosFromEnv(connInfos *ConnexionInfos) {

	if os.Getenv("APPCLACKS_API_ENDPOINT") != "" {
		connInfos.setBaseURL(os.Getenv("APPCLACKS_API_ENDPOINT"))
	}

	if os.Getenv("APPCLACKS_ORGANIZATION_ID") != "" && os.Getenv("APPCLACKS_TOKEN") != "" {
		connInfos.setTokenAuth(os.Getenv("APPCLACKS_ORGANIZATION_ID"), os.Getenv("APPCLACKS_TOKEN"))
		return
	}

	if os.Getenv("APPCLACKS_ACCOUNT_EMAIL") != "" && os.Getenv("APPCLACKS_ACCOUNT_PASSWORD") != ""{
		connInfos.setPasswordAuth(os.Getenv("APPCLACKS_ACCOUNT_EMAIL"), os.Getenv("APPCLACKS_ACCOUNT_PASSWORD"))
		return
	}
	
}

func loadConnInfosFromConfigFile(connInfos *ConnexionInfos) (error) {
	configPath, err := GetConfigFilePath()
	if err != nil {
		return err
	}
	config, err := ReadConfig(configPath)
	if err != nil {
		return err
	}
	if len(config.Profiles) == 0 {
		return fmt.Errorf("The configuration file %s is empty", configPath)
	}
	if connInfos.profile == "" {
		connInfos.profile = config.DefaultProfile
	}
	if connInfos.profile == "" {
		return fmt.Errorf("No profile selected and no default profile specified in the configuration file %s", configPath)
	}
	configProfile, ok := config.Profiles[connInfos.profile]
	if !ok {
		return fmt.Errorf("Profile %s not found in the configuration file %s", connInfos.profile, configPath)
	}
	
	if configProfile.APIToken == "" || configProfile.OrganizationID == "" {
		return fmt.Errorf("Profile %s  in the configuration file %s is not complete. APIToken or OrganizationID is missing", connInfos.profile, configPath)
	}

	connInfos.setTokenAuth(configProfile.OrganizationID, configProfile.APIToken)

	return nil
}

