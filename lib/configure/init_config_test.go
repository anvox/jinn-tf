package configure

import (
	"os"
	"testing"
)

func testLoadInitConfigDefault(t *testing.T) {
	config := LoadInitConfig()
	if config.AwsRegion != "us-east-1" {
		t.Fatal("Failed")
	}
}

func testLoadInitConfigFromFile(t *testing.T) {
	working_dir, _ := os.Getwd()
	f, _ := os.Create(working_dir + "/.jinn")
	defer os.Remove(working_dir + "/.jinn")

	jinn_config := `[Tf]
AwsRegion = "ap-southeast-1"
`
	_, _ = f.WriteString(jinn_config)
	f.Close()
	config := LoadInitConfig()
	if config.AwsRegion != "ap-southeast-1" {
		t.Fatal("Failed")
	}
}

func testLoadInitConfigFromEnvironmentVariables(t *testing.T) {
	os.Setenv("AWS_REGION", "ap-southeast-2")
	defer os.Unsetenv("AWS_REGION")

	config := LoadInitConfig()
	if config.AwsRegion != "ap-southeast-2" {
		t.Fatal("Failed")
	}
}

func TestLoadInitConfig(t *testing.T) {
	testLoadInitConfigDefault(t)
	testLoadInitConfigFromFile(t)
	testLoadInitConfigFromEnvironmentVariables(t)
}

func TestReloadInitConfig(t *testing.T) {
	config := LoadInitConfig()
	if config.AwsRegion != "us-east-1" {
		t.Fatal("Failed")
	}

	config.AwsRegion = "ap-southeast-1"

	CreateInitConfig(config)
	defer os.Remove(getwd() + "/.jinn")

	config = LoadInitConfig()
	if config.AwsRegion != "ap-southeast-1" {
		t.Fatal("Failed")
	}
}
