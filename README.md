# Jinn Terraform

A wrapper of terraform command, to support 1 code base, multiple resource groups/environments for terraform. 

## Usage

```shell
STATE_BUCKET=my-state-bucket STACKS_DIR=./stacks AWS_PROFILE=default AWS_REGION=us-east-1\
  jinn-tf -e <environment> -t <stack> <terraform-subcommand-with-arguments>

# Call with help to have a more verbose guide
jinn-tf -h
jinn-tf --help
```
