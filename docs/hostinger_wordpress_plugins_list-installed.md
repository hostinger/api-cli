## hostinger wordpress plugins list-installed

List installed WordPress plugins

### Synopsis

List plugins installed on a WordPress installation, including their status,
available updates and known vulnerabilities.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress plugins list-installed <username> <software> [flags]
```

### Options

```
      --category string   Filter installed plugins by category. (one of: cache)
  -h, --help              help for list-installed
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress plugins](hostinger_wordpress_plugins.md)	 - Plugins commands

