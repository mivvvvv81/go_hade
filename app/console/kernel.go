package console

import (
	"project/app/console/command/demo"
	"project/framework"
	"project/framework/cobra"
	"project/framework/command"
)

// RunCommand 初始化根command并运行
func RunCommand(container framework.Container) error {
	// 根command
	var rootCmd = &cobra.Command{
		Use:   "hade",
		Short: "hade命令",
		Long:  "hade 框架提供的命令工具，使用这个命令行工具能很方便的执行框自带命令，也很方便编写业务命令",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	rootCmd.SetContainer(container)
	command.AddKernelCommands(rootCmd)
	AddAppCommand(rootCmd)

	//执行rootCommand
	return rootCmd.Execute()

}

func AddAppCommand(rootCmd *cobra.Command) {
	//demo 例子
	rootCmd.AddCommand(demo.InitFoo())
}
