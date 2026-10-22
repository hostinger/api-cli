## hostinger vps docker get

Get project contents

### Synopsis

Retrieves the complete project information including the docker-compose.yml
file contents, project metadata, and current deployment status.

This endpoint provides the full configuration and state details of a specific Docker Compose project.

Use this to inspect project settings, review the compose file, or check the overall project health.

```
hostinger vps docker get <virtual-machine-id> <project-name> [flags]
```

### Examples

```
  hostinger vps docker get 1268054 my-docker-project
```

### Options

```
  -h, --help   help for get
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

