## hostinger wordpress plugins install

Install WordPress plugins

### Synopsis

Install one or more plugins on an existing WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id`
field). Use GET /api/hosting/v1/wordpress/plugins to discover the plugin
slugs available for installation.

This operation is asynchronous: a successful response only means the install
job has been queued, not that the plugins are ready.

```
hostinger wordpress plugins install <username> <software> [flags]
```

### Examples

```
  hostinger wordpress plugins install u123456789 1232456789 --plugins akismet,hello-dolly
```

### Options

```
  -h, --help              help for install
      --plugins strings   Plugin slugs to install. Use GET /api/hosting/v1/wordpress/plugins to discover available slugs.
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

