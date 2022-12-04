package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/appclacks/cli/client"
	apitypes "github.com/appclacks/go-types"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

func createAPITokenCmd(client *client.Client) *cobra.Command {
	var name string
	var description string
	var ttl string
	var permissions []string

	var createAPIToken = &cobra.Command{
		Use:   "create",
		Short: "Create an API token",
		Run: func(cmd *cobra.Command, args []string) {
			payload := apitypes.CreateAPITokenInput{
				Name:        name,
				Description: description,
				TTL:         ttl,
				Permissions: apitypes.Permissions{
					Actions: permissions,
				},
			}
			result, err := client.CreateAPIToken(payload)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(result)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			t := tabby.New()
			t.AddHeader("ID", "Name", "Token", "Description", "TTL", "Permissions")
			t.AddLine(result.ID, result.Name, result.Token, result.Description, result.TTL, result.Permissions.Actions)
			t.Print()
			os.Exit(0)
		},
	}
	createAPIToken.PersistentFlags().StringVar(&name, "name", "", "Token name")
	err := createAPIToken.MarkPersistentFlagRequired("name")
	exitIfError(err)
	createAPIToken.PersistentFlags().StringVar(&description, "description", "", "Token description")
	createAPIToken.PersistentFlags().StringVar(&ttl, "ttl", "72h", "Token TTL")
	createAPIToken.PersistentFlags().StringSliceVar(&permissions, "permission", []string{"*"}, "Attach a permission to this token.")

	return createAPIToken
}

func listAPITokensCmd(client *client.Client) *cobra.Command {
	var listAPITokens = &cobra.Command{
		Use:   "list",
		Short: "List API tokens",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := client.ListAPITokens()
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(result)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			t := tabby.New()
			t.AddHeader("ID", "Name", "Description", "TTL", "Create At", "Permissions")
			for _, token := range result.Result {
				t.AddLine(token.ID, token.Name, token.Description, token.TTL, token.CreatedAt, token.Permissions.Actions)
			}

			t.Print()
			os.Exit(0)
		},
	}

	return listAPITokens
}

func getAPITokenCmd(client *client.Client) *cobra.Command {
	var tokenID string
	var getAPIToken = &cobra.Command{
		Use:   "get",
		Short: "Get an API token",
		Run: func(cmd *cobra.Command, args []string) {
			input := apitypes.GetAPITokenInput{
				ID: tokenID,
			}
			token, err := client.GetAPIToken(input)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(token)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			t := tabby.New()
			t.AddHeader("ID", "Name", "Description", "TTL", "Create At", "Permissions")
			t.AddLine(token.ID, token.Name, token.Description, token.TTL, token.CreatedAt, token.Permissions.Actions)

			t.Print()
			os.Exit(0)
		},
	}
	getAPIToken.PersistentFlags().StringVar(&tokenID, "id", "", "Token ID")
	err := getAPIToken.MarkPersistentFlagRequired("id")
	exitIfError(err)
	return getAPIToken
}

func deleteAPITokenCmd(client *client.Client) *cobra.Command {
	var tokenID string
	var deleteAPIToken = &cobra.Command{
		Use:   "delete",
		Short: "Delete an API token",
		Run: func(cmd *cobra.Command, args []string) {
			input := apitypes.DeleteAPITokenInput{
				ID: tokenID,
			}
			result, err := client.DeleteAPIToken(input)
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
	deleteAPIToken.PersistentFlags().StringVar(&tokenID, "id", "", "Token ID")
	err := deleteAPIToken.MarkPersistentFlagRequired("id")
	exitIfError(err)
	return deleteAPIToken
}
