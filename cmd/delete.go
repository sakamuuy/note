/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"noteapp/db"
	"noteapp/schema"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()

		switch args[0] {
		case schema.Folder.String():
			name, err := flags.GetString("name")
			if err != nil {
				cmd.PrintErr(err)
			}
			db.DeleteFolder(name)

		case schema.File.String():
			name, err := flags.GetString("name")
			if err != nil {
				cmd.PrintErr(err)
			}
			db.DeleteFile(name)

		case schema.Tag.String():
			name, err := flags.GetString("name")
			if err != nil {
				cmd.PrintErr(err)
			}
			db.DeleteTag(name)

		default:
			cmd.PrintErrf("no such command: %v \n", args[0])
		}
	},
}

func init() {
	deleteCmd.Flags().StringP("name", "n", "", "(Folder | File | Tags) name")
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
