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

func createTCPHealthcheckCmd(client *client.Client) *cobra.Command {
	var name string
	var timeout string
	var description string
	var labels []string
	var interval string
	var enabled bool

	var target string
	var port uint
	var shouldFail bool

	var createTCPHealthcheck = &cobra.Command{
		Use:   "create",
		Short: "Create a TCP healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			labelsMap, err := toMap(labels)
			exitIfError(err)

			payload := apitypes.CreateTCPHealthcheckInput{
				Name:        name,
				Description: description,
				Labels:      labelsMap,
				Timeout:     timeout,
				Interval:    interval,
				Enabled:     enabled,
				HealthcheckTCPDefinition: apitypes.HealthcheckTCPDefinition{

					Target: target,

					Port:       port,
					ShouldFail: shouldFail,
				},
			}
			healthcheck, err := client.CreateTCPHealthcheck(payload)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(healthcheck)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			t := tabby.New()
			t.AddHeader("ID", "Name", "Description", "Interval", "Timeout", "Labels", "Enabled", "Definition")
			jsonLabels, err := json.Marshal(healthcheck.Labels)
			exitIfError(err)
			jsonDef, err := json.Marshal(healthcheck.Definition)
			exitIfError(err)
			t.AddLine(healthcheck.ID, healthcheck.Name, healthcheck.Description, healthcheck.Interval, healthcheck.Timeout, string(jsonLabels), healthcheck.Enabled, string(jsonDef))

			t.Print()
			os.Exit(0)
		},
	}

	createTCPHealthcheck.PersistentFlags().StringVar(&name, "name", "", "healthcheck name")
	err := createTCPHealthcheck.MarkPersistentFlagRequired("name")
	exitIfError(err)

	createTCPHealthcheck.PersistentFlags().StringVar(&description, "description", "", "healthcheck description")

	createTCPHealthcheck.PersistentFlags().StringSliceVar(&labels, "labels", []string{}, "healthchecks labels (example: foo=bar)")

	createTCPHealthcheck.PersistentFlags().StringVar(&interval, "interval", "60s", "healthcheck interval (examples: 10s, 3m)")

	createTCPHealthcheck.PersistentFlags().StringVar(&timeout, "timeout", "5s", "healthcheck timeout")

	createTCPHealthcheck.PersistentFlags().BoolVar(&enabled, "enabled", true, "Enable the healthcheck on the appclacks platform")

	createTCPHealthcheck.PersistentFlags().StringVar(&target, "target", "", "Healthcheck target (ip or domain)")
	err = createTCPHealthcheck.MarkPersistentFlagRequired("target")
	exitIfError(err)

	createTCPHealthcheck.PersistentFlags().UintVar(&port, "port", 443, "Healthcheck TCP port")

	createTCPHealthcheck.PersistentFlags().BoolVar(&shouldFail, "should-fail", false, "Consider the healthchek successful if the TCP connection fails")

	return createTCPHealthcheck
}

func updateTCPHealthcheckCmd(client *client.Client) *cobra.Command {
	var id string
	var name string
	var timeout string
	var description string
	var labels []string
	var interval string
	var enabled bool

	var target string
	var port uint
	var shouldFail bool

	var updateTCPHealthcheck = &cobra.Command{
		Use:   "update",
		Short: "Update a TCP healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			labelsMap, err := toMap(labels)
			exitIfError(err)

			payload := apitypes.UpdateTCPHealthcheckInput{
				ID:          id,
				Name:        name,
				Description: description,
				Labels:      labelsMap,
				Timeout:     timeout,
				Interval:    interval,
				Enabled:     enabled,
				HealthcheckTCPDefinition: apitypes.HealthcheckTCPDefinition{
					Target: target,

					Port:       port,
					ShouldFail: shouldFail,
				},
			}
			result, err := client.UpdateTCPHealthcheck(payload)
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
	updateTCPHealthcheck.PersistentFlags().StringVar(&id, "id", "", "healthcheck id")
	err := updateTCPHealthcheck.MarkPersistentFlagRequired("id")
	exitIfError(err)

	updateTCPHealthcheck.PersistentFlags().StringVar(&name, "name", "", "healthcheck name")
	err = updateTCPHealthcheck.MarkPersistentFlagRequired("name")
	exitIfError(err)

	updateTCPHealthcheck.PersistentFlags().StringVar(&description, "description", "", "healthcheck description")

	updateTCPHealthcheck.PersistentFlags().StringSliceVar(&labels, "labels", []string{}, "healthchecks labels (example: foo=bar)")

	updateTCPHealthcheck.PersistentFlags().StringVar(&interval, "interval", "60s", "healthcheck interval (examples: 10s, 3m)")

	updateTCPHealthcheck.PersistentFlags().StringVar(&timeout, "timeout", "5s", "healthcheck timeout")

	updateTCPHealthcheck.PersistentFlags().BoolVar(&enabled, "enabled", true, "Enable the healthcheck on the appclacks platform")

	updateTCPHealthcheck.PersistentFlags().StringVar(&target, "target", "", "Healthcheck target (ip or domain)")
	err = updateTCPHealthcheck.MarkPersistentFlagRequired("target")
	exitIfError(err)

	updateTCPHealthcheck.PersistentFlags().UintVar(&port, "port", 443, "Healthcheck TCP port")

	updateTCPHealthcheck.PersistentFlags().BoolVar(&shouldFail, "should-fail", false, "Consider the healthchek successful if the TCP connection fails")

	return updateTCPHealthcheck
}
