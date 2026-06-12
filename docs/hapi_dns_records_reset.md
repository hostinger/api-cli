## hapi dns records reset

Reset DNS records

### Synopsis

Reset DNS zone to the default records.

Use this endpoint to restore domain DNS to original configuration.

```
hapi dns records reset <domain> [flags]
```

### Options

```
  -h, --help                               help for reset
      --reset-email-records                Determines if email records should be reset (default true)
      --sync                               Determines if operation should be run synchronously (default true)
      --whitelisted-record-types strings   Specifies which record types to not reset
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi dns records](hapi_dns_records.md)	 - Zone commands

