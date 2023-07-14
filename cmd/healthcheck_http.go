package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	apitypes "github.com/appclacks/go-types"
	"github.com/spf13/cobra"
)

func createHTTPHealthcheckCmd() *cobra.Command {
	var name string
	var timeout string
	var description string
	var labels []string
	var interval string
	var disabled bool

	var validStatus []uint
	var target string
	var method string
	var port uint
	var redirect bool
	var body string
	var bodyRegexp []string
	var headers []string
	var query []string
	var protocol string
	var path string
	var host string
	var serverName string
	var insecure bool

	var createHTTPHealthcheck = &cobra.Command{
		Use:   "create",
		Short: "Create a HTTP healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			client := buildClient()
			labelsMap, err := toMap(labels)
			exitIfError(err)
			headersMap, err := toMap(headers)
			exitIfError(err)
			formattedMethod := strings.ToUpper(method)
			queryMap, err := toMap(query)
			exitIfError(err)

			payload := apitypes.CreateHTTPHealthcheckInput{
				Name:        name,
				Description: description,
				Timeout:     timeout,
				Labels:      labelsMap,
				Interval:    interval,
				Enabled:     !disabled,
				HealthcheckHTTPDefinition: apitypes.HealthcheckHTTPDefinition{
					ValidStatus: validStatus,
					Target:      target,
					Method:      formattedMethod,
					Port:        port,
					Query:       queryMap,
					Redirect:    redirect,
					Body:        body,
					BodyRegexp:  bodyRegexp,
					Headers:     headersMap,
					Protocol:    protocol,
					Path:        path,
					Host:        host,
					ServerName:  serverName,
					Insecure:    insecure,
				},
			}
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			healthcheck, err := client.CreateHTTPHealthcheck(ctx, payload)
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

	createHTTPHealthcheck.PersistentFlags().StringVar(&name, "name", "", "healthcheck name")
	err := createHTTPHealthcheck.MarkPersistentFlagRequired("name")
	exitIfError(err)

	createHTTPHealthcheck.PersistentFlags().StringVar(&description, "description", "", "healthcheck description")

	createHTTPHealthcheck.PersistentFlags().StringSliceVar(&labels, "labels", []string{}, "healthchecks labels (example: foo=bar)")

	createHTTPHealthcheck.PersistentFlags().StringVar(&interval, "interval", "60s", "healthcheck interval (examples: 30s, 3m)")

	createHTTPHealthcheck.PersistentFlags().BoolVar(&disabled, "disabled", false, "Disable the healthcheck on the Appclacks platform")

	createHTTPHealthcheck.PersistentFlags().UintSliceVar(&validStatus, "valid-status", []uint{200}, "Expected HTTP response status code")

	createHTTPHealthcheck.PersistentFlags().StringVar(&target, "target", "", "Healthcheck target (ip or domain)")
	err = createHTTPHealthcheck.MarkPersistentFlagRequired("target")

	createHTTPHealthcheck.PersistentFlags().StringVar(&method, "method", "GET", "HTTP method to execute")

	createHTTPHealthcheck.PersistentFlags().UintVar(&port, "port", 443, "Healthcheck HTTP port")

	createHTTPHealthcheck.PersistentFlags().BoolVar(&redirect, "redirect", true, "Enable HTTP redirection on the healthcheck")

	createHTTPHealthcheck.PersistentFlags().StringVar(&body, "body", "", "Body to pass to the healthcheck")

	createHTTPHealthcheck.PersistentFlags().StringSliceVar(&bodyRegexp, "body-regexp", []string{}, "Expected content of the repsonse body")

	createHTTPHealthcheck.PersistentFlags().StringVar(&protocol, "protocol", "https", "Protocol to use for the healthcheck (http or https)")

	createHTTPHealthcheck.PersistentFlags().StringVar(&path, "path", "", "Path to use for the healthcheck")

	createHTTPHealthcheck.PersistentFlags().StringVar(&host, "host", "", "Host header to use for the health check HTTP requests")

	createHTTPHealthcheck.PersistentFlags().StringVar(&timeout, "timeout", "5s", "healthcheck timeout")

	createHTTPHealthcheck.PersistentFlags().StringSliceVar(&headers, "headers", []string{}, "healthchecks http headers (example: foo=bar)")

	createHTTPHealthcheck.PersistentFlags().StringSliceVar(&query, "query", []string{}, "healthchecks http query params (example: foo=bar)")
	createHTTPHealthcheck.PersistentFlags().StringVar(&serverName, "server-name", "", "TLS SNI")
	createHTTPHealthcheck.PersistentFlags().BoolVar(&insecure, "insecure", false, "TLS Insecure")

	exitIfError(err)

	return createHTTPHealthcheck
}

func updateHTTPHealthcheckCmd() *cobra.Command {
	var id string
	var timeout string
	var name string
	var description string
	var labels []string
	var interval string
	var disabled bool

	var validStatus []uint
	var target string
	var method string
	var port uint
	var redirect bool
	var body string
	var bodyRegexp []string
	var headers []string
	var query []string
	var protocol string
	var path string
	var host string
	var serverName string
	var insecure bool

	var updateHTTPHealthcheck = &cobra.Command{
		Use:   "update",
		Short: "Update a HTTP healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			client := buildClient()
			labelsMap, err := toMap(labels)
			exitIfError(err)
			headersMap, err := toMap(headers)
			exitIfError(err)
			formattedMethod := strings.ToUpper(method)
			queryMap, err := toMap(query)
			exitIfError(err)

			payload := apitypes.UpdateHTTPHealthcheckInput{
				ID:          id,
				Name:        name,
				Description: description,
				Labels:      labelsMap,
				Timeout:     timeout,
				Interval:    interval,
				Enabled:     !disabled,
				HealthcheckHTTPDefinition: apitypes.HealthcheckHTTPDefinition{
					ValidStatus: validStatus,
					Target:      target,
					Method:      formattedMethod,
					Port:        port,
					Redirect:    redirect,
					Body:        body,
					Query:       queryMap,
					BodyRegexp:  bodyRegexp,
					Headers:     headersMap,
					Protocol:    protocol,
					Path:        path,
					Host:        host,
					ServerName:  serverName,
					Insecure:    insecure,
				},
			}
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			result, err := client.UpdateHTTPHealthcheck(ctx, payload)
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
	updateHTTPHealthcheck.PersistentFlags().StringVar(&id, "id", "", "healthcheck id")
	err := updateHTTPHealthcheck.MarkPersistentFlagRequired("id")
	exitIfError(err)

	updateHTTPHealthcheck.PersistentFlags().StringVar(&name, "name", "", "healthcheck name")
	err = updateHTTPHealthcheck.MarkPersistentFlagRequired("name")
	exitIfError(err)

	updateHTTPHealthcheck.PersistentFlags().StringVar(&description, "description", "", "healthcheck description")

	updateHTTPHealthcheck.PersistentFlags().StringSliceVar(&labels, "labels", []string{}, "healthchecks labels (example: foo=bar)")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&interval, "interval", "60s", "healthcheck interval (examples: 10s, 3m)")

	updateHTTPHealthcheck.PersistentFlags().BoolVar(&disabled, "disabled", false, "Diasble the healthcheck on the Appclacks platform")

	updateHTTPHealthcheck.PersistentFlags().UintSliceVar(&validStatus, "valid-status", []uint{200}, "Expected HTTP response status code")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&target, "target", "", "Healthcheck target (ip or domain)")
	err = updateHTTPHealthcheck.MarkPersistentFlagRequired("target")
	exitIfError(err)

	updateHTTPHealthcheck.PersistentFlags().StringVar(&method, "method", "GET", "HTTP method to execute")

	updateHTTPHealthcheck.PersistentFlags().UintVar(&port, "port", 443, "Healthcheck HTTP port")

	updateHTTPHealthcheck.PersistentFlags().BoolVar(&redirect, "redirect", true, "Enable HTTP redirection on the healthcheck")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&body, "body", "", "Body to pass to the healthcheck")

	updateHTTPHealthcheck.PersistentFlags().StringSliceVar(&bodyRegexp, "body-regexp", []string{}, "Expected content of the repsonse body")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&protocol, "protocol", "https", "Protocol to use for the healthcheck (http or https)")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&path, "path", "", "Path to use for the healthcheck")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&host, "host", "", "Host header to use for the health check HTTP requests")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&timeout, "timeout", "5s", "healthcheck timeout")

	updateHTTPHealthcheck.PersistentFlags().StringSliceVar(&headers, "headers", []string{}, "healthchecks http headers (example: foo=bar)")

	updateHTTPHealthcheck.PersistentFlags().StringSliceVar(&query, "query", []string{}, "healthchecks http query params (example: foo=bar)")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&serverName, "server-name", "", "TLS SNI")
	updateHTTPHealthcheck.PersistentFlags().BoolVar(&insecure, "insecure", false, "TLS Insecure")

	return updateHTTPHealthcheck
}
