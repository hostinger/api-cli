## hostinger hosting domains verify-ownership

Verify domain ownership

### Synopsis

Verify ownership of a single domain and return the verification status.

Use this endpoint to check if a domain is accessible for you before using it for new websites.
If the domain is accessible, the response will have `is_accessible: true`.
If not, add the given TXT record to your domain's DNS records and try verifying again.
Keep in mind that it may take up to 10 minutes for new TXT DNS records to propagate.

Skip this verification when using Hostinger's free subdomains (*.hostingersite.com).

```
hostinger hosting domains verify-ownership [flags]
```

### Options

```
      --domain string   Domain to verify ownership for
  -h, --help            help for verify-ownership
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger hosting domains](hostinger_hosting_domains.md)	 - Domains commands

