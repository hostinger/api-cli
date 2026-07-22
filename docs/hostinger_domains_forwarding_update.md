## hostinger domains forwarding update

Update domain forwarding

### Synopsis

Update domain forwarding configuration.

Use this endpoint to modify existing redirect configuration for domains.

```
hostinger domains forwarding update <domain> [flags]
```

### Options

```
  -h, --help                   help for update
      --redirect-type string   Redirect type (one of: 301, 302)
      --redirect-url string    URL to forward domain to
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger domains forwarding](hostinger_domains_forwarding.md)	 - Forwarding commands

