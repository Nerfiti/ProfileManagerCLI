package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mws",
	Short: "Profile manager",
	Long:  "Manage YAML profiles stored in current directory",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(newProfileCmd())
}
