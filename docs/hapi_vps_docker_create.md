## hapi vps docker create

Create Docker project

### Synopsis

This endpoint creates and starts a new Docker Compose project on a specified virtual machine.

```
hapi vps docker create <virtual machine ID> [flags]
```

### Options

```
      --content string        URL to a docker-compose.yaml file, a GitHub repository, or raw YAML content of the compose file
      --environment string    Project environment variables
  -h, --help                  help for create
      --project-name string   Docker Compose project name (alphanumeric characters, dashes and underscores only)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps docker](hapi_vps_docker.md)	 - Docker project management

