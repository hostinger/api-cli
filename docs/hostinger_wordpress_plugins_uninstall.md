## hostinger wordpress plugins uninstall

Uninstall WordPress plugins

### Synopsis

Uninstall one or more plugins from a WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

This operation is asynchronous: a successful response only means the uninstall
job has been queued.

```
hostinger wordpress plugins uninstall <username> <software> [flags]
```

### Examples

```
  hostinger wordpress plugins uninstall u123456789 1232456789 --plugins akismet,hello-dolly
```

### Options

```
  -h, --help              help for uninstall
      --plugins strings   Slugs of the installed plugins to uninstall.
  -y, --yes               Skip the confirmation prompt
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

