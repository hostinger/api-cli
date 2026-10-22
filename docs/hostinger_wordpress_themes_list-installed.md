## hostinger wordpress themes list-installed

List installed WordPress themes

### Synopsis

List themes installed on a WordPress installation, including their status,
available updates and known vulnerabilities.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress themes list-installed <username> <software> [flags]
```

### Examples

```
  hostinger wordpress themes list-installed u123456789 1232456789
```

### Options

```
  -h, --help   help for list-installed
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

