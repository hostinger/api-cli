## hapi dns records validate

Validate DNS records

### Synopsis

Validate DNS records prior to update for the selected domain.

If the validation is successful, the response will contain `200 Success` code.
If there is validation error, the response will fail with `422 Validation error` code.

Use this endpoint to verify DNS record validity before applying changes.

```
hapi dns records validate <domain> [flags]
```

### Options

```
  -h, --help             help for validate
      --overwrite true   If true, resource records (RRs) matching name and type will be deleted and new RRs will be created,
                         otherwise resource records' ttl's are updated and new records are appended.
                         If no matching RRs are found, they are created. (default true)
      --zone string       (JSON)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi dns records](hapi_dns_records.md)	 - Zone commands

