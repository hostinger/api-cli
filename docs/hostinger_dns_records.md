## hostinger dns records

Zone commands

```
hostinger dns records [flags]
```

### Options

```
  -h, --help   help for records
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

* [hostinger dns](hostinger_dns.md)	 - DNS commands
* [hostinger dns records delete](hostinger_dns_records_delete.md)	 - Delete DNS records
* [hostinger dns records list](hostinger_dns_records_list.md)	 - Get DNS records
* [hostinger dns records reset](hostinger_dns_records_reset.md)	 - Reset DNS records
* [hostinger dns records update](hostinger_dns_records_update.md)	 - Update DNS records
* [hostinger dns records validate](hostinger_dns_records_validate.md)	 - Validate DNS records

