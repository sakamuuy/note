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
	"fmt"
	"log"
	"noteapp/db"
	"noteapp/schema"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show list of folders, files, tags.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()

		switch args[0] {
		case schema.Folder.String():
			folderNames := db.GetAllFolderName()

			prompt := promptui.Select{
				Label: "Select folder",
				Items: folderNames,
			}

			_, folder, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			actionPrompt := promptui.Select{
				Label: "Select action",
				Items: []string{"show files", "edit", "delete"},
			}

			_, action, err := actionPrompt.Run()

			switch action {
			case "show files":
				showFilePrompt(folder)

			case "edit":
				id := db.GetFolderByName(folder)

				prompt := promptui.Prompt{
					Label: "New name:",
					Validate: func(s string) error {
						return nil
					},
				}
				result, err := prompt.Run()
				if err != nil {
					log.Fatalln(err)
					return
				}

				db.PatchNewNameFolder(id, result)

			case "delete":
				db.DeleteFolder(folder)
			}

		case schema.File.String():
			folderName, err := flags.GetString("folder")
			if err != nil {
				log.Panic(err)
			}
			showFilePrompt(folderName)

		// case schema.Tag.String():
		// 	tagNames := db.GetAllTagName()

		// 	prompt := promptui.Select{
		// 		Label: "Select tag",
		// 		Items: tagNames,
		// 	}

		// 	_, tag, err := prompt.Run()
		// 	if err != nil {
		// 		fmt.Printf("Prompt failed %v\n", err)
		// 		return
		// 	}

		// 	actionPrompt := promptui.Select{
		// 		Label: "Select action",
		// 		Items: []string{"show files", "edit", "delete"},
		// 	}
		// 	_, action, err := actionPrompt.Run()

		default:
			cmd.PrintErrf("no such command: %v \n", args[0])
		}
	},
}

func init() {
	listCmd.Flags().StringP("folder", "f", "", "Name of the folder where the files will be stored.")
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showFilePrompt(folderName string) {
	fileNames := db.GetFilesFolderHas(folderName)

	prompt := promptui.Select{
		Label: "Select file.",
		Items: fileNames,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	actionPrompt := promptui.Select{
		Label: "Select action.",
		Items: []string{"write", "edit", "delete"},
	}

	_, action, err := actionPrompt.Run()

	switch action {
	case "write":
		contentsname := db.GetFileContentsByName(result)
		runVim(strings.Join([]string{"./contents/", contentsname, ".md"}, ""))
	case "edit":

	case "delete":

	default:
		return
	}
}
