## hostinger vps docker restart

Restart project

### Synopsis

Restarts all services in a Docker Compose project by stopping and starting
containers in the correct dependency order.

This operation preserves data volumes and network configurations while refreshing the running containers. 

Use this to apply configuration changes or recover from service failures.

```
hostinger vps docker restart <virtual-machine-id> <project-name> [flags]
```

### Options

```
  -h, --help   help for restart
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger vps docker](hostinger_vps_docker.md)	 - Docker Manager commands

