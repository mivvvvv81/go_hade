package command

import "project/framework/cobra"

func AddKernelCommands(root *cobra.Command) {

	root.AddCommand(DemoCommand)
	root.AddCommand(initAppCommand())
}
