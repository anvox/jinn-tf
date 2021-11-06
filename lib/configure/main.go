package configure

import (
	"log"
	"nullprogram.com/x/optparse"
	"os"
)

type HelpConfig struct {
	Type string
}

type Action int

const (
	Version Action = iota
	Help
	Configure
	Terraform
	Unknown
)

type ActionWithArguments struct {
	Action        Action
	InitConfig    InitConfig
	CommandConfig CommandConfig
}

func GetAction(args []string) ActionWithArguments {
	options := []optparse.Option{
		{"configure", 'c', optparse.KindNone},
		{"environment", 'e', optparse.KindRequired},
		{"stack", 't', optparse.KindRequired},
		{"help", 'h', optparse.KindNone},
		{"version", 'v', optparse.KindNone},
	}

	results, rest, err := optparse.Parse(options, args)
	if err != nil {
		log.Fatalf("Unable to parse command arguments: %+v\n", err)
	}

	action := Terraform
	var environment, stack string
	for _, result := range results {
		switch result.Long {
		case "help":
			action = Help
		case "version":
			action = Version
		case "configure":
			action = Configure
		case "environment":
			environment = result.Optarg
		case "stack":
			stack = result.Optarg
		}
	}

	return ActionWithArguments{
		Action:     action,
		InitConfig: LoadInitConfig(),
		CommandConfig: CommandConfig{
			Environment: environment,
			Stack:       stack,
			Command:     rest,
		},
	}
}

func Validate(conf ActionWithArguments) {

}

func getwd() string {
	if working_dir, err := os.Getwd(); err != nil {
		return "."
	} else {
		return working_dir
	}
}
