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

### Options

```
  -h, --help   help for get
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps docker](hostinger_vps_docker.md)	 - Docker Manager commands

