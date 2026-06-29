## hostinger domains portfolio update-nameservers

Update domain nameservers

### Synopsis

Set nameservers for a specified domain.

Be aware, that improper nameserver configuration can lead to the domain being unresolvable or unavailable.

Use this endpoint to configure custom DNS hosting for domains.

```
hostinger domains portfolio update-nameservers <domain> [flags]
```

### Options

```
  -h, --help         help for update-nameservers
      --ns1 string   First name server
      --ns2 string   Second name server
      --ns3 string   Third name server
      --ns4 string   Fourth name server
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger domains portfolio](hostinger_domains_portfolio.md)	 - Portfolio commands

