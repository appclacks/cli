package cmd

import (
	"github.com/appclacks/go-client"
	"github.com/spf13/cobra"
)

var outputFormat string
var profile string

func Execute() error {
	var rootCmd = &cobra.Command{
		Use:   "appclacks",
		Short: "Appclacks CLI",
	}
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "Command output format (table or json)")

	// healthcheck

	var healthcheck = &cobra.Command{
		Use:   "healthcheck",
		Short: "Manage your healthcheck",
	}
	healthcheck.AddCommand(deleteHealthcheckCmd())
	healthcheck.AddCommand(getHealthcheckCmd())
	healthcheck.AddCommand(listHealthchecksCmd())

	var dns = &cobra.Command{
		Use:   "dns",
		Short: "Manage your DNS healthchecks",
	}
	dns.AddCommand(createDNSHealthcheckCmd())
	dns.AddCommand(updateDNSHealthcheckCmd())

	var http = &cobra.Command{
		Use:   "http",
		Short: "Manage your HTTP healthchecks",
	}
	http.AddCommand(createHTTPHealthcheckCmd())
	http.AddCommand(updateHTTPHealthcheckCmd())

	var tcp = &cobra.Command{
		Use:   "tcp",
		Short: "Manage your TCP healthchecks",
	}
	tcp.AddCommand(createTCPHealthcheckCmd())
	tcp.AddCommand(updateTCPHealthcheckCmd())

	var tls = &cobra.Command{
		Use:   "tls",
		Short: "Manage your TLS healthchecks",
	}
	tls.AddCommand(createTLSHealthcheckCmd())
	tls.AddCommand(updateTLSHealthcheckCmd())

	var command = &cobra.Command{
		Use:   "command",
		Short: "Manage your Command healthchecks",
	}
	command.AddCommand(createCommandHealthcheckCmd())
	command.AddCommand(updateCommandHealthcheckCmd())

	healthcheck.AddCommand(dns)
	healthcheck.AddCommand(command)
	healthcheck.AddCommand(tls)
	healthcheck.AddCommand(tcp)
	healthcheck.AddCommand(http)

	rootCmd.AddCommand(healthcheck)

	return rootCmd.Execute()
}

func buildClient() *client.Client {
	c, err := client.New()
	exitIfError(err)
	return c
}
