## hostinger wordpress maintenance show-status

Show maintenance status

### Synopsis

Show the maintenance mode status for the specified WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress maintenance show-status <username> <software> [flags]
```

### Options

```
  -h, --help   help for show-status
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress maintenance](hostinger_wordpress_maintenance.md)	 - Maintenance commands

