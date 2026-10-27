## hostinger dns records update

Update DNS records

### Synopsis

Update DNS records for the selected domain.

Using `overwrite = true` will replace existing records with the provided ones.
Otherwise existing records will be updated and new records will be added.

Use this endpoint to modify domain DNS configuration.

The domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.

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
hostinger dns records update [domain] [flags]
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
      --config string     Config file (default is $HOME/.hostinger.yaml)
      --domain string     Website domain for commands that act on a website (default: HOSTINGER_DOMAIN, .hostinger/site.json, or default_domain in the config)
      --dry-run           Print the request as a curl command instead of sending it
  -o, --format string     Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string         Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input          Never prompt; fail when an argument, flag or confirmation is missing
      --order-id string   Order for commands that act on an order (default: HOSTINGER_ORDER_ID, .hostinger/site.json for hosting, or a prompt)
      --username string   Hosting account for commands that act on an account (default: HOSTINGER_USERNAME, the website from --domain, or .hostinger/site.json)
      --verbose           Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger dns records](hostinger_dns_records.md)	 - Zone commands

