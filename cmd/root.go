package cmd

import (
	"github.com/CiscoDevNet/icube-cybervision-plugin/beater"
	cmd "github.com/elastic/beats/libbeat/cmd"
)

// Name of this beat
var Name = "icube-cybervision-plugin"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmd(Name, "", beater.New)
