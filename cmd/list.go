/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go-cli/data"

	"github.com/spf13/cobra"
)

<<<<<<< HEAD
var searchword string

=======
>>>>>>> f8a8337 (initial commit)
// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View a list of all the notes you added",
	Long:  `view a list of all the notes you added`,
	Run: func(cmd *cobra.Command, args []string) {
<<<<<<< HEAD
		if searchword != "" {
			data.SearchAllNotes(searchword)

		} else {

			data.DisplayAllNotes()
		}
=======
		data.DisplayAllNotes()
>>>>>>> f8a8337 (initial commit)
	},
}

func init() {
	noteCmd.AddCommand(listCmd)
<<<<<<< HEAD
	listCmd.Flags().StringVarP(&searchword, "Search", "s", "", "Search by definition")
=======
>>>>>>> f8a8337 (initial commit)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
