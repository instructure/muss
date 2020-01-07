package cmd

import (
	"github.com/spf13/cobra"
)

func newStartCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "start",
		Short: "Start services",
		Long:  "Start existing containers.",
		Args:  cobra.ArbitraryArgs,
		// TODO: ArgsInUseLine: "[service...]"
		PreRun: configSavePreRun,
		RunE: func(cmd *cobra.Command, args []string) error {
			return DelegateCmd(
				cmd,
				dockerComposeCmd(cmd, args),
			)
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(newStartCommand())
}