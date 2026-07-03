## hostinger wordpress themes uninstall

Uninstall WordPress themes

### Synopsis

Uninstall one or more themes from a WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

This operation is asynchronous: a successful response only means the uninstall
job has been queued.

```
hostinger wordpress themes uninstall <username> <software> [flags]
```

### Options

```
  -h, --help             help for uninstall
      --themes strings   Slugs of the installed themes to uninstall.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress themes](hostinger_wordpress_themes.md)	 - Themes commands

