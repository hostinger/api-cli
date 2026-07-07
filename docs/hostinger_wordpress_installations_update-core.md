## hostinger wordpress installations update-core

Update WordPress core

### Synopsis

Update the WordPress core for the specified installation (minor update or a
specific version).

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

This operation is asynchronous: a successful response only means the update
job has been queued.

```
hostinger wordpress installations update-core <username> <software> [flags]
```

### Options

```
  -h, --help             help for update-core
      --minor            Update the minor version only.
      --version string   Update to a specific WordPress core version.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress installations](hostinger_wordpress_installations.md)	 - Installations commands

