## hostinger wordpress installations list-core-updates

List available WordPress core updates

### Synopsis

List available WordPress core updates for the specified installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress installations list-core-updates <username> <software> [flags]
```

### Options

```
  -h, --help   help for list-core-updates
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress installations](hostinger_wordpress_installations.md)	 - Installations commands

