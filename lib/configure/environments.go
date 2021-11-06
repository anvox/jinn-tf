package configure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"strings"
)

func ListEnvironments(cfg InitConfig) map[string][]string {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.AwsRegion),
		Credentials: credentials.NewSharedCredentials("", cfg.AwsProfile),
	})
	if err != nil {
		log.Fatal("[ERROR] Unable to connect to backend S3")
	}

	svc := s3.New(sess)

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(cfg.StateBucket),
		Prefix: aws.String(stateFilePrefix()),
	}

	result, err := svc.ListObjectsV2(input)
	if err != nil {
		log.Fatal("[ERROR] Unable to connect to backend S3")
	}

	environments := make(map[string][]string)
	for _, s3_object := range result.Contents {
		stage := strings.TrimPrefix(aws.StringValue(s3_object.Key), stateFilePrefix())
		key_parts := strings.Split(stage, "/")
		environments[key_parts[0]] = append(environments[key_parts[0]], key_parts[1])
	}

	return environments
}
