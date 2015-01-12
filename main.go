package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

var (
	// Global
	debug bool

	// Service
	serveHost  string
	servePort  int
	serveNeo4j string
)

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

func main() {
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

	mainCmd.Execute()
}
