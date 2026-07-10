## hostinger agency-hosting domains change-plan-website

Change Agency Plan website domain

### Synopsis

Changes the primary domain for an Agency Plan website.

Provide the current domain in the path and the new domain in the request body.
Set domain to null to revert to the temporary domain.

```
hostinger agency-hosting domains change-plan-website <website_uid> <from_domain> [flags]
```

### Options

```
      --domain string   New domain to assign to the website. Set to null to revert to the temporary domain.
  -h, --help            help for change-plan-website
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting domains](hostinger_agency-hosting_domains.md)	 - Domains commands

