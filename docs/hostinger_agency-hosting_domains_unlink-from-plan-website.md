## hostinger agency-hosting domains unlink-from-plan-website

Unlink domain from Agency Plan website

### Synopsis

Unlinks a domain from the specified Agency Plan website.

The website stops serving traffic on this domain immediately.

Website files and database are preserved, and any other linked domains remain accessible.

If this is the only domain on the website, unlinking leaves the website without an accessible domain.

```
hostinger agency-hosting domains unlink-from-plan-website <website_uid> <domain> [flags]
```

### Options

```
  -h, --help   help for unlink-from-plan-website
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger agency-hosting domains](hostinger_agency-hosting_domains.md)	 - Domains commands

