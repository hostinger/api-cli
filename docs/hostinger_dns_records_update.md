## hostinger dns records update

Update DNS records

### Synopsis

Update DNS records for the selected domain.

Using `overwrite = true` will replace existing records with the provided ones. 
Otherwise existing records will be updated and new records will be added.

Use this endpoint to modify domain DNS configuration.

```
hostinger dns records update <domain> [flags]
```

### Options

```
  -h, --help             help for update
      --overwrite true   If true, resource records (RRs) matching name and type will be deleted and new RRs will be created,
                         otherwise resource records' ttl's are updated and new records are appended.
                         If no matching RRs are found, they are created. (default true)
      --zone string       (JSON)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger dns records](hostinger_dns_records.md)	 - Zone commands

