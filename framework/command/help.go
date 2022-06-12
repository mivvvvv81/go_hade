package command

import (
	"fmt"
	"project/framework/cobra"
	"project/framework/contract"
)

var DemoCommand = &cobra.Command{
	Use:   "help",
	Short: "demo for framework",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		fmt.Println("app base folder:", appService)
	},
}
