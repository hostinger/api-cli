## hostinger vps docker containers

Get project containers

### Synopsis

Retrieves a list of all containers belonging to a specific Docker Compose project on the virtual machine.

This endpoint returns detailed information about each container including
their current status, port mappings, and runtime configuration.

Use this to monitor the health and state of all services within your Docker Compose project.

```
hostinger vps docker containers <virtual-machine-id> <project-name> [flags]
```

### Examples

```
  hostinger vps docker containers 1268054 my-docker-project
```

### Options

```
  -h, --help   help for containers
```

### Options inherited from parent commands

```
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger vps docker](hostinger_vps_docker.md)	 - Docker Manager commands

