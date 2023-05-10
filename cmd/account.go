package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/appclacks/cli/client"
	apitypes "github.com/appclacks/go-types"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

func changePasswordCmd() *cobra.Command {
	var changePassword = &cobra.Command{
		Use:   "change",
		Short: "Change your password",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("\n* Appclacks Email:\n> ")
			email, err := reader.ReadString('\n')
			exitIfError(err)
			email = strings.TrimSpace(email)
			fmt.Printf("\n* Appclacks current Password:\n> ")
			password, err := reader.ReadString('\n')
			fmt.Println("")
			exitIfError(err)
			password = strings.TrimSpace(password)
			fmt.Printf("\n* Appclacks new Password:\n> ")
			newPassword, err := reader.ReadString('\n')
			fmt.Println("")
			exitIfError(err)
			newPassword = strings.TrimSpace(newPassword)
			cliClient, err := client.New(appclacksURL, client.WithUserPassword(client.AccountEmail(email), client.AccountPassword(password)))
			exitIfError(err)
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()

			payload := apitypes.ChangeAccountPasswordInput{
				NewPassword: newPassword,
			}
			result, err := cliClient.ChangeAccountPassword(ctx, payload)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(result)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			t := tabby.New()
			t.AddHeader("Messages")
			for _, message := range result.Messages {
				t.AddLine(message)
			}
			t.Print()
			os.Exit(0)
		},
	}
	return changePassword
}

func getAccountOrganizationCmd() *cobra.Command {
	var getAccountOrganization = &cobra.Command{
		Use:   "organization",
		Short: "Get the organization for the configured account (email and password)",
		Run: func(cmd *cobra.Command, args []string) {
			client := buildClient()
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			result, err := client.GetOrganizationForAccount(ctx)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(result)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			t := tabby.New()
			t.AddHeader("ID", "Name", "Description")
			desc := ""
			if result.Description != nil {
				desc = *result.Description
			}
			t.AddLine(result.ID, result.Name, desc)
			t.Print()
			os.Exit(0)
		},
	}

	return getAccountOrganization
}
