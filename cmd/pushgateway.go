package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	goclient "github.com/appclacks/go-client"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

func createPushgatewayMetricCmd() *cobra.Command {
	var name string
	var description string
	var labels []string
	var ttl string
	var metricType string
	var value float32

	var createPushgatewayMetric = &cobra.Command{
		Use:   "create",
		Short: "Create (or update) a new metric in the push gateway. ",
		Run: func(cmd *cobra.Command, args []string) {
			client := buildClient()
			labelsMap, err := toMap(labels)
			exitIfError(err)

			payload := goclient.CreateOrUpdatePushgatewayMetricInput{
				Name:        name,
				Description: description,
				Labels:      labelsMap,
				TTL:         ttl,
				Type:        metricType,
				Value:       value,
			}
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			result, err := client.CreateOrUpdatePushgatewayMetric(ctx, payload)
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
	createPushgatewayMetric.PersistentFlags().StringVar(&name, "name", "", "metric name")
	err := createPushgatewayMetric.MarkPersistentFlagRequired("name")
	exitIfError(err)

	createPushgatewayMetric.PersistentFlags().StringVar(&description, "description", "", "metric description")

	createPushgatewayMetric.PersistentFlags().StringSliceVar(&labels, "labels", []string{}, "metric labels (example: foo=bar)")

	createPushgatewayMetric.PersistentFlags().StringVar(&ttl, "ttl", "", "metric timeout")

	createPushgatewayMetric.PersistentFlags().StringVar(&metricType, "type", "", "metric type")

	createPushgatewayMetric.PersistentFlags().Float32Var(&value, "value", 0, "metric value")
	err = createPushgatewayMetric.MarkPersistentFlagRequired("value")
	exitIfError(err)

	return createPushgatewayMetric
}

func listPushgatewayMetricsCmd() *cobra.Command {
	var listPushgatewayMetrics = &cobra.Command{
		Use:   "list",
		Short: "List push gateway metrics",
		Run: func(cmd *cobra.Command, args []string) {
			client := buildClient()
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			result, err := client.ListPushgatewayMetrics(ctx)
			exitIfError(err)
			if outputFormat == "json" {
				json, err := json.Marshal(result)
				exitIfError(err)
				fmt.Println(string(json))
				os.Exit(0)
			}
			t := tabby.New()
			t.AddHeader("ID", "Name", "Description", "Type", "Labels", "TTL", "Created At", "Expires At", "Value")
			for _, metric := range result.Result {
				jsonLabels, err := json.Marshal(metric.Labels)
				exitIfError(err)
				t.AddLine(metric.ID, metric.Name, metric.Description, metric.Type, string(jsonLabels), metric.TTL, metric.CreatedAt, metric.ExpiresAt, metric.Value)
			}
			t.Print()
			os.Exit(0)
		},
	}
	return listPushgatewayMetrics
}

func deletePushgatewayMetricCmd() *cobra.Command {
	var metricID string
	var metricName string
	var deleteMetric = &cobra.Command{
		Use:   "delete",
		Short: "Delete a push gateway metric by name or by ID",
		Run: func(cmd *cobra.Command, args []string) {
			identifier := metricID
			if identifier == "" {
				identifier = metricName
			}
			if identifier == "" {
				exitIfError(errors.New("you should pass the metric name or ID (using --name or --id) you want to delete"))
			}
			input := goclient.DeletePushgatewayMetricInput{
				Identifier: identifier,
			}
			client := buildClient()
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			result, err := client.DeletePushgatewayMetric(ctx, input)
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
	deleteMetric.PersistentFlags().StringVar(&metricID, "id", "", "Metric ID")
	deleteMetric.PersistentFlags().StringVar(&metricName, "name", "", "Metric Name")
	return deleteMetric
}

func deleteAllPushgatewayMetricsCmd() *cobra.Command {
	var deleteMetric = &cobra.Command{
		Use:   "delete-all",
		Short: "Delete all push gateway metrics",
		Run: func(cmd *cobra.Command, args []string) {
			client := buildClient()
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			result, err := client.DeleteAllPushgatewayMetrics(ctx)
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
	return deleteMetric
}
