package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/appclacks/cli/client"
	apitypes "github.com/appclacks/go-types"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

func createHTTPHealthcheckCmd(client *client.Client) *cobra.Command {
	var name string
	var timeout string
	var description string
	var labels []string
	var interval string
	var enabled bool

	var validStatus []uint
	var target string
	var method string
	var port uint
	var redirect bool
	var body string
	var bodyRegexp []string
	var headers []string
	var protocol string
	var path string

	var createHTTPHealthcheck = &cobra.Command{
		Use:   "create",
		Short: "Create a HTTP healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			labelsMap, err := toMap(labels)
			exitIfError(err)
			headersMap, err := toMap(headers)
			exitIfError(err)
			formattedMethod := strings.ToUpper(method)

			payload := apitypes.CreateHTTPHealthcheckInput{
				Name:        name,
				Description: description,
				Timeout:     timeout,
				Labels:      labelsMap,
				Interval:    interval,
				Enabled:     enabled,
				HealthcheckHTTPDefinition: apitypes.HealthcheckHTTPDefinition{
					ValidStatus: validStatus,
					Target:      target,
					Method:      formattedMethod,
					Port:        port,
					Redirect:    redirect,
					Body:        body,
					BodyRegexp:  bodyRegexp,
					Headers:     headersMap,
					Protocol:    protocol,
					Path:        path,
				},
			}
			healthcheck, err := client.CreateHTTPHealthcheck(payload)
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

	createHTTPHealthcheck.PersistentFlags().StringVar(&name, "name", "", "healthcheck name")
	err := createHTTPHealthcheck.MarkPersistentFlagRequired("name")
	exitIfError(err)

	createHTTPHealthcheck.PersistentFlags().StringVar(&description, "description", "", "healthcheck description")

	createHTTPHealthcheck.PersistentFlags().StringSliceVar(&labels, "labels", []string{}, "healthchecks labels (example: foo=bar)")

	createHTTPHealthcheck.PersistentFlags().StringVar(&interval, "interval", "60s", "healthcheck interval (examples: 10s, 3m)")

	createHTTPHealthcheck.PersistentFlags().BoolVar(&enabled, "enabled", true, "Enable the healthcheck on the appclacks platform")

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

	createHTTPHealthcheck.PersistentFlags().StringVar(&timeout, "timeout", "5s", "healthcheck timeout")

	exitIfError(err)

	return createHTTPHealthcheck
}

func updateHTTPHealthcheckCmd(client *client.Client) *cobra.Command {
	var id string
	var timeout string
	var name string
	var description string
	var labels []string
	var interval string
	var enabled bool

	var validStatus []uint
	var target string
	var method string
	var port uint
	var redirect bool
	var body string
	var bodyRegexp []string
	var headers []string
	var protocol string
	var path string

	var updateHTTPHealthcheck = &cobra.Command{
		Use:   "update",
		Short: "Update a HTTP healthcheck",
		Run: func(cmd *cobra.Command, args []string) {
			labelsMap, err := toMap(labels)
			exitIfError(err)
			headersMap, err := toMap(headers)
			exitIfError(err)
			formattedMethod := strings.ToUpper(method)

			payload := apitypes.UpdateHTTPHealthcheckInput{
				ID:          id,
				Name:        name,
				Description: description,
				Labels:      labelsMap,
				Timeout:     timeout,
				Interval:    interval,
				Enabled:     enabled,
				HealthcheckHTTPDefinition: apitypes.HealthcheckHTTPDefinition{
					ValidStatus: validStatus,
					Target:      target,
					Method:      formattedMethod,
					Port:        port,
					Redirect:    redirect,
					Body:        body,
					BodyRegexp:  bodyRegexp,
					Headers:     headersMap,
					Protocol:    protocol,
					Path:        path,
				},
			}
			result, err := client.UpdateHTTPHealthcheck(payload)
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
	updateHTTPHealthcheck.PersistentFlags().StringVar(&id, "id", "", "healthcheck id")
	err := updateHTTPHealthcheck.MarkPersistentFlagRequired("id")
	exitIfError(err)

	updateHTTPHealthcheck.PersistentFlags().StringVar(&name, "name", "", "healthcheck name")
	err = updateHTTPHealthcheck.MarkPersistentFlagRequired("name")
	exitIfError(err)

	updateHTTPHealthcheck.PersistentFlags().StringVar(&description, "description", "", "healthcheck description")

	updateHTTPHealthcheck.PersistentFlags().StringSliceVar(&labels, "labels", []string{}, "healthchecks labels (example: foo=bar)")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&interval, "interval", "60s", "healthcheck interval (examples: 10s, 3m)")

	updateHTTPHealthcheck.PersistentFlags().BoolVar(&enabled, "enabled", true, "Enable the healthcheck on the appclacks platform")

	updateHTTPHealthcheck.PersistentFlags().UintSliceVar(&validStatus, "valid-status", []uint{200}, "Expected HTTP response status code")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&target, "target", "", "Healthcheck target (ip or domain)")
	err = updateHTTPHealthcheck.MarkPersistentFlagRequired("target")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&method, "method", "GET", "HTTP method to execute")

	updateHTTPHealthcheck.PersistentFlags().UintVar(&port, "port", 443, "Healthcheck HTTP port")

	updateHTTPHealthcheck.PersistentFlags().BoolVar(&redirect, "redirect", true, "Enable HTTP redirection on the healthcheck")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&body, "body", "", "Body to pass to the healthcheck")

	updateHTTPHealthcheck.PersistentFlags().StringSliceVar(&bodyRegexp, "body-regexp", []string{}, "Expected content of the repsonse body")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&protocol, "protocol", "https", "Protocol to use for the healthcheck (http or https)")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&path, "path", "", "Path to use for the healthcheck")

	updateHTTPHealthcheck.PersistentFlags().StringVar(&timeout, "timeout", "5s", "healthcheck timeout")

	exitIfError(err)

	return updateHTTPHealthcheck
}
