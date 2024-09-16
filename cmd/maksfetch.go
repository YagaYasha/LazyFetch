package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const (
	Reset        = "\033[0m"
	LightMagenta = "\033[95m"
	LightRed     = "\033[91m"
	Blue         = "\033[34m"
	Green        = "\033[32m"
)

var rootCmd = &cobra.Command{
	Use:   "lazyfetch",
	Short: "The most lazy fetch",
	Run: func(cmd *cobra.Command, args []string) {
		lazyfetch()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func lazyfetch() {

	printinfo()

	fmt.Println("")
	for i := 0; i < 8; i++ {
		fmt.Printf("\033[4%dm", i)
	}
	fmt.Println("")
	for i := 0; i < 8; i++ {
		fmt.Printf("\033[10%dm", i)
	}
}

func printinfo() {
	fmt.Printf("%s%s%s@%s%s%s", LightMagenta, User(), Reset, LightRed, Host(), Reset)
	fmt.Println("")
	fmt.Printf("%sDistribution%s: %s", Blue, Reset, Distro())
	fmt.Println("")
	fmt.Printf("%sKernel%s: %s", Green, Reset, Kernel())
}
