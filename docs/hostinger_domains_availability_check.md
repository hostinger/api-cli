## hostinger domains availability check

Check domain availability

### Synopsis

Check availability of domain names across multiple TLDs.

Multiple TLDs can be checked at once.
If you want alternative domains with response, provide only one TLD and set `with_alternatives` to `true`.
TLDs should be provided without leading dot (e.g. `com`, `net`, `org`).

Endpoint has rate limit of 90 requests per minute.

Use this endpoint to verify domain availability before purchase.

```
hostinger domains availability check [flags]
```

### Examples

```
  hostinger domains availability check --domain mydomain --tlds com,net,org
```

### Options

```
      --domain string       Domain name (without TLD)
  -h, --help                help for check
      --tlds strings        TLDs list
      --with-alternatives   Should response include alternatives
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

* [hostinger domains availability](hostinger_domains_availability.md)	 - Availability commands

