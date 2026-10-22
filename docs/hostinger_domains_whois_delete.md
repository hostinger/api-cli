## hostinger domains whois delete

Delete WHOIS profile

### Synopsis

Delete WHOIS contact profile.

Use this endpoint to remove unused contact profiles from account.

```
hostinger domains whois delete <whois-id> [flags]
```

### Examples

```
  hostinger domains whois delete 564651
```

### Options

```
  -h, --help   help for delete
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

