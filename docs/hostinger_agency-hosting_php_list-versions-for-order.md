## hostinger agency-hosting php list-versions-for-order

List available PHP versions for an order

### Synopsis

Lists the PHP versions available to websites created under an Agency Plan order, determined by the server the order is hosted on. Use this before creating a website; for a website that already exists, call the website-scoped versions endpoint instead.

```
hostinger agency-hosting php list-versions-for-order <order_id> [flags]
```

### Examples

```
  hostinger agency-hosting php list-versions-for-order 123456
```

### Options

```
  -h, --help   help for list-versions-for-order
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

* [hostinger agency-hosting php](hostinger_agency-hosting_php.md)	 - PHP commands

