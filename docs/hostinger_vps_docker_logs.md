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

