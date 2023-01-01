package cmd

import (
	"github.com/appclacks/cli/client"
	"github.com/spf13/cobra"
)

var outputFormat string

func Execute() error {
	var rootCmd = &cobra.Command{
		Use:   "appclacks",
		Short: "AppClacks CLI",
	}
	rootCmd.PersistentFlags().StringVar(&outputFormat, "output", "table", "Command output format (table or json)")
	client := client.New("https://api.appclacks.com")
	var organization = &cobra.Command{
		Use:   "organization",
		Short: "Manage your organization",
	}
	organization.AddCommand(createOrganizationCmd(client))
	// account

	var account = &cobra.Command{
		Use:   "account",
		Short: "Manage your account",
	}
	account.AddCommand(getAccountOrganizationCmd(client))

	// token

	var token = &cobra.Command{
		Use:   "token",
		Short: "Manage your tokens",
	}
	token.AddCommand(createAPITokenCmd(client))
	token.AddCommand(listAPITokensCmd(client))
	token.AddCommand(getAPITokenCmd(client))
	token.AddCommand(deleteAPITokenCmd(client))

	// healthcheck

	var healthcheck = &cobra.Command{
		Use:   "healthcheck",
		Short: "Manage your healthcheck",
	}
	healthcheck.AddCommand(deleteHealthcheckCmd(client))
	healthcheck.AddCommand(getHealthcheckCmd(client))
	healthcheck.AddCommand(listHealthchecksCmd(client))

	var dns = &cobra.Command{
		Use:   "dns",
		Short: "Manage your DNS healthchecks",
	}
	dns.AddCommand(createDNSHealthcheckCmd(client))
	dns.AddCommand(updateDNSHealthcheckCmd(client))

	var http = &cobra.Command{
		Use:   "http",
		Short: "Manage your HTTP healthchecks",
	}
	http.AddCommand(createHTTPHealthcheckCmd(client))
	http.AddCommand(updateHTTPHealthcheckCmd(client))

	var tcp = &cobra.Command{
		Use:   "tcp",
		Short: "Manage your TCP healthchecks",
	}
	tcp.AddCommand(createTCPHealthcheckCmd(client))
	tcp.AddCommand(updateTCPHealthcheckCmd(client))

	var tls = &cobra.Command{
		Use:   "tls",
		Short: "Manage your TLS healthchecks",
	}
	tls.AddCommand(createTLSHealthcheckCmd(client))
	tls.AddCommand(updateTLSHealthcheckCmd(client))

	var command = &cobra.Command{
		Use:   "command",
		Short: "Manage your Command healthchecks",
	}
	command.AddCommand(createCommandHealthcheckCmd(client))
	command.AddCommand(updateCommandHealthcheckCmd(client))

	var result = &cobra.Command{
		Use:   "result",
		Short: "Manage healthchecks results",
	}

	result.AddCommand(listHealthckecksResultsCmd(client))

	var metrics = &cobra.Command{
		Use:   "metrics",
		Short: "Manage healthchecks metrics",
	}

	metrics.AddCommand(getHealthchecksMetricsCmd(client))
	healthcheck.AddCommand(metrics)
	healthcheck.AddCommand(result)
	healthcheck.AddCommand(dns)
	healthcheck.AddCommand(command)
	healthcheck.AddCommand(tls)
	healthcheck.AddCommand(tcp)
	healthcheck.AddCommand(http)

	rootCmd.AddCommand(organization)
	rootCmd.AddCommand(account)
	rootCmd.AddCommand(token)
	rootCmd.AddCommand(healthcheck)

	return rootCmd.Execute()
}
