package cmd

import (
	"os"
	"strings"
)

func Kernel() string {
	kernel, _ := os.ReadFile("/proc/version")
	return strings.Split(string(kernel), " ")[2]
}

func Userhost() string {
	host, _ := os.Hostname()
	return os.Getenv("USER") + "@" + host
}

//Host: os.Hostname()

func Distro() string {
	distro, _ := os.ReadFile("/etc/os-release")
	return strings.Split(string(distro), "\"")[1] //I dunno
}
