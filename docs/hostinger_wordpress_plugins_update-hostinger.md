## hostinger wordpress plugins update-hostinger

Update Hostinger WordPress plugin

### Synopsis

Update a Hostinger plugin to its latest version on a WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

This operation is asynchronous: a successful response only means the update job
has been queued.

```
hostinger wordpress plugins update-hostinger <username> <software> [flags]
```

### Examples

```
  hostinger wordpress plugins update-hostinger u123456789 1232456789 --slug hostinger-affiliate-plugin
```

### Options

```
  -h, --help          help for update-hostinger
      --slug string   Slug of the Hostinger plugin to update to its latest version. (one of: hostinger, hostinger-ai-assistant, hostinger-affiliate-plugin, hostinger-easy-onboarding, hostinger-reach)
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

