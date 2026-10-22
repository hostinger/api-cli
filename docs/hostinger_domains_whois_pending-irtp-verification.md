## hostinger domains whois pending-irtp-verification

Get pending IRTP verification

### Synopsis

Retrieve a pending IRTP verification for a domain.

Both the old and new registrant must confirm it before the WHOIS change takes effect.

Use this endpoint to check the status of a WHOIS change awaiting registrant confirmation.

```
hostinger domains whois pending-irtp-verification <domain> [flags]
```

### Examples

```
  hostinger domains whois pending-irtp-verification mydomain.tld
```

### Options

```
  -h, --help   help for pending-irtp-verification
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

