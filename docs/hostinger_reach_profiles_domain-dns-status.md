## hostinger reach profiles domain-dns-status

Get profile domain DNS status

### Synopsis

Retrieve the DNS configuration status for a profile's domain.

This endpoint reports the state of MX, SPF, DKIM and DMARC records, including the
actual records found and the suggested records required for correct email delivery.

```
hostinger reach profiles domain-dns-status <profile-uuid> [flags]
```

### Examples

```
  hostinger reach profiles domain-dns-status 550e8400-e09b-41d4-a716-400055000000
```

### Options

```
  -h, --help   help for domain-dns-status
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

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

