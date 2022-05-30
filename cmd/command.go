package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "application for purpose xxx",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("welcome to the application, check commands")
	},
}

func init() {
	addCommands()
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func addCommands() {
	// add all command declaration here
	// will be called from main before root command executes
	// check example - version command
	rootCmd.AddCommand(versionCmd)
}
