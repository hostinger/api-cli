## hapi domains forwarding create

Create domain forwarding

### Synopsis

Create domain forwarding configuration.

Use this endpoint to set up domain redirects to other URLs.

```
hapi domains forwarding create [flags]
```

### Options

```
      --domain string          Domain name
  -h, --help                   help for create
      --redirect-type string   Redirect type (one of: 301, 302)
      --redirect-url string    URL to forward domain to
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi domains forwarding](hapi_domains_forwarding.md)	 - Forwarding commands

