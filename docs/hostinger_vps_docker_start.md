## hostinger vps docker start

Start project

### Synopsis

Starts all services in a Docker Compose project that are currently stopped.

This operation brings up containers in the correct dependency order as defined in the compose file.

Use this to resume a project that was previously stopped or to start services after a system reboot.

```
hostinger vps docker start <virtual-machine-id> <project-name> [flags]
```

### Examples

```
  hostinger vps docker start 1268054 my-docker-project
```

### Options

```
  -h, --help                    help for start
      --wait                    Wait until the action finishes
      --wait-timeout duration   Maximum time to wait with --wait (default 10m0s)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger vps docker](hostinger_vps_docker.md)	 - Docker Manager commands

