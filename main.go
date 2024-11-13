/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"go-cli/cmd"
	"go-cli/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
