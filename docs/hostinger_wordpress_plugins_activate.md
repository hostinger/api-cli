## hostinger wordpress plugins activate

Activate WordPress plugin

### Synopsis

Activate an installed plugin on a WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

This operation is asynchronous: a successful response only means the activation
job has been queued.

```
hostinger wordpress plugins activate <username> <software> [flags]
```

### Options

```
  -h, --help            help for activate
      --plugin string   Slug of the installed plugin to activate.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress plugins](hostinger_wordpress_plugins.md)	 - Plugins commands

