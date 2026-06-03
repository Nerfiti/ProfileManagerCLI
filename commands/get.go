package commands

import (
	"mws/profile"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newGetCmd() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Show a profile in YAML format",
		RunE: func(cmd *cobra.Command, args []string) error {
			p, err := profile.Get(name)
			if err != nil {
				return err
			}
			data, err := yaml.Marshal(p)
			if err != nil {
				return err
			}
			cmd.OutOrStdout().Write(data)
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "profile name")
	cmd.MarkFlagRequired("name")
	return cmd
}
