## hostinger dns records validate

Validate DNS records

### Synopsis

Validate DNS records prior to update for the selected domain.

If the validation is successful, the response will contain `200 Success` code.
If there is validation error, the response will fail with `422 Validation error` code.

Use this endpoint to verify DNS record validity before applying changes.

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
hostinger dns records validate <domain> [flags]
```

### Examples

```
  hostinger dns records validate mydomain.tld --zone @zone.json
```

### Options

```
  -h, --help             help for validate
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

