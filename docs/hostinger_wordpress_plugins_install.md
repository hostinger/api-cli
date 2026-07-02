## hostinger wordpress plugins install

Install WordPress plugins

### Synopsis

Install one or more plugins on an existing WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id`
field). Use GET /api/hosting/v1/wordpress/plugins to discover the plugin
slugs available for installation.

This operation is asynchronous: a successful response only means the install
job has been queued, not that the plugins are ready.

```
hostinger wordpress plugins install <username> <software> [flags]
```

### Options

```
  -h, --help              help for install
      --plugins strings   Plugin slugs to install. Use GET /api/hosting/v1/wordpress/plugins to discover available slugs.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress plugins](hostinger_wordpress_plugins.md)	 - Plugins commands

