## hapi vps docker update

Update project

### Synopsis

Updates a Docker Compose project by pulling the latest image versions and
recreating containers with new configurations.

This operation preserves data volumes while applying changes from the compose file. 

Use this to deploy application updates, apply configuration changes, or
refresh container images to their latest versions.

```
hapi vps docker update <virtual-machine-id> <project-name> [flags]
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

* [hapi vps docker](hapi_vps_docker.md)	 - Docker Manager commands

