package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/appclacks/cli/client"
	apitypes "github.com/appclacks/go-types"
	"github.com/spf13/cobra"
)

func createCommandHealthcheckCmd(client *client.Client) *cobra.Command {
	var name string
	var timeout string
	var description string
	var labels []string
	var interval string
	var disabled bool
	var command string
	var arguments []string

	var createCommandHealthcheck = &cobra.Command{
		Use:   "create",
		Short: "Create a command healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			labelsMap, err := toMap(labels)
			exitIfError(err)

			payload := apitypes.CreateCommandHealthcheckInput{
				Name:        name,
				Description: description,
				Labels:      labelsMap,
				Timeout:     timeout,
				Interval:    interval,
				Enabled:     !disabled,
				HealthcheckCommandDefinition: apitypes.HealthcheckCommandDefinition{
					Command:   command,
					Arguments: arguments,
				},
			}
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			healthcheck, err := client.CreateCommandHealthcheck(ctx, payload)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(healthcheck)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			printHealthcheckTab([]apitypes.Healthcheck{healthcheck})
			os.Exit(0)
		},
	}

	createCommandHealthcheck.PersistentFlags().StringVar(&name, "name", "", "healthcheck name")
	err := createCommandHealthcheck.MarkPersistentFlagRequired("name")
	exitIfError(err)

	createCommandHealthcheck.PersistentFlags().StringVar(&description, "description", "", "healthcheck description")

	createCommandHealthcheck.PersistentFlags().StringSliceVar(&labels, "labels", []string{}, "healthchecks labels (example: foo=bar)")

	createCommandHealthcheck.PersistentFlags().StringVar(&interval, "interval", "60s", "healthcheck interval (examples: 30s, 3m)")

	createCommandHealthcheck.PersistentFlags().StringVar(&timeout, "timeout", "5s", "healthcheck timeout")

	createCommandHealthcheck.PersistentFlags().BoolVar(&disabled, "disabled", false, "Disable the healthcheck on the Appclacks platform")

	createCommandHealthcheck.PersistentFlags().StringVar(&command, "command", "", "Healthcheck command")
	err = createCommandHealthcheck.MarkPersistentFlagRequired("command")
	exitIfError(err)

	createCommandHealthcheck.PersistentFlags().StringSliceVar(&arguments, "arguments", []string{}, "Healthcheck command arguments")

	return createCommandHealthcheck
}

func updateCommandHealthcheckCmd(client *client.Client) *cobra.Command {
	var id string
	var timeout string
	var name string
	var description string
	var labels []string
	var interval string
	var disabled bool
	var command string
	var arguments []string

	var updateCommandHealthcheck = &cobra.Command{
		Use:   "update",
		Short: "Update a command healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			labelsMap, err := toMap(labels)
			exitIfError(err)
			payload := apitypes.UpdateCommandHealthcheckInput{
				ID:          id,
				Name:        name,
				Description: description,
				Labels:      labelsMap,
				Timeout:     timeout,
				Interval:    interval,
				Enabled:     !disabled,
				HealthcheckCommandDefinition: apitypes.HealthcheckCommandDefinition{
					Command:   command,
					Arguments: arguments,
				},
			}
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			result, err := client.UpdateCommandHealthcheck(ctx, payload)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(result)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			printHealthcheckTab([]apitypes.Healthcheck{result})
			os.Exit(0)
		},
	}
	updateCommandHealthcheck.PersistentFlags().StringVar(&id, "id", "", "healthcheck id")
	err := updateCommandHealthcheck.MarkPersistentFlagRequired("id")
	exitIfError(err)

	updateCommandHealthcheck.PersistentFlags().StringVar(&name, "name", "", "healthcheck name")
	err = updateCommandHealthcheck.MarkPersistentFlagRequired("name")
	exitIfError(err)

	updateCommandHealthcheck.PersistentFlags().StringVar(&description, "description", "", "healthcheck description")

	updateCommandHealthcheck.PersistentFlags().StringVar(&timeout, "timeout", "5s", "healthcheck timeout")

	updateCommandHealthcheck.PersistentFlags().StringSliceVar(&labels, "labels", []string{}, "healthchecks labels (example: foo=bar)")

	updateCommandHealthcheck.PersistentFlags().StringVar(&interval, "interval", "60s", "healthcheck interval (examples: 10s, 3m)")

	updateCommandHealthcheck.PersistentFlags().BoolVar(&disabled, "disabled", false, "Disable the healthcheck on the Appclacks platform")

	updateCommandHealthcheck.PersistentFlags().StringVar(&command, "command", "", "Healthcheck command")
	err = updateCommandHealthcheck.MarkPersistentFlagRequired("command")
	exitIfError(err)

	updateCommandHealthcheck.PersistentFlags().StringSliceVar(&arguments, "arguments", []string{}, "Healthcheck command arguments")

	return updateCommandHealthcheck
}
