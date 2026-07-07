## hostinger wordpress maintenance toggle

Toggle maintenance mode

### Synopsis

Enable or disable maintenance mode for the specified WordPress installation,
based on the `enabled` flag.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress maintenance toggle <username> <software> [flags]
```

### Options

```
      --enabled   Enable (true) or disable (false) maintenance mode for the WordPress installation.
  -h, --help      help for toggle
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress maintenance](hostinger_wordpress_maintenance.md)	 - Maintenance commands

