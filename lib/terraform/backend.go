package terraform

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"github.com/anvox/jinn-tf/lib/configure"
)

const (
	BACKEND_TEMPLATE_PATH = "./lib/templates/backend.tf"
)

type BackendContext struct {
	StateBucket   string
	Environment   string
	Stack         string
	StateFilePath string
	AwsRegion     string
}

func prepareBackend(initConfig configure.InitConfig, tfConfig configure.CommandConfig) {
	backend_template, err := template.New(backendFile()).Parse(BACKEND_TEMPLATE)
	if err != nil {
		log.Fatal("Unable to generate backend template.")
	}

	backendContext := BackendContext{
		StateBucket:   initConfig.StateBucket,
		StateFilePath: tfConfig.StateFilePath(),
		AwsRegion:     initConfig.AwsRegion,
		Environment:   tfConfig.Environment,
		Stack:         tfConfig.Stack,
	}
	buf := new(bytes.Buffer)
	err = backend_template.Execute(buf, backendContext)
	if err != nil {
		log.Fatal("Unable to generate backend template 2.")
	}
	backend_path := initConfig.StacksPath() + tfConfig.Stack + "/" + backendFile()
	log.Printf("Writing content to: %s\n-----------\n%+v\n----------\n", backend_path, buf.String())

	f, err := os.Create(backend_path)
	if err != nil {
		log.Fatal("Unable to generate backend file.")
	}
	defer f.Close()

	buf.WriteTo(f)
}

func backendFile() string {
	return "backend.tf"
}
