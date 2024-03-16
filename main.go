package main

import (
	"github.com/fatih/color"
	"strings"
	"os"
)

func main() {
	osRelease, _ := os.ReadFile("/etc/os-release")
	productName, _ := os.ReadFile("/sys/devices/virtual/dmi/id/product_name")
	hostName, _ := os.ReadFile("/etc/hostname")

	lines := strings.Split(string(osRelease), "\n")
	username, _ := os.LookupEnv("USER")

	for i := range lines {
		if strings.HasPrefix(lines[i], "PRETTY") {
			color.Blue("\n  You are in fucking " + lines[i][13:len(lines[i]) - 1])
			color.Green("  On a piece of shit named " + string(productName))
			color.Yellow("  Your wretched hostname is " + string(hostName)[:len(string(hostName)) - 1])
			color.Red("  Your stupid username is " + username)
		}
	}
}
