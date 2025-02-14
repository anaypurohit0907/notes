/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"notes/data"
	"strings"
)

var searchword, searchwordCategory string

// listCmd represents the list command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: " Search through all the notes you added",
	Long:  `Search through the list of all the notes you added`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if searchword != "" {
			data.SearchAllNotes(searchword)

		} else if searchwordCategory != "" {

			data.SearchByCategory(searchwordCategory)
		} else {
			temp := string(strings.Join(args, ""))
			data.SearchAllNotes(temp)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&searchword, "Search by definition", "d", "", "Search by definition")
	searchCmd.Flags().StringVarP(&searchwordCategory, "Search by category", "c", "", "Search by category")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
