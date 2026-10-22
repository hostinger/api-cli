## hostinger domains whois cancel-pending-irtp-verification

Cancel pending IRTP verification

### Synopsis

Cancel a pending IRTP verification.

Use this endpoint to back out of a WHOIS change that is stuck waiting on registrant confirmation,
for example when the confirmation email cannot be received, without waiting out the 5-day expiry.

```
hostinger domains whois cancel-pending-irtp-verification <domain> [flags]
```

### Examples

```
  hostinger domains whois cancel-pending-irtp-verification mydomain.tld
```

### Options

```
  -h, --help   help for cancel-pending-irtp-verification
  -y, --yes    Skip the confirmation prompt
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

