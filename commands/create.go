package commands

import (
	"fmt"

	"mws/profile"

	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var name, user, project string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := profile.Create(name, user, project); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "created: %s\n", name)
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "profile name")
	cmd.Flags().StringVar(&user, "user", "", "user name")
	cmd.Flags().StringVar(&project, "project", "", "project name")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("user")
	cmd.MarkFlagRequired("project")

	return cmd
}
