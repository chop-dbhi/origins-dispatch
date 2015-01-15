package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const version = "0.1.0"

// The main command describes the service and defaults to printing the
// help message.
var mainCmd = &cobra.Command{
	Use:   "origins-dispatch",
	Short: "Event dispatch service for Origins.",
	Long: `An HTTP service that consumes events produced by Origins and dispatches
them to subscribers.

Learn more about Origins: https://github.com/chop-dbhi/origins/`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

// The version command prints this service.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version.",
	Long:  "The version of the origins-dispatch service.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	mainCmd.AddCommand(versionCmd)

	viper.SetEnvPrefix("ORIGINS_DISPATCH")
	viper.AutomaticEnv()

	flags := mainCmd.Flags()

	flags.Bool("debug", false, "Turn on debugging.")
	flags.String("addr", "localhost:5002", "Address of the service")
	flags.String("neo4j", "http://localhost:7474/db/data/", "URI of the Neo4j HTTP server")
	flags.String("smtp-addr", "localhost:25", "Address of the SMTP server")
	flags.String("smtp-user", "", "User to authenticate with the SMTP server")
	flags.String("smtp-password", "", "Password to authenticate with the SMTP server")
	flags.String("email-from", "noreply@example.com", "The from email address.")

	viper.BindPFlag("debug", flags.Lookup("debug"))
	viper.BindPFlag("addr", flags.Lookup("addr"))
	viper.BindPFlag("neo4j", flags.Lookup("neo4j"))
	viper.BindPFlag("smtp_addr", flags.Lookup("smtp-addr"))
	viper.BindPFlag("smtp_user", flags.Lookup("smtp-user"))
	viper.BindPFlag("smtp_password", flags.Lookup("smtp-password"))
	viper.BindPFlag("email_from", flags.Lookup("email-from"))
}

func main() {
	mainCmd.Execute()
}
