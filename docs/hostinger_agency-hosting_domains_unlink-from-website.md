## hostinger agency-hosting domains unlink-from-website

Unlink domain from website

### Synopsis

Unlinks a domain from the specified Agency Plan website.

The website stops serving traffic on this domain immediately.

Website files and database are preserved, and any other linked domains remain accessible.

If this is the only domain on the website, unlinking leaves the website without an accessible domain.

```
hostinger agency-hosting domains unlink-from-website <website_uid> <domain> [flags]
```

### Examples

```
  hostinger agency-hosting domains unlink-from-website zpwlGlp19 mydomain.tld
```

### Options

```
  -h, --help   help for unlink-from-website
  -y, --yes    Skip the confirmation prompt
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

* [hostinger agency-hosting domains](hostinger_agency-hosting_domains.md)	 - Domains commands

