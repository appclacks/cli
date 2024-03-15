package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	apitypes "github.com/appclacks/go-types"
	"github.com/spf13/cobra"
)

func createDNSHealthcheckCmd() *cobra.Command {
	var name string
	var timeout string
	var description string
	var labels []string
	var interval string
	var disabled bool
	var domain string
	var expectedIPs []string

	var createDNSHealthcheck = &cobra.Command{
		Use:   "create",
		Short: "Create a DNS healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			client := buildClient()
			labelsMap, err := toMap(labels)
			exitIfError(err)

			payload := apitypes.CreateDNSHealthcheckInput{
				Name:        name,
				Description: description,
				Labels:      labelsMap,
				Timeout:     timeout,
				Interval:    interval,
				Enabled:     !disabled,
				HealthcheckDNSDefinition: apitypes.HealthcheckDNSDefinition{
					Domain:      domain,
					ExpectedIPs: expectedIPs,
				},
			}
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			healthcheck, err := client.CreateDNSHealthcheckWithContext(ctx, payload)
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

	createDNSHealthcheck.PersistentFlags().StringVar(&name, "name", "", "healthcheck name")
	err := createDNSHealthcheck.MarkPersistentFlagRequired("name")
	exitIfError(err)

	createDNSHealthcheck.PersistentFlags().StringVar(&description, "description", "", "healthcheck description")

	createDNSHealthcheck.PersistentFlags().StringSliceVar(&labels, "labels", []string{}, "healthchecks labels (example: foo=bar)")

	createDNSHealthcheck.PersistentFlags().StringVar(&interval, "interval", "60s", "healthcheck interval (examples: 30s, 3m)")

	createDNSHealthcheck.PersistentFlags().StringVar(&timeout, "timeout", "5s", "healthcheck timeout")

	createDNSHealthcheck.PersistentFlags().BoolVar(&disabled, "disabled", false, "Disable the healthcheck on the Appclacks platform")

	createDNSHealthcheck.PersistentFlags().StringVar(&domain, "domain", "", "healthcheck domain")
	err = createDNSHealthcheck.MarkPersistentFlagRequired("domain")
	exitIfError(err)

	createDNSHealthcheck.PersistentFlags().StringSliceVar(&expectedIPs, "expected-ips", []string{}, "DNS resolution expected IPs")

	return createDNSHealthcheck
}

func updateDNSHealthcheckCmd() *cobra.Command {
	var id string
	var timeout string
	var name string
	var description string
	var labels []string
	var interval string
	var disabled bool
	var domain string
	var expectedIPs []string

	var updateDNSHealthcheck = &cobra.Command{
		Use:   "update",
		Short: "Update a DNS healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			client := buildClient()
			labelsMap, err := toMap(labels)
			exitIfError(err)
			payload := apitypes.UpdateDNSHealthcheckInput{
				ID:          id,
				Name:        name,
				Description: description,
				Labels:      labelsMap,
				Timeout:     timeout,
				Interval:    interval,
				Enabled:     !disabled,
				HealthcheckDNSDefinition: apitypes.HealthcheckDNSDefinition{
					Domain:      domain,
					ExpectedIPs: expectedIPs,
				},
			}
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			result, err := client.UpdateDNSHealthcheckWithContext(ctx, payload)
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
	updateDNSHealthcheck.PersistentFlags().StringVar(&id, "id", "", "healthcheck id")
	err := updateDNSHealthcheck.MarkPersistentFlagRequired("id")
	exitIfError(err)

	updateDNSHealthcheck.PersistentFlags().StringVar(&name, "name", "", "healthcheck name")
	err = updateDNSHealthcheck.MarkPersistentFlagRequired("name")
	exitIfError(err)

	updateDNSHealthcheck.PersistentFlags().StringVar(&description, "description", "", "healthcheck description")

	updateDNSHealthcheck.PersistentFlags().StringVar(&timeout, "timeout", "5s", "healthcheck timeout")

	updateDNSHealthcheck.PersistentFlags().StringSliceVar(&labels, "labels", []string{}, "healthchecks labels (example: foo=bar)")

	updateDNSHealthcheck.PersistentFlags().StringVar(&interval, "interval", "60s", "healthcheck interval (examples: 120s, 3m)")

	updateDNSHealthcheck.PersistentFlags().BoolVar(&disabled, "disabled", false, "Disable the healthcheck on the Appclacks platform")

	updateDNSHealthcheck.PersistentFlags().StringVar(&domain, "domain", "", "healthcheck domain")
	err = updateDNSHealthcheck.MarkPersistentFlagRequired("domain")
	exitIfError(err)

	updateDNSHealthcheck.PersistentFlags().StringSliceVar(&expectedIPs, "expected-ips", []string{}, "DNS resolution expected IPs")

	return updateDNSHealthcheck
}
