## hostinger dns records reset

Reset DNS records

### Synopsis

Reset DNS zone to the default records.

Use this endpoint to restore domain DNS to original configuration.

```
hostinger dns records reset <domain> [flags]
```

### Examples

```
  hostinger dns records reset mydomain.tld
```

### Options

```
  -h, --help                               help for reset
      --reset-email-records                Determines if email records should be reset (default true)
      --sync                               Determines if operation should be run synchronously (default true)
      --whitelisted-record-types strings   Specifies which record types to not reset
  -y, --yes                                Skip the confirmation prompt
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

* [hostinger dns records](hostinger_dns_records.md)	 - Zone commands

