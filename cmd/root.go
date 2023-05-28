package cmd

import (
	"github.com/appclacks/cli/client"
	"github.com/spf13/cobra"
)

var outputFormat string
var profile string
var appclacksURL = "https://api.appclacks.com"

func Execute() error {
	var rootCmd = &cobra.Command{
		Use:   "appclacks",
		Short: "Appclacks CLI",
	}
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "Command output format (table or json)")
	rootCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "", "Profile to use in the configuration file")

	// login
	rootCmd.AddCommand(loginCmd())

	// organization
	var organization = &cobra.Command{
		Use:   "organization",
		Short: "Manage your organization",
	}
	organization.AddCommand(createOrganizationCmd())
	// account

	var account = &cobra.Command{
		Use:   "account",
		Short: "Manage your account",
	}
	account.AddCommand(getAccountOrganizationCmd())

	var password = &cobra.Command{
		Use:   "password",
		Short: "Manage your account password",
	}
	password.AddCommand(changePasswordCmd())
	password.AddCommand(sendResetPasswordLink())
	account.AddCommand(password)

	// token

	var token = &cobra.Command{
		Use:   "token",
		Short: "Manage your tokens",
	}
	token.AddCommand(createAPITokenCmd())
	token.AddCommand(listAPITokensCmd())
	token.AddCommand(getAPITokenCmd())
	token.AddCommand(deleteAPITokenCmd())

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

	var result = &cobra.Command{
		Use:   "result",
		Short: "Manage healthchecks results",
	}

	result.AddCommand(listHealthckecksResultsCmd())

	var metrics = &cobra.Command{
		Use:   "metrics",
		Short: "Manage healthchecks metrics",
	}

	metrics.AddCommand(getHealthchecksMetricsCmd())
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

func buildClient() *client.Client {
	client, err := client.New(appclacksURL, client.WithProfile(profile))
	exitIfError(err)
	return client
}
