package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
	"gopkg.in/yaml.v3"

	apitypes "github.com/appclacks/go-types"
	"github.com/spf13/cobra"
)

func loginCmd() *cobra.Command {
	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Log in to the Appclacks Cloud platform",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			appclacksConfigDir, err := GetConfigDirPath()
			exitIfError(err)
			err = os.MkdirAll(appclacksConfigDir, os.ModePerm)
			exitIfError(err)
			filepath, err := GetConfigFilePath()
			exitIfError(err)

			fp, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0600)
			exitIfError(err)
			fp.Close()

			profiles, err := ReadConfig(filepath)
			exitIfError(err)

			fmt.Printf("The new profile will be written into the file located at: %s\n\n", filepath)
			fmt.Printf("* Profile name:\n> ")
			profileName, err := reader.ReadString('\n')
			exitIfError(err)
			profileName = strings.TrimSpace(profileName)
			fmt.Printf("\n* Appclacks Email:\n> ")
			email, err := reader.ReadString('\n')
			exitIfError(err)
			email = strings.TrimSpace(email)
			fmt.Printf("\n* Appclacks Password:\n> ")
			bytePassword, err := term.ReadPassword(int(syscall.Stdin))
			exitIfError(err)
			fmt.Println("")
			password := strings.TrimSpace(string(bytePassword))

			cliClient := buildClientByPassword(email, password)

			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)

			org, err := cliClient.GetOrganizationForAccountWithContext(ctx)
			
			exitIfError(err)
			cancel()

			ctx, cancel = context.WithTimeout(context.Background(), defaultTimeout)

			tokenName := fmt.Sprintf("%s-cli", profileName)
			payload := apitypes.CreateAPITokenInput{
				Name:        tokenName,
				Description: fmt.Sprintf("Token generated using the Appclacks CLI login command for account %s", email),
				TTL:         "1051200h",
				Permissions: apitypes.Permissions{
					Actions: []string{"*"},
				},
			}
			apiToken, err := cliClient.CreateAPITokenWithContext(ctx, payload)
			exitIfError(err)
			cancel()

			profile := ConfigFileProfile{
				OrganizationID: org.ID,
				APIToken:       apiToken.Token,
			}
			if profiles.Profiles == nil {
				profiles.Profiles = make(map[string]ConfigFileProfile)
			}

			profiles.Profiles[profileName] = profile
			if profiles.DefaultProfile == "" {
				profiles.DefaultProfile = profileName
			}
			configFileData, err := yaml.Marshal(&profiles)
			exitIfError(err)

			err = os.WriteFile(filepath, configFileData, 0)
			exitIfError(err)

			fmt.Printf("Profile %s successfully created. You're now ready to use the CLI !\n", profileName)
		},
	}

	return loginCmd
}
