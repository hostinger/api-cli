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

### Options

```
  -h, --help   help for logs
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps docker](hostinger_vps_docker.md)	 - Docker Manager commands

