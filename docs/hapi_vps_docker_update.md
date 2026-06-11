## hapi vps docker update

Update Docker project

### Synopsis

This endpoint updates a specified Docker Compose project on a virtual machine, pulling the latest images and recreating the containers.

```
hapi vps docker update <virtual machine ID> <project name> [flags]
```

### Options

```
  -h, --help   help for update
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps docker](hapi_vps_docker.md)	 - Docker project management

