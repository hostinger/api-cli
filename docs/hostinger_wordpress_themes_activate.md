## hostinger wordpress themes activate

Activate WordPress theme

### Synopsis

Activate an installed theme on a WordPress installation.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

This operation is asynchronous: a successful response only means the activation
job has been queued.

```
hostinger wordpress themes activate <username> <software> [flags]
```

### Examples

```
  hostinger wordpress themes activate u123456789 1232456789 --theme twentytwentyone
```

### Options

```
  -h, --help           help for activate
      --theme string   Slug of the installed theme to activate.
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

