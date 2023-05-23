package cmd

import (
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "shorty",
	}

	cmd.AddCommand(
		newStartCmd(),
	)

	return cmd
}

func Execute() error {
	return newRootCmd().Execute()
}
