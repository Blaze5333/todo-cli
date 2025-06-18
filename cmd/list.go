/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Blaze5333/todo-cli/internal/bubbletea"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// name := user.CheckSession()
		// tasks, err := todo.GetTasks(name)
		// if err != nil {
		// 	utils.ShowErrorMessage("Error fetching tasks:")
		// 	return
		// }
		// tbl := table.New(os.Stdout)
		// tbl.SetHeaders("ID", "Title", "Description", "Completed", "Priority", "Created At", "Updated At")

		// tbl.SetHeaderStyle(table.StyleBold)
		// tbl.SetLineStyle(table.StyleMagenta)
		// tbl.SetDividers(table.UnicodeRoundedDividers)

		// for index, task := range tasks {

		// 	completed := "❌"
		// 	if task.Done {
		// 		completed = "✅"
		// 	}
		// 	var color string
		// 	switch task.Priority {
		// 	case 0:
		// 		color = "high 🔴"
		// 	case 1:
		// 		color = "medium 🟠"
		// 	case 2:
		// 		color = "low 🟡"
		// 	default:
		// 		color = "⚪️"
		// 	}
		// 	tbl.AddRow(
		// 		strconv.Itoa(index),
		// 		task.Title,
		// 		task.Description,
		// 		completed,
		// 		color,
		// 		task.CreatedAt.Format("2006-01-02 15:04:05"),
		// 		task.UpdatedAt.Format("2006-01-02 15:04:05"),
		// 	)
		// }
		// tbl.Render()
		bubbletea.Start()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
