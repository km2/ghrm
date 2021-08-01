package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/km2/ghrm"
	"github.com/spf13/cobra"
)

const tokenGenURL = "https://github.com/settings/tokens/new?description=ghrm&scopes=delete_repo" //nolint:gosec

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Set token via prompt",
	Long:  "Set token via prompt.",
	RunE:  runToken,
}

func runToken(cmd *cobra.Command, args []string) error {
	fmt.Fprintln(os.Stdout, tokenGenURL)

	var token ghrm.Token

	prompt := &survey.Input{
		Message: "Input your token:",
	}
	if err := survey.AskOne(prompt, &token.Token); err != nil {
		return fmt.Errorf("failed to ask token: %w", err)
	}

	bytes, err := json.Marshal(token)
	if err != nil {
		return fmt.Errorf("failed to marshal: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(ghrm.DefaultTokenPath()), os.ModePerm); err != nil {
		return fmt.Errorf("failed to mkdir: %w", err)
	}

	if err := os.WriteFile(ghrm.DefaultTokenPath(), bytes, os.ModePerm); err != nil {
		return fmt.Errorf("failed to write token to JSON: %w", err)
	}

	fmt.Fprintf(os.Stdout, "Your token is stored in %s\n", ghrm.DefaultTokenPath())

	return nil
}

func init() {
	rootCmd.AddCommand(tokenCmd)
}
