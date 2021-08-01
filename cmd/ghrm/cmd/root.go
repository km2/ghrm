package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/km2/ghrm"
	"github.com/spf13/cobra"
)

type repository struct {
	owner string
	repo  string
}

var (
	repositories []repository
	version      bool
)

var rootCmd = &cobra.Command{
	Use:   "ghrm",
	Short: "Just remove GitHub repositories",
	Long:  "Just remove GitHub repositories.",
	Args: func(cmd *cobra.Command, args []string) error {
		for _, arg := range args {
			ss := strings.Split(arg, "/")
			if len(ss) != 2 {
				return fmt.Errorf("invalid format: %s", arg)
			}

			repositories = append(repositories, repository{
				owner: ss[0],
				repo:  ss[1],
			})
		}

		return nil
	},
	RunE: runRoot,
}

func runRoot(cmd *cobra.Command, args []string) error {
	if version {
		return runVersion(cmd, args)
	}

	token, err := ghrm.ReadToken(ghrm.DefaultTokenPath())
	if err != nil {
		return fmt.Errorf("failed to read token: %w", err)
	}

	cli := ghrm.New(token)
	for _, repository := range repositories {
		if err := cli.RemoveRepository(repository.owner, repository.repo); err != nil {
			return fmt.Errorf("failed to remove repository: %w", err)
		}

		fmt.Fprintf(os.Stdout, "%s/%s was removed successfully\n", repository.owner, repository.repo)
	}

	return nil
}

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "V", false, "show version")
}

func Execute() {
	rootCmd.Execute() //nolint:errcheck
}
