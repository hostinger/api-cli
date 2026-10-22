## hostinger hosting redirects create-website

Create website redirect

### Synopsis

Creates a redirect from a URL on the selected website to another URL or IP address.

```
hostinger hosting redirects create-website <username> <domain> [flags]
```

### Examples

```
  hostinger hosting redirects create-website u123456789 mydomain.tld --from https://example.com/old-page --to https://example.com/new-page
```

### Options

```
      --from string   Source URL on the selected website
  -h, --help          help for create-website
      --to string     Destination URL or IP address
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

* [hostinger hosting redirects](hostinger_hosting_redirects.md)	 - Redirects commands

