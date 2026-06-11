## hapi billing catalog list

Get catalog item list

### Synopsis

This endpoint retrieves a list of available catalog items that can be ordered.

Prices in the response are displayed in cents.

```
hapi billing catalog list [flags]
```

### Options

```
      --category string   Filter catalog items by category (one of: DOMAIN, VPS)
  -h, --help              help for list
      --name string       Filter catalog items by name. Use * for wildcard search, e.g. .COM* to find .com domain
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi billing catalog](hapi_billing_catalog.md)	 - Catalog management

