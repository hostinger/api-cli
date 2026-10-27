## hostinger vps docker

Docker Manager commands

```
hostinger vps docker [flags]
```

### Options

```
  -h, --help   help for docker
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

* [hostinger vps](hostinger_vps.md)	 - VPS commands
* [hostinger vps docker containers](hostinger_vps_docker_containers.md)	 - Get project containers
* [hostinger vps docker create](hostinger_vps_docker_create.md)	 - Create new project
* [hostinger vps docker delete](hostinger_vps_docker_delete.md)	 - Delete project
* [hostinger vps docker get](hostinger_vps_docker_get.md)	 - Get project contents
* [hostinger vps docker list](hostinger_vps_docker_list.md)	 - Get project list
* [hostinger vps docker logs](hostinger_vps_docker_logs.md)	 - Get project logs
* [hostinger vps docker restart](hostinger_vps_docker_restart.md)	 - Restart project
* [hostinger vps docker start](hostinger_vps_docker_start.md)	 - Start project
* [hostinger vps docker stop](hostinger_vps_docker_stop.md)	 - Stop project
* [hostinger vps docker update](hostinger_vps_docker_update.md)	 - Update project

