## hostinger domains portfolio renewal-information

Get domain renewal information

### Synopsis

Retrieve renewal information for a specified domain, including its status and current
expiration date.

Use this endpoint to build renewal automation and expiry monitoring for a single domain.

```
hostinger domains portfolio renewal-information <domain> [flags]
```

### Examples

```
  hostinger domains portfolio renewal-information mydomain.tld
```

### Options

```
  -h, --help   help for renewal-information
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

* [hostinger domains portfolio](hostinger_domains_portfolio.md)	 - Portfolio commands

