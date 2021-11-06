package help

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"

	"github.com/anvox/jinn-tf/lib/configure"
)

func PrintUserGuide() {
	fmt.Println(GENERAL_HELP_TEXT)
}

func PrintVersion() {
	fmt.Printf("Current version: %s\n", Version)
}

func PrintStacks(conf configure.InitConfig) {
	fmt.Println(STACK_HELP_TEXT)

	if conf.IsValid() {
		fmt.Println("Available stacks:")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Stack", "Description"})
		table.SetAutoWrapText(false)
		table.SetBorder(false)

		stacks := configure.ListStacks(conf)
		for _, stack := range stacks {
			row := []string{stack, " "}
			table.Append(row)
		}

		table.Render()
	} else {
		fmt.Println("Please run with --configure to setup to fetch stacks")
	}
}

func PrintEnvironments(conf configure.InitConfig) {
	fmt.Println(ENVIRONMENT_HELP_TEXT)

	if conf.IsValid() {
		fmt.Println("Available environments:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Environment", "Stack", "Description"})
		table.SetAutoWrapText(false)
		table.SetBorder(false)

		environments := configure.ListEnvironments(conf)
		for environment, stacks := range environments {
			for _, stack := range stacks {
				row := []string{environment, stack, " "}
				table.Append(row)
			}
		}

		table.Render()
	} else {
		fmt.Println("Please run with --configure to setup to fetch environments from remote states")
	}
}

func PrintHelp(conf configure.ActionWithArguments) {
	if len(conf.CommandConfig.Command) == 1 {
		switch conf.CommandConfig.Command[0] {
		case "e", "environment", "environments":
			PrintEnvironments(conf.InitConfig)
		case "t", "stack", "stacks":
			PrintStacks(conf.InitConfig)
		}
		return
	}

	if len(conf.CommandConfig.Command) == 0 {
		PrintUserGuide()
		return
	}
}

func Validate(conf configure.ActionWithArguments) {
	if conf.Action == configure.Help {
		if len(conf.CommandConfig.Command) == 0 {
			return
		}
		if len(conf.CommandConfig.Command) == 1 {
			for _, supportedArg := range supportedArguments() {
				if supportedArg == conf.CommandConfig.Command[0] {
					return
				}
			}
		}
	} else if conf.Action == configure.Version {
		if len(conf.CommandConfig.Command) == 0 {
			return
		}
	}

	log.Fatal("Invalid arguments")
}

func supportedArguments() []string {
	return []string{
		"e",
		"environment",
		"environments",
		"t",
		"stack",
		"stacks",
	}
}
