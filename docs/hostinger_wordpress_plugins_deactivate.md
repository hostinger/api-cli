## hostinger wordpress plugins deactivate

Deactivate WordPress plugin

### Synopsis

Deactivate an installed plugin on a WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

This operation is asynchronous: a successful response only means the
deactivation job has been queued.

```
hostinger wordpress plugins deactivate <username> <software> [flags]
```

### Options

```
  -h, --help            help for deactivate
      --plugin string   Slug of the installed plugin to deactivate.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress plugins](hostinger_wordpress_plugins.md)	 - Plugins commands

