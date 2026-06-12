## hapi domains availability check

Check domain availability

### Synopsis

Check availability of domain names across multiple TLDs.

Multiple TLDs can be checked at once.
If you want alternative domains with response, provide only one TLD and set `with_alternatives` to `true`.
TLDs should be provided without leading dot (e.g. `com`, `net`, `org`).

Endpoint has rate limit of 10 requests per minute.

Use this endpoint to verify domain availability before purchase.

```
hapi domains availability check [flags]
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
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi domains availability](hapi_domains_availability.md)	 - Availability commands

