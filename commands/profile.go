package commands

import (
	"github.com/spf13/cobra"
)

func newProfileCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profile",
		Short: "Manage profiles",
	}

	cmd.AddCommand(
		newCreateCmd(),
		newGetCmd(),
		newListCmd(),
		newDeleteCmd(),
	)

	return cmd
}
