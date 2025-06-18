/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/Blaze5333/todo-cli/internal/todo"
	"github.com/Blaze5333/todo-cli/internal/user"
	"github.com/Blaze5333/todo-cli/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := user.CheckSession()
		validate := func(input string) error {
			if input == "" {
				return fmt.Errorf("task ID cannot be empty")
			}
			_, err := strconv.Atoi(input)
			if err != nil {
				return fmt.Errorf("task ID must be a number")
			}
			return nil

		}
		prompt := promptui.Prompt{
			Label:    "Enter task ID to delete",
			Validate: validate,
		}
		taskID, err := prompt.Run()
		if err != nil {
			utils.ShowErrorMessage("Error reading task ID: " + err.Error())
			return
		}
		tk, _ := strconv.Atoi(taskID)
		err = todo.DeleteTask(name, tk)
		if err != nil {
			utils.ShowErrorMessage("Error deleting task: " + err.Error())
			return
		}
		utils.ShowSuccessMessage("Task deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
