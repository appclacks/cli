package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	apitypes "github.com/appclacks/go-types"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

func listHealthckecksResultsCmd() *cobra.Command {
	var startDate string
	var endDate string
	var healthcheckID string
	var page uint
	var successOnly bool
	var errorOnly bool

	var listHealthchecksResults = &cobra.Command{
		Use:   "list",
		Short: "List healthchecks results",
		Run: func(cmd *cobra.Command, args []string) {
			client := buildClient()
			var err error
			start := time.Now().UTC().Add(-5 * time.Minute)
			if startDate != "" {
				start, err = time.Parse(time.RFC3339, fmt.Sprintf("%sZ", startDate))
				exitIfError(err)
			}
			end := time.Now().UTC()
			if endDate != "" {
				end, err = time.Parse(time.RFC3339, fmt.Sprintf("%sZ", endDate))
				exitIfError(err)
			}
			var trueP = true
			var falseP = false
			var success *bool
			if successOnly {
				success = &trueP
			}
			if errorOnly {
				success = &falseP
			}
			input := apitypes.ListHealthchecksResultsInput{
				StartDate:     start,
				EndDate:       end,
				HealthcheckID: healthcheckID,
				Page:          int(page),
				Success:       success,
			}
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			result, err := client.ListHealthchecksResults(ctx, input)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(result)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			t := tabby.New()
			t.AddHeader("ID", "Created At", "Success", "Duration (ms)", "Summary", "Message", "Healthcheck ID", "Labels")
			for i := len(result.Result) - 1; i >= 0; i-- {
				result := result.Result[i]
				jsonLabels, err := json.Marshal(result.Labels)
				exitIfError(err)
				t.AddLine(result.ID, result.CreatedAt, result.Success, result.Duration, result.Summary, result.Message, result.HealthcheckID, string(jsonLabels))
			}
			t.Print()
			os.Exit(0)
		},
	}
	listHealthchecksResults.PersistentFlags().StringVar(&startDate, "start-date", "", "Start date for results retrieval in UTC (Example: 2006-01-02T15:04:05). Default to 5 minutes ago")

	listHealthchecksResults.PersistentFlags().StringVar(&endDate, "end-date", "", "End date for results retrieval in UTC (Example: 2006-01-02T15:04:05). Default to the current time")

	listHealthchecksResults.PersistentFlags().StringVar(&healthcheckID, "healthcheck-id", "", "Get result for a specific healthcheck")
	listHealthchecksResults.PersistentFlags().UintVar(&page, "page", 1, "Result page to retrieve")
	listHealthchecksResults.PersistentFlags().BoolVar(&successOnly, "success", false, "Retrieve only successful healthchecks results")
	listHealthchecksResults.PersistentFlags().BoolVar(&errorOnly, "error", false, "Retrieve only failed healthchecks results")

	return listHealthchecksResults
}
