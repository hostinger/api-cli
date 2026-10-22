## hostinger wordpress installations detect

Detect WordPress installations

### Synopsis

Trigger a background scan to detect WordPress installations for the account.

This operation is asynchronous: a successful response only means the scan has
been queued. Poll GET /api/hosting/v1/wordpress/installations to fetch the
detected installations once the scan completes.

```
hostinger wordpress installations detect <username> [flags]
```

### Examples

```
  hostinger wordpress installations detect u123456789
```

### Options

```
  -h, --help   help for detect
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

