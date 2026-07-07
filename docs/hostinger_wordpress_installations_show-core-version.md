## hostinger wordpress installations show-core-version

Show WordPress core version

### Synopsis

Show the WordPress core version for the specified installation, along with
known vulnerabilities affecting it.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress installations show-core-version <username> <software> [flags]
```

### Options

```
  -h, --help   help for show-core-version
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress installations](hostinger_wordpress_installations.md)	 - Installations commands

