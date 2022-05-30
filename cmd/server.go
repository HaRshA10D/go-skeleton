package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/user-name/skeleton-name/app/server"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "starts the http server",
	RunE:  startServer,
}

func startServer(cmd *cobra.Command, args []string) error {

	app, err := server.InitializeApp()
	if err != nil {
		log.Printf("failed to create app: %s\n", err)
	}
	app.Start()
	return nil
}
