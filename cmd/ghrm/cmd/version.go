package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show version.",
	RunE:  runVersion,
}

func runVersion(cmd *cobra.Command, args []string) error {
	fmt.Fprintf(os.Stdout, "v%s\n", Version)

	return nil
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
