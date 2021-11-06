package terraform

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/anvox/jinn-tf/lib/configure"
)

func buildCommand(initConfig configure.InitConfig, tfConfig configure.CommandConfig) *exec.Cmd {
	command := exec.Command("terraform", tfConfig.Command...)
	command.Dir = initConfig.StacksPath() + tfConfig.Stack
	return command
}

func executeCommandNoop(initConfig configure.InitConfig, tfConfig configure.CommandConfig) {
	command := buildCommand(initConfig, tfConfig)
	fmt.Printf("%+v\n", command)
}

func ExecuteCommand(initConfig configure.InitConfig, tfConfig configure.CommandConfig, isNoop bool) {
	if isNoop {
		executeCommandNoop(initConfig, tfConfig)
		return
	}

	prepareBackend(initConfig, tfConfig)

	command := buildCommand(initConfig, tfConfig)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	err := command.Run()
	if err != nil {
		log.Fatalf("[ERROR] Failed to run: %+v\n", err)
	}
}
