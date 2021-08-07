package cmd

import (
	"noteapp/db"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize db of noteapp.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		db.Open()
		defer db.Close()

		db.Initialize()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
