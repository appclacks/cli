package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/appclacks/cli/client"
	apitypes "github.com/appclacks/go-types"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

func listHealthckecksResultsCmd(client *client.Client) *cobra.Command {
	var startDate string
	var endDate string
	var healthcheckID string
	var page uint

	var listHealthchecksResults = &cobra.Command{
		Use:   "list",
		Short: "List healthchecks results",
		Run: func(cmd *cobra.Command, args []string) {
			s, err := time.Parse(time.RFC3339, fmt.Sprintf("%sZ", startDate))
			exitIfError(err)
			e, err := time.Parse(time.RFC3339, fmt.Sprintf("%sZ", endDate))
			exitIfError(err)
			input := apitypes.ListHealthchecksResultsInput{
				StartDate:     s,
				EndDate:       e,
				HealthcheckID: healthcheckID,
				Page:          int(page),
			}
			result, err := client.ListHealthchecksResults(input)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(result)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			t := tabby.New()
			t.AddHeader("ID", "Created At", "Success", "Summary", "Message", "Healthcheck ID", "Labels")
			for _, result := range result.Result {
				jsonLabels, err := json.Marshal(result.Labels)
				exitIfError(err)
				t.AddLine(result.ID, result.CreatedAt, result.Success, result.Summary, result.Message, result.HealthcheckID, jsonLabels)
			}
			t.Print()
			os.Exit(0)
		},
	}
	listHealthchecksResults.PersistentFlags().StringVar(&startDate, "start-date", "", "Start date for results retrieval (Example: 2006-01-02T15:04:05)")
	err := listHealthchecksResults.MarkPersistentFlagRequired("start-date")
	exitIfError(err)
	listHealthchecksResults.PersistentFlags().StringVar(&endDate, "end-date", "", "End date for results retrieval (Example: 2006-01-02T15:04:05)")
	err = listHealthchecksResults.MarkPersistentFlagRequired("end-date")
	exitIfError(err)
	listHealthchecksResults.PersistentFlags().StringVar(&healthcheckID, "healthcheck-id", "", "Get result for a specific healthcheck")
	listHealthchecksResults.PersistentFlags().UintVar(&page, "page", 1, "Result page to retrieve")
	exitIfError(err)
	return listHealthchecksResults
}
