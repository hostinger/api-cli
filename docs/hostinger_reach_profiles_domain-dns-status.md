## hostinger reach profiles domain-dns-status

Get profile domain DNS status

### Synopsis

Retrieve the DNS configuration status for a profile's domain.

This endpoint reports the state of MX, SPF, DKIM and DMARC records, including the
actual records found and the suggested records required for correct email delivery.

```
hostinger reach profiles domain-dns-status <profile-uuid> [flags]
```

### Options

```
  -h, --help   help for domain-dns-status
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger reach profiles](hostinger_reach_profiles.md)	 - Profiles commands

