## hostinger domains whois set-as-default

Set WHOIS profile as default

### Synopsis

Set WHOIS contact profile as default.

The default profile is pre-selected for the TLD it belongs to when registering new domains.

Use this endpoint to avoid picking contact information for every registration.

```
hostinger domains whois set-as-default <whois-id> [flags]
```

### Examples

```
  hostinger domains whois set-as-default 564651
```

### Options

```
  -h, --help   help for set-as-default
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

* [hostinger domains whois](hostinger_domains_whois.md)	 - WHOIS commands

