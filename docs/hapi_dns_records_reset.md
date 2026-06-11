## hapi dns records reset

Reset DNS records

### Synopsis

This endpoint resets DNS zone records to the default state for a specific domain.

```
hapi dns records reset <domain> [flags]
```

### Options

```
  -h, --help                                   help for reset
      --reset-email-records                    Determines if email records should be reset
      --sync                                   Determines if operation should be run synchronously
      --whitelisted-record-types stringArray   Record types to not reset (repeatable)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi dns records](hapi_dns_records.md)	 - DNS zone record management

