package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	goclient "github.com/appclacks/go-client"
	"github.com/spf13/cobra"
)

func createTLSHealthcheckCmd() *cobra.Command {
	var name string
	var timeout string
	var description string
	var labels []string
	var interval string
	var disabled bool

	var target string
	var port uint
	var key string
	var cert string
	var cacert string
	var serverName string
	var insecure bool
	var expirationDelay string

	var createTLSHealthcheck = &cobra.Command{
		Use:   "create",
		Short: "Create a TLS healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			client := buildClient()
			labelsMap, err := toMap(labels)
			exitIfError(err)

			payload := goclient.CreateTLSHealthcheckInput{
				Name:        name,
				Description: description,
				Labels:      labelsMap,
				Timeout:     timeout,
				Interval:    interval,
				Enabled:     !disabled,
				HealthcheckTLSDefinition: goclient.HealthcheckTLSDefinition{

					Target:          target,
					Key:             key,
					Cert:            cert,
					Cacert:          cacert,
					ServerName:      serverName,
					Insecure:        insecure,
					ExpirationDelay: expirationDelay,
					Port:            port,
				},
			}
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			healthcheck, err := client.CreateTLSHealthcheck(ctx, payload)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(healthcheck)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			printHealthcheckTab([]goclient.Healthcheck{healthcheck})
			os.Exit(0)
		},
	}

	createTLSHealthcheck.PersistentFlags().StringVar(&name, "name", "", "healthcheck name")
	err := createTLSHealthcheck.MarkPersistentFlagRequired("name")
	exitIfError(err)

	createTLSHealthcheck.PersistentFlags().StringVar(&description, "description", "", "healthcheck description")

	createTLSHealthcheck.PersistentFlags().StringSliceVar(&labels, "labels", []string{}, "healthchecks labels (example: foo=bar)")

	createTLSHealthcheck.PersistentFlags().StringVar(&interval, "interval", "60s", "healthcheck interval (examples: 30s, 3m)")

	createTLSHealthcheck.PersistentFlags().StringVar(&timeout, "timeout", "5s", "healthcheck timeout")

	createTLSHealthcheck.PersistentFlags().BoolVar(&disabled, "disabled", false, "Disable the healthcheck on the Appclacks platform")

	createTLSHealthcheck.PersistentFlags().StringVar(&target, "target", "", "Healthcheck target (ip or domain)")
	err = createTLSHealthcheck.MarkPersistentFlagRequired("target")
	exitIfError(err)

	createTLSHealthcheck.PersistentFlags().UintVar(&port, "port", 443, "Healthcheck TLS port")

	createTLSHealthcheck.PersistentFlags().StringVar(&key, "key", "", "TLS key file")
	createTLSHealthcheck.PersistentFlags().StringVar(&cert, "cert", "", "TLS cert file")
	createTLSHealthcheck.PersistentFlags().StringVar(&cacert, "cacert", "", "TLS cacert file")
	createTLSHealthcheck.PersistentFlags().StringVar(&serverName, "server-name", "", "TLS SNI")
	createTLSHealthcheck.PersistentFlags().StringVar(&expirationDelay, "expiration-delay", "", "TLS certificate expiration delay")
	createTLSHealthcheck.PersistentFlags().BoolVar(&insecure, "insecure", false, "TLS Insecure")

	return createTLSHealthcheck
}

func updateTLSHealthcheckCmd() *cobra.Command {
	var id string
	var name string
	var timeout string
	var description string
	var labels []string
	var interval string
	var disabled bool

	var target string
	var port uint
	var key string
	var cert string
	var cacert string
	var serverName string
	var insecure bool
	var expirationDelay string

	var updateTLSHealthcheck = &cobra.Command{
		Use:   "update",
		Short: "Update a TLS healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			client := buildClient()
			labelsMap, err := toMap(labels)
			exitIfError(err)

			payload := goclient.UpdateTLSHealthcheckInput{
				ID:          id,
				Name:        name,
				Description: description,
				Labels:      labelsMap,
				Timeout:     timeout,
				Interval:    interval,
				Enabled:     !disabled,
				HealthcheckTLSDefinition: goclient.HealthcheckTLSDefinition{
					Target:          target,
					Key:             key,
					Cert:            cert,
					Cacert:          cacert,
					ServerName:      serverName,
					Insecure:        insecure,
					ExpirationDelay: expirationDelay,
					Port:            port,
				},
			}
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			result, err := client.UpdateTLSHealthcheck(ctx, payload)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(result)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			printHealthcheckTab([]goclient.Healthcheck{result})
			os.Exit(0)
		},
	}
	updateTLSHealthcheck.PersistentFlags().StringVar(&id, "id", "", "healthcheck id")
	err := updateTLSHealthcheck.MarkPersistentFlagRequired("id")
	exitIfError(err)

	updateTLSHealthcheck.PersistentFlags().StringVar(&name, "name", "", "healthcheck name")
	err = updateTLSHealthcheck.MarkPersistentFlagRequired("name")
	exitIfError(err)

	updateTLSHealthcheck.PersistentFlags().StringVar(&description, "description", "", "healthcheck description")

	updateTLSHealthcheck.PersistentFlags().StringSliceVar(&labels, "labels", []string{}, "healthchecks labels (example: foo=bar)")

	updateTLSHealthcheck.PersistentFlags().StringVar(&interval, "interval", "60s", "healthcheck interval (examples: 120s, 3m)")

	updateTLSHealthcheck.PersistentFlags().StringVar(&timeout, "timeout", "5s", "healthcheck timeout")

	updateTLSHealthcheck.PersistentFlags().BoolVar(&disabled, "disabled", false, "Disable the healthcheck on the appclacks platform")

	updateTLSHealthcheck.PersistentFlags().StringVar(&target, "target", "", "Healthcheck target (ip or domain)")
	err = updateTLSHealthcheck.MarkPersistentFlagRequired("target")
	exitIfError(err)

	updateTLSHealthcheck.PersistentFlags().UintVar(&port, "port", 443, "Healthcheck TLS port")
	updateTLSHealthcheck.PersistentFlags().StringVar(&key, "key", "", "TLS key file")
	updateTLSHealthcheck.PersistentFlags().StringVar(&cert, "cert", "", "TLS cert file")
	updateTLSHealthcheck.PersistentFlags().StringVar(&cacert, "cacert", "", "TLS cacert file")
	updateTLSHealthcheck.PersistentFlags().StringVar(&serverName, "server-name", "", "TLS SNI")
	updateTLSHealthcheck.PersistentFlags().StringVar(&expirationDelay, "expiration-delay", "", "TLS certificate expiration delay")
	updateTLSHealthcheck.PersistentFlags().BoolVar(&insecure, "insecure", false, "TLS Insecure")

	return updateTLSHealthcheck
}
