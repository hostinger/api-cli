## hostinger wordpress themes update

Update WordPress themes

### Synopsis

Update one or more installed themes to their latest version on a WordPress
installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

This operation is asynchronous: a successful response only means the update job
has been queued.

```
hostinger wordpress themes update <username> <software> [flags]
```

### Options

```
  -h, --help             help for update
      --themes strings   Slugs of the installed themes to update to their latest version.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress themes](hostinger_wordpress_themes.md)	 - Themes commands

