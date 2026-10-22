## hostinger domains forwarding create

Create domain forwarding

### Synopsis

Create domain forwarding configuration.

Use this endpoint to set up domain redirects to other URLs.

```
hostinger domains forwarding create [flags]
```

### Examples

```
  hostinger domains forwarding create --domain mydomain.tld --redirect-type 301 --redirect-url https://forward.to.my.url
```

### Options

```
      --domain string          Domain name
  -h, --help                   help for create
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

