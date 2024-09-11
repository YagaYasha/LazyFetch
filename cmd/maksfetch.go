package cmd

import (
	"fmt"
	"os"

	cobra "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "maksfetch",
	Short: "Yet another fetch",
	Run: func(cmd *cobra.Command, args []string) {
		maksfetch()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func maksfetch() {
	fmt.Printf(
		" %s\n Distribution: %s\n Kernel: %s\n", Userhost(), Distro(), Kernel(),
	//Userhost(),
	//Distro(),
	//Kernel(),
	)
}
