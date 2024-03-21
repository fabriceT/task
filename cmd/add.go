// Package cmd provides functionality related to the command line.
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run:   add,
}

func init() {
	RootCommand.AddCommand(addCommand)
}

func add(_ *cobra.Command, args []string) {
	task := strings.Join(args, " ")
	fmt.Printf("\"%s\" to your task list. \n", task)
}
