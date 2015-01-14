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

// The serve command runs the service.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs the service.",
	Long:  "The serve command runs the service.",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	mainCmd.AddCommand(versionCmd)
	mainCmd.AddCommand(serveCmd)

	viper.SetEnvPrefix("ORIGINS_DISPATCH")
	viper.AutomaticEnv()

	flags := mainCmd.PersistentFlags()
	flags.Bool("debug", false, "Turn on debugging.")

	viper.BindPFlag("debug", flags.Lookup("debug"))

	flags = serveCmd.Flags()
	flags.String("host", "localhost", "Host address to bind to.")
	flags.Int("port", 5002, "Host port to bind to.")
	flags.String("neo4j", "http://localhost:7474/db/data/", "URI of the Neo4j server.")

	viper.BindPFlag("serve_host", flags.Lookup("host"))
	viper.BindPFlag("serve_port", flags.Lookup("port"))
	viper.BindPFlag("serve_neo4j", flags.Lookup("neo4j"))
}

func main() {
	mainCmd.Execute()
}
