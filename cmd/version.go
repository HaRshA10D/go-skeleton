package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "tells you the app version",
	RunE:  showVersion,
}

func showVersion(cmd *cobra.Command, args []string) error {

	fileData, err := os.ReadFile(".version")
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("app version: v%s", string(fileData)))
	return nil
}
