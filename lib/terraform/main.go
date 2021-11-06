package terraform

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"os"

	"github.com/anvox/jinn-tf/lib/configure"
)

func Validate(conf configure.ActionWithArguments) {
	if !conf.InitConfig.IsValid() {
		log.Fatalf("Please configure before use. Run with --help for user guide.")
	}

	if !conf.CommandConfig.IsValid() {
		log.Fatalf("Please configure before use. Run with --help for user guide.")
	}

	validateStack(conf.InitConfig.StacksPath(), conf.CommandConfig.Stack)
	validateEnvironment(conf.InitConfig, conf.CommandConfig)
}

func validateStack(stacksPath string, stack string) {
	_, reserved := reservedStackName()[stack]
	if reserved {
		log.Fatalf("[ERROR] Invalid stack name [%s]", stack)
	}

	f, err := os.Open(stacksPath)
	if err != nil {
		log.Fatal("[ERROR] Cannot find stacks path")
	}

	files, err := f.Readdir(-1)
	if err != nil {
		log.Fatal("[ERROR] Failed to read stacks path")
	}

	for _, file := range files {
		if file.IsDir() && file.Name() == stack {
			return
		}
	}

	log.Fatalf("[ERROR] Cannot find stack [%s] definition in [%s]", stack, stacksPath)
}

func reservedStackName() map[string]bool {
	result := make(map[string]bool)

	result["modules"] = true

	return result
}

func validateEnvironment(initConfig configure.InitConfig, commandConfig configure.CommandConfig) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(initConfig.AwsRegion),
		Credentials: credentials.NewSharedCredentials("", initConfig.AwsProfile),
	})
	if err != nil {
		log.Fatal("[ERROR] Unable to connect to backend S3")
	}

	svc := s3.New(sess)
	input := &s3.HeadObjectInput{
		Bucket: aws.String(initConfig.StateBucket),
		Key:    aws.String(commandConfig.StateFilePath()),
	}

	_, backend_err := svc.HeadObject(input)
	if backend_err != nil {
		log.Printf("[WARN] State file for environment [%s], stack [%s] is not available.",
			commandConfig.Environment,
			commandConfig.Stack,
		)
	}
}
