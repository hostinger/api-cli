## hapi dns records update

Update DNS records

### Synopsis

This endpoint updates DNS records for the selected domain.

Using `overwrite = true` in the request will replace existing records matching the name and type with the provided
records. When `false`, provided records are appended and existing records' TTLs are updated.

```
hapi dns records update <domain> [flags]
```

### Options

```
      --content stringArray   Content of the record (repeatable to assign multiple records to the name)
  -h, --help                  help for update
      --name string           Name of the record (use @ for wildcard name)
      --overwrite             Replace records matching name and type instead of appending
      --ttl int               TTL (Time-To-Live) of the record
      --type string           Type of the record (one of: A, AAAA, ALIAS, CAA, CNAME, MX, NS, SOA, SRV, TXT)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi dns records](hapi_dns_records.md)	 - DNS zone record management

