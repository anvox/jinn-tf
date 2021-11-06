package help

const GENERAL_HELP_TEXT = `NAME
	jinn-tf - A terraform wrapper. Dynamically gen backend config, to make flexible in using 1 code to provision multiple stages, follow some conventions.

SYNOPSIS
	jinn-tf [< -e | --environment environment > < -t | --stack stack > <terraform-subcommand>] [-h [stack | enviroment]]

DESCRIPTION
	jinn-tf is a wrapper of terraform command.
	Support multiple resource groups/environments for terraform in one codebase.
	This repo includes default config for AWS infrastructure and conventions to build and manage AWS resources.

	The options are as follows:

	-e environment, --environment environment
		Environment to run on, could be: production, preproduction, staging
		Use jinn-tf -h environment to see current available environments

	-s stack, --stack stack
		Stack to  to run on. Under folder ./stacks as default, each folder is blueprint of stack.
		Use jinn-tf -h stack to see current available stack.
		Customize stacks dir by env STACK_DIR

	terraform-subcommand
		Could be any terraform command with their arguments: init, init -reconfigure, plan, apply, import ...

	-h, --help [< t | stack | stacks > | < e | environment | environments >]
		Show this help text.
		If more options are passed:
		* t | stack | stacks : Show stack help info
		* e | environment | environments : Show environment help info

	Configuration:

	STATE_BUCKET are required to build backend file. The rest is set to default.

	Configuration could set by environment variables. It's recommended to use direnv to avoid missing them.
	If don't want to use direnv, we could use a ".jinn" file on the working directory. Run "golem-tf --configure" to init a new file.

	STATE_BUCKET
		Default: <current_dir>-remote-state
		S3 bucket to store terraform state

	AWS_REGION
		Default: us-east-1
		Config default region for terraform backend.

	AWS_PROFILE
		Default: default
		See AWS CLI configuration.
	STACK_DIR
		Default: ./stacks
		Path to directory contains stacks, from current directory.
`

const STACK_HELP_TEXT = `
Stack is blueprint of group of AWS resources, written in terraform script.
One stack could be used to provision multiple groups of AWS resources, for different environments.
`
const ENVIRONMENT_HELP_TEXT = `
Environment is an isolated space for services, applications, each used for specific purpose.
Like "production" is the most important space which used by our client.
"staging", "preproduction", ... is mostly for development.
`
