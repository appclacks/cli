package cmd

import (
	"fmt"
	"os"

	"github.com/appclacks/cli/client"
	"github.com/spf13/cobra"
)

func getHealthchecksMetricsCmd(client *client.Client) *cobra.Command {
	var getHealthchecksMetricsCmd = &cobra.Command{
		Use:   "get",
		Short: "Get healthchecks metrics in Prometheus format",
		Run: func(cmd *cobra.Command, args []string) {
			metrics, err := client.GetHealthchecksMetrics()
			exitIfError(err)
			fmt.Print(metrics)
			os.Exit(0)
		},
	}
	return getHealthchecksMetricsCmd
}
