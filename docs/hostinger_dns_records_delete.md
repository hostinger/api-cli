## hostinger dns records delete

Delete DNS records

### Synopsis

Delete DNS records for the selected domain.

To filter which records to delete, add the `name` of the record and `type` to the filter.
Multiple filters can be provided with single request.

If you have multiple records with the same name and type, and you want to delete only part of them,
refer to the `Update zone records` endpoint.

Use this endpoint to remove specific DNS records from domains.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --filters
    [
      {
        "name": "@",
        "type": "A"
      }
    ]

```
hostinger dns records delete <domain> [flags]
```

### Examples

```
  hostinger dns records delete mydomain.tld --filters '[{"name":"@","type":"A"}]'
```

### Options

```
      --filters string   Filter records for deletion (JSON, @file or -; shape in the help above)
  -h, --help             help for delete
  -y, --yes              Skip the confirmation prompt
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

