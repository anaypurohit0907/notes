/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"notes/data"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"os"
)

//var availableCategories = data.GetAllCategory()

type promptContent struct {
	errorMsg string
	label    string
}

var NewEntry string
var NewCategory string
var NewDefinition string

// newCmd represents the new command
var newCmd = &cobra.Command{

	Use:   "new",
	Short: "creates a new note.",
	Long:  `Use this to add a new command to your memory.`,
	//	ValidArgs: availableCategories,
	Run: func(cmd *cobra.Command, args []string) {
		if NewEntry != "" {
			fmt.Print("test")
			data.InsertNote(NewEntry, NewDefinition, NewCategory)

		} else {

			createNewNote()
		}
	},
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func createNewNote() {
	if NewEntry != "" && NewCategory != "" && NewDefinition != "" {
		data.InsertNote(NewEntry, NewDefinition, NewCategory)
	} else {

		wordPromptContent := promptContent{
			"Please provide a command or binding you learnt.",
			"What command would you like to make a note of?",
		}
		word := promptGetInput(wordPromptContent)

		definitionPromptContent := promptContent{
			"Please provide a definition.",
			fmt.Sprintf("What is the definition of %s?", word),
		}
		definition := promptGetInput(definitionPromptContent)
		categoryPromptContent := promptContent{
			"Please provide a category.",
			fmt.Sprintf("What category does %s belong to?", word),
		}
		category := promptGetSelect(categoryPromptContent)

		data.InsertNote(word, definition, category)
	}
}

func promptGetSelect(pc promptContent) string {
	items := []string{"vim keybind", "extention keybind", "shortcut"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&NewEntry, "New", "n", "", "make a new entry enter name of note")
	newCmd.Flags().StringVarP(&NewCategory, "Category", "c", "", "Category of note")
	newCmd.Flags().StringVarP(&NewDefinition, "Definition", "d", "", "Definition of note")
	newCmd.MarkFlagsRequiredTogether("New", "Category", "Definition")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
