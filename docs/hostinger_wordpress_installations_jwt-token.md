## hostinger wordpress installations jwt-token

Get installation JWT token

### Synopsis

Return a JWT token used to authenticate requests against the specified
WordPress installation, including its MCP (Model Context Protocol) endpoint.

Provide the WordPress installation (software) identifier in the path. It can
be obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).

```
hostinger wordpress installations jwt-token <username> <software> [flags]
```

### Examples

```
  hostinger wordpress installations jwt-token u123456789 1232456789
```

### Options

```
  -h, --help   help for jwt-token
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

