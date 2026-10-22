## hostinger domains forwarding update

Update domain forwarding

### Synopsis

Update domain forwarding configuration.

Use this endpoint to modify existing redirect configuration for domains.

```
hostinger domains forwarding update <domain> [flags]
```

### Examples

```
  hostinger domains forwarding update mydomain.tld --redirect-type 301 --redirect-url https://forward.to.my.url
```

### Options

```
  -h, --help                   help for update
      --redirect-type string   Redirect type (one of: 301, 302)
      --redirect-url string    URL to forward domain to
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

* [hostinger domains forwarding](hostinger_domains_forwarding.md)	 - Forwarding commands

