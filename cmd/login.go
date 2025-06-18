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

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
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
		loggedInUser, err := user.Login(username, password)
		if err != nil {
			utils.ShowErrorMessage("Error logging in: " + err.Error())
			return
		}
		err = user.SaveSession(loggedInUser.Username)
		if err != nil {
			utils.ShowErrorMessage("Error saving session: " + err.Error())
			return
		}
		utils.ShowSuccessMessage(fmt.Sprintf("User %s logged in successfully!", loggedInUser.Username))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
