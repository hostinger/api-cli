## hostinger hosting redirects list-website

List website redirects

### Synopsis

Returns a paginated list of redirects configured for the selected website.

```
hostinger hosting redirects list-website <username> <domain> [flags]
```

### Examples

```
  hostinger hosting redirects list-website u123456789 mydomain.tld
```

### Options

```
      --all            Fetch every page and print all items
  -h, --help           help for list-website
      --page int       Page number
      --per-page int   Number of items per page (default 25)
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

