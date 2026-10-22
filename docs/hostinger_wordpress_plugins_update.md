## hostinger wordpress plugins update

Update WordPress plugins

### Synopsis

Update one or more installed plugins to their latest version on a WordPress
installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

This operation is asynchronous: a successful response only means the update job
has been queued.

```
hostinger wordpress plugins update <username> <software> [flags]
```

### Examples

```
  hostinger wordpress plugins update u123456789 1232456789 --plugins akismet,hello-dolly
```

### Options

```
  -h, --help              help for update
      --plugins strings   Slugs of the installed plugins to update to their latest version.
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

