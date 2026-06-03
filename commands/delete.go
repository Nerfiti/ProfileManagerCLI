package commands

import (
	"fmt"

	"mws/profile"

	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := profile.Delete(name); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "deleted: %s\n", name)
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "profile name")
	cmd.MarkFlagRequired("name")
	return cmd
}
