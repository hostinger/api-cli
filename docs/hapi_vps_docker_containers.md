## hapi vps docker containers

Get project containers

### Synopsis

Retrieves a list of all containers belonging to a specific Docker Compose project on the virtual machine. 

This endpoint returns detailed information about each container including
their current status, port mappings, and runtime configuration.

Use this to monitor the health and state of all services within your Docker Compose project.

```
hapi vps docker containers <virtual-machine-id> <project-name> [flags]
```

### Options

```
  -h, --help   help for containers
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps docker](hapi_vps_docker.md)	 - Docker Manager commands

