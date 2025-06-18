package utils

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// Success message
func ShowSuccessMessage(message string) {
	fmt.Println(promptui.IconGood + promptui.Styler(promptui.FGGreen)(" Success: "+message))
}

// Error message
func ShowErrorMessage(message string) {
	fmt.Println(promptui.IconBad + promptui.Styler(promptui.FGRed)(message))
}

// Info message
func ShowInfoMessage(message string) {
	fmt.Println(promptui.IconWarn + promptui.Styler(promptui.FGBlue)(" Info: "+message))
}
