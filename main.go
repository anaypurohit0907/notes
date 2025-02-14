/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"notes/cmd"
	"notes/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
