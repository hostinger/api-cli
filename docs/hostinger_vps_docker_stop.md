## hostinger vps docker stop

Stop project

### Synopsis

Stops all running services in a Docker Compose project while preserving
container configurations and data volumes.

This operation gracefully shuts down containers in reverse dependency order.

Use this to temporarily halt a project without removing data or configurations.

```
hostinger vps docker stop <virtual-machine-id> <project-name> [flags]
```

### Examples

```
  hostinger vps docker stop 1268054 my-docker-project
```

### Options

```
  -h, --help                    help for stop
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

