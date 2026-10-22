## hostinger vps docker list

Get project list

### Synopsis

Retrieves a list of all Docker Compose projects currently deployed on the virtual machine.

This endpoint returns basic information about each project including name,
status, file path and list of containers with details about their names,
image, status, health and ports. Container stats are omitted in this
endpoint. If you need to get detailed information about container with
stats included, use the `Get project containers` endpoint.

Use this to get an overview of all Docker projects on your VPS instance.

```
hostinger vps docker list <virtual-machine-id> [flags]
```

### Examples

```
  hostinger vps docker list 1268054
```

### Options

```
  -h, --help   help for list
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

