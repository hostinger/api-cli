## hostinger wordpress themes activate

Activate WordPress theme

### Synopsis

Activate an installed theme on a WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

This operation is asynchronous: a successful response only means the activation
job has been queued.

```
hostinger wordpress themes activate <username> <software> [flags]
```

### Options

```
  -h, --help           help for activate
      --theme string   Slug of the installed theme to activate.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress themes](hostinger_wordpress_themes.md)	 - Themes commands

