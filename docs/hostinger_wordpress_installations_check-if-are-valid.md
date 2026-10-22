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

### Examples

```
  hostinger wordpress installations check-if-are-valid u123456789 --software-ids 123,456
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
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger wordpress installations](hostinger_wordpress_installations.md)	 - Installations commands

