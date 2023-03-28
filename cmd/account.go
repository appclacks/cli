package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

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
