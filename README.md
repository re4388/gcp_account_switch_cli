# GCP Account Switcher

```bash

## need to have below file ready in the config path


.config/gcloud/configurations
➜ ls                                                                                      
config_default  config_default-bk  config_default_personal  config_default_wemo

## modify the path for you need
- say I have another gcp config, just create another one, like `config_default_XXX`
- I do need to update my code to include this new path `config_default_personal`, but it shall be easy to add 


## TODO:
- update code to having a strut for config and use slice to handle, so I don't need to update the code


```

`./gcpSwticher -h`

```
NAME:
   GCP Account Switcher - ./gcpSwticher wemo  or ./gcpSwticher personal

USAGE:
   GCP Account Switcher [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
done



```