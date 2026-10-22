## hostinger wordpress themes update

Update WordPress themes

### Synopsis

Update one or more installed themes to their latest version on a WordPress
installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

This operation is asynchronous: a successful response only means the update job
has been queued.

```
hostinger wordpress themes update <username> <software> [flags]
```

### Examples

```
  hostinger wordpress themes update u123456789 1232456789 --themes twentytwenty,twentytwentyone
```

### Options

```
  -h, --help             help for update
      --themes strings   Slugs of the installed themes to update to their latest version.
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

* [hostinger wordpress themes](hostinger_wordpress_themes.md)	 - Themes commands

