package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.9.0"

var mainCmd = &cobra.Command{
	Use:   "origins-webhooks-service",
	Short: "Origins Webhooks service",
	Long:  "A service that broadcasts event payloads to registered webhooks.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  "The version of the Origins Webhooks service.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

var (
	// Global
	debug bool

	// Service
	serveHost  string
	servePort  int
	serveNeo4j string
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs the service",
	Long:  "Runs the service",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	// Main flags
	mainFlags := serveCmd.PersistentFlags()
	mainFlags.BoolVar(&debug, "debug", false, "Turn on debugging.")

	// Serve command flags
	serveFlags := serveCmd.Flags()
	serveFlags.StringVar(&serveHost, "host", "localhost", "Host address to bind to.")
	serveFlags.IntVar(&servePort, "port", 5002, "Host port to bind to.")
	serveFlags.StringVar(&serveNeo4j, "neo4j", "http://localhost:7474/db/data/", "URI of the Neo4j server.")

	// Register subcommands
	mainCmd.AddCommand(versionCmd)
	mainCmd.AddCommand(serveCmd)
}

func main() {
	mainCmd.Execute()
}
