## hostinger wordpress plugins list

List available WordPress plugins

### Synopsis

List plugins recommended for installation on a WordPress installation that are
not yet installed.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress plugins list <username> <software> [flags]
```

### Examples

```
  hostinger wordpress plugins list u123456789 1232456789
```

### Options

```
  -h, --help   help for list
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

* [hostinger wordpress plugins](hostinger_wordpress_plugins.md)	 - Plugins commands

