package configure

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

const CONF_FILE = ".jinn"

type InitConfig struct {
	StateBucket string
	AwsRegion   string
	StackDir    string
	AwsProfile  string
}

func (cfg InitConfig) IsValid() bool {
	return cfg.AwsRegion != "" &&
		cfg.StateBucket != "" &&
		cfg.StackDir != ""
}

func (cfg InitConfig) StacksPath() string {
	return getwd() + "/" + cfg.StackDir + "/"
}

func CreateInitConfig(cfg InitConfig) bool {
	if _, err := os.Stat(getwd() + CONF_FILE); os.IsNotExist(err) {
		if file, fErr := os.Create(getwd() + "/" + CONF_FILE); fErr == nil {
			defer file.Close()

			wrapper := TfConfigWrapper{
				Tf: cfg,
			}
			if string_config, tomErr := toml.Marshal(wrapper); tomErr == nil {
				file.WriteString(string(string_config))
				return true
			}
		}
	}

	return false
}

func LoadInitConfig() InitConfig {
	return loadInitConfigFromEnv(
		loadInitConfigFromFile(defaultInitConfig()),
	)
}

func defaultInitConfig() InitConfig {
	return InitConfig{
		StateBucket: "",
		AwsRegion:   "us-east-1",
		StackDir:    "stacks",
		AwsProfile:  "default",
	}
}

type TfConfigWrapper struct {
	Tf InitConfig
}

func loadInitConfigFromFile(config InitConfig) InitConfig {
	if content, err := os.ReadFile(getwd() + "/" + CONF_FILE); err == nil {
		wrapper := TfConfigWrapper{
			Tf: config,
		}
		if unmarshalErr := toml.Unmarshal(content, &wrapper); unmarshalErr == nil {
			return wrapper.Tf
		}
	}

	return config
}

func loadInitConfigFromEnv(config InitConfig) InitConfig {
	if stateBucket := os.Getenv("STATE_BUCKET"); stateBucket != "" {
		config.StateBucket = stateBucket
	}
	if awsRegion := os.Getenv("AWS_REGION"); awsRegion != "" {
		config.AwsRegion = awsRegion
	}
	if stackDir := os.Getenv("STACK_DIR"); stackDir != "" {
		config.StackDir = stackDir
	}
	if awsProfile := os.Getenv("AWS_PROFILE"); awsProfile != "" {
		config.AwsProfile = awsProfile
	}

	return config
}
