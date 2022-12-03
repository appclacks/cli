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

func createOrganizationCmd(client *client.Client) *cobra.Command {
	var orgName string
	var orgDescription string
	var accountFirstName string
	var accountLastName string
	var accountPassword string
	var accountEmail string

	var createOrganization = &cobra.Command{
		Use:   "create",
		Short: "Create an organization",
		Run: func(cmd *cobra.Command, args []string) {
			payload := apitypes.CreateOrganizationInput{
				Organization: apitypes.CreateOrganizationOrg{
					Name:        orgName,
					Description: orgDescription,
				},
				Account: apitypes.CreateOrganizationAccount{
					FirstName: accountFirstName,
					LastName:  accountLastName,
					Password:  accountPassword,
					Email:     accountEmail,
				},
			}
			result, err := client.CreateOrganization(payload)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(result)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			t := tabby.New()
			t.AddHeader("Organization ID", "Account ID")
			t.AddLine(result.Organization.ID, result.Account.ID)
			t.Print()
			os.Exit(0)
		},
	}
	createOrganization.PersistentFlags().StringVar(&orgName, "organization-name", "", "Organization name")
	err := createOrganization.MarkPersistentFlagRequired("organization-name")
	exitIfError(err)

	createOrganization.PersistentFlags().StringVar(&orgDescription, "organization-description", "", "Organization description")
	createOrganization.PersistentFlags().StringVar(&accountFirstName, "account-first-name", "", "account first name")
	err = createOrganization.MarkPersistentFlagRequired("account-first-name")
	exitIfError(err)

	createOrganization.PersistentFlags().StringVar(&accountLastName, "account-last-name", "", "account last name")
	err = createOrganization.MarkPersistentFlagRequired("account-last-name")
	exitIfError(err)

	createOrganization.PersistentFlags().StringVar(&accountPassword, "account-password", "", "account password")
	err = createOrganization.MarkPersistentFlagRequired("account-password")
	exitIfError(err)

	createOrganization.PersistentFlags().StringVar(&accountEmail, "account-email", "", "account email")
	err = createOrganization.MarkPersistentFlagRequired("account-email")
	exitIfError(err)

	return createOrganization
}
