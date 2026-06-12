## hapi vps docker stop

Stop project

### Synopsis

Stops all running services in a Docker Compose project while preserving
container configurations and data volumes.

This operation gracefully shuts down containers in reverse dependency order. 

Use this to temporarily halt a project without removing data or configurations.

```
hapi vps docker stop <virtual-machine-id> <project-name> [flags]
```

### Options

```
  -h, --help   help for stop
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps docker](hapi_vps_docker.md)	 - Docker Manager commands

