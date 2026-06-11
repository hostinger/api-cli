## hapi dns records delete

Delete DNS records

### Synopsis

This endpoint deletes DNS zone records for a specific domain, filtered by record name and type.

```
hapi dns records delete <domain> [flags]
```

### Options

```
      --filter stringArray   Record to delete in format <name>:<type> (repeatable). Type one of: A, AAAA, ALIAS, CAA, CNAME, MX, NS, SOA, SRV, TXT
  -h, --help                 help for delete
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi dns records](hapi_dns_records.md)	 - DNS zone record management

