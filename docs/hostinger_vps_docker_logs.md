## hostinger vps docker logs

Get project logs

### Synopsis

Retrieves aggregated log entries from all services within a Docker Compose project.

This endpoint returns recent log output from each container, organized by service name with timestamps.
The response contains the last 300 log entries across all services.

Use this for debugging, monitoring application behavior, and
troubleshooting issues across your entire project stack.

```
hostinger vps docker logs <virtual-machine-id> <project-name> [flags]
```

### Examples

```
  hostinger vps docker logs 1268054 my-docker-project
```

### Options

```
  -h, --help   help for logs
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

