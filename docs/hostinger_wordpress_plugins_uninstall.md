## hostinger wordpress plugins uninstall

Uninstall WordPress plugins

### Synopsis

Uninstall one or more plugins from a WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

This operation is asynchronous: a successful response only means the uninstall
job has been queued.

```
hostinger wordpress plugins uninstall <username> <software> [flags]
```

### Options

```
  -h, --help              help for uninstall
      --plugins strings   Slugs of the installed plugins to uninstall.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress plugins](hostinger_wordpress_plugins.md)	 - Plugins commands

