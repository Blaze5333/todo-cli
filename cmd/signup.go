/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Blaze5333/todo-cli/internal/user"
	"github.com/Blaze5333/todo-cli/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// signupCmd represents the signup command
var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Prompt{Label: "Username"}
		username, err := prompt.Run()
		if err != nil {
			utils.ShowErrorMessage("Error reading username: " + err.Error())
			return
		}
		prompt = promptui.Prompt{Label: "Password", Mask: '*'}
		password, err1 := prompt.Run()
		if err1 != nil {
			utils.ShowErrorMessage("Error reading password: " + err1.Error())
			return
		}
		userRegistered, err := user.Register(username, password)
		if err != nil {
			utils.ShowErrorMessage("Error registering user: " + err.Error())
			return
		}
		utils.ShowSuccessMessage(fmt.Sprintf("User %s registered successfully! ,You can now Login ", userRegistered.Username))

	},
}

func init() {
	rootCmd.AddCommand(signupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
