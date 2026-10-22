## hostinger domains portfolio disable-privacy-protection

Disable privacy protection

### Synopsis

Disable privacy protection for the domain.

When privacy protection is disabled, domain owner's personal information is visible in public WHOIS database.

Use this endpoint to make domain owner's information publicly visible.

```
hostinger domains portfolio disable-privacy-protection <domain> [flags]
```

### Examples

```
  hostinger domains portfolio disable-privacy-protection mydomain.tld
```

### Options

```
  -h, --help   help for disable-privacy-protection
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

* [hostinger domains portfolio](hostinger_domains_portfolio.md)	 - Portfolio commands

