package configure

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

func Reconfigure() {
	conf := LoadInitConfig()

	defaultAwsRegion := ""
	var awsRegionOptions []string
	for key, region := range allAwsRegions() {
		awsRegionOptions = append(awsRegionOptions, key)
		if defaultAwsRegion == "" {
			defaultAwsRegion = key
		}
		if region == conf.AwsRegion {
			defaultAwsRegion = key
		}
	}

	bucketMessage := "Please enter state bucket name:"
	bucketValidate := survey.Required
	if conf.StateBucket != "" {
		bucketMessage = "Please enter state bucket name: [" + conf.StateBucket + "]"
		bucketValidate = nil
	}
	var qs = []*survey.Question{
		{
			Name: "AwsRegion",
			Prompt: &survey.Select{
				Message: "Please chose region to store S3 backend:",
				Options: awsRegionOptions,
				Default: defaultAwsRegion,
			},
		},
		{
			Name:   "AwsProfile",
			Prompt: &survey.Input{Message: "Please AWS default profile to use with terraform: [" + conf.AwsProfile + "]"},
		},
		{
			Name:     "StateBucket",
			Prompt:   &survey.Input{Message: bucketMessage},
			Validate: bucketValidate,
		},
		{
			Name:   "StackDir",
			Prompt: &survey.Input{Message: "Please enter stacks directory: [" + conf.StackDir + "]"},
		},
	}

	conf2 := InitConfig{}
	err := survey.Ask(qs, &conf2)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		conf2.AwsRegion = allAwsRegions()[conf2.AwsRegion]
		// Merge for default value
		if conf2.StackDir == "" {
			conf2.StackDir = conf.StackDir
		}
		if conf2.AwsProfile == "" {
			conf2.AwsProfile = conf.AwsProfile
		}
		if conf2.StateBucket == "" {
			conf2.StateBucket = conf.StateBucket
		}

		verifyInitConfig(conf2)
		if ok := CreateInitConfig(conf2); ok {
			fmt.Println("jinn-tf is configured. You could use it now.")
		} else {
			fmt.Println("Something went wrong. jinn-tf is not configured yet. " +
				"Please check if .jinn file exists or additional errors.")
		}
	}
}

// Ref https://docs.aws.amazon.com/general/latest/gr/rande.html
func allAwsRegions() map[string]string {
	return map[string]string{
		"us-east-2 - US East (Ohio)":                "us-east-2",
		"us-east-1 - US East (N. Virginia)":         "us-east-1",
		"us-west-1 - US West (N. California)":       "us-west-1",
		"us-west-2 - US West (Oregon)":              "us-west-2",
		"af-south-1 - Africa (Cape Town)":           "af-south-1",
		"ap-east-1 - Asia Pacific (Hong Kong)":      "ap-east-1",
		"ap-south-1 - Asia Pacific (Mumbai)":        "ap-south-1",
		"ap-northeast-3 - Asia Pacific (Osaka)":     "ap-northeast-3",
		"ap-northeast-2 - Asia Pacific (Seoul)":     "ap-northeast-2",
		"ap-southeast-1 - Asia Pacific (Singapore)": "ap-southeast-1",
		"ap-southeast-2 - Asia Pacific (Sydney)":    "ap-southeast-2",
		"ap-northeast-1 - Asia Pacific (Tokyo)":     "ap-northeast-1",
		"ca-central-1 - Canada (Central)":           "ca-central-1",
		"cn-north-1 - China (Beijing)":              "cn-north-1",
		"cn-northwest-1 - China (Ningxia)":          "cn-northwest-1",
		"eu-central-1 - Europe (Frankfurt)":         "eu-central-1",
		"eu-west-1 - Europe (Ireland)":              "eu-west-1",
		"eu-west-2 - Europe (London)":               "eu-west-2",
		"eu-south-1 - Europe (Milan)":               "eu-south-1",
		"eu-west-3 - Europe (Paris)":                "eu-west-3",
		"eu-north-1 - Europe (Stockholm)":           "eu-north-1",
		"me-south-1 - Middle East (Bahrain)":        "me-south-1",
		"sa-east-1 - South America (São Paulo)":     "sa-east-1",
	}
}

func verifyInitConfig(cfg InitConfig) {
	// TODO: [AV] Implement
	// Verify state bucket is ready
	// Able to use profile
	// Able to list state  bucket
	// Verify stack dir is available
}
