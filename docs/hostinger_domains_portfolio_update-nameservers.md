## hostinger domains portfolio update-nameservers

Update domain nameservers

### Synopsis

Set nameservers for a specified domain.

Be aware, that improper nameserver configuration can lead to the domain being unresolvable or unavailable.

Use this endpoint to configure custom DNS hosting for domains.

```
hostinger domains portfolio update-nameservers <domain> [flags]
```

### Examples

```
  hostinger domains portfolio update-nameservers mydomain.tld --ns1 ns1.some-nameserver.tld --ns2 ns2.some-nameserver.tld
```

### Options

```
  -h, --help         help for update-nameservers
      --ns1 string   First name server
      --ns2 string   Second name server
      --ns3 string   Third name server
      --ns4 string   Fourth name server
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

