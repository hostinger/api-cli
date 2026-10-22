## hostinger hosting php update-extensions

Update PHP extensions

### Synopsis

Enables or disables PHP extensions (modules) for the website.

Use the Get PHP details endpoint to check the current extension states before changing them.

```
hostinger hosting php update-extensions <username> <domain> [flags]
```

### Examples

```
  hostinger hosting php update-extensions u123456789 mydomain.tld
```

### Options

```
      --disable strings   PHP extensions to disable.
      --enable strings    PHP extensions to enable.
  -h, --help              help for update-extensions
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

* [hostinger hosting php](hostinger_hosting_php.md)	 - PHP commands

