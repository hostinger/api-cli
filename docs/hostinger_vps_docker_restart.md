## hostinger vps docker restart

Restart project

### Synopsis

Restarts all services in a Docker Compose project by stopping and starting
containers in the correct dependency order.

This operation preserves data volumes and network configurations while refreshing the running containers.

Use this to apply configuration changes or recover from service failures.

```
hostinger vps docker restart <virtual-machine-id> <project-name> [flags]
```

### Examples

```
  hostinger vps docker restart 1268054 my-docker-project
```

### Options

```
  -h, --help                    help for restart
      --wait                    Wait until the action finishes
      --wait-timeout duration   Maximum time to wait with --wait (default 10m0s)
  -y, --yes                     Skip the confirmation prompt
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

