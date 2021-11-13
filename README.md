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

### Configuration

It's recommended to use direnv to avoid missing environment variables configuration. Or use configure flag to create a `.jinn` file in working directory, `jinn-tf` will read this file first, then overwrite by values in environment variables. 

```shell
jinn-tf --configure
```

## How

`jinn-tf` enforce run terraform with variables named `environment` and `stack`. It change working directory to target `stack` under `STACKS_DIR`, create a backend file follow template and run terraform. 

The stacks must satisfy some conditions:

* Ignore `backend.tf` file, ignore tf backed part.
* Stack must keep resource name unique itself, like auto-gen by tf, or inject `environment` into it. 

### Note 🚨🚨🚨

Until https://github.com/anvox/jinn-tf/issues/1 is resolved, must run terraform `init -reconfig` manually when switching between environments. 

### Why not terraform workspace

I didn't have chance to use tf workspace, I'll try to use it. But this wrapper build before I known workspace, and it's fun to write it. 
