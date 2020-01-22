package main

import (
	"os"
	"github.com/CiscoDevNet/icube-cybervision-plugin/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
