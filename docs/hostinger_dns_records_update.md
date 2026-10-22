## hostinger dns records update

Update DNS records

### Synopsis

Update DNS records for the selected domain.

Using `overwrite = true` will replace existing records with the provided ones.
Otherwise existing records will be updated and new records will be added.

Use this endpoint to modify domain DNS configuration.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --zone
    [
      {
        "name": "www",
        "records": [
          {
            "content": "mydomain.tld."
          }
        ],
        "ttl": 14400,
        "type": "A"
      }
    ]

```
hostinger dns records update <domain> [flags]
```

### Examples

```
  hostinger dns records update mydomain.tld --zone @zone.json
```

### Options

```
  -h, --help             help for update
      --overwrite true   If true, resource records (RRs) matching name and type will be deleted and new RRs will be created, otherwise resource records' ttl's are updated and new records are appended. If no matching RRs are found, they are created. (default true)
      --zone string      (JSON, @file or -; shape in the help above)
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

