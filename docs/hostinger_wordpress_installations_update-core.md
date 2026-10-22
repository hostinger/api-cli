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

### Examples

```
  hostinger wordpress installations update-core u123456789 1232456789
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
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger wordpress installations](hostinger_wordpress_installations.md)	 - Installations commands

