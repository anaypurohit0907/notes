/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-cli/data"
)

// noteCmd represents the note command

var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "Note a command you just learned along with a small description of what it does .",
	Long:  `Note a command you just learned along with a small description of what it does .`,
	Run: func(cmd *cobra.Command, args []string) {
		if NewEntry != "" {
			fmt.Print("test")
			data.InsertNote(NewEntry, NewDefinition, NewCategory)

		}

	},
}

func init() {
	rootCmd.AddCommand(noteCmd)

	noteCmd.Flags().StringVarP(&NewEntry, "New", "n", "", "make a new entry enter name of note")
	noteCmd.Flags().StringVarP(&NewCategory, "Category", "c", "", "Category of note")
	noteCmd.Flags().StringVarP(&NewDefinition, "Definition", "d", "", "Definition of note")
	noteCmd.MarkFlagsRequiredTogether("New", "Category", "Definition")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// noteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// noteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
