package cmd

import (
	"fmt"
)

const (
	Reset        = "\033[0m"
	LightMagenta = "\033[95m"
	LightRed     = "\033[91m"
	Blue         = "\033[34m"
	Green        = "\033[32m"
)

func Info() {
	fmt.Printf("%s%s%s@%s%s%s", LightMagenta, User(), Reset, LightRed, Host(), Reset)
	fmt.Println("")
	fmt.Printf("%sDistribution%s: %s", Blue, Reset, Distro())
	fmt.Println("")
	fmt.Printf("%sKernel%s: %s \n", Green, Reset, Kernel())
}
