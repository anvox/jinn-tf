package configure

import (
	"fmt"
)

type CommandConfig struct {
	Environment string
	Stack       string
	Command     []string
}

func (cfg CommandConfig) StateFilePath() string {
	return fmt.Sprintf(
		"%s%s/%s/terraform.tfstate",
		stateFilePrefix(),
		cfg.Environment,
		cfg.Stack,
	)
}

func (cfg CommandConfig) IsValid() bool {
	return cfg.Environment != "" &&
		cfg.Stack != "" &&
		len(cfg.Command) > 0
}

func stateFilePrefix() string {
	return "jinn/infrastructure/"
}
