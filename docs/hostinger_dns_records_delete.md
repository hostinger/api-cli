## hostinger dns records delete

Delete DNS records

### Synopsis

Delete DNS records for the selected domain.

To filter which records to delete, add the `name` of the record and `type` to the filter. 
Multiple filters can be provided with single request.

If you have multiple records with the same name and type, and you want to delete only part of them,
refer to the `Update zone records` endpoint.

Use this endpoint to remove specific DNS records from domains.

```
hostinger dns records delete <domain> [flags]
```

### Options

```
      --filters string   Filter records for deletion (JSON)
  -h, --help             help for delete
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger dns records](hostinger_dns_records.md)	 - Zone commands

