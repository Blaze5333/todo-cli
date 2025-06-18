package cmd

import (
	"fmt"

	"github.com/Blaze5333/todo-cli/internal/todo"
	"github.com/Blaze5333/todo-cli/internal/user"
	"github.com/Blaze5333/todo-cli/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := user.CheckSession()

		prompt := promptui.Prompt{
			Label: "Enter task Title",
		}
		title, err := prompt.Run()
		if err != nil {
			utils.ShowErrorMessage("Error reading title: " + err.Error())
			return
		}
		if title == "" {
			utils.ShowErrorMessage("Title cannot be empty")
			return
		}
		prompt = promptui.Prompt{
			Label: "Enter task Description",
		}
		description, err := prompt.Run()
		if err != nil {
			utils.ShowErrorMessage("Error reading description: " + err.Error())
			return
		}
		if description == "" {
			utils.ShowErrorMessage("Description cannot be empty")
			return
		}
		prompt1 := promptui.Select{
			Label: "Enter the priority of the task (low, medium, high)",
			Items: []string{"high 🔴", "medium 🟠", "low 🟡"},
		}

		priority, _, err := prompt1.Run()
		if err != nil {
			utils.ShowErrorMessage("Error reading priority: " + err.Error())
			return
		}
		fmt.Println("Priority selected:", priority)
		task, err := todo.AddTask(name, title, description, priority)
		if err != nil {
			utils.ShowErrorMessage("Error adding task: " + err.Error())
			return
		}
		utils.ShowSuccessMessage(fmt.Sprintf("task '%s' added successfully with priority %d!", task.Title, priority))

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
