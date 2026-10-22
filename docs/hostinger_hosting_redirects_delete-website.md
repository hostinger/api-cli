## hostinger hosting redirects delete-website

Delete website redirect

### Synopsis

Permanently deletes the redirect identified by its source URL.

Pass the `from` value exactly as returned by the list redirects endpoint.

```
hostinger hosting redirects delete-website <username> <domain> [flags]
```

### Examples

```
  hostinger hosting redirects delete-website u123456789 mydomain.tld --from https://example.com/old-page
```

### Options

```
      --from string   Source URL returned by the list redirects endpoint.
  -h, --help          help for delete-website
  -y, --yes           Skip the confirmation prompt
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

