package main

import (
	"log"
	"os"

	"github.com/anvox/jinn-tf/lib/configure"
	"github.com/anvox/jinn-tf/lib/help"
	"github.com/anvox/jinn-tf/lib/terraform"
)

func main() {
	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")

	action := configure.GetAction(os.Args)
	switch action.Action {
	case configure.Version:
		help.Validate(action)
		help.PrintVersion()
	case configure.Help:
		help.Validate(action)
		help.PrintHelp(action)
	case configure.Configure:
		configure.Validate(action)
		configure.Reconfigure()
	case configure.Terraform:
		terraform.Validate(action)
		config := action.CommandConfig
		log.Printf("Running command \"%s\" with environment=\"%s\" stack=\"%s\"", config.Command, config.Environment, config.Stack)
		terraform.ExecuteCommand(action.InitConfig, action.CommandConfig, isNoop())
	}
}

func isNoop() bool {
	if os.Getenv("NOOP") == "1" {
		return true
	}

	return false
}
