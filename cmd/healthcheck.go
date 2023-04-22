package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	apitypes "github.com/appclacks/go-types"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

const (
	defaultTimeout = 10 * time.Second
)

func toMap(slice []string) (map[string]string, error) {
	sliceMap := make(map[string]string)
	for _, l := range slice {
		splitted := strings.Split(l, "=")
		if len(splitted) != 2 {
			return nil, fmt.Errorf("Invalid label %s", l)
		}
		sliceMap[splitted[0]] = splitted[1]
	}
	return sliceMap, nil

}

func getHealthcheckCmd() *cobra.Command {
	var healthcheckID string
	var healthcheckName string
	var getHealthcheck = &cobra.Command{
		Use:   "get",
		Short: "Get an API healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			client := buildClient()
			identifier := healthcheckID
			if identifier == "" {
				identifier = healthcheckName
			}
			input := apitypes.GetHealthcheckInput{
				Identifier: identifier,
			}
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			healthcheck, err := client.GetHealthcheck(ctx, input)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(&healthcheck)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			printHealthcheckTab([]apitypes.Healthcheck{healthcheck})
			os.Exit(0)
		},
	}
	getHealthcheck.PersistentFlags().StringVar(&healthcheckID, "id", "", "Healthcheck ID")
	getHealthcheck.PersistentFlags().StringVar(&healthcheckName, "name", "", "Healthcheck Name")

	return getHealthcheck
}

func deleteHealthcheckCmd() *cobra.Command {
	var tokenID string
	var deleteHealthcheck = &cobra.Command{
		Use:   "delete",
		Short: "Delete an healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			client := buildClient()
			input := apitypes.DeleteHealthcheckInput{
				ID: tokenID,
			}
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			result, err := client.DeleteHealthcheck(ctx, input)
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
	deleteHealthcheck.PersistentFlags().StringVar(&tokenID, "id", "", "Token ID")
	err := deleteHealthcheck.MarkPersistentFlagRequired("id")
	exitIfError(err)

	return deleteHealthcheck
}

func listHealthchecksCmd() *cobra.Command {
	var listHealthchecks = &cobra.Command{
		Use:   "list",
		Short: "List API healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			client := buildClient()
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			result, err := client.ListHealthchecks(ctx)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(result)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			printHealthcheckTab(result.Result)
			os.Exit(0)
		},
	}

	return listHealthchecks
}

func printHealthcheckTab(healthchecks []apitypes.Healthcheck) {
	t := tabby.New()
	t.AddHeader("ID", "Name", "Description", "Interval", "Timeout", "Labels", "Enabled", "Definition")
	for _, healthcheck := range healthchecks {
		jsonLabels, err := json.Marshal(healthcheck.Labels)
		exitIfError(err)
		jsonDef, err := json.Marshal(healthcheck.Definition)
		exitIfError(err)
		t.AddLine(healthcheck.ID, healthcheck.Name, healthcheck.Description, healthcheck.Interval, healthcheck.Timeout, string(jsonLabels), healthcheck.Enabled, string(jsonDef))
	}
	t.Print()
}
