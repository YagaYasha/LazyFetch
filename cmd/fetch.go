package cmd

import (
	"os"
	"strings"
)

func Kernel() string {
	kernel, _ := os.ReadFile("/proc/version")
	return strings.Split(string(kernel), " ")[2]
}

func Host() string {
	host, _ := os.Hostname()
	return host
}

func User() string {
	return os.Getenv("USER")
}

func Distro() string {
	distro, _ := os.ReadFile("/etc/os-release")
	return strings.Split(string(distro), "\"")[1]
}
