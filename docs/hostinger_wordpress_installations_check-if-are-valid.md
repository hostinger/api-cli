## hostinger wordpress installations check-if-are-valid

Check if WordPress installations are valid

### Synopsis

Check whether one or more WordPress installations are valid and working
correctly. Detects broken installations caused by missing files, broken
plugins, themes and similar issues.

Provide the WordPress installation (software) identifiers in the body. They
can be obtained from GET /api/hosting/v1/wordpress/installations (the `id`
field).

```
hostinger wordpress installations check-if-are-valid <username> [flags]
```

### Options

```
      --force                  Force fresh validation without cache. Preferable for troubleshooting purposes.
  -h, --help                   help for check-if-are-valid
      --software-ids strings   WordPress installation (software) identifiers to validate.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger wordpress installations](hostinger_wordpress_installations.md)	 - Installations commands

